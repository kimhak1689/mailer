package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()

	// Add basic authentication middleware
	r.Use(gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	}))

	r.POST("/send_email", func(c *gin.Context) {
		// Load email credentials and settings from environment variables
		username := os.Getenv("SMTP_USERNAME")
		password := os.Getenv("SMTP_PASSWORD")
		host := os.Getenv("SMTP_HOST")
		port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
		fromAddress := os.Getenv("SMTP_FROM_ADDRESS")
		fromName := os.Getenv("SMTP_FROM_NAME") // Add this line

		toAddress := c.PostForm("to_address")
		subject := c.PostForm("subject")
		body := c.PostForm("body")

		m := gomail.NewMessage()
		m.SetAddressHeader("From", fromAddress, fromName) // Use SetAddressHeader to set both email and name
		m.SetHeader("To", toAddress)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)

		d := gomail.NewDialer(host, port, username, password)

		if err := d.DialAndSend(m); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	})

	r.Run(":8080")
}
