package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/hmajid2301/banterbus/internal/banterbustest"
	"gitlab.com/hmajid2301/banterbus/internal/service"
)

func setupSubtest(t *testing.T) (*sql.DB, func()) {
	ctx := context.Background()
	db, err := banterbustest.CreateDB(ctx)
	require.NoError(t, err)

	return db, func() {
		db.Close()
	}
}

func createRoom(ctx context.Context, srv *service.LobbyService) (service.Lobby, error) {
	newPlayer := service.NewHostPlayer{
		ID: defaultHostPlayerID,
	}
	lobby, err := srv.Create(ctx, "fibbing_it", newPlayer)
	return lobby, err
}

func lobbyWithTwoPlayers(ctx context.Context, srv *service.LobbyService) (service.Lobby, error) {
	newPlayer := service.NewHostPlayer{
		ID:       defaultHostPlayerID,
		Nickname: defaultHostNickname,
	}
	lobby, err := srv.Create(ctx, "fibbing_it", newPlayer)
	if err != nil {
		return lobby, err
	}

	lobby, err = srv.Join(ctx, lobby.Code, defaultOtherPlayerID, defaultOtherPlayerNickname)
	return lobby, err
}

// func startGame(
// 	ctx context.Context,
// 	lobbySrv *service.LobbyService,
// 	playerSrv *service.PlayerService,
// ) (service.Lobby, error) {
// 	newPlayer := service.NewHostPlayer{
// 		ID:       defaultHostPlayerID,
// 		Nickname: defaultHostNickname,
// 	}
// 	lobby, err := lobbySrv.Create(ctx, "fibbing_it", newPlayer)
// 	if err != nil {
// 		return lobby, err
// 	}
//
// 	lobby, err = lobbySrv.Join(ctx, lobby.Code, defaultOtherPlayerID, defaultOtherPlayerNickname)
// 	if err != nil {
// 		return lobby, err
// 	}
//
// 	_, err = playerSrv.TogglePlayerIsReady(ctx, defaultHostPlayerID)
// 	if err != nil {
// 		return lobby, err
// 	}
//
// 	_, err = playerSrv.TogglePlayerIsReady(ctx, defaultOtherPlayerID)
// 	if err != nil {
// 		return lobby, err
// 	}
//
// 	return lobby, err
// }