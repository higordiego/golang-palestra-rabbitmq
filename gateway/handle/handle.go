package handle

import (
	"encoding/json"
	"net/http"

	rabbit "github.com/higordiego/rabbitmq-palestra/gateway/integration"
	"github.com/higordiego/rabbitmq-palestra/gateway/models"
	"github.com/labstack/echo"
)

// WhatsappPostBack - handler facebook postback receive
func WhatsappPostBack(c echo.Context) error {
	var whatsapp models.WhatsappPostBack
	var receive models.WhatsappMessage
	if err := c.Bind(&whatsapp); err != nil {
		return err
	}

	receive.ReceiveParsePostBack(&whatsapp)

	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&receive)

	conn.PublishRabbitMQ("messages-whatsapp", []byte(jsonString), "whatsapp")

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
