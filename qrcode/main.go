package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	r := gin.Default()
	r.POST("/qr", generate)
	r.Run("localhost:8000")
}

type qrModel struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func generate(c *gin.Context) {
	var qr qrModel

	if err := c.BindJSON(&qr); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid payload."})
		return
	}

	qrImage, err := generateQR(qr.URL)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"image": qrImage})
}

func generateQR(url string) ([]byte, error) {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}

	return png, nil
}

/*
References:
https://pkg.go.dev/github.com/skip2/go-qrcode#section-readme
*/
