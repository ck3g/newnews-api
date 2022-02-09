package mockdb

import "github.com/ck3g/newnews-api/data"

func New() data.Models {
	return data.Models{
		Items:        &ItemModel{},
		Users:        &UserModel{},
		AuthSessions: &AuthSessionModel{},
	}
}
