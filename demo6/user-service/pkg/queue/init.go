package queue

import (
	"fmt"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/config"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection

const USER_EVNETS_QUEUE_NAME = "userEvents"

func init() {
	var err error
	mqConfig := config.GetAppConfig().Application.Rabbitmq
	mqURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	conn, err = amqp.Dial(mqURL)
	if err != nil {
		panic(err)
	}
}
