package service

import (
	"gitlab.com/hmajid2301/banterbus/internal/entities"

	sqlc "gitlab.com/hmajid2301/banterbus/internal/store/db"
)

func getRoom(playerRows []sqlc.GetAllPlayersInRoomRow, roomCode string) entities.Room {
	var players []entities.Player
	for _, player := range playerRows {
		p := entities.Player{
			ID:       player.ID,
			Nickname: player.Nickname,
			Avatar:   string(player.Avatar),
			IsReady:  player.IsReady.Bool,
		}

		players = append(players, p)
	}

	room := entities.Room{
		Code:    roomCode,
		Players: players,
	}
	return room
}
