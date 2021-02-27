package context

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
	"github.com/vivian-hua/civic-qa/service/account/internal/repository/session"
	"github.com/vivian-hua/civic-qa/services/common/config"
	common "github.com/vivian-hua/civic-qa/services/common/model"
)

//===============================
//|AUTH						    |
//===============================

func TestHashPassword(t *testing.T) {
	cases := []string{
		"testpassword",
		"p4ssw0rd! ",
		"        ",
		"",
		"123456789|123456789|123456789|123456789|123456789|",
	}

	for i, testCase := range cases {
		_, err := hashPassword(testCase)
		if err != nil {
			t.Fatalf("Case %d: %v", i, err)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	cases := []struct {
		user       *common.User
		login      model.LoginRequest
		shouldPass bool
	}{
		{&common.User{PassHash: mustHash("password")}, model.LoginRequest{Password: "password"}, true},
		{&common.User{PassHash: mustHash("password")}, model.LoginRequest{Password: "wrong"}, false},
		{&common.User{PassHash: mustHash("")}, model.LoginRequest{Password: ""}, true},
	}

	for i, testCase := range cases {
		err := checkPassword(testCase.user, testCase.login)
		if err != nil && testCase.shouldPass {
			t.Fatalf("Case %d Unexpected error: %v", i, err)
		} else if err == nil && !testCase.shouldPass {
			t.Fatalf("Case %d Unexpected pass", i)
		}
	}
}

func mustHash(password string) []byte {
	hash, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	return hash
}

func TestAddAuthorizationHeader(t *testing.T) {
	cases := []struct {
		w        http.ResponseWriter
		token    session.Token
		expected string
	}{
		{httptest.NewRecorder(), session.Token("12345"), "Bearer 12345"},
		{httptest.NewRecorder(), session.Token("123456789|123456789|123456789|123456789|123456789|"), "Bearer 123456789|123456789|123456789|123456789|123456789|"},
	}

	for i, testCase := range cases {
		addAuthorizationHeader(testCase.w, testCase.token)
		if got := testCase.w.Header().Get("Authorization"); got != testCase.expected {
			t.Fatalf("Case %d unexpected header: expected %s, got %s", i, testCase.expected, got)
		}
	}
}

func TestGetAuthorizationToken(t *testing.T) {
	cases := []struct {
		r          *http.Request
		expected   session.Token
		shouldPass bool
	}{
		{testReqWHeaders(map[string]string{"Authorization": "Bearer 12345"}), session.Token("12345"), true},
		{testReqWHeaders(map[string]string{"Authorization": "Bearer mytoken"}), session.Token("mytoken"), true},
		{testReqWHeaders(map[string]string{"Authorization": "Bearer "}), session.Token(""), true},
		{testReqWHeaders(map[string]string{"Authorization": "Bearer"}), session.InvalidSessionToken, false},
		{testReqWHeaders(map[string]string{"Authorization": "Plain token"}), session.InvalidSessionToken, false},
		{testReqWHeaders(nil), session.InvalidSessionToken, false},
		{testReqWHeaders(map[string]string{"Auth": "bad"}), session.InvalidSessionToken, false},
	}

	for i, testCase := range cases {
		token, err := getAuthorizationToken(testCase.r)
		if err != nil && testCase.shouldPass {
			t.Fatalf("Case %d Unexpected error: %v", i, err)
		} else if err == nil && !testCase.shouldPass {
			t.Fatalf("Case %d Unexpected pass", i)
		}

		if err == nil && token != testCase.expected {
			t.Fatalf("Case %d Unexpected token: got %s, expected %s", i, token, testCase.expected)
		}
	}
}

func testReqWHeaders(headers map[string]string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/test/url", nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}

//===============================
//|CONTEXT					    |
//===============================

func TestBuildContext(t *testing.T) {
	cases := []struct {
		cfg        config.Provider
		shouldPass bool
	}{
		{&config.MapProvider{Data: map[string]string{"DB_DSN": ":memory:"}}, true},
		{&config.MapProvider{Data: map[string]string{"DB_IMPL": "sqlite", "DB_DSN": ":memory:"}}, true},
		{&config.MapProvider{Data: map[string]string{"SESS_IMPL": "memory", "DB_DSN": ":memory:"}}, true},
		{&config.MapProvider{Data: map[string]string{"SESS_IMPL": "** unknown **", "DB_DSN": ":memory:"}}, false},
		{&config.MapProvider{Data: map[string]string{"DB_IMPL": "** unknown **", "DB_DSN": ":memory:"}}, false},
	}

	for i, testCase := range cases {
		_, err := BuildContext(testCase.cfg)
		if err != nil && testCase.shouldPass {
			t.Fatalf("Case %d Unexpected error: %v", i, err)
		} else if err == nil && !testCase.shouldPass {
			t.Fatalf("Case %d Unexpected pass", i)
		}
	}
}

//===============================
//|HANDLERS					    |
//===============================

func TestHandleSignup(t *testing.T) {
	cases := []struct {
		r              *http.Request
		headers        map[string]string
		w              *httptest.ResponseRecorder
		expectedStatus int
	}{
		{
			// Valid user signup
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email@example.com",
				Password:        "Password!",
				PasswordConfirm: "Password!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusCreated,
		},
		{
			// email in use
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email@example.com",
				Password:        "Password!",
				PasswordConfirm: "Password!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusConflict,
		},
		{
			// Invalid email
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "bad-email.com",
				Password:        "Password!",
				PasswordConfirm: "Password!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			// Bad password
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email1@example.com",
				Password:        "short",
				PasswordConfirm: "short",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			// non-matching passwords
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email2@example.com",
				Password:        "ValidPassword!",
				PasswordConfirm: "DifferentSecret!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			// Bad Method
			r: httptest.NewRequest("GET", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email3@example.com",
				Password:        "Password!",
				PasswordConfirm: "Password!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			// Bad content-type
			r: httptest.NewRequest("POST", "/signup", newUserReader(&model.NewUserRequest{
				Email:           "email4@example.com",
				Password:        "Password!",
				PasswordConfirm: "Password!",
				FirstName:       "testfname",
				LastName:        "testlname",
			})),
			headers:        map[string]string{"content-type": "something/else"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusUnsupportedMediaType,
		},
	}

	cfg := &config.MapProvider{Data: map[string]string{"DB_DSN": ":memory:"}}
	ctx, err := BuildContext(cfg)
	if err != nil {
		t.Fatalf("Failed to build handler context: %v", err)
	}

	for i, testCase := range cases {
		// add request headers
		for k, v := range testCase.headers {
			testCase.r.Header.Set(k, v)
		}
		// test the handler
		ctx.HandleSignup(testCase.w, testCase.r)
		if status := testCase.w.Result().StatusCode; status != testCase.expectedStatus {
			t.Fatalf("Case %d Unexpected status code: got %d, expected %d", i, status, testCase.expectedStatus)
		}
	}
}

// newUserReader returns a io.Reader over a NewUserRequest,
// panics on failure
func newUserReader(newUser *model.NewUserRequest) io.Reader {
	bodyBytes, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(bodyBytes)
}

func TestHandleLogin(t *testing.T) {
	preCreatedUsers := []*model.NewUserRequest{
		{
			Email:           "email@example.com",
			Password:        "validpassword",
			PasswordConfirm: "validpassword",
			FirstName:       "firstname",
			LastName:        "lastname",
		},
		{
			Email:           "another@example.com",
			Password:        "validpassword",
			PasswordConfirm: "validpassword",
			FirstName:       "firstname",
			LastName:        "lastname",
		},
	}

	cases := []struct {
		r              *http.Request
		headers        map[string]string
		w              *httptest.ResponseRecorder
		expectedStatus int
	}{
		{
			// valid login for first pre-created
			r: httptest.NewRequest("POST", "/login", loginReader(&model.LoginRequest{
				Email:    "email@example.com",
				Password: "validpassword",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
		},
		{
			// valid login for second pre-created user
			r: httptest.NewRequest("POST", "/login", loginReader(&model.LoginRequest{
				Email:    "another@example.com",
				Password: "validpassword",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
		},
		{
			// invalid email
			r: httptest.NewRequest("POST", "/login", loginReader(&model.LoginRequest{
				Email:    "invalid@example.com",
				Password: "validpassword",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			// invalid password
			r: httptest.NewRequest("POST", "/login", loginReader(&model.LoginRequest{
				Email:    "another@example.com",
				Password: "bad-password",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			// invalid method
			r: httptest.NewRequest("GET", "/login", loginReader(&model.LoginRequest{
				Email:    "another@example.com",
				Password: "validpassword",
			})),
			headers:        map[string]string{"content-type": "application/json"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			// Invalid content type
			r: httptest.NewRequest("POST", "/login", loginReader(&model.LoginRequest{
				Email:    "email@example.com",
				Password: "validpassword",
			})),
			headers:        map[string]string{"content-type": "something/else"},
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusUnsupportedMediaType,
		},
	}

	// create handler context
	cfg := &config.MapProvider{Data: map[string]string{"DB_DSN": ":memory:"}}
	ctx, err := BuildContext(cfg)
	if err != nil {
		t.Fatalf("Failed to build handler context: %v", err)
	}

	// precreate users
	for i, newUser := range preCreatedUsers {
		r := httptest.NewRequest("POST", "/signup", newUserReader(newUser))
		r.Header.Add("content-type", "application/json")
		w := httptest.NewRecorder()
		ctx.HandleSignup(w, r)
		if w.Result().StatusCode != http.StatusCreated {
			t.Fatalf("Failed to pre-create user %d: status %d", i, w.Result().StatusCode)
		}
	}

	// test logins
	for i, testCase := range cases {
		// add request headers
		for k, v := range testCase.headers {
			testCase.r.Header.Set(k, v)
		}
		// test the handler
		ctx.HandleLogin(testCase.w, testCase.r)
		if status := testCase.w.Result().StatusCode; status != testCase.expectedStatus {
			t.Fatalf("Case %d Unexpected status code: got %d, expected %d", i, status, testCase.expectedStatus)
		}
	}
}

func loginReader(login *model.LoginRequest) io.Reader {
	bodyBytes, err := json.Marshal(login)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(bodyBytes)
}

//===============================
//|VALIDATION				    |
//===============================

func TestValidateNewUser(t *testing.T) {
	cases := []struct {
		newUser    model.NewUserRequest
		shouldPass bool
	}{
		{model.NewUserRequest{}, false},
		{model.NewUserRequest{Email: "mail"}, false},
		{model.NewUserRequest{Email: "email@mail.com"}, false},
		{model.NewUserRequest{Email: "email@mail.com", Password: "A", PasswordConfirm: "A"}, false},
		{model.NewUserRequest{Email: "email@mail.com", Password: "abcdefgh", PasswordConfirm: "123456678"}, false},
		{model.NewUserRequest{Email: "email@mail.com", Password: "Password", PasswordConfirm: "Password"}, true},
	}

	for i, testCase := range cases {
		err := validateNewUser(testCase.newUser)
		if err != nil && testCase.shouldPass {
			t.Fatalf("Case %d Unexpected error: %v", i, err)
		} else if err == nil && !testCase.shouldPass {
			t.Fatalf("Case %d Unexpected pass", i)
		}
	}
}
