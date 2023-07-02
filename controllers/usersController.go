package controllers

import (
	"go-gin/initializers"
	"go-gin/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass of request body
	var body struct {
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Failed to hash password",
		})
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Failed to create user",
		})
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {

	// Get the email and pass of request body
	var body struct {
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Failed to read body",
		})
		return
	}

	// Look up request
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid Email or Password",
		})
	}

	// Compare sent in pass with saved user password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid Email or Password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), 
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Failed to create token",
		})
		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H {})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}