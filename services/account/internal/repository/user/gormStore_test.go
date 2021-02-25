package user

import (
	"testing"
	"time"

	"github.com/vivian-hua/civic-qa/services/common/model"
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
			PassHash:  []byte{0xa, 0xb},
			FirstName: "fname",
			LastName:  "lname",
			CreatedOn: time.Now(),
		}, false}, // None null passes
		{model.User{
			Email:     "test@mail.com",
			PassHash:  []byte{0xa, 0xb},
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
