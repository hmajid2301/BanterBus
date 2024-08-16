package service

import (
	"context"
	"fmt"

	"gitlab.com/hmajid2301/banterbus/internal/entities"
	"gitlab.com/hmajid2301/banterbus/internal/store"
)

type PlayerService struct {
	store      store.Store
	randomizer UserRandomizer
}

func NewPlayerService(store store.Store, randomizer UserRandomizer) *PlayerService {
	return &PlayerService{store: store, randomizer: randomizer}
}

func (p *PlayerService) UpdateNickname(ctx context.Context, nickname string, playerID string) (entities.Room, error) {
	playerRows, err := p.store.UpdateNickname(ctx, nickname, playerID)
	if err != nil {
		return entities.Room{}, err
	}

	if len(playerRows) == 0 {
		return entities.Room{}, fmt.Errorf("no players in room")
	}

	room := getRoom(playerRows, playerRows[0].RoomCode)
	return room, err
}

func (p *PlayerService) GenerateNewAvatar(ctx context.Context, playerID string) (entities.Room, error) {
	avatar := p.randomizer.GetAvatar()

	playerRows, err := p.store.UpdateAvatar(ctx, avatar, playerID)
	if err != nil {
		return entities.Room{}, err
	}

	if len(playerRows) == 0 {
		return entities.Room{}, fmt.Errorf("no players in room")
	}

	room := getRoom(playerRows, playerRows[0].RoomCode)
	return room, err
}
