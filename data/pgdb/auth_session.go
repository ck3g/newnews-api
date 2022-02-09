package pgdb

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/ck3g/newnews-api/data"
	"github.com/jackc/pgx/v4"
)

type AuthSessionModel struct {
	DB *pgx.Conn
}

func (m *AuthSessionModel) GenerateForUserID(userID int64) (string, error) {
	models := New(m.DB)
	_, err := models.Users.Find(userID)
	if err != nil {
		return "", data.ErrUserDoesNotExist
	}

	token := generateSecureToken()
	expiredAt := time.Now().Add(time.Hour * 24 * 14)
	query := "INSERT INTO auth_sessions (token, user_id, expired_at) VALUES($1, $2, $3)"
	_, err = m.DB.Exec(context.Background(), query, token, userID, expiredAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateSecureToken() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
