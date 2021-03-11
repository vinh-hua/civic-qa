package context

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/team-ravl/civic-qa/services/common/helpers"
	common "github.com/team-ravl/civic-qa/services/common/model"

	"github.com/team-ravl/civic-qa/service/account/internal/model"
	"github.com/team-ravl/civic-qa/service/account/internal/repository/session"
	"github.com/team-ravl/civic-qa/service/account/internal/repository/user"
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

	// check if the email is already being used
	inUse, err := ctx.UserStore.EmailInUse(newUser.Email)
	if err != nil {
		log.Printf("Could not query EmailInUse: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if inUse {
		http.Error(w, "Email already in use", http.StatusConflict)
		return
	}

	// Hash the users password
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Printf("Failed to hash users password, error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
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
	token, err := ctx.SessionStore.Create(common.SessionState{
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
	token, err := ctx.SessionStore.Create(common.SessionState{
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
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Logged In!")
}

// HandleLogout handles a delete post to logout
func (ctx *Context) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// check method, we don't care about content-type
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// get the authorization token from the request header
	token, err := getAuthorizationToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// attempt to delete the session
	err = ctx.SessionStore.Delete(token)
	if err != nil {
		if err == session.ErrStateNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Printf("Error deleting session: %v", err)
		http.Error(w, "Internal Server Errror", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Logged Out!")
}

// HandleGetSession handles a get request to retrieve a users SessionState
func (ctx *Context) HandleGetSession(w http.ResponseWriter, r *http.Request) {
	// check method, we don't care about content-type
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// get the authorization token from the request header
	token, err := getAuthorizationToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// retrieve the session state
	state, err := ctx.SessionStore.Get(token)
	if err != nil {
		if err == session.ErrStateNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Printf("Error getting session: %v", err)
		http.Error(w, "Internal Server Errror", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(state)
	if err != nil {
		log.Printf("Failed to encode user, error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
