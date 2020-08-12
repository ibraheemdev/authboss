package authboss

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthBossInit(t *testing.T) {
	t.Parallel()

	ab := New()
	err := ab.Init()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}

func TestAuthbossUpdatePassword(t *testing.T) {
	t.Parallel()

	user := &mockUser{}
	storer := newMockServerStorer()

	ab := New()
	ab.Config.Storage.Server = storer

	if err := ab.UpdatePassword(context.Background(), user, "hello world"); err != nil {
		t.Error(err)
	}

	if len(user.Password) == 0 {
		t.Error("password was not updated")
	}
}

type testRedirector struct {
	Opts RedirectOptions
}

func (r *testRedirector) Redirect(w http.ResponseWriter, req *http.Request, ro RedirectOptions) error {
	r.Opts = ro
	if len(ro.RedirectPath) == 0 {
		panic("no redirect path on redirect call")
	}
	http.Redirect(w, req, ro.RedirectPath, ro.Code)
	return nil
}

func TestAuthbossMiddleware(t *testing.T) {
	t.Parallel()

	ab := New()
	ab.Core.Logger = mockLogger{}
	ab.Storage.Server = &mockServerStorer{
		Users: map[string]*mockUser{
			"test@test.com": {},
		},
	}

	setupMore := func(mountPathed bool, failResponse MWRespondOnFailure, requirements MWRequirements) (*httptest.ResponseRecorder, bool, bool) {
		r := httptest.NewRequest("GET", "/super/secret", nil)
		rec := httptest.NewRecorder()
		w := ab.NewResponse(rec)

		var err error
		r, err = ab.LoadClientState(w, r)
		if err != nil {
			t.Fatal(err)
		}

		var mid func(http.Handler) http.Handler
		if !mountPathed {
			mid = Middleware(ab, requirements, failResponse)
		} else {
			mid = MountedMiddleware(ab, true, requirements, failResponse)
		}
		var called, hadUser bool
		server := mid(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			hadUser = r.Context().Value(CTXKeyUser) != nil
			w.WriteHeader(http.StatusOK)
		}))

		server.ServeHTTP(w, r)

		return rec, called, hadUser
	}

	t.Run("Accept", func(t *testing.T) {
		ab.Storage.SessionState = mockClientStateReadWriter{
			state: mockClientState{SessionKey: "test@test.com"},
		}

		var reqs MWRequirements
		_, called, hadUser := setupMore(false, RespondNotFound, reqs)

		if !called {
			t.Error("should have been called")
		}
		if !hadUser {
			t.Error("should have had user")
		}
	})
	t.Run("AcceptHalfAuth", func(t *testing.T) {
		ab.Storage.SessionState = mockClientStateReadWriter{
			state: mockClientState{SessionKey: "test@test.com", SessionHalfAuthKey: "true"},
		}

		var reqs MWRequirements
		_, called, hadUser := setupMore(false, RespondNotFound, reqs)

		if !called {
			t.Error("should have been called")
		}
		if !hadUser {
			t.Error("should have had user")
		}
	})
	t.Run("Reject404", func(t *testing.T) {
		ab.Storage.SessionState = mockClientStateReadWriter{}

		var reqs MWRequirements
		rec, called, hadUser := setupMore(false, RespondNotFound, reqs)

		if rec.Code != http.StatusNotFound {
			t.Error("wrong code:", rec.Code)
		}
		if called {
			t.Error("should not have been called")
		}
		if hadUser {
			t.Error("should not have had user")
		}
	})
	t.Run("Reject401", func(t *testing.T) {
		ab.Storage.SessionState = mockClientStateReadWriter{}

		r := httptest.NewRequest("GET", "/super/secret", nil)
		rec := httptest.NewRecorder()
		w := ab.NewResponse(rec)

		var err error
		r, err = ab.LoadClientState(w, r)
		if err != nil {
			t.Fatal(err)
		}

		var mid func(http.Handler) http.Handler
		mid = Middleware(ab, RequireNone, RespondUnauthorized)
		var called, hadUser bool
		server := mid(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			hadUser = r.Context().Value(CTXKeyUser) != nil
			w.WriteHeader(http.StatusOK)
		}))

		server.ServeHTTP(w, r)

		if rec.Code != http.StatusUnauthorized {
			t.Error("wrong code:", rec.Code)
		}
		if called {
			t.Error("should not have been called")
		}
		if hadUser {
			t.Error("should not have had user")
		}
	})
	t.Run("RejectRedirect", func(t *testing.T) {
		redir := &testRedirector{}
		ab.Config.Core.Redirector = redir

		ab.Storage.SessionState = mockClientStateReadWriter{}

		var reqs MWRequirements
		_, called, hadUser := setupMore(false, RespondRedirect, reqs)

		if redir.Opts.Code != http.StatusTemporaryRedirect {
			t.Error("code was wrong:", redir.Opts.Code)
		}
		if redir.Opts.RedirectPath != "/auth/login?redir=%2Fsuper%2Fsecret" {
			t.Error("redirect path was wrong:", redir.Opts.RedirectPath)
		}
		if called {
			t.Error("should not have been called")
		}
		if hadUser {
			t.Error("should not have had user")
		}
	})
	t.Run("RejectMountpathedRedirect", func(t *testing.T) {
		redir := &testRedirector{}
		ab.Config.Core.Redirector = redir

		ab.Storage.SessionState = mockClientStateReadWriter{}

		var reqs MWRequirements
		_, called, hadUser := setupMore(true, RespondRedirect, reqs)

		if redir.Opts.Code != http.StatusTemporaryRedirect {
			t.Error("code was wrong:", redir.Opts.Code)
		}
		if redir.Opts.RedirectPath != "/auth/login?redir=%2Fauth%2Fsuper%2Fsecret" {
			t.Error("redirect path was wrong:", redir.Opts.RedirectPath)
		}
		if called {
			t.Error("should not have been called")
		}
		if hadUser {
			t.Error("should not have had user")
		}
	})
	t.Run("RejectHalfAuth", func(t *testing.T) {
		ab.Storage.SessionState = mockClientStateReadWriter{
			state: mockClientState{SessionKey: "test@test.com", SessionHalfAuthKey: "true"},
		}

		rec, called, hadUser := setupMore(false, RespondNotFound, RequireFullAuth)

		if rec.Code != http.StatusNotFound {
			t.Error("wrong code:", rec.Code)
		}
		if called {
			t.Error("should not have been called")
		}
		if hadUser {
			t.Error("should not have had user")
		}
	})
}
