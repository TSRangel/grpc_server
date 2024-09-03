package mappers

import (
	"github.com/TSRangel/grpc_server/application/pb"
	"github.com/TSRangel/grpc_server/domain/entities"
)

func CreateGameInputToEntity(input *pb.CreateGameInput) *entities.Game {
	return entities.NewGame(input.Name, input.Platform, input.Price)
}

func EntityToBaseGameOutput(game *entities.Game) *pb.BaseGameOutput {
	return &pb.BaseGameOutput{ID: game.ID, Name: game.Name, Platform: game.Platform, Price: game.Price}
}