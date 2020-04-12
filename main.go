package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"bufio"
	"os"
	"encoding/base64"
	"io/ioutil"
	"mime/multipart"
  	mailjet "github.com/mailjet/mailjet-apiv3-go"
)


// EmailRequest é o objeto recebido na requisição
type EmailRequest struct {
	From string `json:"from"`
	To string `json:"to"`
}

func main() {
	r := gin.Default()
	r.GET("/health", getHealth)
	r.POST("/send-mail", sendEmail)
	r.Run()
}

func getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

func sendEmail(c *gin.Context) {
	fh, _ := c.FormFile("file")
	from := c.PostForm("from")
	to := c.PostForm("to")
	
	encoded := convertToBase64(fh)
	
	_, err := callMailJetAPI(from, to, encoded)
	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(200, gin.H{
		"message": "email enviado",
	})
}

func convertToBase64(fh *multipart.FileHeader) string {
	file, _ := fh.Open()
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	return base64.StdEncoding.EncodeToString(content)
}

func callMailJetAPI(from string, to string, pdf string) (*mailjet.ResultsV31, error) {

	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	if publicKey == "" || secretKey == "" {
		log.Fatal("Precisa definir MJ_APIKEY_PUBLIC e MJ_APIKEY_PRIVATE")
	}

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
		From: &mailjet.RecipientV31{
			Email: from,
		},
		To: &mailjet.RecipientsV31{
			mailjet.RecipientV31 {
			Email: to,
			},
		},
		Subject: "Certificado de Conclusão de Curso",
		HTMLPart: "<h3>Parabéns! Segue em anexo o seu certificado de conclusão de curso</h3>",
		Attachments: &mailjet.AttachmentsV31{
			mailjet.AttachmentV31{
				ContentType: "application/pdf",
				Base64Content: pdf,
				Filename: "certificado.pdf",
			},
		},
		CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	return mj.SendMailV31(&messages)

}
