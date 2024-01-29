package controller

import (
	"brewery/database"
	"brewery/model"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

// InitializeDB initializes the database and returns the DB connection.
func InitializeDB() *gorm.DB {
	DB := database.InitializeDB()
	return DB
}

var DB *gorm.DB = InitializeDB()

var Otp string

func HandlePostUser(c *gin.Context) {

	var user model.User

	c.BindJSON(&user)

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	var namecount int64

	DB.Model(&model.User{}).Where("first_name = ?", user.FirstName).Count(&namecount)

	if namecount > 0 {
		c.JSON(404, gin.H{
			"message": "This Username Already Exists",
		})
		return
	}

	var count int64

	DB.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)

	if count > 0 {
		c.JSON(404, gin.H{
			"message": "This Email Already Exists",
		})
		return
	}

	err := model.ValidateUser(&user)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	if user.DateOfBirth.After(time.Now()) {
		c.JSON(404, gin.H{
			"message": "please provide valid date of birth",
		})
		return
	}

	resultUser, err := handlePostUser(&user)

	if err != nil {

		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return

	}
	c.JSON(200, gin.H{
		"user": resultUser,
	})

}

func Handletoken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	params := mux.Vars(r)
	id := params["id"]
	DB.Find(&user, id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.FirstName,
		"userID":   user.ID,
	})

	fmt.Println(token)

	tokenString, err := token.SignedString([]byte("secret-key"))

	fmt.Println(tokenString)

	if err != nil {
		return
	}

	w.Write([]byte(tokenString))
}

// handlePostUser handles creating a new user.
func handlePostUser(user *model.User) (*model.User, error) {

	DB.AutoMigrate(&model.User{})

	if err := DB.Create(user).Error; err != nil {
		return nil, errors.New("please provide valid input")
	}

	return user, nil
}

func HandleSendEmail(c *gin.Context) {

	id := c.Param("id")

	var user model.User

	DB.Where("id = ?", id).First(&user)

	handleSendEmail(user.Email)

	c.JSON(200, gin.H{
		"message": "email sent successfully",
	})

}

func generateOtp() {

	rand.Seed(time.Now().UnixNano())

	var otp string = ""

	for i := 0; i < 6; i++ {

		otp += strconv.Itoa(rand.Intn(10))
	}

	Otp = otp
}

func handleSendEmail(email string) error {

	m := gomail.NewMessage()

	generateOtp()
	m.SetHeader("From", "venkatakamesh.sarvan@thoughtclan.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Otp Verification")
	m.SetBody("text/plain", fmt.Sprintf("This is otp for verification : %s ", Otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, "venkatakamesh.sarvan@thoughtclan.com", "kamesh3939A$123")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func HandleVerifyOtp(c *gin.Context) {

	otp := c.Query("otp")

	fmt.Println(otp, Otp)

	val := validateOtp(otp)

	if val {
		c.JSON(200, gin.H{
			"message": "User Verified Successfully",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Please Enter Valid Otp",
	})
}

func validateOtp(otp string) bool {

	return otp == Otp
}

func HandleLoginUser(c *gin.Context) {

	var user model.User

	c.BindJSON(&user)

	val, err := handleLoginUser(user)

	if err != nil || !val {
		c.JSON(404, gin.H{
			"error": model.ApiError{Message: "Please Provide Valid Username and Password", StatusCode: http.StatusBadRequest},
		})
		return
	}

	c.JSON(200, gin.H{"message": "User got Logged In....."})
}

func handleLoginUser(user model.User) (bool, error) {

	var existingUser model.User

	err := DB.Where("first_name = ?", user.FirstName).Find(&existingUser).Error

	if err != nil {
		return false, err
	}

	if user.Password == existingUser.Password {
		return true, nil
	}

	return false, nil

}

// HandleGetUser handles getting a user by ID.
func HandleGetUser(c *gin.Context) {

	id := c.Param("id")

	user, err := handleGetUser(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": model.ApiError{Message: "x", StatusCode: 404},
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

// handleGetUser gets a user by ID.
func handleGetUser(id string) (*model.User, error) {

	var user model.User
	err := DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// HandleUpdateUser handles updating a user.
func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	updatedUser, err := handleUpdateUser(user, id)

	if err != nil {
		json.NewEncoder(w).Encode(model.ApiError{Message: "please provide valid input", StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

// handleUpdateUser handles updating a user.
func handleUpdateUser(user model.User, id string) (*model.User, error) {
	var existingUser model.User
	err := DB.First(&existingUser, id).Error

	if err != nil {
		return nil, err
	}

	existingUser.FirstName = user.FirstName
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.LastName = user.LastName
	existingUser.DateOfBirth = user.DateOfBirth

	err = DB.Save(&existingUser).Error

	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func HandlePartialUpdateUser(c *gin.Context) {
	//w.Header().Set("Content-Type", "application/json")

	var user model.User

	if err := c.BindJSON(&user); err != nil {

		c.JSON(200, gin.H{
			"error": err,
		})

		return
	}

	params := mux.Vars(r)
	id := params["id"]

	updatedItem, err := handlePartialUpdateUser(user, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ApiError{Message: "please provide valid input", StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(updatedItem)
}

func handlePartialUpdateUser(user model.User, id string) (*model.User, error) {

	var existingUser model.User

	err := DB.First(&existingUser, id).Error

	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}

	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	err = DB.Save(&existingUser).Error

	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

// HandleDeleteItem handles deleting a Item.

// HandleDeleteUser handles deleting a user.
func HandleDeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := handleDeleteUser(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please enter valid id",
		})
	}

	c.JSON(200, gin.H{
		"message": "user got deleted",
	})
}

// handleDeleteUser handles deleting a user.
func handleDeleteUser(id string) error {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return fmt.Errorf("DB not initialized")
	}

	err := DB.Unscoped().Delete(&model.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// HandleGetAllUsers handles getting all users.
func HandleGetAllUsers(c *gin.Context) {

	users, err := handleGetAllUsers()

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

// handleGetAllUsers gets all users.
func handleGetAllUsers() ([]model.User, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var users []model.User
	err := DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
