package mockdb

import "github.com/ck3g/newnews-api/data"

type AuthSessionModel struct{}

func (m *AuthSessionModel) Authenticate(id int64) (string, error) {
	um := UserModel{}
	_, err := um.Find(id)
	if err != nil {
		return "", data.ErrUserDoesNotExist
	}

	return "fake-token", nil
}
