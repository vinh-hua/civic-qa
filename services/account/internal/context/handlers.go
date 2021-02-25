package context

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/vivian-hua/civic-qa/services/common/helpers"
	common "github.com/vivian-hua/civic-qa/services/common/model"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
	"github.com/vivian-hua/civic-qa/service/account/internal/repository/user"
)

// HandleSignup handles a post request to create a new user
// if successful, the response will have the new users Authorization header
func (ctx *Context) HandleSignup(w http.ResponseWriter, r *http.Request) {
	// verify method and content-type
	status, err := helpers.ExpectMethodAndContentType(r, http.MethodPost, "application/json")
	if err != nil {
		http.Error(w, err.Error(), status)
	}

	// parse request body
	var newUser model.NewUserRequest
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Malformed Request Body", http.StatusBadRequest)
		return
	}

	// validate request
	err = validateNewUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if email is already in use
	_, err = ctx.UserStore.GetByEmail(newUser.Email)
	if err != user.ErrUserNotFound {
		// if we had no error at all, email is in use
		if err == nil {
			http.Error(w, "Email already in use", http.StatusConflict)
			return
		}
		// otherwise, unknow error
		log.Printf("Could not GetByEmail: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Hash the users password
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Printf("Failed to hash users password, error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// Create the User and insert into UserStore
	createdUser := common.User{
		Email:     newUser.Email,
		PassHash:  hashedPassword,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		CreatedOn: time.Now(),
	}
	err = ctx.UserStore.Create(&createdUser)
	if err != nil {
		log.Printf("Could not create user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Begin a session and return
	token, err := ctx.SessionStore.Create(model.SessionState{
		User:      createdUser,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Printf("Could not create session: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add the authorization header
	addAuthorizationHeader(w, token)

	// Send response
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Account Created!")
}

// HandleLogin handles a post request to log in
// if successful, the response will have the new users Authorization header
func (ctx *Context) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// verify method and content-type
	status, err := helpers.ExpectMethodAndContentType(r, http.MethodPost, "application/json")
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	// parse request body
	var loginRequest model.LoginRequest
	err = json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Malformed Request Body", http.StatusBadRequest)
		return
	}

	// retrieve the user from the database
	reqUser, err := ctx.UserStore.GetByEmail(loginRequest.Email)
	if err != nil {
		if err == user.ErrUserNotFound {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		log.Printf("Could not GetByEmail: %v", err)
		return
	}

	// check that the password is correct
	err = checkPassword(reqUser, loginRequest)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	// Begin a session and return
	token, err := ctx.SessionStore.Create(model.SessionState{
		User:      *reqUser,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Printf("Could not create session: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add the authorization header
	addAuthorizationHeader(w, token)

	// Send response
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Logged In!")

}

func (ctx *Context) HandleLogout(w http.ResponseWriter, r *http.Request) {

}

func (ctx *Context) HandleGetSession(w http.ResponseWriter, r *http.Request) {

}
