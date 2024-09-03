package gameusecases

import (
	"context"
	"io"

	"github.com/TSRangel/grpc_server/application/mappers"
	"github.com/TSRangel/grpc_server/application/pb"
	repositoriesinterfaces "github.com/TSRangel/grpc_server/domain/repositories_interfaces"
	"google.golang.org/grpc"
)

type GameUseCase struct {
	pb.UnimplementedGameServiceServer
	Repository repositoriesinterfaces.GameRepositoryInterface
}

func NewGameUseCase(repository repositoriesinterfaces.GameRepositoryInterface) *GameUseCase {
	return &GameUseCase{Repository: repository}
}

func (gu *GameUseCase) CreateGame(ctx context.Context, input *pb.CreateGameInput) (*pb.BaseGameOutput, error) {
	newGame := mappers.CreateGameInputToEntity(input)
	createdGame, err := gu.Repository.Create(newGame)
	if err != nil {
		return &pb.BaseGameOutput{}, err
	}
	return mappers.EntityToBaseGameOutput(createdGame), nil
}

func (gu *GameUseCase) CreateGameStream(
	stream grpc.ClientStreamingServer[pb.CreateGameInput, pb.ListAllGamesOutput]) error {
	var games pb.ListAllGamesOutput

	for {
		game, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&games)
		}
		if err != nil {
			return err
		}
		newGame, err := gu.Repository.Create(mappers.CreateGameInputToEntity(game))
		if err != nil {
			return err
		}
		games.Games = append(games.Games, mappers.EntityToBaseGameOutput(newGame))
	}
}

func (gu *GameUseCase) CreateGameStreamBidirectional(
	stream grpc.BidiStreamingServer[pb.CreateGameInput, pb.BaseGameOutput]) error {
	for {
		game, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		newGame, err := gu.Repository.Create(mappers.CreateGameInputToEntity(game))
		if err != nil {
			return err
		}
		err = stream.Send(mappers.EntityToBaseGameOutput(newGame))
		if err != nil {
			return err
		}
	}
}

func (gu *GameUseCase) ListAllGames(ctx context.Context,input *pb.BlankInput) (*pb.ListAllGamesOutput, error) {
	var output pb.ListAllGamesOutput
	listedGames, err := gu.Repository.ListAll()
	if err != nil {
		return &output, err
	}

	for _, game := range listedGames {
		output.Games = append(output.Games, mappers.EntityToBaseGameOutput(game))
	}
	return &output, nil
}

func (gu *GameUseCase) GetGame(ctx context.Context, input *pb.GetGameInput) (*pb.BaseGameOutput, error) {
	game, err := gu.Repository.GetGame(input.Name)
	if err != nil {
		return &pb.BaseGameOutput{}, err
	}
	return mappers.EntityToBaseGameOutput(game), nil
}