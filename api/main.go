package main

import (
	"fmt"
	"log"
	"nirun/pkg/database"
	"nirun/pkg/hook"

	//"nirun/pkg/models"
	"nirun/pkg/linebot"

	"github.com/gin-gonic/gin"
)

func main() {
	// เรียกใช้ InitLineBot เพื่อสร้าง instance ของ LINE Bot
	linebot.InitLineBot()

	router := gin.Default()
	// ผูกเส้นทาง /webhook กับฟังก์ชัน HandleLineWebhook
	router.POST("/webhook", hook.HandleLineWebhook)
	router.Run(":8080") // รันเซิร์ฟเวอร์ที่พอร์ต 8080 หรือพอร์ตที่คุณต้องการ

	// เรียก ConnectToDB เพื่อเชื่อมต่อฐานข้อมูลโดยตรง
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully!")

	
}
