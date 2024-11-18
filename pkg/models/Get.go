package models

import (
	"database/sql"
	"fmt"
	"log"
	//"nirun/pkg/flex"

	"github.com/line/line-bot-sdk-go/linebot"
)

// UserSession เก็บสถานะของการโต้ตอบกับผู้ใช้
var UserSession = make(map[string]string)

// ฟังก์ชันสำหรับดึงข้อมูลผู้ป่วยตาม Name
func GetPatientInfoByName(db *sql.DB, name_ string) (*PatientInfo, error) {
	query := `SELECT name_, patiet_id, age, sex, blood, phone_numbers FROM patient_info WHERE name_ = ?`
	row := db.QueryRow(query, name_)

	var patientInfo PatientInfo
	err := row.Scan(&patientInfo.Name, &patientInfo.PatientID, &patientInfo.Age, &patientInfo.Sex, &patientInfo.Blood, &patientInfo.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ไม่พบข้อมูลผู้ป่วยที่มีชื่อ %s", name_)
		}
		return nil, err
	}
	log.Println("ข้อมูลผู้ป่วยที่ดึงมา:", &patientInfo)
	return &patientInfo, nil
}

// handlePatientInfo ฟังก์ชันจัดการการดึงและส่งข้อมูลผู้ป่วยผ่าน Line Bot
// public
/*func HandlePatientInfo(bot *linebot.Client, replyToken string, db *sql.DB, name_ string) {
	data, err := GetPatientInfoByName(db, name_) // ใช้ฟังก์ชันจาก patient_info
	if err != nil {
		log.Println("ข้อผิดพลาดในการดึงข้อมูล: ", err)
		return
	}

	replyMessage := FormatPatientInfo(data)
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Println("ข้อผิดพลาดในการตอบกลับ:", err)
	}
}*/

// handlePatientInfo ฟังก์ชันจัดการการดึงและส่งข้อมูลผู้ป่วยผ่าน Line Bot
func HandlePatientInfo(client *linebot.Client, replyToken string, db *sql.DB, name_ string) {
	data, err := GetPatientInfoByName(db, name_) // ดึงข้อมูลผู้ป่วย
	if err != nil {
		log.Println("Error fetching patient info:", err)
		return
	}

	// เรียกใช้ฟังก์ชันใน flex_patientInfo.go
	flexMessage := flex.CreatePatientFlexMessage(data)

	// ส่ง flex message ไปยัง Line Bot
	if _, err := client.ReplyMessage(replyToken, linebot.NewImageMessage("ข้อมูลผู้ป่วย",flexMessage)).Do(); err != nil {
		log.Println("Error sending flex message:", err)
	}
}

// formatPatientInfo ฟังก์ชันเพื่อจัดรูปแบบข้อมูลผู้ป่วยให้เหมาะสมสำหรับการแสดงผล
/*func FormatPatientInfo(patient *PatientInfo) string {
	return fmt.Sprintf("ข้อมูลผู้ป่วย:\nชื่อ: %s\nรหัสผู้ป่วย: %s\nอายุ: %d\nเพศ: %s\nหมู่เลือด: %s\nหมายเลขโทรศัพท์: %s",
		patient.Name, patient.PatientID, patient.Age, patient.Sex, patient.Blood, patient.PhoneNumber)
}*/

// replyErrorFormat ส่งข้อความแจ้งรูปแบบการใช้งานที่ถูกต้อง
func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
	bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage("กรุณากรอกรูปแบบ 'ข้อมูลผู้สูงอายุ []'"),
		linebot.NewTextMessage("กรุณากรอกรูปแบบ 'ข้อมูลผู้กิจกรรม []'"),
	)
}

// // GetPatientInfoByActivity ค้นหาข้อมูลกิจกรรมตามชื่อกิจกรรมโดยใช้ LIKE
// func GetPatientInfoByActivity(db *sql.DB, activity string) (*ServiceInfo, error) {
// 	query := `SELECT
// 		p.name_, p.patient_id, p.age, p.sex, p.blood, p.phone_numbers,
// 		s.activity, s.service_code, s.into_number, s.location, s.sine, s.end_, s.period
// 	FROM
// 		patient_info p
// 	JOIN
// 		service_info s ON p.patient_id = s.patient_id
// 	WHERE
// 		s.activity LIKE ?`

// 	row := db.QueryRow(query, "%"+activity+"%") // ใช้ LIKE เพื่อค้นหาคำที่คล้ายกัน

// 	var patientInfo PatientInfo
// 	var serviceInfo ServiceInfo

// 	// Scan ข้อมูลจาก query result ไปยัง patientInfo และ serviceInfo
// 	err := row.Scan(
// 		&patientInfo.Name, &patientInfo.PatientID, &patientInfo.Age, &patientInfo.Sex, &patientInfo.Blood, &patientInfo.PhoneNumber,
// 		&serviceInfo.Activity, &serviceInfo.ServiceCode, &serviceInfo.IntoNumber, &serviceInfo.Location, &serviceInfo.Sine, &serviceInfo.End, &serviceInfo.Period,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("ไม่พบข้อมูลกิจกรรมที่มีชื่อ %s", activity)
// 		}

