package repositoriesinterfaces

import "github.com/TSRangel/grpc_server/domain/entities"

type GameRepositoryInterface  interface {
	Create(game *entities.Game) (*entities.Game, error)
	ListAll() ([]*entities.Game, error)
	GetGame(name string) (*entities.Game, error)
}