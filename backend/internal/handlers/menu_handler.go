package handlers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GET /api/menu -> Buat Customer (Cuma yang available)
func GetAllMenus(c *gin.Context) {
	var menus []models.MenuItem
	database.DB.Where("is_available = ?", true).Find(&menus)
	c.JSON(http.StatusOK, gin.H{"data": menus})
}

// GET /api/admin/menu -> Buat Owner (Lihat semua, termasuk yang di-sold out)
func GetAllMenusAdmin(c *gin.Context) {
	var menus []models.MenuItem
	database.DB.Find(&menus)
	c.JSON(http.StatusOK, gin.H{"data": menus})
}

// POST /api/admin/menu -> Tambah menu baru
func CreateMenu(c *gin.Context) {
	var input models.MenuItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// PUT /api/admin/menu/:id -> Update menu (misal: ganti harga, atau sold out)
func UpdateMenu(c *gin.Context) {
	var menu models.MenuItem
	id := c.Param("id")

	if err := database.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	var input models.MenuItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&menu).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": menu})
}

// DELETE /api/admin/menu/:id -> Hapus menu
func DeleteMenu(c *gin.Context) {
	var menu models.MenuItem
	id := c.Param("id")

	if err := database.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	database.DB.Delete(&menu)
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
