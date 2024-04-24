package services

import (
	"context"
	"time"

	pb "order-service/genproto/order"
	"order-service/internal/entity"
	"order-service/internal/usecase"
	"order-service/internal/usecase/event"

	"go.uber.org/zap"
)

type orderRPC struct {
	logger         *zap.Logger
	order          usecase.Order
	brokerProducer event.BrokerProducer
}

func UserRPC(logger *zap.Logger, orderUsecase usecase.Order, brokerProducer event.BrokerProducer) pb.OrderServiceServer {
	return &orderRPC{
		logger:         logger,
		order:          orderUsecase,
		brokerProducer: brokerProducer,
	}
}

func (u orderRPC) CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {

	req_order := entity.Order{
		Id:         order.Id,
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
		Id:         res_ordered.Id,
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

	reqMap := make(map[string]int64)
	reqMap["id"] = id.Id

	res, err := u.order.GetOrder(ctx, reqMap)

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
func (u orderRPC) DeleteOrder(ctx context.Context, id *pb.Id) (*pb.Order, error) {

	reqMap := make(map[string]int64)
	reqMap["id"] = id.Id

	orderReq, err := u.order.GetOrder(ctx, reqMap)

	if err != nil {
		u.logger.Error("get order error", zap.Error(err))
		return nil, err
	}

	err = u.order.DeleteOrder(ctx, id.Id)
	if err != nil {
		u.logger.Error("delete order error", zap.Error(err))
		return &pb.Order{}, nil
	}

	return &pb.Order{
		Id:         orderReq.Id,
		WorkerId:   orderReq.WorkerId,
		ProductId:  orderReq.ProductId,
		Tax:        orderReq.Tax,
		Discount:   orderReq.Discount,
		TotalPrice: orderReq.TotalPrice,
		CreatedAt:  orderReq.CreatedAt.String(),
		UpdatedAt:  orderReq.UpdatedAt.String(),
	}, nil
}
func (u orderRPC) UpdateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {

	updated_order := entity.Order{
		Id:         order.Id,
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
		Id:         updated_ordered.Id,
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
	offset := req.Limit * (req.Page - 1)

	res_orders, err := u.order.GetOrders(ctx, uint64(req.Limit), uint64(offset), map[string]string{})

	if err != nil {
		u.logger.Error("get all orders error", zap.Error(err))
		return nil, err
	}

	var orders pb.GetAllResponse

	for _, in := range res_orders {
		orders.Orders = append(orders.Orders, &pb.Order{
			Id:         in.Id,
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
