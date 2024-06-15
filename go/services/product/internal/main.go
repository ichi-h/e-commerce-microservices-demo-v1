package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	v1product "product/pkg/api/v1"
	"product/pkg/api/v1/v1connect"
)

type ProductServer struct{}

func (s *ProductServer) GetProducts(
	ctx context.Context,
	req *connect.Request[v1product.GetProductsRequest],
) (*connect.Response[v1product.GetProductsResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1product.GetProductsResponse{
		Products: []*v1product.Product{
			{
				Id:    "f7fc1e30-c830-572e-6a80-e5b78ebac367",
				Name:  "Product 1",
				Price: 1000,
			},
			{
				Id:    "cd407a1a-4ec7-05da-49e4-73c18a5f4c3a",
				Name:  "Product 2",
				Price: 2000,
			},
			{
				Id:    "f7fc1e30-c830-572e-6a80-e5b78ebac367",
				Name:  "Product 3",
				Price: 3000,
			},
		},
	})
	res.Header().Set("GetProducts-Version", "v1")
	return res, nil
}

func main() {
	server := &ProductServer{}
	mux := http.NewServeMux()
	path, handler := v1connect.NewProductServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"0.0.0.0:1324",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
