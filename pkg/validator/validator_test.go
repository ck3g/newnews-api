package validator

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	v := New()

	if !v.Valid() {
		t.Error("expected validator to be valid when it doesn't contain any errors")
	}

	v.AddError("name", "cannot be blank")

	if v.Valid() {
		t.Error("expected validator to be invalid when it contains errors")
	}
}

func TestErrorMessages(t *testing.T) {
	v := New()
	v.AddError("username", "cannot be blank")
	v.AddError("username", "cannot contain spaces")
	v.AddError("password", "is too short")

	got := v.ErrorMessages()
	want := []map[string][]string{
		{"message": []string{"username cannot be blank", "username cannot contain spaces"}},
		{"message": []string{"password is too short"}},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong error messages returned; want %+v; got %+v", want, got)
	}
}

func TestAddError(t *testing.T) {
	v := New()

	if len(v.Errors) != 0 {
		t.Fatal()
	}

	v.AddError("name", "cannot be blank")

	if len(v.Errors) != 1 {
		t.Errorf("wrong number of errors; want %d; got %d", 1, len(v.Errors))
	}

	if len(v.Errors["name"]) != 1 {
		t.Errorf("wrong number of errors per field; want %d; got %d", 1, len(v.Errors["name"]))
	}

	if v.Errors["name"][0] != "cannot be blank" {
		t.Errorf("wrong error message; want %s; got %s", "cannot be blank", v.Errors["name"][0])
	}

	v.AddError("name", "some other error")

	if len(v.Errors["name"]) != 2 {
		t.Errorf("wrong number of errors per field; want %d; got %d", 2, len(v.Errors["name"]))
	}

	if v.Errors["name"][1] != "some other error" {
		t.Errorf("wrong error message; want %s; got %s", "some other error", v.Errors["name"][1])
	}

	v.AddError("password", "cannot be less than 6 characters")

	if len(v.Errors) != 2 {
		t.Errorf("wrong number of errors; want %d; got %d", 2, len(v.Errors))
	}
}

func TestValidatePresenceOf(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
		valid bool
		msg   string
	}{
		{"valid", "name", "John", true, ""},
		{"empty string", "name", "", false, "cannot be blank"},
		{"string of spaces", "name", "  ", false, "cannot be blank"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.ValidatePresenceOf(tt.field, tt.value)

			assertValidation(t, v, tt.field, tt.valid, tt.msg)
		})
	}
}

func TestValidateLengthOf(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
		min   int
		max   int
		valid bool
		msg   string
	}{
		{"valid", "name", "John", 3, 10, true, ""},
		{"blank", "name", "", 3, 10, true, ""},
		{"too short", "name", "John", 6, 10, false, "is too short (minimum is 6 characters)"},
		{"too long", "name", "John", 1, 3, false, "is too long (maximum is 3 characters)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.ValidateLengthOf(tt.field, tt.value, tt.min, tt.max)

			assertValidation(t, v, tt.field, tt.valid, tt.msg)
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
		valid bool
		msg   string
	}{
		{"valid email", "email", "user@example.com", true, ""},
		{"blank", "email", "", true, ""},
		{"plain string", "email", "user", false, "is invalid"},
		{"no user", "email", "@example", false, "is invalid"},
		{"no domain", "email", "user@", false, "is invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.ValidateEmail(tt.field, tt.value)

			assertValidation(t, v, tt.field, tt.valid, tt.msg)
		})
	}
}

func assertValidation(t *testing.T, v Validator, field string, wantValid bool, wantMsg string) {
	valid := v.Valid()

	if wantValid {
		if !valid {
			t.Error("should be valid, but it's not")
		}
		return
	}

	if len(v.Errors[field]) == 0 {
		t.Errorf("should have at least 1 error")
		return
	}

	msg := v.Errors[field][0]
	if msg != wantMsg {
		t.Errorf("%s should have validation error %s, but has %s", wantMsg, wantMsg, msg)
	}
}
