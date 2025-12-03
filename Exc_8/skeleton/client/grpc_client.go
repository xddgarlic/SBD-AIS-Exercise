package client

import (
	"context"
	"fmt"

	"exc8/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

// NewGrpcClient connects to the server and returns a GrpcClient
func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.Dial(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %v", err)
	}

	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

// Run executes the client workflow: list drinks, order drinks, get totals
func (c *GrpcClient) Run() error {
	ctx := context.Background()
	fmt.Println("🚀 Client Run started")

	// -----------------------------
	// 1 List drinks
	// -----------------------------
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("GetDrinks failed: %v", err)
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%d  description:%q\n",
			d.Id, d.Name, d.Price, d.Description)
	}

	// Map to lookup names by ID
	drinkNames := make(map[int32]string)
	for _, d := range drinksResp.Drinks {
		drinkNames[d.Id] = d.Name
	}

	// -----------------------------
	// 2 Order first round: 2x each
	// -----------------------------
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 2 x %s\n", d.Name)
		_, err := c.client.OrderDrink(ctx, &pb.OrderDrinkRequest{
			DrinkId:  d.Id,
			Quantity: 2,
		})
		if err != nil {
			return fmt.Errorf("OrderDrink failed: %v", err)
		}
	}

	// -----------------------------
	// 3 Order second round: 6x each
	// -----------------------------
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 6 x %s\n", d.Name)
		_, err := c.client.OrderDrink(ctx, &pb.OrderDrinkRequest{
			DrinkId:  d.Id,
			Quantity: 6,
		})
		if err != nil {
			return fmt.Errorf("OrderDrink failed: %v", err)
		}
	}

	// -----------------------------
	// 4 Get aggregated orders
	// -----------------------------
	fmt.Println("Getting the bill 💹💹💹")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("GetOrders failed: %v", err)
	}

	for _, o := range ordersResp.Orders {
		name := drinkNames[o.DrinkId]
		fmt.Printf("\t> Total: %d x %s\n", o.Quantity, name)
	}

	fmt.Println("Orders complete!")
	return nil
}
