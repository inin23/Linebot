package models

import (
	// "database/sql"
	// "fmt"
	// "log"

	// //"nirun/pkg/models"

	// "github.com/line/line-bot-sdk-go/linebot"
)

type PatientInfo struct {
	ID          int    `json:"id"`
	PatientID   string `json:"patiet_id"`
	Image       []byte `json:"image"`
	Name        string `json:"name_"`
	PhoneNumber string `json:"phone_numbers"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CardID      string `json:"card_id"`
	Religion    string `json:"religion"`
	Sex         string `json:"sex"`
	Blood       string `json:"blood"`
	DateOfBirth string `json:"date_of_birth"`
	Age         int    `json:"age"`
}

// // ฟังก์ชันสำหรับดึงข้อมูลผู้ป่วยตาม Name
// func GetPatientInfoByName(db *sql.DB, name_ string) (*PatientInfo, error) {
// 	query := `SELECT name_, patiet_id, age, sex, blood, phone_numbers FROM patient_info WHERE name_ = ?`
// 	row := db.QueryRow(query, name_)

// 	var patientInfo PatientInfo
// 	err := row.Scan(&patientInfo.Name, &patientInfo.PatientID, &patientInfo.Age, &patientInfo.Sex, &patientInfo.Blood, &patientInfo.PhoneNumber)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("ไม่พบข้อมูลผู้ป่วยที่มีชื่อ %s", name_)
// 		}
// 		return nil, err
// 	}
// 	log.Println("ข้อมูลผู้ป่วยที่ดึงมา:", &patientInfo)
// 	return &patientInfo, nil
// }

// // handlePatientInfo ฟังก์ชันจัดการการดึงและส่งข้อมูลผู้ป่วยผ่าน Line Bot
// // public
// func HandlePatientInfo(bot *linebot.Client, replyToken string, db *sql.DB, name_ string) {
// 	data, err := GetPatientInfoByName(db, name_) // ใช้ฟังก์ชันจาก patient_info
// 	if err != nil {
// 		log.Println("ข้อผิดพลาดในการดึงข้อมูล: ", err)
// 		return
// 	}

// 	replyMessage := FormatPatientInfo(data)
// 	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
// 		log.Println("ข้อผิดพลาดในการตอบกลับ:", err)
// 	}
// }

// // formatPatientInfo ฟังก์ชันเพื่อจัดรูปแบบข้อมูลผู้ป่วยให้เหมาะสมสำหรับการแสดงผล
// func FormatPatientInfo(patient *PatientInfo) string {
// 	return fmt.Sprintf("ข้อมูลผู้ป่วย:\nชื่อ: %s\nรหัสผู้ป่วย: %s\nอายุ: %d\nเพศ: %s\nหมู่เลือดเลือด: %s\nหมายเลขโทรศัพท์: %s",
// 		patient.Name, patient.PatientID, patient.Age, patient.Sex, patient.Blood, patient.PhoneNumber)
// }

// // replyErrorFormat ส่งข้อความแจ้งรูปแบบการใช้งานที่ถูกต้อง
// func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
// 	bot.ReplyMessage(
// 		replyToken,
// 		linebot.NewTextMessage("กรุณากรอกรูปแบบ 'ข้อมูลผู้สูงอายุ []'"),
// 	)
// }
