package database

import (
	"log"

	"backend/internal/models"
)

func SeedDatabase() {
	// Cek apakah tabel menu_items sudah ada isinya
	// Ini biar kalau aplikasi di-restart, data nggak numpuk/duplikat
	var count int64
	DB.Model(&models.MenuItem{}).Count(&count)

	if count > 0 {
		log.Println("ℹ️ Database already seeded. Skipping seeder...")
		return
	}

	log.Println("🌱 Seeding database with initial coffee menus...")

	// Data dummy dengan konsep "Menu Jujur"
	menus := []models.MenuItem{
		{
			Name:        "Es Kopi Susu Gula Aren",
			Description: "Susu segar lokal, gula aren asli Bogor yang kita masak sendiri. Bukan sirup pabrikan!",
			Price:       18000,
			IsAvailable: true,
		},
		{
			Name:        "Hot Americano Gayo",
			Description: "Single origin Gayo, diseduh manual. Bold, earthy, dan bikin melek tanpa bikin deg-degan.",
			Price:       15000,
			IsAvailable: true,
		},
		{
			Name:        "Kopi Susu Gula Merah (Less Sugar)",
			Description: "Versi lebih sehat dari ES Kopi Susu kita. Gula merahnya dikurangi 30%, rasa kopinya lebih menonjol.",
			Price:       18000,
			IsAvailable: true,
		},
		{
			Name:        "Cappuccino",
			Description: "Espresso shot ganda dengan microfoam susu yang creamy. Classic comfort di pagi hari.",
			Price:       20000,
			IsAvailable: true,
		},
		{
			Name:        "Air Mineral (Buat Netralin Lidah)",
			Description: "Kadang kita butuh jeda. Air mineral dingin buat bilas mulut sebelum seduhan berikutnya.",
			Price:       5000,
			IsAvailable: true,
		},
	}

	// Loop dan masukkan ke database
	for _, menu := range menus {
		if err := DB.Create(&menu).Error; err != nil {
			log.Printf("❌ Failed to seed menu '%s': %v\n", menu.Name, err)
		}
	}

	log.Println("✅ Database seeded successfully! Siap jualan, bro!")
}
