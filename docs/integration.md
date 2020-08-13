# Integration


### Modules

Authboss consists of 9 modules. You can use as many and as few of these modules as you want (within reason). These are:

* [Database Authenticatable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authenticatable?tab=doc): hashes and stores a password in the database to validate the authenticity of a user while signing in.
* [Logoutable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/logoutable?tab=doc): implements user logout functionality
* [OAuthable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/oauthable?tab=doc): adds OAuth support.
* [Confirmable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/confirmable?tab=doc): sends emails with confirmation instructions and verifies whether an account is already confirmed during sign in.
* [Recoverable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/recoverable?tab=doc): resets the user password and sends reset instructions.
* [Registerable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/registerable?tab=doc): handles signing up users through a registration process, also allowing them to edit and destroy their account.
* [Rememberable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/rememberable?tab=doc): manages generating and clearing a token for remembering the user from a saved cookie.
* [Timeoutable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/timeoutable?tab=doc): expires sessions that have not been active in a specified period of time.
* [Lockable](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/lockable?tab=doc): locks an account after a specified number of failed sign-in attempts.

To use a module, you can simply use a blank import:

```go
import _"github.com/ibraheemdev/authboss/pkg/authenticatable"
```

This will call the module's init function, which registers the module with authboss. Most modules are accompanied with middlewares to protect certain routes, or provide other useful functionality. See [Middlewares](middlewares.md) for more information.

### Configuration

Authboss uses interfaces to make almost everything configurable. It includes many defaults, that will fit most usecases. For a full list of configuration options, refer to [Configuration](config.md).

### Storage and Core implementations

Everything under Config.Storage and Config.Core are required variables. However, you can optionally use default implementations from the [defaults package](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss/defaults?tab=doc).
This also provides an easy way to share implementations of certain stack pieces (like HTML Form Parsing).
As you saw in the quick start example above these can be easily initialized with the `SetCore` method in that
package.

### User implementation

Users in Authboss are represented by the
[User interface](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#User). The user
interface is a flexible notion, because it can be upgraded to suit the needs of the various modules.

Initially the User must only be able to Get/Set a `PID` or primary identifier. This allows the authboss
modules to know how to refer to him in the database. The `ServerStorer` also makes use of this
to save/load users.

As mentioned, it can be upgraded, for example suppose now we want to use the `confirm` module,
in that case the e-mail address now becomes a requirement. So the `confirm` module will attempt
to upgrade the user (and panic if it fails) to a
[ConfirmableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#ConfirmableUser)
which supports retrieving and setting of confirm tokens, e-mail addresses, and a confirmed state.

Your `User` implementation does not need to implement all these additional user interfaces unless you're
using a module that requires it. See the [Functionality](functionality.md) documentation to know what the
requirements are.

### BodyReader implementation

The BodyReader interface in the Config returns Validator implementations which can be used to validate user input. But much like the storer and user it can be upgraded to add different capabilities.

A typical BodyReader (like the one in the defaults package) implementation parses the request body and produces a struct that has the ability to Validate() it's data as well as functions to retrieve the data necessary for the particular valuer required by the module.

Many modules use upgraded "valuers". For example the confirmable package uses a ConfirmValuer which allows it to pull out the confirmation token from the request.

Your body reader implementation does not need to implement all valuer types unless you're using a module that requires it. See the [Functionality](functionality.md) documentation to know what the requirements are.

# What's Not Included

Authboss has a lot of built in functionality. However, some important things are not included as they would be difficult to abstract into a solution that would fit the many go frameworks and toolkits. These are included below:

### CSRF Protection

Authboss does not deal with csrf protection.
You should apply a middleware that will protect your application from crsf attacks or you may be vulnerable. Some popular libraries for dealing with this are [gorilla csrf](https://github.com/gorilla/csrf), and [justinas's nosurf package](https://github.com/justinas/nosurf).

### Request Throttling

Currently Authboss is vulnerable to brute force attacks because there are no protections on
it's endpoints. This again is left up to the creator of the website to protect the whole application
at once (as well as Authboss) from these sorts of attacks.