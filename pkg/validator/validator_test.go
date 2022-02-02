package validator

import (
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
