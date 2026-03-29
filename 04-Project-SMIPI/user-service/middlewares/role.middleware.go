package middlewares

import (
	"net/http"
	"strings" // 1. Tambahkan import strings

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			// ... (kode error tetap sama)
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			// ... (kode error tetap sama)
			return
		}

		isAllowed := false
		for _, role := range allowedRoles {
			// 2. GUNAKAN EqualFold untuk perbandingan Case-Insensitive
			if strings.EqualFold(roleStr, role) {
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
