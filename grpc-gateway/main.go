package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/klim0v/benchmark-go-swagger-and-grpc-gateway/grpc-gateway/shop"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type service struct {
}

func (s *service) GetProduct(_ context.Context, _ *shop.ProductRequest) (*shop.ProductResponse, error) {
	return &shop.ProductResponse{
		Title:    "A",
		Favorite: true,
		Price:    1999.99,
		Count:    200,
		Image:    "https://image",
		Vote: []*shop.ProductResponse_Vote{
			{
				Name: "A",
				Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
			},
		},
	}, nil
}

func (s *service) ListProducts(_ context.Context, _ *empty.Empty) (*shop.ListProductsResponse, error) {
	return &shop.ListProductsResponse{
		Products: []*shop.ProductResponse{
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    200,
				Image:    "https://image",
				Vote: []*shop.ProductResponse_Vote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
		},
	}, nil
}

func main() {
	addrGRPC := ":9090"
	addrApi := ":8080"
	lis, err := net.Listen("tcp", addrGRPC)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	shop.RegisterShopServer(grpcServer, &service{})
	var group errgroup.Group
	group.Go(func() error { return grpcServer.Serve(lis) })
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = shop.RegisterShopHandlerFromEndpoint(context.Background(), gwmux, addrGRPC, opts)
	if err != nil {
		panic(err)
	}
	group.Go(func() error { return http.ListenAndServe(addrApi, gwmux) })
	log.Fatal(group.Wait())
}
