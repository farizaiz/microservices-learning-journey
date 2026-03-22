package controllers

import (
	"net/http"
	"os"
	"time"

	"user-service/config"
	"user-service/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ==========================================
// 1. FUNGSI REGISTER
// ==========================================
func RegisterUser(c *gin.Context) {
	var userBaru models.User

	if err := c.ShouldBindJSON(&userBaru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "GAGAL", "pesan": "Format data tidak valid: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userBaru.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "GAGAL", "pesan": "Gagal memproses keamanan password"})
		return
	}

	userBaru.Password = string(hashedPassword)

	if err := config.DB.Create(&userBaru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "GAGAL", "pesan": "Gagal menyimpan data. Email/Nomor Induk mungkin sudah terdaftar."})
		return
	}

	userBaru.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"status": "BERHASIL",
		"pesan":  "Akun berhasil didaftarkan!",
		"data":   userBaru,
	})
}

// ==========================================
// 2. FUNGSI LOGIN
// ==========================================
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "GAGAL", "pesan": "Format input salah"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Email tidak terdaftar!"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Password salah!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "GAGAL", "pesan": "Gagal membuat token keamanan"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Login sukses!",
		"token":  tokenString,
		"data":   user,
	})
}

// ==========================================
// 3. FUNGSI GET PROFIL (RAHASIA)
// ==========================================
func GetProfil(c *gin.Context) {
	// Ambil ID yang tadi dititipkan oleh Middleware ke dalam Context
	userID, _ := c.Get("user_id")

	var user models.User
	// Cari data user tersebut di database
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "GAGAL", "pesan": "Data pengguna tidak ditemukan"})
		return
	}

	// ⚠️ PENTING: Kosongkan password lagi agar tidak bocor!
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Ini adalah area rahasia profil Anda",
		"data":   user,
	})
}
