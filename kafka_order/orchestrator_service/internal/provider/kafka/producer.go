package provider

import (
	"orchestrator_service/internal/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	*kafka.Writer
	*kafka.Conn
}


func NewKafkaProducer() (*KafkaProducer, error) {
	return &KafkaProducer{
		
	}
}