package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRoles adalah satpam lapis kedua khusus untuk mengecek Hak Akses (RBAC)
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"status": "GAGAL", "pesan": "Role tidak ditemukan!"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR", "pesan": "Kesalahan membaca role"})
			c.Abort()
			return
		}

		isAllowed := false
		for _, role := range allowedRoles {
			if roleStr == role {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "GAGAL",
				"pesan":  "Akses ditolak. Tingkat otorisasi Anda (" + roleStr + ") tidak mencukupi.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
