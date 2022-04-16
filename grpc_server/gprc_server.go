package grpc_server

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCMain() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		logs.Sugar.Fatalf("listen error:%v", err)
		panic(err)
	}
	s := grpc.NewServer()
	grpc.WithTransportCredentials(insecure.NewCredentials())
	pb_gen.RegisterGraduateDesignApiServer(s, &server{})
	httpServer := provideHTTP(":9000", s)
	if err := httpServer.Serve(lis); err != nil {
		logs.Sugar.Fatalf("server error:%v", err)
		panic(err)
	}
}

func provideHTTP(endpoint string, s *grpc.Server) *http.Server {
	ctx := context.Background()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := pb_gen.RegisterGraduateDesignApiHandlerFromEndpoint(ctx, gwmux, endpoint, opts)
	if err != nil {
		logs.Sugar.Fatalf("RegisterGraduateDesignApiHandlerFromEndpoint error:%v", err)
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	return &http.Server{
		Addr:    endpoint,
		Handler: grpcHandleFunc(s, mux),
	}
}

func grpcHandleFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
