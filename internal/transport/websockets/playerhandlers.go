package websockets

import (
	"context"
	"fmt"

	"gitlab.com/hmajid2301/banterbus/internal/entities"
	"gitlab.com/hmajid2301/banterbus/internal/store"
)

type PlayerServicer interface {
	UpdateNickname(ctx context.Context, nickname string, playerID string) (entities.Lobby, error)
	GenerateNewAvatar(ctx context.Context, playerID string) (entities.Lobby, error)
	TogglePlayerIsReady(ctx context.Context, playerID string) (entities.Lobby, error)
	GetRoomState(ctx context.Context, playerID string) (store.RoomState, error)
	GetLobby(ctx context.Context, playerID string) (entities.Lobby, error)
	GetGameState(ctx context.Context, playerID string) (entities.GameState, error)
}

func (u *UpdateNickname) Handle(ctx context.Context, client *client, sub *Subscriber) error {
	updatedRoom, err := sub.playerService.UpdateNickname(ctx, u.PlayerNickname, client.playerID)
	if err != nil {
		errStr := "failed to update nickname"
		if err == entities.ErrNicknameExists {
			errStr = err.Error()
		}
		clientErr := sub.updateClientAboutErr(ctx, client.playerID, errStr)
		return fmt.Errorf("%w: %w", err, clientErr)
	}

	err = sub.updateClientsAboutLobby(ctx, updatedRoom)
	return err
}

func (g *GenerateNewAvatar) Handle(ctx context.Context, client *client, sub *Subscriber) error {
	updatedRoom, err := sub.playerService.GenerateNewAvatar(ctx, client.playerID)
	if err != nil {
		errStr := "failed to generate new avatar"
		clientErr := sub.updateClientAboutErr(ctx, client.playerID, errStr)
		return fmt.Errorf("%w: %w", err, clientErr)
	}

	err = sub.updateClientsAboutLobby(ctx, updatedRoom)
	return err
}

// TODO: join errors in all handlers rather than fmt them
func (t *TogglePlayerIsReady) Handle(ctx context.Context, client *client, sub *Subscriber) error {
	updatedRoom, err := sub.playerService.TogglePlayerIsReady(ctx, client.playerID)
	if err != nil {
		errStr := "failed to update ready status"
		clientErr := sub.updateClientAboutErr(ctx, client.playerID, errStr)
		return fmt.Errorf("%w: %w", err, clientErr)
	}

	err = sub.updateClientsAboutLobby(ctx, updatedRoom)
	return err
}
