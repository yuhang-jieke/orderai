package server

import (
	"context"
	"errors"

	__ "github.com/yuhang-jieke/orderai/srv/proto"
	"github.com/yuhang-jieke/orderai/srv/user-server/basic/config"
	"github.com/yuhang-jieke/orderai/srv/user-server/model"
)

type Server struct {
	__.UnimplementedEcommerceServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) AddOrders(_ context.Context, in *__.AddOrdersReq) (*__.AddOrdersResp, error) {
	order := model.Orders{
		Name:  in.Name,
		Num:   int(in.Num),
		Price: in.Price,
	}
	err := order.OrderAdd(config.DB)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &__.AddOrdersResp{
		Message: "添加成功",
	}, nil
}
func (s *Server) UpdateOrders(_ context.Context, in *__.UpdateOrdersReq) (*__.UpdateOrdersResp, error) {
	var order model.Orders

	err := order.UpdateId(config.DB, in)
	if err != nil {
		return nil, errors.New("修改失败")
	}
	return &__.UpdateOrdersResp{
		Message: "修改成功",
	}, nil
}
func (s *Server) DelOrders(_ context.Context, in *__.DelOrdersReq) (*__.DelOrdersResp, error) {
	var order model.Orders

	err := order.DelId(config.DB, in)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &__.DelOrdersResp{
		Message: "删除成功",
	}, nil
}
func (s *Server) GetOrdersById(_ context.Context, in *__.GetOrdersByIdReq) (*__.GetOrdersByIdResp, error) {
	var order model.Orders
	id, err := order.GetId(config.DB, in)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	return &__.GetOrdersByIdResp{
		Orders: &__.Orders{
			Name:  id.Name,
			Num:   int64(id.Num),
			Price: id.Price,
			Id:    int64(id.ID),
		},
	}, nil
}
func (s *Server) SearchOrders(_ context.Context, in *__.SearchOrdersReq) (*__.SearchOrdersResp, error) {
	var order model.Orders
	var search []model.Orders
	search, err := order.Search(config.DB, in)
	if err != nil {
		return nil, errors.New("搜索失败")
	}
	var list []*__.Orders
	for _, orders := range search {
		list = append(list, &__.Orders{
			Name:  orders.Name,
			Num:   int64(orders.Num),
			Price: orders.Price,
			Id:    int64(orders.ID),
		})
	}
	return &__.SearchOrdersResp{
		Orders: list,
	}, nil
}
