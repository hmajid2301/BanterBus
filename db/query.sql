-- name: AddRoom :one
INSERT INTO rooms (id, game_name, host_player, room_code, room_state)  VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: AddPlayer :one
INSERT INTO players (id, avatar, nickname) VALUES (?, ?, ?) RETURNING *;

-- name: AddRoomPlayer :one
INSERT INTO rooms_players (room_id, player_id) VALUES (?, ?) RETURNING *;

-- name: RemovePlayerFromRoom :one
UPDATE rooms_players SET room_id = "" WHERE player_id = ? RETURNING *;

-- name: GetAllPlayersInRoom :many
SELECT p.id, p.created_at, p.updated_at, p.avatar, p.nickname, p.is_ready, r.room_code, r.host_player
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
WHERE rp.room_id = (
    SELECT rp_inner.room_id
    FROM rooms_players rp_inner
    WHERE rp_inner.player_id = ?
);

-- name: GetAllPlayerByRoomCode :many
SELECT p.id, p.created_at, p.updated_at, p.avatar, p.nickname, p.is_ready, r.room_code, r.host_player
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
WHERE rp.room_id = (
    SELECT r_inner.id
    FROM rooms r_inner
    WHERE r_inner.room_code = ? AND (r_inner.room_state = "CREATED" OR r_inner.room_state = "PLAYING")
)
ORDER BY p.created_at;

-- name: GetRoomByPlayerID :one
SELECT r.* FROM rooms r JOIN rooms_players rp ON r.id = rp.room_id WHERE rp.player_id = ?;

-- name: GetRoomByCode :one
SELECT * FROM rooms WHERE room_code = ?;

-- name: UpdateRoomState :one
UPDATE rooms SET room_state = ? WHERE id = ? RETURNING *;

-- name: GetPlayerByID :one
SELECT * FROM players WHERE id = ?;

-- name: UpdateNickname :one
UPDATE players SET nickname = ? WHERE id = ? RETURNING *;

-- name: UpdateAvatar :one
UPDATE players SET avatar = ? WHERE id = ? RETURNING *;

-- name: UpdateIsReady :one
UPDATE players SET is_ready = ? WHERE id = ? RETURNING *;

-- name: AddFibbingItRound :one
INSERT INTO fibbing_it_rounds (id, round_type, round, fibber_question_id, normal_question_id, room_id, game_state_id) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: AddGameState :one
INSERT INTO game_state (id, room_id, submit_deadline, state) VALUES (?, ?, ?, ?) RETURNING *;

-- name: AddFibbingItAnswer :one
INSERT INTO fibbing_it_answers (id, answer, round_id, player_id) VALUES (?, ?, ?, ?) RETURNING *;

-- name: AddFibbingItRole :one
INSERT INTO fibbing_it_player_roles (id, player_role, round_id, player_id) VALUES (?, ?, ?, ?) RETURNING *;

-- name: AddFibbingItVoting :one
INSERT INTO fibbing_it_voting (id, player_id, round_id, votes) VALUES (?, ?, ?, ?) RETURNING *;

-- name: AddQuestion :one
INSERT INTO questions (id, game_name, round, question, language_code, group_id) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: AddQuestionsGroup :one
INSERT INTO questions_groups (id, group_name, group_type) VALUES (?, ?, ?) RETURNING *;

-- name: GetRandomQuestionByRound :one
SELECT * FROM questions WHERE game_name = ? AND round = ? AND language_code = ? AND enabled = TRUE ORDER BY RANDOM() LIMIT 1;

-- name: GetRandomQuestionInGroup :one
SELECT *
FROM questions q
JOIN questions_groups qg ON q.group_id = qg.id
WHERE qg.group_type = 'questions'
  AND q.group_id = ?
  AND q.enabled = TRUE
  AND q.id != ?
ORDER BY RANDOM()
LIMIT 1;

-- name: GetLatestRoundByPlayerID :one
SELECT fir.*, gs.submit_deadline
FROM fibbing_it_rounds fir
JOIN rooms_players rp ON fir.room_id = rp.room_id
JOIN game_state gs ON fir.room_id = gs.room_id
WHERE rp.player_id = ?
ORDER BY fir.created_at DESC
LIMIT 1;

-- name: GetGameStateByPlayerID :one
SELECT
    p.id,
    p.nickname,
    fr.round,
    fr.round_type,
    r.room_code,
    fpr.player_role,
    fq1.question AS fibber_question,
    fq2.question AS normal_question,
    p.avatar
FROM players p
JOIN rooms_players rp ON p.id = rp.player_id
JOIN rooms r ON rp.room_id = r.id
JOIN fibbing_it_rounds fr ON r.id = fr.room_id
LEFT JOIN questions fq1 ON fr.fibber_question_id = fq1.id
LEFT JOIN questions fq2 ON fr.normal_question_id = fq2.id
LEFT JOIN fibbing_it_player_roles fpr ON p.id = fpr.player_id AND fr.id = fpr.round_id
WHERE rp.room_id = (
    SELECT rp_inner.room_id
    FROM rooms_players rp_inner
    WHERE rp_inner.player_id = ?
)
ORDER BY p.created_at;
