package lead

import (
	"strconv"

	"github.com/Jdsatashi/GoFiber01/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	//search := c.Query("search", "")
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	db := database.DBConn
	var leads []Lead
	results := db.Limit(intLimit).Offset(offset).Find(&leads)
	return c.JSON(fiber.Map{
		"message": "Get Leads successfully.",
		"Code":    200,
		"data":    results,
	})
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func UpdateLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	if lead.Name == "" {
		return c.Status(404).SendString("No lead found with ID")
	}
	db.Save(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(404).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead deleted")
}
