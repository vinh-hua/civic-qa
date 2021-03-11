package session

import (
	"reflect"
	"testing"
	"time"

	"github.com/team-ravl/civic-qa/services/common/model"
)

func TestMemoryStore(t *testing.T) {
	cases := []model.SessionState{
		{},
		{User: model.User{}, CreatedAt: time.Now()},
		{User: model.User{ID: 1, Email: "test@mail.com"}, CreatedAt: time.Now()},
		{User: model.User{
			ID:        1,
			Email:     "test@mail.com",
			PassHash:  []byte{},
			FirstName: "A",
			LastName:  "B",
		}, CreatedAt: time.Now()},
	}

	store := NewMemoryStore()

	for _, state := range cases {
		// Test create
		token, err := store.Create(state)
		if err != nil {
			t.Fatal(err)
		}

		// Test Get
		stateOut, err := store.Get(token)
		if err != nil {
			t.Fatal(err)
		}

		// must use DeepEqual as state.User.PassHash is a byte slice,
		// which doesn't implement operator ==
		if !reflect.DeepEqual(state, *stateOut) {
			t.Fatalf("State did not match: got %v, expected %v", stateOut, state)
		}

		// Test delete
		err = store.Delete(token)
		if err != nil {
			t.Fatal(err)
		}

		// Make sure state was really deleted
		_, err = store.Get(token)
		if err != ErrStateNotFound {
			t.Fatalf("State still found after delete for token: %s", token)
		}
	}
}
