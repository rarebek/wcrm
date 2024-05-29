package services

import (
	"context"
	pb "projects/order-service/genproto/order"
	"projects/order-service/internal/entity"
	"projects/order-service/internal/pkg/otlp"
	"projects/order-service/internal/usecase"
	"projects/order-service/internal/usecase/event"
	"time"

	"github.com/k0kubun/pp"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

const (
	orderTableName      = "orders"
	orderServiceName    = "orderService"
	orderSpanRepoPrefix = "orderRepo"
)

type orderRPC struct {
	logger         *zap.Logger
	order          usecase.Order
	brokerProducer event.BrokerProducer
	pb.UnimplementedOrderServiceServer
}

func UserRPC(logger *zap.Logger, orderUsecase usecase.Order, brokerProducer event.BrokerProducer) pb.OrderServiceServer {
	return &orderRPC{
		logger:         logger,
		order:          orderUsecase,
		brokerProducer: brokerProducer,
	}
}

func (u orderRPC) CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Create")
	span.SetAttributes(attribute.String("create", "order"))
	defer span.End()

	reqOrder := entity.Order{
		Id:         order.Id,
		WorkerId:   order.WorkerId,
		ProductIds: order.ProductIds,
		Tax:        order.Tax,
		TotalPrice: order.TotalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resOrdered, err := u.order.CreateOrder(ctx, &reqOrder)
	if err != nil {
		u.logger.Error("Create order error", zap.Error(err))
		return &pb.Order{}, nil
	}

	return &pb.Order{
		Id:         resOrdered.Id,
		WorkerId:   resOrdered.WorkerId,
		ProductIds: resOrdered.ProductIds,
		Tax:        resOrdered.Tax,
		TotalPrice: resOrdered.TotalPrice,
		CreatedAt:  resOrdered.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  resOrdered.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u orderRPC) GetOrder(ctx context.Context, id *pb.OrderId) (*pb.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Get")
	span.SetAttributes(attribute.String("get", "order"))
	defer span.End()

	res, err := u.order.GetOrder(ctx, id.Id)
	if err != nil {
		u.logger.Error("get order error", zap.Error(err))
		return &pb.Order{}, nil
	}

	return &pb.Order{
		Id:         res.Id,
		WorkerId:   res.WorkerId,
		ProductIds: res.ProductIds,
		Tax:        res.Tax,
		TotalPrice: res.TotalPrice,
		CreatedAt:  res.CreatedAt.String(),
		UpdatedAt:  res.UpdatedAt.String(),
	}, nil
}

func (u orderRPC) DeleteOrder(ctx context.Context, id *pb.OrderId) (*pb.Empty, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Delete")
	span.SetAttributes(attribute.String("delete", "order"))
	defer span.End()

	err := u.order.DeleteOrder(ctx, id.Id)
	if err != nil {
		u.logger.Error("delete order error", zap.Error(err))
		return &pb.Empty{}, nil
	}

	return &pb.Empty{}, nil
}

func (u orderRPC) UpdateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Update")
	span.SetAttributes(attribute.String("update", "order"))
	defer span.End()

	updated_order := entity.Order{
		Id:         order.Id,
		WorkerId:   order.WorkerId,
		ProductIds: order.ProductIds,
		Tax:        order.Tax,
		TotalPrice: order.TotalPrice,
		UpdatedAt:  time.Now(),
	}

	updated_ordered, err := u.order.UpdateOrder(ctx, &updated_order)
	if err != nil {
		u.logger.Error("update order error", zap.Error(err))
		return nil, err
	}

	return &pb.Order{
		Id:         updated_order.Id,
		WorkerId:   updated_ordered.WorkerId,
		ProductIds: updated_ordered.ProductIds,
		Tax:        updated_ordered.Tax,
		TotalPrice: updated_ordered.TotalPrice,
		CreatedAt:  updated_ordered.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  updated_ordered.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u orderRPC) GetOrders(ctx context.Context, req *pb.GetAllOrderRequest) (*pb.GetAllOrderResponse, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Gets")
	span.SetAttributes(attribute.String("gets", "order"))
	defer span.End()

	offset := req.Limit * (req.Page - 1)
	pp.Println(offset, req.Limit, req.Page)

	res_orders, err := u.order.GetOrders(ctx, uint64(req.Limit), uint64(offset), map[string]string{})

	if err != nil {
		u.logger.Error("get all orders error", zap.Error(err))
		return nil, err
	}

	var orders pb.GetAllOrderResponse

	for _, in := range res_orders {
		orders.Orders = append(orders.Orders, &pb.Order{
			Id:         in.Id,
			WorkerId:   in.WorkerId,
			ProductIds: in.ProductIds,
			Tax:        in.Tax,
			TotalPrice: in.TotalPrice,
			CreatedAt:  in.CreatedAt.String(),
			UpdatedAt:  in.UpdatedAt.String(),
		})
	}

	return &orders, nil
}
