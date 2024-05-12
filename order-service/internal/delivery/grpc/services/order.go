package services

import (
	"context"
	"github.com/k0kubun/pp"
	pb "projects/order-service/genproto/order"
	"projects/order-service/internal/entity"
	"projects/order-service/internal/pkg/otlp"
	"projects/order-service/internal/usecase"
	"projects/order-service/internal/usecase/event"
	"strconv"
	"time"

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

	req_order := entity.Order{
		Id:         int(order.Id),
		WorkerId:   order.WorkerId,
		ProductId:  order.ProductId,
		Tax:        order.Tax,
		Discount:   order.Discount,
		TotalPrice: order.TotalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	res_ordered, err := u.order.CreateOrder(ctx, &req_order)
	if err != nil {
		u.logger.Error("Create order error", zap.Error(err))
		return &pb.Order{}, nil
	}

	return &pb.Order{
		Id:         int64(res_ordered.Id),
		WorkerId:   res_ordered.WorkerId,
		ProductId:  res_ordered.ProductId,
		Tax:        res_ordered.Tax,
		Discount:   res_ordered.Discount,
		TotalPrice: res_ordered.TotalPrice,
		CreatedAt:  res_ordered.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  res_ordered.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u orderRPC) GetOrder(ctx context.Context, id *pb.Id) (*pb.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Get")
	span.SetAttributes(attribute.String("get", "order"))
	defer span.End()

	res, err := u.order.GetOrder(ctx, strconv.FormatInt(id.Id, 10))

	if err != nil {
		u.logger.Error("get order error", zap.Error(err))
		return &pb.Order{}, nil
	}

	return &pb.Order{
		Id:         id.Id,
		WorkerId:   res.WorkerId,
		ProductId:  res.ProductId,
		Tax:        res.Tax,
		Discount:   res.Discount,
		TotalPrice: res.TotalPrice,
		CreatedAt:  res.CreatedAt.String(),
		UpdatedAt:  res.UpdatedAt.String(),
	}, nil
}
func (u orderRPC) DeleteOrder(ctx context.Context, id *pb.Id) (*pb.DeleteStatus, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Delete")
	span.SetAttributes(attribute.String("delete", "order"))
	defer span.End()

	err := u.order.DeleteOrder(ctx, strconv.FormatInt(id.Id, 10))
	if err != nil {
		u.logger.Error("delete order error", zap.Error(err))
		return &pb.DeleteStatus{Status: false}, nil
	}

	return &pb.DeleteStatus{Status: true}, nil
}
func (u orderRPC) UpdateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Update")
	span.SetAttributes(attribute.String("update", "order"))
	defer span.End()

	updated_order := entity.Order{
		Id:         int(order.Id),
		WorkerId:   order.WorkerId,
		ProductId:  order.ProductId,
		Tax:        order.Tax,
		Discount:   order.Discount,
		TotalPrice: order.TotalPrice,
		UpdatedAt:  time.Now(),
	}

	updated_ordered, err := u.order.UpdateOrder(ctx, &updated_order)
	if err != nil {
		u.logger.Error("update order error", zap.Error(err))
		return nil, err
	}

	return &pb.Order{
		Id:         int64(updated_ordered.Id),
		WorkerId:   updated_ordered.WorkerId,
		ProductId:  updated_ordered.ProductId,
		Tax:        updated_ordered.Tax,
		Discount:   updated_ordered.Discount,
		TotalPrice: updated_ordered.TotalPrice,
		CreatedAt:  updated_ordered.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  updated_ordered.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u orderRPC) GetOrders(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
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

	var orders pb.GetAllResponse

	for _, in := range res_orders {
		orders.Orders = append(orders.Orders, &pb.Order{
			Id:         int64(in.Id),
			WorkerId:   in.WorkerId,
			ProductId:  in.ProductId,
			Tax:        in.Tax,
			Discount:   in.Discount,
			TotalPrice: in.TotalPrice,
			CreatedAt:  in.CreatedAt.String(),
			UpdatedAt:  in.UpdatedAt.String(),
		})
	}

	return &orders, nil
}
