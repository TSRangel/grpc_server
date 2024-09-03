package main

import (
	"net"

	"github.com/TSRangel/grpc_server/application/pb"
	gameusecases "github.com/TSRangel/grpc_server/application/use_cases/game_use_cases"
	"github.com/TSRangel/grpc_server/infrastructure/repositories"
	"github.com/TSRangel/grpc_server/pkg/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dbConn, err := tools.NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	gameRepository := repositories.NewGameRepository(dbConn)
	gameUseCases := gameusecases.NewGameUseCase(gameRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterGameServiceServer(grpcServer, gameUseCases)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}