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
	"nirun/pkg/flex"
	"nirun/pkg/models".

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
			log.Println("Invalid signature error:", err)
			c.Writer.WriteHeader(http.StatusBadRequest)
		} else {
			log.Println("Error parsing request:", err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		log.Println("Event received:", event)
		if event.Type == linebot.EventTypeMessage {
			if message, ok := event.Message.(*linebot.TextMessage); ok {
				name_ := strings.TrimSpace(message.Text)
				log.Println("Patient name received:", name_)

				// เชื่อมต่อฐานข้อมูล
				db, err := database.ConnectToDB()
				if err != nil {
					log.Println("Database connection error:", err)
					replyMessage := "เกิดข้อผิดพลาดในการเชื่อมต่อฐานข้อมูล กรุณาลองใหม่อีกครั้ง"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					return
				}
				defer db.Close()

				// ดึงข้อมูลผู้ป่วยตามชื่อ
				patientInfo, err := models.GetPatientInfoByName(db, name_)
				if err != nil {
					log.Println("Error fetching patient info:", err)
					replyMessage := "ไม่พบข้อมูลของผู้ป่วยชื่อ: " + name_
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					return
				}

				// สร้าง Flex Message และส่งกลับ
				flexMessage := linebot.NewFlexMessage("ข้อมูลผู้ป่วย", flex.CreatePatientFlexMessage(patientInfo))
				if _, err := bot.ReplyMessage(event.ReplyToken, flexMessage).Do(); err != nil {
					log.Println("Error sending reply message:", err)
				} else {
					log.Println("Reply message sent successfully")
				}

			}
		}
	}

	// ส่งสถานะ 200 OK หลังการประมวลผลสำเร็จ
	c.Writer.WriteHeader(http.StatusOK)
	log.Println("Webhook response sent with status 200")
}

// handleEvent เป็น placeholder สำหรับการจัดการ event อื่น ๆ
func HandleEvent(event *linebot.Event) {
	log.Println("Event handler placeholder called for:", event)
	// TODO: เพิ่มการจัดการสำหรับ Event อื่น ๆ ที่ไม่ได้เกี่ยวข้องกับข้อความ
}
