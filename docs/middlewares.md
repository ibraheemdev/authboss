# Middlewares

The only middleware that's truly required is the `LoadClientStateMiddleware`, and that's because it
enables session and cookie handling for Authboss. Without that, it's not a very useful piece of
software.

The remaining middlewares are either the implementation of an entire module (like expire),
or a key part of a module. You can use a specific module's middleware on part of your application to prevent unauthorized users from accessing those routes:

Name | Description
---- | -----------
[Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#Middleware) | Prevents unauthenticated users from accessing routes.
[LoadClientStateMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#Authboss.LoadClientStateMiddleware) | **Required** Enables cookie and session handling
[ModuleListMiddleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=doc#Authboss.ModuleListMiddleware) | Inserts a list of loaded modules into the view data. This can be useful for showing specific routes depending on which modules are being used, or for debugging an application.
[confirm.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/authboss?tab=docconfirm/#Middleware) | Prevents unconfirmed users from accessing routes.
[expire.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/expire/#Middleware) | **Required** with expire. Expires user sessions after an inactive period
[lock.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/lock/#Middleware) | Prevents locked users from accessing routes.
[remember.Middleware](https://pkg.go.dev/github.com/ibraheemdev/authboss/pkg/remember/#Middleware) | Logs a user in from a remember cookie

Note: Middlewares that load a user context will panic if there is no current user. An example of this is the lock middleware, which requires a user context in order to check if the current user is locked or not. You should no wrap your entire app in these middlewares, and only use them on specific routes. You can use a middleware to recover from unexpected panics, such as [gorilla recovery handler](https://godoc.org/github.com/gorilla/handlers#RecoveryHandler).