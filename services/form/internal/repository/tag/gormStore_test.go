package tag

import (
	"testing"

	common "github.com/team-ravl/civic-qa/services/common/model"
	"github.com/team-ravl/civic-qa/services/form/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestEverything(t *testing.T) {
	// shared memory database
	dsn := "file:memdb1?mode=memory&cache=shared"

	db := sqlite.Open(dsn)

	store, err := NewGormStore(db, &gorm.Config{})
	if err != nil {
		t.Fatalf("Error making gorm store: %v", err)
	}

	// pre-create a user
	user := &common.User{Email: "test@example.com", PassHash: []byte{0xa}, FirstName: "Rafi", LastName: "Bayer"}

	err = store.db.Create(user).Error
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// create a form
	form := &model.Form{Name: "test", UserID: user.ID}

	err = store.db.Create(form).Error
	if err != nil {
		t.Fatalf("Failed to create form: %v", err)
	}

	// create a response
	resp := &model.FormResponse{
		Name:         "test",
		EmailAddress: "resp@email.com",
		InquiryType:  "general",
		Subject:      "subject",
		Body:         "body",
		FormID:       form.ID,
	}
	err = store.db.Create(resp).Error
	if err != nil {
		t.Fatalf("Failed to create response: %v", err)
	}

	// create 2 tags
	err = store.Create(user.ID, resp.ID, "tag1")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	err = store.Create(user.ID, resp.ID, "tag2")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	// get all the tags
	all, err := store.GetAll(user.ID)
	if err != nil {
		t.Fatalf("Failed to get all: %v", err)
	}

	if len(all) != 2 {
		t.Fatalf("Expected 2 tags, got: %d", len(all))
	}

	// ensure we only get tags for the correct user
	all, err = store.GetAll(1234)
	if err != nil {
		t.Fatalf("Failed to get all: %v", err)
	}

	if len(all) != 0 {
		t.Fatalf("Expected 0 tags, got: %d", len(all))
	}

	// create another response
	resp2 := &model.FormResponse{
		Name:         "test2",
		EmailAddress: "resp2@email.com",
		InquiryType:  "general",
		Subject:      "subject2",
		Body:         "body2",
		FormID:       form.ID,
	}
	err = store.db.Create(resp2).Error
	if err != nil {
		t.Fatalf("Failed to create response: %v", err)
	}

	// create a tag
	err = store.Create(user.ID, resp2.ID, "tag3")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	// ensure we only get tags for the correct response
	byID, err := store.GetByResponseID(user.ID, resp2.ID)
	if err != nil {
		t.Fatalf("Failed to GetByResponseID: %v", err)
	}

	if len(byID) != 1 {
		t.Fatalf("Expected 1 tag, got: %d", len(all))
	}

	all, err = store.GetAll(user.ID)
	if err != nil {
		t.Fatalf("Failed to get all: %v", err)
	}

	// delete all the tags
	for _, tag := range all {
		err = store.Delete(user.ID, tag.FormResponseID, tag.Value)
		if err != nil {
			t.Fatalf("Failed to delete: %v", err)
		}
	}

	// ensure all tags were deleted
	all, err = store.GetAll(user.ID)
	if err != nil {
		t.Fatalf("Failed to get all: %v", err)
	}

	if len(all) != 0 {
		t.Fatalf("Expected 0 tags, got: %d", len(all))
	}

}
