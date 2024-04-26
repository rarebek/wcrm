package kafka

// import (
// 	"evrone_service/api_gateway/api/models"
// 	configpkg "evrone_service/api_gateway/internal/pkg/config"
// 	"context"
// 	"encoding/json"

// 	"github.com/segmentio/kafka-go"
// 	"go.uber.org/zap"
// )

// type producer struct {
// 	logger     *zap.Logger
// 	userCreate *kafka.Writer
// }

// func NewProducer(config *configpkg.Config, logger *zap.Logger) *producer {
// 	return &producer{
// 		logger: logger,
// 		userCreate: &kafka.Writer{
// 			Addr:                   kafka.TCP(config.Kafka.Address...),
// 			Topic:                  config.Kafka.Topic.UserCreateTopic,
// 			Balancer:               &kafka.Hash{},
// 			RequiredAcks:           kafka.RequireAll,
// 			AllowAutoTopicCreation: true,
// 			Async:                  true,
// 			Completion: func(messages []kafka.Message, err error) {
// 				if err != nil {
// 					logger.Error("kafka userCreated", zap.Error(err))
// 				}
// 				for _, message := range messages {
// 					logger.Sugar().Info(
// 						"kafka investmentCreated message",
// 						zap.Int("partition", message.Partition),
// 						zap.Int64("offset", message.Offset),
// 						zap.String("key", string(message.Key)),
// 						zap.String("value", string(message.Value)),
// 					)
// 				}
// 			},
// 		},
// 	}
// }

// func (p *producer) ProduceUserToCreate(ctx context.Context, key string, value *models.User) error {
// 	byteValue, err := json.Marshal(&value)
// 	if err != nil {
// 		return err
// 	}

// 	message := p.buildMessage(key, byteValue)

// 	return p.userCreate.WriteMessages(ctx, message)

// }

// func (p *producer) buildMessage(key string, value []byte) kafka.Message {
// 	return kafka.Message{
// 		Key:   []byte(key),
// 		Value: value,
// 	}
// }

// func (p *producer) Close() {
// 	if err := p.userCreate.Close(); err != nil {
// 		p.logger.Error("error during close writer investmentCreated", zap.Error(err))
// 	}
// }
