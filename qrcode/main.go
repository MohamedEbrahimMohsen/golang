package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	if err := connectDB(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/qr", generate)
	r.Run("localhost:8000")
}

type qrModel struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	URL    string `json:"url"`
	Title  string `json:"title"`
	QRCode []byte `json:"qrcode"`
}

func generate(c *gin.Context) {
	var model qrModel

	if err := c.BindJSON(&model); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid payload."})
		return
	}

	qr, err := generateQR(model.URL)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	model.QRCode = qr
	addQRToDb(&model)
	c.IndentedJSON(http.StatusCreated, model)
}

func generateQR(url string) ([]byte, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}

	return qr, nil
}

func addQRToDb(model *qrModel) {
	db.Create(model)
}

func connectDB() error {
	conn := "host=database-qrcode.c1j6msw5gk6k.us-east-1.rds.amazonaws.com user=postgres password=postgres dbname=database-qrcode port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&qrModel{})
	return nil
}

/*
References:
https://pkg.go.dev/github.com/skip2/go-qrcode#section-readme
*/
