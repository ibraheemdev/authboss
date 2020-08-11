# Authboss

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4)](https://pkg.go.dev/mod/github.com/ibraheemdev/authboss)
[![Go Report Card](https://goreportcard.com/badge/github.com/ibraheemdev/authboss)](https://goreportcard.com/report/github.com/ibraheemdev/authboss)
[![Gopherbadger](https://img.shields.io/badge/Go%20Coverage-85.8%25-brightgreen.svg?longCache=true&style=flat)](https://github.com/jpoles1/gopherbadger)
[![Maintainability](https://api.codeclimate.com/v1/badges/9d7f1698687e79cf9ebf/maintainability)](https://codeclimate.com/github/ibraheemdev/authboss/maintainability)

Authboss is a flexible authentication solution for Go Web Applications. It makes it easy 
to plug in authentication to an application and get a lot of functionality with little
effort.

It is composed of 9 modules:

* [Database Authenticatable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/authenticatable?tab=doc): hashes and stores a password in the database to validate the authenticity of a user while signing in.
* [Logoutable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/logoutable?tab=doc): implements user logout functionality
* [OAuthable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/oauthable?tab=doc): adds OAuth support.
* [Confirmable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/confirmable?tab=doc): sends emails with confirmation instructions and verifies whether an account is already confirmed during sign in.
* [Recoverable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/recoverable?tab=doc): resets the user password and sends reset instructions.
* [Registerable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/registerable?tab=doc): handles signing up users through a registration process, also allowing them to edit and destroy their account.
* [Rememberable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/rememberable?tab=doc): manages generating and clearing a token for remembering the user from a saved cookie.
* [Timeoutable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/timeoutable?tab=doc): expires sessions that have not been active in a specified period of time.
* [Lockable](https://pkg.go.dev/github.com/ibraheemdev/authboss@v0.0.0-20200810134708-7f71c73afdd9/pkg/lockable?tab=doc): locks an account after a specified number of failed sign-in attempts.


### Why use Authboss?

Every time you'd like to start a new web project, you really want to get to the heart of what you're
trying to accomplish very quickly and it would be a sure bet to say one of the systems you're excited
about implementing and innovating on is not authentication. In fact it's very much the opposite: it's
one of those things that you have to do and one of those things you loathe to do. Authboss is supposed
to remove a lot of the tedium that comes with this, as well as a lot of the chances to make mistakes.
This allows you to care about what you're intending to do, rather than care about ancillary support
systems required to make what you're intending to do happen.

Here are a few bullet point reasons you might like to try it out:

* Saves you time (Authboss integration time should be less than re-implementation time)
* Saves you mistakes (at least using Authboss, people can bug fix as a collective and all benefit)
* Should integrate with or without any web framework
