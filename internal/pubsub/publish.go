package pubsub

import (
	"encoding/json"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	encodedJson, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("error marshalling value for publishing")
	}

	ch.PublishWithContext(
		context.Backgroun(),
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: encodedJson,
		}
	)

	return
}
