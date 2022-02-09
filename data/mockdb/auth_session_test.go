package mockdb

import (
	"testing"

	"github.com/ck3g/newnews-api/data"
)

func TestAuthSession_Authenticate(t *testing.T) {
	um := UserModel{}
	existingID, _ := um.Create("username", "password")
	nonExistingID := int64(-1)
	tests := []struct {
		name      string
		userID    int64
		wantToken string
		wantError error
	}{
		{"when user exists", existingID, "fake-token", nil},
		{"when user does not exist", nonExistingID, "", data.ErrUserDoesNotExist},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authSession := AuthSessionModel{}
			token, err := authSession.Authenticate(tt.userID)
			if token != tt.wantToken {
				t.Errorf("wrong token returned; want %s; got %s", tt.wantToken, token)
			}

			if tt.wantError != nil && err != tt.wantError {
				t.Errorf("wrong error returned; want %v; got %v", tt.wantError, err)
			}
		})
	}
}
