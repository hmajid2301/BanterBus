// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addFibbingItAnswer = `-- name: AddFibbingItAnswer :one
INSERT INTO fibbing_it_answers (id, answer, round_id, player_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, answer, player_id, round_id, is_ready
`

type AddFibbingItAnswerParams struct {
	ID       uuid.UUID
	Answer   string
	RoundID  uuid.UUID
	PlayerID uuid.UUID
}

func (q *Queries) AddFibbingItAnswer(ctx context.Context, arg AddFibbingItAnswerParams) (FibbingItAnswer, error) {
	row := q.db.QueryRow(ctx, addFibbingItAnswer,
		arg.ID,
		arg.Answer,
		arg.RoundID,
		arg.PlayerID,
	)
	var i FibbingItAnswer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Answer,
		&i.PlayerID,
		&i.RoundID,
		&i.IsReady,
	)
	return i, err
}

const addFibbingItRole = `-- name: AddFibbingItRole :one
INSERT INTO fibbing_it_player_roles (id, player_role, round_id, player_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, player_role, round_id, player_id
`

type AddFibbingItRoleParams struct {
	ID         uuid.UUID
	PlayerRole string
	RoundID    uuid.UUID
	PlayerID   uuid.UUID
}

func (q *Queries) AddFibbingItRole(ctx context.Context, arg AddFibbingItRoleParams) (FibbingItPlayerRole, error) {
	row := q.db.QueryRow(ctx, addFibbingItRole,
		arg.ID,
		arg.PlayerRole,
		arg.RoundID,
		arg.PlayerID,
	)
	var i FibbingItPlayerRole
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PlayerRole,
		&i.RoundID,
		&i.PlayerID,
	)
	return i, err
}

const addFibbingItRound = `-- name: AddFibbingItRound :one
INSERT INTO fibbing_it_rounds (id, round_type, round, fibber_question_id, normal_question_id, room_id, game_state_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at, round_type, round, fibber_question_id, normal_question_id, room_id, game_state_id
`

type AddFibbingItRoundParams struct {
	ID               uuid.UUID
	RoundType        string
	Round            int32
	FibberQuestionID uuid.UUID
	NormalQuestionID uuid.UUID
	RoomID           uuid.UUID
	GameStateID      uuid.UUID
}

func (q *Queries) AddFibbingItRound(ctx context.Context, arg AddFibbingItRoundParams) (FibbingItRound, error) {
	row := q.db.QueryRow(ctx, addFibbingItRound,
		arg.ID,
		arg.RoundType,
		arg.Round,
		arg.FibberQuestionID,
		arg.NormalQuestionID,
		arg.RoomID,
		arg.GameStateID,
	)
	var i FibbingItRound
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoundType,
		&i.Round,
		&i.FibberQuestionID,
		&i.NormalQuestionID,
		&i.RoomID,
		&i.GameStateID,
	)
	return i, err
}

const addGameState = `-- name: AddGameState :one
INSERT INTO game_state (id, room_id, submit_deadline, state) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, room_id, submit_deadline, state
`

type AddGameStateParams struct {
	ID             uuid.UUID
	RoomID         uuid.UUID
	SubmitDeadline pgtype.Timestamp
	State          string
}

func (q *Queries) AddGameState(ctx context.Context, arg AddGameStateParams) (GameState, error) {
	row := q.db.QueryRow(ctx, addGameState,
		arg.ID,
		arg.RoomID,
		arg.SubmitDeadline,
		arg.State,
	)
	var i GameState
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoomID,
		&i.SubmitDeadline,
		&i.State,
	)
	return i, err
}

const addPlayer = `-- name: AddPlayer :one
INSERT INTO players (id, avatar, nickname) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at, avatar, nickname, is_ready
`

type AddPlayerParams struct {
	ID       uuid.UUID
	Avatar   []byte
	Nickname string
}

func (q *Queries) AddPlayer(ctx context.Context, arg AddPlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, addPlayer, arg.ID, arg.Avatar, arg.Nickname)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.IsReady,
	)
	return i, err
}

const addQuestion = `-- name: AddQuestion :one
INSERT INTO questions (id, game_name, round, question, language_code, group_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at, game_name, round, enabled, question, language_code, group_id
`

type AddQuestionParams struct {
	ID           uuid.UUID
	GameName     string
	Round        string
	Question     string
	LanguageCode string
	GroupID      uuid.UUID
}

func (q *Queries) AddQuestion(ctx context.Context, arg AddQuestionParams) (Question, error) {
	row := q.db.QueryRow(ctx, addQuestion,
		arg.ID,
		arg.GameName,
		arg.Round,
		arg.Question,
		arg.LanguageCode,
		arg.GroupID,
	)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.Round,
		&i.Enabled,
		&i.Question,
		&i.LanguageCode,
		&i.GroupID,
	)
	return i, err
}

const addQuestionsGroup = `-- name: AddQuestionsGroup :one
INSERT INTO questions_groups (id, group_name, group_type) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at, group_name, group_type
`

type AddQuestionsGroupParams struct {
	ID        uuid.UUID
	GroupName string
	GroupType string
}

func (q *Queries) AddQuestionsGroup(ctx context.Context, arg AddQuestionsGroupParams) (QuestionsGroup, error) {
	row := q.db.QueryRow(ctx, addQuestionsGroup, arg.ID, arg.GroupName, arg.GroupType)
	var i QuestionsGroup
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GroupName,
		&i.GroupType,
	)
	return i, err
}

const addRoom = `-- name: AddRoom :one
INSERT INTO rooms (id, game_name, host_player, room_code, room_state) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, game_name, host_player, room_state, room_code
`

type AddRoomParams struct {
	ID         uuid.UUID
	GameName   string
	HostPlayer uuid.UUID
	RoomCode   string
	RoomState  string
}

func (q *Queries) AddRoom(ctx context.Context, arg AddRoomParams) (Room, error) {
	row := q.db.QueryRow(ctx, addRoom,
		arg.ID,
		arg.GameName,
		arg.HostPlayer,
		arg.RoomCode,
		arg.RoomState,
	)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.HostPlayer,
		&i.RoomState,
		&i.RoomCode,
	)
	return i, err
}

const addRoomPlayer = `-- name: AddRoomPlayer :one
INSERT INTO rooms_players (room_id, player_id) VALUES ($1, $2) RETURNING room_id, player_id, created_at, updated_at
`

type AddRoomPlayerParams struct {
	RoomID   uuid.UUID
	PlayerID uuid.UUID
}

func (q *Queries) AddRoomPlayer(ctx context.Context, arg AddRoomPlayerParams) (RoomsPlayer, error) {
	row := q.db.QueryRow(ctx, addRoomPlayer, arg.RoomID, arg.PlayerID)
	var i RoomsPlayer
	err := row.Scan(
		&i.RoomID,
		&i.PlayerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllPlayerAnswerIsReady = `-- name: GetAllPlayerAnswerIsReady :one
SELECT
    COUNT(*) = SUM(CASE WHEN COALESCE(fa.is_ready, FALSE) THEN 1 ELSE 0 END) AS all_players_ready
FROM rooms_players rp
LEFT JOIN fibbing_it_answers fa ON fa.player_id = rp.player_id AND fa.round_id = (
    SELECT fir.id
    FROM fibbing_it_rounds fir
    WHERE fir.room_id = rp.room_id
    ORDER BY fir.created_at DESC
    LIMIT 1
)
WHERE rp.room_id = (
    SELECT room_id
    FROM rooms_players rp
    WHERE rp.player_id = $1
    LIMIT 1
)
`

func (q *Queries) GetAllPlayerAnswerIsReady(ctx context.Context, playerID uuid.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, getAllPlayerAnswerIsReady, playerID)
	var all_players_ready bool
	err := row.Scan(&all_players_ready)
	return all_players_ready, err
}

const getAllPlayerByRoomCode = `-- name: GetAllPlayerByRoomCode :many
SELECT p.id, p.created_at, p.updated_at, p.avatar, p.nickname, p.is_ready, r.room_code, r.host_player
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
WHERE rp.room_id = (
    SELECT r_inner.id
    FROM rooms r_inner
    WHERE r_inner.room_code = $1 AND (r_inner.room_state = 'CREATED' OR r_inner.room_state = 'PLAYING')
)
ORDER BY p.created_at
`

type GetAllPlayerByRoomCodeRow struct {
	ID         uuid.UUID
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	Avatar     []byte
	Nickname   string
	IsReady    pgtype.Bool
	RoomCode   string
	HostPlayer uuid.UUID
}

func (q *Queries) GetAllPlayerByRoomCode(ctx context.Context, roomCode string) ([]GetAllPlayerByRoomCodeRow, error) {
	rows, err := q.db.Query(ctx, getAllPlayerByRoomCode, roomCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPlayerByRoomCodeRow
	for rows.Next() {
		var i GetAllPlayerByRoomCodeRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Avatar,
			&i.Nickname,
			&i.IsReady,
			&i.RoomCode,
			&i.HostPlayer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPlayersInRoom = `-- name: GetAllPlayersInRoom :many
SELECT p.id, p.created_at, p.updated_at, p.avatar, p.nickname, p.is_ready, r.room_code, r.host_player
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
WHERE rp.room_id = (
    SELECT rp_inner.room_id
    FROM rooms_players rp_inner
    WHERE rp_inner.player_id = $1
)
`

type GetAllPlayersInRoomRow struct {
	ID         uuid.UUID
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	Avatar     []byte
	Nickname   string
	IsReady    pgtype.Bool
	RoomCode   string
	HostPlayer uuid.UUID
}

func (q *Queries) GetAllPlayersInRoom(ctx context.Context, playerID uuid.UUID) ([]GetAllPlayersInRoomRow, error) {
	rows, err := q.db.Query(ctx, getAllPlayersInRoom, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPlayersInRoomRow
	for rows.Next() {
		var i GetAllPlayersInRoomRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Avatar,
			&i.Nickname,
			&i.IsReady,
			&i.RoomCode,
			&i.HostPlayer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCurrentQuestionByPlayerID = `-- name: GetCurrentQuestionByPlayerID :one
SELECT
    gs.id AS game_state_id,
    fr.round,
    fr.round_type,
    r.room_code,
    gs.submit_deadline,
    p.id AS player_id,
    p.nickname,
    fpr.player_role AS role,
    COALESCE(fq1.question, fq2.question) AS question,
    p.avatar,
    COALESCE(fia.is_ready, FALSE) AS is_answer_ready
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
JOIN fibbing_it_rounds fr ON r.id = fr.room_id
JOIN game_state gs ON fr.game_state_id = gs.id
LEFT JOIN questions fq1 ON fr.fibber_question_id = fq1.id
LEFT JOIN questions fq2 ON fr.normal_question_id = fq2.id
LEFT JOIN fibbing_it_player_roles fpr ON p.id = fpr.player_id AND fr.id = fpr.round_id
LEFT JOIN fibbing_it_answers fia ON p.id = fia.player_id AND fr.id = fia.round_id
WHERE p.id = $1
ORDER BY p.created_at
LIMIT 1
`

type GetCurrentQuestionByPlayerIDRow struct {
	GameStateID    uuid.UUID
	Round          int32
	RoundType      string
	RoomCode       string
	SubmitDeadline pgtype.Timestamp
	PlayerID       uuid.UUID
	Nickname       string
	Role           pgtype.Text
	Question       string
	Avatar         []byte
	IsAnswerReady  bool
}

func (q *Queries) GetCurrentQuestionByPlayerID(ctx context.Context, id uuid.UUID) (GetCurrentQuestionByPlayerIDRow, error) {
	row := q.db.QueryRow(ctx, getCurrentQuestionByPlayerID, id)
	var i GetCurrentQuestionByPlayerIDRow
	err := row.Scan(
		&i.GameStateID,
		&i.Round,
		&i.RoundType,
		&i.RoomCode,
		&i.SubmitDeadline,
		&i.PlayerID,
		&i.Nickname,
		&i.Role,
		&i.Question,
		&i.Avatar,
		&i.IsAnswerReady,
	)
	return i, err
}

const getGameState = `-- name: GetGameState :one
SELECT
    gs.id,
    gs.created_at,
    gs.updated_at,
    gs.room_id,
    gs.submit_deadline,
    gs.state
FROM game_state gs
WHERE gs.id = $1
`

func (q *Queries) GetGameState(ctx context.Context, id uuid.UUID) (GameState, error) {
	row := q.db.QueryRow(ctx, getGameState, id)
	var i GameState
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoomID,
		&i.SubmitDeadline,
		&i.State,
	)
	return i, err
}

const getGameStateByPlayerID = `-- name: GetGameStateByPlayerID :one
SELECT
    gs.id,
    gs.created_at,
    gs.updated_at,
    gs.room_id,
    gs.submit_deadline,
    gs.state
FROM game_state gs
JOIN rooms_players rp ON gs.room_id = rp.room_id
WHERE rp.player_id = $1
`

func (q *Queries) GetGameStateByPlayerID(ctx context.Context, playerID uuid.UUID) (GameState, error) {
	row := q.db.QueryRow(ctx, getGameStateByPlayerID, playerID)
	var i GameState
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoomID,
		&i.SubmitDeadline,
		&i.State,
	)
	return i, err
}

const getLatestRoundByPlayerID = `-- name: GetLatestRoundByPlayerID :one
SELECT fir.id, fir.created_at, fir.updated_at, fir.round_type, fir.round, fir.fibber_question_id, fir.normal_question_id, fir.room_id, fir.game_state_id, gs.submit_deadline
FROM fibbing_it_rounds fir
JOIN rooms_players rp ON fir.room_id = rp.room_id
JOIN game_state gs ON fir.room_id = gs.room_id
WHERE rp.player_id = $1
ORDER BY fir.created_at DESC
LIMIT 1
`

type GetLatestRoundByPlayerIDRow struct {
	ID               uuid.UUID
	CreatedAt        pgtype.Timestamp
	UpdatedAt        pgtype.Timestamp
	RoundType        string
	Round            int32
	FibberQuestionID uuid.UUID
	NormalQuestionID uuid.UUID
	RoomID           uuid.UUID
	GameStateID      uuid.UUID
	SubmitDeadline   pgtype.Timestamp
}

func (q *Queries) GetLatestRoundByPlayerID(ctx context.Context, playerID uuid.UUID) (GetLatestRoundByPlayerIDRow, error) {
	row := q.db.QueryRow(ctx, getLatestRoundByPlayerID, playerID)
	var i GetLatestRoundByPlayerIDRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoundType,
		&i.Round,
		&i.FibberQuestionID,
		&i.NormalQuestionID,
		&i.RoomID,
		&i.GameStateID,
		&i.SubmitDeadline,
	)
	return i, err
}

const getPlayerByID = `-- name: GetPlayerByID :one
SELECT id, created_at, updated_at, avatar, nickname, is_ready FROM players WHERE id = $1
`

func (q *Queries) GetPlayerByID(ctx context.Context, id uuid.UUID) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.IsReady,
	)
	return i, err
}

const getRandomQuestionByRound = `-- name: GetRandomQuestionByRound :one
SELECT id, created_at, updated_at, game_name, round, enabled, question, language_code, group_id FROM questions WHERE game_name = $1 AND round = $2 AND language_code = $3 AND enabled = TRUE ORDER BY RANDOM() LIMIT 1
`

type GetRandomQuestionByRoundParams struct {
	GameName     string
	Round        string
	LanguageCode string
}

func (q *Queries) GetRandomQuestionByRound(ctx context.Context, arg GetRandomQuestionByRoundParams) (Question, error) {
	row := q.db.QueryRow(ctx, getRandomQuestionByRound, arg.GameName, arg.Round, arg.LanguageCode)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.Round,
		&i.Enabled,
		&i.Question,
		&i.LanguageCode,
		&i.GroupID,
	)
	return i, err
}

const getRandomQuestionInGroup = `-- name: GetRandomQuestionInGroup :one
SELECT q.id, q.created_at, q.updated_at, game_name, round, enabled, question, language_code, group_id, qg.id, qg.created_at, qg.updated_at, group_name, group_type
FROM questions q
JOIN questions_groups qg ON q.group_id = qg.id
WHERE qg.group_type = 'questions'
  AND q.group_id = $1
  AND q.enabled = TRUE
  AND q.id != $2
ORDER BY RANDOM()
LIMIT 1
`

type GetRandomQuestionInGroupParams struct {
	GroupID uuid.UUID
	ID      uuid.UUID
}

type GetRandomQuestionInGroupRow struct {
	ID           uuid.UUID
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
	GameName     string
	Round        string
	Enabled      pgtype.Bool
	Question     string
	LanguageCode string
	GroupID      uuid.UUID
	ID_2         uuid.UUID
	CreatedAt_2  pgtype.Timestamp
	UpdatedAt_2  pgtype.Timestamp
	GroupName    string
	GroupType    string
}

func (q *Queries) GetRandomQuestionInGroup(ctx context.Context, arg GetRandomQuestionInGroupParams) (GetRandomQuestionInGroupRow, error) {
	row := q.db.QueryRow(ctx, getRandomQuestionInGroup, arg.GroupID, arg.ID)
	var i GetRandomQuestionInGroupRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.Round,
		&i.Enabled,
		&i.Question,
		&i.LanguageCode,
		&i.GroupID,
		&i.ID_2,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.GroupName,
		&i.GroupType,
	)
	return i, err
}

const getRoomByCode = `-- name: GetRoomByCode :one
SELECT id, created_at, updated_at, game_name, host_player, room_state, room_code FROM rooms WHERE room_code = $1
`

func (q *Queries) GetRoomByCode(ctx context.Context, roomCode string) (Room, error) {
	row := q.db.QueryRow(ctx, getRoomByCode, roomCode)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.HostPlayer,
		&i.RoomState,
		&i.RoomCode,
	)
	return i, err
}

const getRoomByPlayerID = `-- name: GetRoomByPlayerID :one
SELECT r.id, r.created_at, r.updated_at, r.game_name, r.host_player, r.room_state, r.room_code FROM rooms r JOIN rooms_players rp ON r.id = rp.room_id WHERE rp.player_id = $1
`

func (q *Queries) GetRoomByPlayerID(ctx context.Context, playerID uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoomByPlayerID, playerID)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.HostPlayer,
		&i.RoomState,
		&i.RoomCode,
	)
	return i, err
}

const getVotingState = `-- name: GetVotingState :many
SELECT 
    fir.round AS round,
    q.question,
    gs.submit_deadline,
    p.id AS player_id,
    p.nickname,
    p.avatar,
    COALESCE(COUNT(fv.id), 0) AS votes,
    fia.answer
FROM fibbing_it_rounds fir
JOIN questions q ON fir.fibber_question_id = q.id
JOIN game_state gs ON fir.game_state_id = gs.id
JOIN rooms_players rp ON rp.room_id = fir.room_id
JOIN players p ON p.id = rp.player_id
LEFT JOIN fibbing_it_answers fia ON fia.round_id = fir.id AND fia.player_id = p.id
LEFT JOIN fibbing_it_votes fv ON fv.round_id = fir.id AND fv.voted_for_player_id = p.id
WHERE fir.id = $1
GROUP BY
    fir.round,
    q.question,
    gs.submit_deadline,
    p.id,
    p.nickname,
    p.avatar,
    fia.answer
ORDER BY votes DESC, p.nickname
`

type GetVotingStateRow struct {
	Round          int32
	Question       string
	SubmitDeadline pgtype.Timestamp
	PlayerID       uuid.UUID
	Nickname       string
	Avatar         []byte
	Votes          interface{}
	Answer         pgtype.Text
}

func (q *Queries) GetVotingState(ctx context.Context, id uuid.UUID) ([]GetVotingStateRow, error) {
	rows, err := q.db.Query(ctx, getVotingState, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVotingStateRow
	for rows.Next() {
		var i GetVotingStateRow
		if err := rows.Scan(
			&i.Round,
			&i.Question,
			&i.SubmitDeadline,
			&i.PlayerID,
			&i.Nickname,
			&i.Avatar,
			&i.Votes,
			&i.Answer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removePlayerFromRoom = `-- name: RemovePlayerFromRoom :one
DELETE FROM rooms_players WHERE player_id = $1 RETURNING room_id, player_id, created_at, updated_at
`

func (q *Queries) RemovePlayerFromRoom(ctx context.Context, playerID uuid.UUID) (RoomsPlayer, error) {
	row := q.db.QueryRow(ctx, removePlayerFromRoom, playerID)
	var i RoomsPlayer
	err := row.Scan(
		&i.RoomID,
		&i.PlayerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const toggleAnswerIsReady = `-- name: ToggleAnswerIsReady :one
UPDATE fibbing_it_answers SET is_ready = NOT is_ready WHERE player_id = $1 RETURNING id, created_at, updated_at, answer, player_id, round_id, is_ready
`

func (q *Queries) ToggleAnswerIsReady(ctx context.Context, playerID uuid.UUID) (FibbingItAnswer, error) {
	row := q.db.QueryRow(ctx, toggleAnswerIsReady, playerID)
	var i FibbingItAnswer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Answer,
		&i.PlayerID,
		&i.RoundID,
		&i.IsReady,
	)
	return i, err
}

const togglePlayerIsReady = `-- name: TogglePlayerIsReady :one
UPDATE players SET is_ready = NOT is_ready WHERE id = $1 RETURNING id, created_at, updated_at, avatar, nickname, is_ready
`

func (q *Queries) TogglePlayerIsReady(ctx context.Context, id uuid.UUID) (Player, error) {
	row := q.db.QueryRow(ctx, togglePlayerIsReady, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.IsReady,
	)
	return i, err
}

const updateAvatar = `-- name: UpdateAvatar :one
UPDATE players SET avatar = $1 WHERE id = $2 RETURNING id, created_at, updated_at, avatar, nickname, is_ready
`

type UpdateAvatarParams struct {
	Avatar []byte
	ID     uuid.UUID
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (Player, error) {
	row := q.db.QueryRow(ctx, updateAvatar, arg.Avatar, arg.ID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.IsReady,
	)
	return i, err
}

const updateGameState = `-- name: UpdateGameState :one
UPDATE game_state SET state = $1, submit_deadline = $2 WHERE id = $3 RETURNING id, created_at, updated_at, room_id, submit_deadline, state
`

type UpdateGameStateParams struct {
	State          string
	SubmitDeadline pgtype.Timestamp
	ID             uuid.UUID
}

func (q *Queries) UpdateGameState(ctx context.Context, arg UpdateGameStateParams) (GameState, error) {
	row := q.db.QueryRow(ctx, updateGameState, arg.State, arg.SubmitDeadline, arg.ID)
	var i GameState
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoomID,
		&i.SubmitDeadline,
		&i.State,
	)
	return i, err
}

const updateNickname = `-- name: UpdateNickname :one
UPDATE players SET nickname = $1 WHERE id = $2 RETURNING id, created_at, updated_at, avatar, nickname, is_ready
`

type UpdateNicknameParams struct {
	Nickname string
	ID       uuid.UUID
}

func (q *Queries) UpdateNickname(ctx context.Context, arg UpdateNicknameParams) (Player, error) {
	row := q.db.QueryRow(ctx, updateNickname, arg.Nickname, arg.ID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.IsReady,
	)
	return i, err
}

const updateRoomState = `-- name: UpdateRoomState :one
UPDATE rooms SET room_state = $1 WHERE id = $2 RETURNING id, created_at, updated_at, game_name, host_player, room_state, room_code
`

type UpdateRoomStateParams struct {
	RoomState string
	ID        uuid.UUID
}

func (q *Queries) UpdateRoomState(ctx context.Context, arg UpdateRoomStateParams) (Room, error) {
	row := q.db.QueryRow(ctx, updateRoomState, arg.RoomState, arg.ID)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.HostPlayer,
		&i.RoomState,
		&i.RoomCode,
	)
	return i, err
}

const upsertFibbingItVote = `-- name: UpsertFibbingItVote :exec
INSERT INTO fibbing_it_votes (id, created_at, updated_at, player_id, voted_for_player_id, round_id)
VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $2, $3, $4)
ON CONFLICT(id) DO UPDATE SET
    updated_at = EXCLUDED.updated_at,
    player_id = EXCLUDED.player_id,
    voted_for_player_id = EXCLUDED.voted_for_player_id,
    round_id = EXCLUDED.round
RETURNING id, created_at, updated_at, player_id, voted_for_player_id, round_id
`

type UpsertFibbingItVoteParams struct {
	ID               uuid.UUID
	PlayerID         uuid.UUID
	VotedForPlayerID uuid.UUID
	RoundID          uuid.UUID
}

func (q *Queries) UpsertFibbingItVote(ctx context.Context, arg UpsertFibbingItVoteParams) error {
	_, err := q.db.Exec(ctx, upsertFibbingItVote,
		arg.ID,
		arg.PlayerID,
		arg.VotedForPlayerID,
		arg.RoundID,
	)
	return err
}
