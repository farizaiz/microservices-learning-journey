package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware adalah satpam pencegat request
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Cek apakah ada surat pengantar (Header Authorization)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Akses ditolak. Token VIP tidak ditemukan!"})
			c.Abort() // Hentikan proses!
			return
		}

		// 2. Cek apakah formatnya benar: "Bearer <token_panjang_anda>"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Format token salah. Gunakan format 'Bearer <token>'"})
			c.Abort()
			return
		}

		// 3. Pisahkan kata "Bearer " untuk mengambil murni tokennya saja
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Validasi keaslian Token menggunakan Sandi Rahasia dari .env
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan algoritma penandatanganannya sesuai
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode enkripsi tidak valid")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Token palsu atau sudah kedaluwarsa!"})
			c.Abort()
			return
		}

		// 5. JIKA TOKEN ASLI: Ambil data ID User di dalamnya dan simpan di memori sementara (Context)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])

			// PINTU DIBUKA! Silakan masuk ke Controller
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "GAGAL", "pesan": "Gagal membaca data dari token"})
			c.Abort()
			return
		}
	}
}
