package controllers

import (
	"fmt"
	"net/http"

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
		return
	}
	//converter string para byte of string
	var pass = []byte(input.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pass"})
		return
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

type Claims struct {
	ID         uint `json:"ID"`
	Authorized bool `json:"authorized"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	privateKey := "oasijdoaidsoijs"
	var userEntry UserAuth
	if err := c.BindJSON(&userEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	//trata se o id nao existe
	fmt.Println(userEntry.Email)
	if err := models.DB.Where("email = ?", userEntry.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println(user.Email)
	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userEntry.Password))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	//atClaims := &Claims{
	//	ID:         user.ID,
	//	Authorized: true,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	//	},
	//}
	at := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(at)
	token, erro := at.SignedString([]byte(privateKey))
	fmt.Println(erro)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": token})

}
