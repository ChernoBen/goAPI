package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/ChernoBen/bookstore/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	//Validação de entrada
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//converter string para byte of string
	var pass = []byte(input.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pass"})
	}
	//criação
	user := models.User{Email: input.Email, Password: string(hash)} //convertendo byte of string em string
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

type UserAuth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var userEntry UserAuth
	if err := c.BindJSON(&userEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user models.User
	//trata se o id nao existe
	if err := models.DB.Where("email = ?", userEntry.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println(user.Email)
	var hashed = []byte(user.Password)
	er := bcrypt.CompareHashAndPassword(hashed, []byte(userEntry.Password))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)
	token, err := at.SignedString([]byte("chernoBen"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	c.JSON(http.StatusOK, gin.H{"data": token})
}
