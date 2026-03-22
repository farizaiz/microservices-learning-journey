// File: routes/dokumenroutes.go
package routes

import (
	"eval-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Fitur canggih Gin: Routing Group
	// Kita kelompokkan semua URL yang berawalan "/api/dokumen"
	ruteDokumen := r.Group("/api/dokumen")
	{
		// Jika metode GET, panggil fungsi TampilSemuaDokumen
		ruteDokumen.GET("/", controllers.TampilSemuaDokumen)

		// Jika metode POST, panggil fungsi TambahDokumen
		ruteDokumen.POST("/", controllers.TambahDokumen)

		// Jika metode DELETE (dengan tambahan parameter :id)
		ruteDokumen.DELETE("/:id", controllers.HapusDokumen)
	}
}
