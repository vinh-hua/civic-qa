package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
	common "github.com/vivian-hua/civic-qa/services/common/model"
)

func TestSessionStore(t *testing.T) {
	cases := []model.SessionState{
		{},
		{User: common.User{}, CreatedAt: time.Now()},
		{User: common.User{ID: 1, Email: "test@mail.com"}, CreatedAt: time.Now()},
		{User: common.User{
			ID:        1,
			Email:     "test@mail.com",
			PassHash:  []byte{},
			FirstName: "A",
			LastName:  "B",
		}, CreatedAt: time.Now()},
	}

	store := NewMemorySessionStore(0)

	for _, state := range cases {
		token, err := store.Create(state)
		if err != nil {
			t.Fatal(err)
		}

		stateOut, err := store.Get(token)
		if err != nil {
			t.Fatal(err)
		}

		// not sure why we need reflect here but can't compile otherwise
		if !reflect.DeepEqual(state, *stateOut) {
			t.Fatalf("State did not match: got %v, expected %v", stateOut, state)
		}
	}
}
