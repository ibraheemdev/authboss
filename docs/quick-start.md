# Quick Start

Authboss is a standard go module. You can install it by running:

```bash
go get github.com/ibraheemdev/authboss/...
```
You can start by generating a user model. Note: This command **will overwrite the file** if it already exists:

```bash
$ authboss generate:user -d ./user.go
```

The generated file will contain a user model implementing the `Authable`, `Recoverable`, `Confirmable`, `Lockable`, `OAuthable`, and `Rememberable` modules. It is a generic implementation that uses an in memory database: 

```go
// User database model
type User struct {
  ID   string
  Name string

  // Authable
  Email    string
  Password string

  // Recoverable
  RecoverSelector    string
  RecoverVerifier    string
  RecoverTokenExpiry time.Time

  // Confirmable
  ConfirmSelector string
  ConfirmVerifier string
  Confirmed       bool

  // Lockable
  AttemptCount int
  LastAttempt  time.Time
  Locked       time.Time

  // OAuthable
  OAuth2UID          string
  OAuth2Provider     string
  OAuth2AccessToken  string
  OAuth2RefreshToken string
  OAuth2Expiry       time.Time
    
  // Rememberable
  RememberTokens []string
}

// DB : The in memory database instance
var DB *InMemDB

// Authboss interface implementation methods ...
```
You can now edit the methods to suit your database.

*Want to help out? Create a pull request including a user model for a popular ORM!*

You can generate the default templates using the build in authboss generator:

```bash
authboss generate:templates -d ./destination_path
```

Here's a bit of starter code:

```go
ab := authboss.New()

ab.Config.Storage.Server = myDatabaseImplementation
ab.Config.Storage.SessionState = mySessionImplementation
ab.Config.Storage.CookieState = myCookieImplementation

ab.Config.Paths.Mount = "/authboss"
ab.Config.Paths.RootURL = "https://www.example.com/"

// This is using the renderer from: github.com/volatiletech/authboss-renderer
ab.Config.Core.ViewRenderer = abrenderer.NewHTML("/auth", "ab_views")
ab.Config.Core.MailRenderer = abrenderer.NewEmail("/auth", "ab_views")

// This instantiates and uses every default implementation
// in the Config.Core area that exist in the defaults package.
 defaults.SetCore(&ab.Config, false, false)

if err := ab.Init(); err != nil {
    panic(err)
}

// Mount the router to a path (this should be the same as the Mount path above)
// mux in this example is a chi router
mux.Mount("/authboss", http.StripPrefix("/authboss", ab.Config.Core.Router))
```

For a more in-depth look, refer to the the authboss sample to see what a full implementation looks like. This will probably help you more than any of this documentation.

[https://github.com/volatiletech/authboss-sample](https://github.com/volatiletech/authboss-sample)
