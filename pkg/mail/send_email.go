package mail

import (
	"context"
	"github.com/trycourier/courier-go/v2"
	"log"
)

func SendProjectToNewsLetter() {

}

func SendEmail(to string, template string) {
	templateType := "MTPRTRG6NKMH11G45S2AFP3CJAEY"
	switch template {
	case "JOIN":
		templateType = "MTPRTRG6NKMH11G45S2AFP3CJAEY"

	}
	client := courier.CreateClient("pk_prod_VZBZM4377R44KPPY70Y9AB39SA5M", nil)

	requestID, err := client.SendMessage(
		context.Background(),
		courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"to":       map[string]string{"email": to},
				"template": templateType,
				"data": map[string]string{
					"variables": "awesomeness",
				},
			},
		},
	)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
