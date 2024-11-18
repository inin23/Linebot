package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"nirun/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// ConnectToDB เชื่อมต่อกับฐานข้อมูลโดยใช้ข้อมูลจาก config.yaml
func ConnectToDB() (*sql.DB, error) {
	// โหลด config จากไฟล์ config.yaml
	var config models.Config
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
		return nil, err
	}

	// รูปแบบการเชื่อมต่อฐานข้อมูล (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name,
	)

	// เปิดการเชื่อมต่อฐานข้อมูล
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ฟังก์ชัน LoadConfig เพื่อโหลดค่าการตั้งค่าจาก config.yaml
func LoadConfig() (models.Config, error) {
	var config models.Config

	// อ่านไฟล์ config.yaml
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("Error reading config file:", err)
		return config, err
	}

	// แปลงข้อมูลจาก YAML เป็น struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Println("Error parsing config file:", err)
		return config, err
	}

	return config, nil
}
