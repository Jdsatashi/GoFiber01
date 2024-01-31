package main

import (
	"os"

	"github.com/Jdsatashi/GoFiber01/database"
	"github.com/Jdsatashi/GoFiber01/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Post("/api/v1/lead/:id", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		// panic("Failed connect to database.")
		log.Info("Failed to connect to database:", err)
		os.Exit(1)
	}
	log.Info("Connected to database.")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Info("Migrating database.")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Listen(":3000")
}
