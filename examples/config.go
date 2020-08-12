package examples

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ibraheemdev/authboss/pkg/authboss"
	"github.com/ibraheemdev/authboss/pkg/authboss/defaults"
)

// SetupAuthboss :
func SetupAuthboss() {
	ab := authboss.New()
	c := ab.Config

	// ************** Core Config **************

	// Router is the entity that controls all routing to authboss routes
	// modules will register their routes with it.
	c.Core.Router = defaults.NewRouter()

	// ErrorHandler wraps http requests with centralized error handling.
	c.Core.ErrorHandler = defaults.NewErrorHandler(defaults.NewLogger(os.Stdout))

	// ViewRenderer loads the templates for the application.
	// You can also use defaults.JSONRenderer for api usage
	// MUST have trailing slash
	c.Core.ViewRenderer = defaults.NewHTMLRenderer("/auth", "./web/templates/authboss/")

	// Responder takes a generic response from a controller and prepares
	// the response, uses a renderer to create the body, and replies to the
	// http request.
	c.Core.Responder = defaults.NewResponder(c.Core.ViewRenderer)

	// Redirector can redirect a response, similar to Responder but
	// responsible only for redirection.
	c.Core.Redirector = defaults.NewRedirector(c.Core.ViewRenderer, authboss.FormValueRedirect)

	// BodyReader reads validatable data from the body of a request to
	// be able to get data from the user's client.
	c.Core.BodyReader = defaults.NewHTTPBodyReader(false, false)

	// Mailer is the mailer being used to send e-mails out via smtp
	c.Core.Mailer = defaults.NewLogMailer(os.Stdout)

	// Logger implies just a few log levels for use, can optionally
	// also implement the ContextLogger to be able to upgrade to a
	// request specific logger.
	c.Core.Logger = defaults.NewLogger(os.Stdout)

	// MailRenderer loads the templates for mail. If this is nil, it will
	// fall back to using the Renderer created from the ViewLoader instead.
	c.Core.MailRenderer = c.Core.ViewRenderer

	// ************** Storage Config **************

	// Storer is the interface through which Authboss accesses the web apps
	// database for user operations.
	//
	// The in memory database from the user model
	// c.Storage.Server = DB

	// SessionState must be defined to provide an interface capable of
	// storing session-only values for the given response, and reading them
	// from the request.
	//
	// Must implement the authboss ClientStateReadWriter interface
	// c.Storage.SessionState = yourSessionState

	// CookieState must be defined to provide an interface capapable of
	// storing cookies for the given response, and reading them from the
	// request.
	//
	// Must implement the authboss ClientStateReadWriter interface
	// c.Storage.CookieState = yourCookieStore

	// SessionStateWhitelistKeys are set to preserve keys in the session
	// when authboss.DelAllSession is called. A correct implementation
	// of ClientStateReadWriter will delete ALL session key-value pairs
	// unless that key is whitelisted here.
	c.Storage.SessionStateWhitelistKeys = []string{}

	// ************** Paths Config **************

	// Mount is the path to mount authboss's routes at (eg /auth).
	c.Paths.Mount = "/auth"

	// NotAuthorized is the default URL to kick users back to when
	// they attempt an action that requires them to be logged in and
	// they're not auth'd
	c.Paths.NotAuthorized = "/"

	// AuthLoginOK is the redirect path after a successful authentication.
	c.Paths.AuthLoginOK = "/"

	// AuthLoginOK is the redirect path after a successful authentication.
	c.Paths.ConfirmOK = "/"

	// ConfirmNotOK is used by the middleware, when a user is still supposed
	// to confirm their account, this is where they should be redirected to.
	c.Paths.ConfirmNotOK = "/auth/login"

	// LockNotOK is a path to go to when the user gets locked out
	c.Paths.LockNotOK = "/auth/login"

	// LogoutOK is the redirect path after a log out.
	c.Paths.LogoutOK = "/"

	// OAuth2LoginOK is the redirect path after a successful oauth2 login
	c.Paths.LogoutOK = "/"

	// OAuth2LoginOK is the redirect path after a successful oauth2 login
	c.Paths.OAuth2LoginOK = "/"

	// OAuth2LoginOK is the redirect path after a unsuccessful oauth2 login
	c.Paths.OAuth2LoginNotOK = "/"

	// OAuth2LoginOK is the redirect path after a successful oauth2 login
	c.Paths.OAuth2LoginNotOK = "/"

	// OAuth2LoginOK is the redirect path after a successful oauth2 login
	c.Paths.RecoverOK = "/"

	// OAuth2LoginOK is the redirect path after a successful oauth2 login
	c.Paths.RegisterOK = "/"

	// RootURL is the scheme+host+port of the web application
	// No trailing slash.
	c.Paths.RootURL = "http://localhost:8080"

	// TwoFactorEmailAuthRequired forces users to first confirm they have
	// access to their e-mail with the current device by clicking a link
	// and confirming a token stored in the session.
	c.Paths.TwoFactorEmailAuthNotOK = "/"

	// ************** Modules Config **************

	// BCryptCost is the cost of the bcrypt password hashing function.
	c.Modules.BCryptCost = bcrypt.DefaultCost

	// ConfirmMethod IS DEPRECATED! See MailRouteMethod instead.
	//
	// ConfirmMethod controls which http method confirm expects.
	// This is because typically this is a GET request since it's a link
	// from an e-mail, but in api-like cases it needs to be able to be a
	// post since there's data that must be sent to it.
	c.Modules.ConfirmMethod = http.MethodGet

	// ExpireAfter controls the time an account is idle before being
	// logged out by the ExpireMiddleware.
	c.Modules.ExpireAfter = time.Hour

	// LockAfter this many tries.
	c.Modules.LockAfter = 3

	// LockWindow is the waiting time before the number of attempts are reset.
	c.Modules.LockWindow = 5 * time.Minute

	// LockDuration is how long an account is locked for.
	c.Modules.LockDuration = 12 * time.Hour

	// LogoutMethod is the method the logout route should use
	// (default should be DELETE)
	c.Modules.LogoutMethod = "DELETE"

	// MailRouteMethod is used to set the type of request that's used for
	// routes that require a token from an e-mail link's query string.
	// This is things like confirm and two factor e-mail auth.
	//
	// You should probably set this to POST if you are building an API
	// so that the user goes to the frontend with their link & token
	// and the front-end calls the API with the token in a POST JSON body.
	//
	// This configuration setting deprecates ConfirmMethod.
	// If ConfirmMethod is set to the default value (GET) then
	// MailRouteMethod is used. If ConfirmMethod is not the default value
	// then it is used until Authboss v3 when only MailRouteMethod will be
	// used.
	c.Modules.MailRouteMethod = http.MethodGet

	// MailNoGoroutine is used to prevent the mailer from being launched
	// in a goroutine by the Authboss modules.
	//
	// It's important that this is the case if you are using contexts
	// as the http request context will be cancelled by the Go http server
	// and it may interrupt your use of the context that the Authboss module
	// is passing to you, preventing proper use
	c.Modules.MailNoGoroutine = false

	// RegisterPreserveFields are fields used with registration that are
	// to be rendered when post fails in a normal way
	// (for example validation errors), they will be passed back in the
	// data of the response under the key DataPreserve which
	// will be a map[string]string. This way the user does not have to
	// retype these whitelisted fields.
	//
	// All fields that are to be preserved must be able to be returned by
	// the ArbitraryValuer.GetValues()
	//
	// This means in order to have a field named "address" you would need
	// to have that returned by the ArbitraryValuer.GetValues() method and
	// then it would be available to be whitelisted by this
	// configuration variable.
	c.Modules.RegisterPreserveFields = []string{}

	// RecoverTokenDuration controls how long a token sent via
	// email for password recovery is valid for.
	c.Modules.RecoverTokenDuration = 24 * time.Hour

	// RecoverLoginAfterRecovery says for the recovery module after a
	// user has successfully recovered the password, are they simply
	// logged in, or are they redirected to the login page with an
	// "updated password" message.
	c.Modules.RecoverLoginAfterRecovery = false

	// OAuth2Providers lists all providers that can be used. See
	// OAuthProvider documentation for more details.
	c.Modules.OAuth2Providers = map[string]authboss.OAuth2Provider{}

	// OAuth2Providers lists all providers that can be used. See
	// OAuthProvider documentation for more details.
	c.Modules.TwoFactorEmailAuthRequired = true

	// ************** Mail Config **************

	// RootURL is a full path to an application that is hosting a front-end
	// Typically using a combination of Paths.RootURL and Paths.Mount
	// MailRoot will be assembled if not set.
	// Typically looks like: https://our-front-end.com/authenication
	// No trailing slash.
	// Defaults to the mount url if empty
	c.Mail.RootURL = ""

	// From is the email address authboss e-mails come from.
	c.Mail.From = "authboss@example.org"

	// From is the email address authboss e-mails come from.
	c.Mail.FromName = "authboss"

	// SubjectPrefix is used to add something to the front of the authboss
	// email subjects.
	c.Mail.SubjectPrefix = ""
}
