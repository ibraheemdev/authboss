# Functionality

###  Get Current User

CurrentUser can be retrieved by calling
[Authboss.CurrentUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#Authboss.CurrentUser)
but a pre-requisite is that
[Authboss.LoadClientState](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#Authboss.LoadClientState)
has been called first to load the client state into the request context.
This is typically achieved by using the
[Authboss.LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#Authboss.LoadClientStateMiddleware), but can
be done manually as well.

###  Reset Password

Updating a user's password is non-trivial for several reasons:

1. The bcrypt algorithm must have the correct cost, and also be being used.
1. The user's remember me tokens should all be deleted so that previously authenticated sessions are invalid
1. Optionally the user should be logged out (**not taken care of by UpdatePassword**)

In order to do this, we can use the
[Authboss.UpdatePassword](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#Authboss.UpdatePassword)
method. This ensures the above facets are taken care of which the exception of the logging out part.

If it's also desirable to have the user logged out, please use the following methods to erase
all known sessions and cookies from the user.

* [authboss.DelKnownSession](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#DelKnownSession)
* [authboss.DelKnownCookie](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#DelKnownCookie)

*Note: DelKnownSession has been deprecated for security reasons*

###  User Auth via Password

| Info and Requirements |          |
| --------------------- | -------- |
Module        | auth
Pages         | login
Routes        | /login
Emails        | _None_
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware)
ClientStorage | Session and Cookie
ServerStorer  | [ServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#ServerStorer)
User          | [AuthableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#AuthableUser)
Values        | [UserValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#UserValuer)
Mailer        | _None_

To enable this side-effect import the auth module, and ensure that the requirements above are met.
It's very likely that you'd also want to enable the logout module in addition to this.

Direct a user to `GET /login` to have them enter their credentials and log in.


###  User Auth via OAuth2

| Info and Requirements |          |
| --------------------- | -------- |
Module        | oauth2
Pages         | _None_
Routes        | /oauth2/{provider}, /oauth2/callback/{provider}
Emails        | _None_
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#Authboss.LoadClientStateMiddleware)
ClientStorage | Session
ServerStorer  | [OAuth2ServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#OAuth2ServerStorer)
User          | [OAuth2User](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#OAuth2User)
Values        | _None_
Mailer        | _None_

This is a tougher implementation than most modules because there's a lot going on. In addition to the
requirements stated above, you must also configure the `OAuth2Providers` in the config struct.

The providers require an oauth2 configuration that's typical for the Go oauth2 package, but in addition
to that they need a `FindUserDetails` method which has to take the token that's retrieved from the oauth2
provider, and call an endpoint that retrieves details about the user (at LEAST user's uid).
These parameters are returned in `map[string]string` form and passed into the `OAuth2ServerStorer`.

Please see the following documentation for more details:

* [Package docs for oauth2](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/oauth2/)
* [authboss.OAuth2Provider](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#OAuth2Provider)
* [authboss.OAuth2ServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss#OAuth2ServerStorer)

### User Registration

| Info and Requirements |          |
| --------------------- | -------- |
Module        | register
Pages         | register
Routes        | /register
Emails        | _None_
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware)
ClientStorage | Session
ServerStorer  | [CreatingServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#CreatingServerStorer)
User          | [AuthableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/#AuthableUser), optionally [ArbitraryUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ArbitraryUser)
Values        | [UserValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#UserValuer), optionally also [ArbitraryValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ArbitraryValuer)
Mailer        | _None_

Users can self-register for a service using this module. You may optionally want them to confirm
themselves, which can be done using the confirm module.

The complicated part in implementing registrations are around the `RegisterPreserveFields`. This is to
help in the case where a user fills out many fields, and then say enters a password
which doesn't meet minimum requirements and it fails during validation. These preserve fields should
stop the user from having to type in all that data again (it's a whitelist). This **must** be used
in conjuction with `ArbitraryValuer` and although it's not a hard requirement `ArbitraryUser`
should be used otherwise the arbitrary values cannot be stored in the database.

When the register module sees arbitrary data from an `ArbitraryValuer`, it sets the data key
`authboss.DataPreserve` with a `map[string]string` in the data for when registration fails.
This means the (whitelisted) values entered by the user previously will be accessible in the
templates by using `.preserve.field_name`. Preserve may be empty or nil so use
`{{with ...}}` to make sure you don't have template errors.

There is additional [Godoc documentation](https://pkg.go.dev/mod/github.com/ibraheemdev/authboss/v3#Config) on the `RegisterPreserveFields` config option as well as
the `ArbitraryUser` and `ArbitraryValuer` interfaces themselves.

### Confirming Registrations

| Info and Requirements |          |
| --------------------- | -------- |
Module        | confirm
Pages         | confirm
Routes        | /confirm
Emails        | confirm_html, confirm_txt
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware), [confirm.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/confirm/#Middleware)
ClientStorage | Session
ServerStorer  | [ConfirmingServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ConfirmingServerStorer)
User          | [ConfirmableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ConfirmableUser)
Values        | [ConfirmValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ConfirmValuer)
Mailer        | Required

Confirming registrations via e-mail can be done with this module (whether or not done via the register
module).

A hook on register kicks off the start of a confirmation which sends an e-mail with a token for the user.
When the user re-visits the page, the `BodyReader` must read the token and return a type that returns
the token.

Confirmations carry two values in the database to prevent a timing attack. The selector and the
verifier, always make sure in the ConfirmingServerStorer you're searching by the selector and
not the verifier.

### Password Recovery

| Info and Requirements |          |
| --------------------- | -------- |
Module        | recover
Pages         | recover_start, recover_middle (not used for renders, only values), recover_end
Routes        | /recover, /recover/end
Emails        | recover_html, recover_txt
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware)
ClientStorage | Session
ServerStorer  | [RecoveringServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RecoveringServerStorer)
User          | [RecoverableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RecoverableUser)
Values        | [RecoverStartValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RecoverStartValuer), [RecoverMiddleValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RecoverMiddleValuer), [RecoverEndValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RecoverEndValuer)
Mailer        | Required

The flow for password recovery is that the user is initially shown a page that wants their `PID` to
be entered. The `RecoverStartValuer` retrieves that on `POST` to `/recover`.

An e-mail is sent out, and the user clicks the link inside it and is taken back to `/recover/end`
as a `GET`, at this point the `RecoverMiddleValuer` grabs the token and will insert it into the data
to be rendered.

They enter their password into the form, and `POST` to `/recover/end` which sends the token and
the new password which is retrieved by `RecoverEndValuer` which sets their password and saves them.

Password recovery has two values in the database to prevent a timing attack. The selector and the
verifier, always make sure in the RecoveringServerStorer you're searching by the selector and
not the verifier.

### Remember Me

| Info and Requirements |          |
| --------------------- | -------- |
Module        | remember
Pages         | _None_
Routes        | _None_
Emails        | _None_
Middlewares   | LoadClientStateMiddleware,
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware), [remember.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/remember/#Middleware)
ClientStorage | Session, Cookies
ServerStorer  | [RememberingServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RememberingServerStorer)
User          | User
Values        | [RememberValuer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#RememberValuer) (not a Validator)
Mailer        | _None_

Remember uses cookie storage to log in users without a session via the `remember.Middleware`.
Because of this this middleware should be used high up in the stack, but it also needs to be after
the `LoadClientStateMiddleware` so that client state is available via the authboss mechanisms.

There is an intricacy to the `RememberingServerStorer`, it doesn't use the `User` struct at all,
instead it simply instructs the storer to save tokens to a pid and recall them just the same. Typically
in most databases this will require a separate table, though you could implement using pg arrays
or something as well.

A user who is logged in via Remember tokens is also considered "half-authed" which is a session
key (`authboss.SessionHalfAuthKey`) that you can query to check to see if a user should have
full rights to more sensitive data, if they are half-authed and they want to change their user
details for example you may want to force them to go to the login screen and put in their
password to get a full auth first. The `authboss.Middleware` has a boolean flag to `forceFullAuth`
which prevents half-authed users from using that route.

### Locking Users

| Info and Requirements |          |
| --------------------- | -------- |
Module        | lock
Pages         | _None_
Routes        | _None_
Emails        | _None_
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware), [lock.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/lock/#Middleware)
ClientStorage | Session
ServerStorer  | [ServerStorer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#ServerStorer)
User          | [LockableUser](https://pkg.go.dev/github.com/ibraheemdev/authboss/#LockableUser)
Values        | _None_
Mailer        | _None_

Lock ensures that a user's account becomes locked if authentication (both auth, oauth2, otp) are
failed enough times.

The middleware protects resources from locked users, without it, there is no point to this module.
You should put in front of any resource that requires a login to function.

### Expiring User Sessions

| Info and Requirements |          |
| --------------------- | -------- |
Module        | expire
Pages         | _None_
Routes        | _None_
Emails        | _None_
Middlewares   | [LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Authboss.LoadClientStateMiddleware), [expire.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/expire/#Middleware)
ClientStorage | Session
ServerStorer  | _None_
User          | [User](https://pkg.go.dev/github.com/ibraheemdev/authboss/#User)
Values        | _None_
Mailer        | _None_

Expire simply uses sessions to track when the last action of a user is, if that action is longer
than configured then the session is deleted and the user removed from the request context.

This middleware should be inserted at a high level (closer to the request) in the middleware chain
to ensure that "activity" is logged properly, as well as any middlewares down the chain do not
attempt to do anything with the user before it's removed from the request context.