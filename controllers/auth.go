package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
	var body models.CreateUserInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//Create User
	user := models.User{
		PhoneNumber: body.PhoneNumber,
		Password:    string(hash),
		Lastname:    body.Lastname,
		Firstname:   body.Firstname,
		Role:        body.Role,
		Gender:      body.Gender,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{"data": "User created!"})
}

func SignIn(c *gin.Context) {
	var body models.LoginUserInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "phone_number = ?", body.PhoneNumber)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phoneNumber",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func Validate(c *gin.Context) {
	user, err := c.Get("user")

	if !err {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetUserDetailsFromToken(c *gin.Context) {
	// retrieve the token from the cookie
	tokenCookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		// handle error
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "auth token not found"})
		return
	}
	tokenString := tokenCookie.Value

	// verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// add validation for signature and issuer
		// you can use a constant secret key or use a dynamic one loaded from the environment variable
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid auth token"})
		return
	}

	// decode the token and access the user details
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// do something with the user details
		c.JSON(http.StatusOK, gin.H{"claims": claims})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid auth token"})
		return
	}
}
