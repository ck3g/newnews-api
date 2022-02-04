package validator

import (
	"fmt"
	"regexp"
	"strings"
)

// EmailRX parses and compiles an Email regular expression
//
// Doing this once at runtime, and storing the compiled regular expression
// object in a variable, is more performant than re-compiling the pattern with
// every request.
//
// The pattern is recommended by the W3C and Web Hypertext Application Technology Working Group.
// https://html.spec.whatwg.org/multipage/forms.html#valid-e-mail-address
// https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
var emailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	Errors map[string][]string
}

func New() Validator {
	return Validator{
		Errors: make(map[string][]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(field, msg string) {
	v.Errors[field] = append(v.Errors[field], msg)
}

func (v *Validator) ErrorMessages() []map[string][]string {
	messages := []map[string][]string{}

	for field, errors := range v.Errors {
		msg := make(map[string][]string)
		for _, err := range errors {
			msg["message"] = append(msg["message"], fmt.Sprintf("%s %s", field, err))
		}
		messages = append(messages, msg)
	}

	return messages
}

func (v *Validator) ValidatePresenceOf(field, value string) {
	if strings.Trim(value, " ") == "" {
		v.AddError(field, "cannot be blank")
	}
}

func (v *Validator) ValidateLengthOf(field, value string, min, max int) {
	if strings.Trim(value, " ") == "" {
		return
	}

	if len(value) < min {
		v.AddError(field, fmt.Sprintf("is too short (minimum is %d characters)", min))
	}

	if len(value) > max {
		v.AddError(field, fmt.Sprintf("is too long (maximum is %d characters)", max))
	}
}

func (v *Validator) ValidateEmail(field, value string) {
	if strings.Trim(value, " ") == "" {
		return
	}

	if !emailRX.MatchString(value) {
		v.AddError(field, "is invalid")
	}
}
