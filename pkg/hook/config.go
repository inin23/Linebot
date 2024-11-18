// pkg/linebot/hook/config.go
package hook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"nirun/pkg/database"
	"nirun/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"

	// ปรับ import path ตามโครงสร้างโปรเจคของคุณ
	linebotConfig "nirun/pkg/linebot"
)

type Config struct {
	LineBot struct {
		Webhook_url string `yaml:"webhook_url"`
	} `yaml:"line_bot"`
}

// LoadConfig ฟังก์ชันอ่านไฟล์ YAML
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// VerifySignature ทำการตรวจสอบลายเซ็น
func VerifySignature(channelSecret string, body []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(channelSecret))
	mac.Write(body)
	expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

// HandleLineWebhook จัดการ webhook requests จาก LINE
func HandleLineWebhook(c *gin.Context) {
	bot := linebotConfig.GetLineBot()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(http.StatusBadRequest)
		} else {
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		log.Println("Event received:", event)
		if event.Type == linebot.EventTypeMessage {
			message, ok := event.Message.(*linebot.TextMessage)
			if ok {
				// ใช้ข้อความที่ผู้ใช้ส่งมาเป็นชื่อผู้ป่วย
				name_ := strings.TrimSpace(message.Text)
				log.Println("Patient name received:", name_)

				// เชื่อมต่อฐานข้อมูล
				db, err := database.ConnectToDB()
				if err != nil {
					log.Println("Database connection error:", err)
					return
				}
				defer db.Close()

				// ดึงข้อมูลผู้ป่วยตามชื่อ
				patientInfo, err := models.GetPatientInfoByName(db, name_)
				if err != nil {
					log.Println("Error fetching patient info:", err)
					models.ReplyErrorFormat(bot, event.ReplyToken) // ส่งข้อความแจ้งว่าหาไม่พบ
					return
				}

				// ส่งข้อมูลผู้ป่วยกลับไปยังผู้ใช้
				replyMessage := models.FormatPatientInfo(patientInfo)
_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
if err != nil {
	log.Println("Error replying message:", err) // ตรวจสอบว่าเกิด error ในการส่งข้อความกลับหรือไม่
} else {
	log.Println("Reply message sent successfully") // Log ว่าส่งข้อความสำเร็จ
}

			}
		}
	}

	// ส่งสถานะ 200 OK กลับไปหลังจากประมวลผล webhook เสร็จสิ้น
	c.Writer.WriteHeader(http.StatusOK)
	log.Println("Webhook response sent with status 200")
}

// handleEvent จัดการกับแต่ละ event ที่ได้รับจาก LINE
func HandleEvent(event *linebot.Event) {
	// TODO: เพิ่มโค้ดจัดการ event ตามที่ต้องการ
}
