package handlers

import (
	"context"
	"encoding/json"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/kafka"
	"user-service/internal/pkg/config"
	"user-service/internal/usecase"
	"user-service/internal/usecase/event"

	"go.uber.org/zap"
)

type ownerCreateHandler struct {
	config         *config.Config
	brokerConsumer event.BrokerConsumer
	logger         *zap.Logger
	ownerUsecase    usecase.Owner
}

func NewUserCreateHandler(config *config.Config,
	brokerConsumer event.BrokerConsumer,
	logger *zap.Logger,
	ownerUsecase usecase.Owner) *ownerCreateHandler {
	return &ownerCreateHandler{
		config:         config,
		brokerConsumer: brokerConsumer,
		logger:         logger,
		ownerUsecase:    ownerUsecase,
	}
}

func (h *ownerCreateHandler) HandlerEvents() error {
	consumerConfig := kafka.NewConsumerConfig(
		h.config.Kafka.Address,
		"api.owner.create",
		"1",
		func(ctx context.Context, key, value []byte) error {
			var owner *entity.Owner

			if err := json.Unmarshal(value, &owner); err != nil {
				return err
			}

			if _, err := h.ownerUsecase.CreateOwner(ctx, owner); err != nil {
				return err
			}

			return nil
		},
	)

	h.brokerConsumer.RegisterConsumer(consumerConfig)
	h.brokerConsumer.Run()

	return nil

}
