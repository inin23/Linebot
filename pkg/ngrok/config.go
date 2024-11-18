package Ngrok

import (
    "fmt"
    "log"
    "nirun/pkg/database" // เส้นทางที่ถูกต้องสำหรับการนำเข้า database package
)

// ฟังก์ชันเพื่อโหลดค่า authtoken จาก config.yaml
func GetNgrokAuthToken() string {
    // โหลดการตั้งค่าจาก config.yaml
    config, err := database.LoadConfig()
    if err != nil {
        log.Fatalf("ไม่สามารถโหลดไฟล์ config ได้: %v", err)
    }

    // ดึงค่า authtoken จาก config ที่โหลดได้
    authtoken := config.Agent.Authtoken
    fmt.Println("Ngrok Auth Token:", authtoken)
    return authtoken
}
