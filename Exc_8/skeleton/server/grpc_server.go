package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"exc8/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	drinks []*pb.Drink
	orders map[int32]int32
	mu     sync.Mutex
}

func StartGrpcServer() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	fmt.Println("✅ gRPC server listening on :4000")

	srv := grpc.NewServer()

	grpcService := &GRPCService{
		drinks: []*pb.Drink{
			{Id: 1, Name: "Blaufränkischer", Price: 2, Description: "Better than Zweigelt."},
			{Id: 2, Name: "Beer", Price: 3, Description: "The elixir of life."},
			{Id: 3, Name: "Espresso", Description: "Is it that sweet, I guess so."},
		},
		orders: make(map[int32]int32),
	}

	pb.RegisterOrderServiceServer(srv, grpcService)

	if err := srv.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

// GetDrinks returns the list of available drinks
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.GetDrinksResponse, error) {
	fmt.Println("📋 GetDrinks called")
	return &pb.GetDrinksResponse{
		Drinks: s.drinks,
	}, nil
}

// OrderDrink stores ordered drinks in memory
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderDrinkRequest) (*pb.OrderDrinkResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[req.DrinkId] += req.Quantity
	fmt.Printf("🍻 OrderDrink called: drink_id=%d quantity=%d\n", req.DrinkId, req.Quantity)

	return &pb.OrderDrinkResponse{}, nil
}

// GetOrders returns aggregated orders
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var orders []*pb.Order
	for drinkID, qty := range s.orders {
		orders = append(orders, &pb.Order{
			DrinkId:  drinkID,
			Quantity: qty,
		})
	}

	fmt.Println("💹 GetOrders called")
	return &pb.GetOrdersResponse{
		Orders: orders,
	}, nil
}
