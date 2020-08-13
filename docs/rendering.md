# Rendering Views

The authboss rendering system is simple. It's defined by one interface: [Renderer](https://pkg.go.dev/github.com/ibraheemdev/authboss/#Renderer)

The renderer knows how to load templates, and how to render them with some data and that's it.
So let's examine the most common view types that you might want to use.

### HTML Views

When your app is a traditional web application it is generating it's HTML
serverside using templates. You can use the `defaults.NewHTMLRenderer()` for this setup. 
It takes a url mount path, the path to your templates, and the path to your application layout template.
If you need additional functionality, you can set `c.Core.ViewRenderer` to your custom renderer. Take a look at the [default renderer's source code](https://github.com/ibraheemdev/authboss/blob/master/pkg/authboss/defaults/html.go#L15) for more information.

### JSON Views

If you're building an API that's mostly backed by a javascript front-end, then you'll probably
want to use a renderer that converts the data to JSON. For this, you can use set the `c.Core.ViewRenderer` to 
[`defaults.JSONRenderer`](https://github.com/ibraheemdev/authboss/blob/v0.6.0/pkg/authboss/defaults/json.go#L14), or customize it to your preference.

### Data

The most important part about this interface is the data that you have to render.
There are several keys that are used throughout authboss that you'll want to render in your views.

They're in the file [html_data.go](https://github.com/ibraheemdev/authboss/blob/master/pkg/authboss/html_data.go#L9)
and are constants prefixed with `Data`. See the documentation in that file for more information on
which keys exist and what they contain.

The default [responder](https://pkg.go.dev/github.com/ibraheemdev/authboss/defaults/#Responder)
also happens to collect data from the Request context, and hence this is a great place to inject
data you'd like to render (for example data for your html layout, or csrf tokens).

There is also a useful configuration option called `c.Modules.RegisterPreserveFields`. It takes a slice of strings that will be preserved after a failed request. This way the fields will be sent back to the user, and he will not have to retype them again. The default only preserves the email field.