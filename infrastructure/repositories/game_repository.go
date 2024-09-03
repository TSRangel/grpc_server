package repositories

import (
	"database/sql"

	"github.com/TSRangel/grpc_server/domain/entities"
)

type GameRepository struct {
	DB *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{DB: db}
}

func (gr *GameRepository) Create(game *entities.Game) (*entities.Game, error) {
	stmt, err := gr.DB.Prepare("INSERT INTO games (id, name, platform, price) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return &entities.Game{}, nil
	}
	defer stmt.Close()
	if _, err := stmt.Exec(game.ID, game.Name, game.Platform, game.Price);
	err != nil {
		return &entities.Game{}, nil
	}

	return game, nil
}

func (gr *GameRepository) ListAll() ([]*entities.Game, error) {
	var games []*entities.Game
	rows, err := gr.DB.Query("SELECT id, name, platform, price FROM games")
	if err != nil {
		return games, err
	}
	
	for rows.Next(){
		var game entities.Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Platform, &game.Price);
		err != nil {
			return games, err
		}
		games = append(games, &game)	
	}

	return games, nil
}

func (gr *GameRepository) GetGame(name string) (*entities.Game, error) {
	var game entities.Game
	stmt, err := gr.DB.Prepare("SELECT id, name, platform, price FROM games WHERE name = $1")
	if err != nil {
		return &game, err
	}
	defer stmt.Close()
	if err := stmt.QueryRow(name).Scan(&game.ID, &game.Name, &game.Platform, &game.Price);
	err != nil {
		return &game, err
	}
	return &game, nil
}