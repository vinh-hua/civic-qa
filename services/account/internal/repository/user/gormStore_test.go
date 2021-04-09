package user

import (
	"reflect"
	"testing"
	"time"

	"github.com/team-ravl/civic-qa/services/common/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	store, err := NewGormStore(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		user        model.User
		errExpected bool
	}{
		{model.User{}, true},                        // All null fails
		{model.User{Email: "email@mail.com"}, true}, // Some null fails
		{model.User{
			Email:     "test@mail.com",
			PassHash:  []byte("hashed"),
			FirstName: "fname",
			LastName:  "lname",
			CreatedOn: time.Now(),
		}, false}, // None null passes
		{model.User{
			Email:     "test@mail.com",
			PassHash:  []byte("hashed"),
			FirstName: "namename",
			LastName:  "thename",
			CreatedOn: time.Now(),
		}, true}, // duplicate email fails

	}

	for i, testCase := range cases {
		err = store.Create(&testCase.user)
		if err != nil && !testCase.errExpected {
			t.Fatalf("(%d)Unexpected error: %v", i, err)
		} else if err == nil && testCase.errExpected {
			t.Fatalf("(%d) Expected error but got none", i)
		}
	}
}

func TestAll(t *testing.T) {
	store, err := NewGormStore(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a user
	u := &model.User{
		Email:     "test@example.com",
		PassHash:  []byte("hashed"),
		FirstName: "testfname",
		LastName:  "testlname",
	}

	err = store.Create(u)
	if err != nil {
		t.Fatal(err)
	}

	// get by id
	out, err := store.GetByID(u.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(u, out) {
		t.Fatalf("User did not match after creation, before: %v, after: %v", u, out)
	}

	// get some non-existant user
	_, err = store.GetByID(out.ID + 1)
	if err != ErrUserNotFound {
		t.Fatalf("Expected ErrUserNotFound, got: %v", err)
	}

	// get by email
	out3, err := store.GetByEmail(u.Email)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(u, out3) {
		t.Fatalf("User did not match after creation, before: %v, after: %v", u, out)
	}

	// get some non-existant user
	_, err = store.GetByEmail("fakeemail@fakemail.com")
	if err != ErrUserNotFound {
		t.Fatalf("Expected ErrUserNotFound, got: %v", err)
	}

	// email in use
	inUse, err := store.EmailInUse(u.Email)
	if err != nil {
		t.Fatal(err)
	}

	if !inUse {
		t.Fatal("Expected email in use, but inUse = false")
	}

	inUse2, err := store.EmailInUse("fakemail@fakemail.com")
	if err != nil {
		t.Fatal(err)
	}

	if inUse2 {
		t.Fatal("Expected email not in use, but inUse = true")
	}
}
