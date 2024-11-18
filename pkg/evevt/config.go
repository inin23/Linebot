package event

import (
	"fmt"
	"log"
	"nirun/pkg/database"              // แก้ไขตาม path ที่ถูกต้องของ database package
	linebotConfig "nirun/pkg/linebot" // แก้ไขตาม path ที่ถูกต้องของ linebotConfig package
	"nirun/pkg/models"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleEvent(event *linebot.Event) {
	bot := linebotConfig.GetLineBot()

	if event.Type != linebot.EventTypeMessage {
		return
	}

	message, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		return
	}

	text := message.Text
	log.Println("Received message:", text)

	db, err := database.ConnectToDB()
	if err != nil {
		log.Println("Database connection error:", err)
		ReplyDatabaseError(bot, event.ReplyToken)
		return
	}
	defer db.Close()

	if strings.HasPrefix(text, "ข้อมูลผู้ป่วย") {
		// ค้นหาผู้ป่วย
		name := strings.TrimSpace(strings.TrimPrefix(text, "ข้อมูลผู้ป่วย "))
		log.Println("Searching for patient with name:", name)

		patientInfo, err := models.GetPatientInfoByName(db, name)
		if err != nil {
			log.Println("Error fetching patient info:", err)
			ReplyDataNotFound(bot, event.ReplyToken)
			return
		}

		// ส่งข้อมูลผู้ป่วยกลับไปยังผู้ใช้
		replyMessage := fmt.Sprintf(
			"ข้อมูลผู้ป่วย:\nชื่อ: %s\nรหัสผู้ป่วย: %s\nอายุ: %d\nเพศ: %s\nหมู่เลือด: %s\nเบอร์โทรศัพท์: %s",
			patientInfo.Name, patientInfo.PatientID, patientInfo.Age, patientInfo.Sex, patientInfo.Blood, patientInfo.PhoneNumber,
		)
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			log.Println("Error replying patient info message:", err)
		}
	}
}

// 	} else if strings.HasPrefix(text, "ข้อมูลกิจกรรม") {
// 		// ค้นหากิจกรรม
// 		activity := strings.TrimSpace(strings.TrimPrefix(text, "ข้อมูลกิจกรรม"))
// 		log.Println("Searching for patient with activity:", activity)

// 		serviceInfo, err := models.GetPatientInfoByActivity(db, activity)
// 		if err != nil {
// 			log.Println("Error fetching patient info:", err)
// 			ReplyDataNotFound(bot, event.ReplyToken)
// 			return
// 		}

// 		// ส่งข้อมูลผู้ป่วยกลับไปยังผู้ใช้
// 		replyMessage := fmt.Sprintf(
// 			"ข้อมูลกิจกรรม:\nกิจกรรม: %s\nรหัสบริการ: %s\nจำนวนที่เข้ารับบริการ: %d\nสถานที่: %s\nตั้งแต่: %s\nจนถึง: %s\nระยะเวลา: %s",
// 			serviceInfo.Activity, serviceInfo.ServiceCode, serviceInfo.IntoNumber, serviceInfo.Location, serviceInfo.Sine, serviceInfo.End, serviceInfo.Period)
// 		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
// 			log.Println("Error replying patient info message:", err)
// 		}
// 	}
// }

// ฟังก์ชัน replyErrorFormat ใช้ส่งข้อความแสดงข้อผิดพลาดกลับไปยังผู้ใช้
func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
	errorMessage := "กรุณากรอกรูปแบบข้อความให้ถูกต้อง เช่น 'ข้อมูลผู้ป่วย นางสมหวัง สดใส'"
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(errorMessage)).Do(); err != nil {
		log.Println("Error sending error message:", err)
	}
}

// ฟังก์ชัน replyDataNotFound ใช้ส่งข้อความเมื่อไม่พบข้อมูลผู้ป่วย
func ReplyDataNotFound(bot *linebot.Client, replyToken string) {
	notFoundMessage := "ไม่พบข้อมูลผู้สูงอายุตามชื่อ กรุณาลองใหม่อีกครั้ง"
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(notFoundMessage)).Do(); err != nil {
		log.Println("Error sending not found message:", err)
	}
}

// ฟังก์ชัน replyDatabaseError ใช้ส่งข้อความเมื่อเกิดข้อผิดพลาดในการเชื่อมต่อฐานข้อมูล
func ReplyDatabaseError(bot *linebot.Client, replyToken string) {
	dbErrorMessage := "เกิดข้อผิดพลาดในการเชื่อมต่อฐานข้อมูล กรุณาลองใหม่อีกครั้งภายหลัง"
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(dbErrorMessage)).Do(); err != nil {
		log.Println("Error sending database error message:", err)
	}
}
