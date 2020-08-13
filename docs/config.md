# Config

The config struct is an important part of Authboss. It's the key to making Authboss do what you
want with the implementations you want. Please look at it's code definition below:

[Config Struct Documentation](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#Config)

To view all the configuration options, you can generate the default config file:

```bash
authboss generate:config ./config.go
```

### Paths

Paths are the paths that should be redirected to or used in whatever circumstance they describe.
Two special paths that are required are `Mount` and `RootURL` without which certain authboss
modules will not function correctly. Most modules redirect to the root page `/`, such as after login or register success
or when a user is locked out of their account.

### Modules

Modules are module specific configuration options. They mostly control the behavior of modules.
For example `RegisterPreserveFields` decides a whitelist of fields to allow back into the data
to be re-rendered so the user doesn't have to type them in again.

### Mail

Mail sending related options, including the RootURL, the Name and Email address of the sender, as well as a subject prefix that can be used to add text infront of the email subject.

### Storage

These are the implementations of how storage on the server and the client are done in your
application. There are implementations for the CookieStore and the SessionStore, but not the ServerStorer. The ServerStorer must be manually implemented depending on which database your app uses. For a sample in memory server store, you can run the user model generator:

Note: This command **will overwrite the file** if it already exists:

```bash
$ authboss generate:user ./user.go
```

### Core

These are the implementations of the HTTP stack for your app. How do responses render? How are
they redirected? How are errors handled?

There are default implementations from the
[defaults package](https://github.com/ibraheemdev/authboss/tree/master/pkg/authboss/defaults) available for all of the core config options.