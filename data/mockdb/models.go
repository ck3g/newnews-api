package mockdb

import "github.com/ck3g/newnews-api/data"

func NewMock() data.Models {
	return data.Models{
		Items: &MockItemModel{},
		Users: &MockUserModel{},
	}
}
