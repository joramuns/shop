package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (p *Producer) SendMessage(ctx context.Context, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	if err := p.writer.WriteMessages(ctx, kafka.Message{Value: data}); err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	log.Println("Message sent:", data)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
