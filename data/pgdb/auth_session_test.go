package pgdb

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ck3g/newnews-api/data"
)

func TestAuthSession_GenerateForUserID(t *testing.T) {
	db := newTestDB(t)

	model := AuthSessionModel{DB: db}

	var userID int64

	_, err := model.GenerateForUserID(userID)
	if !errors.Is(err, data.ErrUserDoesNotExist) {
		t.Error("expected an error when user does not exist, but got nothing")
	}

	userModel := UserModel{DB: db}
	userID, _ = userModel.Create("user", "password")

	token, err := model.GenerateForUserID(userID)
	if err != nil {
		t.Errorf("expected no error when user exists, but got one %v", err)
	}

	if len(token) != 32 {
		t.Errorf("returned a token of wrong length; want 32; got %d", len(token))
	}

	query := "SELECT token, expired_at FROM auth_sessions WHERE user_id = $1"
	var persistedToken string
	var expiredAt time.Time
	err = db.QueryRow(context.Background(), query, userID).Scan(
		&persistedToken, &expiredAt,
	)
	if err != nil {
		t.Fatal("expected auth session to be peristed, but it doesn't")
	}

	if persistedToken != token {
		t.Errorf("wrong persisted token; want %s; got %s", token, persistedToken)
	}

	twoWeeksFromNow := time.Now().Add(time.Hour * 24 * 14)
	minuteAfter := twoWeeksFromNow.Add(1 * time.Minute)
	minuteBefore := twoWeeksFromNow.Add(-1 * time.Minute)
	between := expiredAt.After(minuteBefore) && expiredAt.Before(minuteAfter)
	if !between {
		t.Errorf("expected token to expire in 14 days")
	}
}
