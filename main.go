package main

import (
	"database/sql"
	"fmt"
	"log"

	connectDB "newgo-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	_ "github.com/rizalgowandy/go-swag-sample/docs/fibersimple"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Person struct {
	Rid     *int `json:"r_id"`
	Ragid   *int `json:"r_agent_id"`
	Rstatus *int `json:"r_status"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// เชื่อมต่อฐานข้อมูล
	db, err := connectDB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Route เพื่อ query ข้อมูลและแสดงผลบน Swagger
	app.Get("/users", getUsersHandler(db))

	// Route Swagger
	app.Get("/sc/docs/*", fiberSwagger.FiberWrapHandler()) // เปลี่ยนให้ใช้ middleware Swagger จาก fiber-swagger/v2

	// เริ่มเซิร์ฟเวอร์
	app.Listen(":3000")
}

// ตัวอย่าง query ข้อมูลและแสดงผลบน path /users
func getUsersHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ตัวอย่าง query ข้อมูลจากตาราง users
		rows, err := db.Query("SELECT r_id , r_agent_id , r_status FROM bom_tb_agent")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		var people []Person

		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Rid, &p.Ragid, &p.Rstatus)
			if err != nil {
				panic(err)
			}
			people = append(people, p)
		}

		fmt.Println(people)

		// แสดงผลเป็น JSON บน Swagger
		return c.JSON(fiber.Map{
			"users": people,
		})
	}
}
