package handlers

import "github.com/ck3g/newnews-api/pkg/validator"

func validateUsernameAndPassword(v validator.Validator, username, password string) {
	v.ValidatePresenceOf("username", username)
	v.ValidateLengthOf("username", username, 3, 20)

	v.ValidatePresenceOf("password", password)
	v.ValidateLengthOf("password", password, 6, 128)
}