// 	}
// 	log.Println("ข้อมูลผู้ป่วยที่ดึงมา:", &serviceInfo)
// 	return &serviceInfo, nil
// }

// // handleServiceInfo ฟังก์ชันจัดการการดึงและส่งข้อมูลผู้ป่วยผ่าน Line Bot
// // public
// func HandleServiceInfo(bot *linebot.Client, replyToken string, db *sql.DB, activity string) {
// 	data, err := GetPatientInfoByActivity(db, activity) // ใช้ฟังก์ชัน
// 	if err != nil {
// 		log.Println("ข้อผิดพลาดในการดึงข้อมูล: ", err)
// 		return
// 	}

// 	replyMessage := FormatServiceInfo(data)
// 	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
// 		log.Println("ข้อผิดพลาดในการตอบกลับ:", err)
// 	}
// }

// // formatPatientInfo ฟังก์ชันเพื่อจัดรูปแบบข้อมูลผู้ป่วยให้เหมาะสมสำหรับการแสดงผล
// func FormatServiceInfo(service *ServiceInfo) string {
// 	return fmt.Sprintf("ข้อมูลกิจกรรม:\nกิจกรรม: %s\nรหัสบริการ: %s\nจำนวนที่เข้ารับบริการ: %d\nสถานที่: %s\nตั้งแต่: %s\nจนถึง: %s\nระยะเวลา: %s",
// 		service.Activity, service.ServiceCode, service.IntoNumber, service.Location, service.Sine, service.End, service.Period)
// }

// // replyErrorFormat ส่งข้อความแจ้งรูปแบบการใช้งานที่ถูกต้อง
// func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
// 	bot.ReplyMessage(
// 		replyToken,
// 		linebot.NewTextMessage("กรุณากรอกรูปแบบ 'ข้อมูลกิจกรรม []'"),
// 	)
// }

// 	// รวมข้อมูล patientInfo เข้ากับ serviceInfo
// 	serviceInfo.PatientInfo = patientInfo

// 	log.Println("ข้อมูลผู้ป่วยและกิจกรรมที่ดึงมา:", &serviceInfo)
// 	return &serviceInfo, nil
// }

// //handlePatientInfo ฟังก์ชันจัดการการดึงและส่งข้อมูลผู้ป่วยผ่าน Line Bot
// //public
// func HandlePatientInfo(bot *linebot.Client, replyToken string, db *sql.DB, name_ string) {
// 	data, err := GetPatientInfoByName(db, name_) // ใช้ฟังก์ชันจาก patient_info
// 	if err != nil {
// 		log.Println("ข้อผิดพลาดในการดึงข้อมูล: ", err)
// 		bot.ReplyMessage(replyToken, linebot.NewTextMessage(fmt.Sprintf("ไม่พบข้อมูลผู้ป่วยที่มีชื่อ %s", name_))).Do()
// 		return
// 	}

// 	replyMessage := FormatPatientInfo(data)
// 	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
// 		log.Println("ข้อผิดพลาดในการตอบกลับ:", err)
// 	}
// }

// // formatPatientInfo ฟังก์ชันเพื่อจัดรูปแบบข้อมูลผู้ป่วยให้เหมาะสมสำหรับการแสดงผล
// func FormatPatientInfo(patient *PatientInfo) string {
// 	return fmt.Sprintf("ข้อมูลผู้ป่วย:\nชื่อ: %s\nรหัสผู้ป่วย: %s\nอายุ: %d\nเพศ: %s\nหมู่เลือด: %s\nหมายเลขโทรศัพท์: %s",
// 		patient.Name, patient.PatientID, patient.Age, patient.Sex, patient.Blood, patient.PhoneNumber)
// }

// // FormatServiceInfo ฟังก์ชันเพื่อจัดรูปแบบข้อมูลกิจกรรมให้เหมาะสมสำหรับการแสดงผล
// func FormatServiceInfo(service *ServiceInfo) string {
// 	return fmt.Sprintf("ข้อมูลกิจกรรม:\nกิจกรรม: %s\nรหัสบริการ: %s\nจำนวนที่เข้ารับบริการ: %d\nสถานที่: %s\nตั้งแต่: %s\nจนถึง: %s\nระยะเวลา: %s",
// 		service.Activity, service.ServiceCode, service.IntoNumber, service.Location, service.Sine, service.End, service.Period)
// }

// // replyErrorFormat ส่งข้อความแจ้งรูปแบบการใช้งานที่ถูกต้อง
// func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
// 	bot.ReplyMessage(
// 		replyToken,
// 		linebot.NewTextMessage("กรุณากรอกรูปแบบ 'ข้อมูลผู้สูงอายุ []'"),
// 	)
// }
