# xREL API Package

[![GoDoc](https://godoc.org/github.com/hashworks/go-xREL-API/xrel?status.svg)](https://godoc.org/github.com/hashworks/go-xREL-API/xrel)

A golang package to authorize with and access the complete xrel.to API.

Import this way:
```go
import github.com/hashworks/go-xREL-API/xrel
```

All methods may return a `types.Error` struct, which implements the normal `error` type.
Additional to the `Error()` function this struct contains the variables `Type` and `Code`.
Errors with type `api` are [xREL.to errors](https://www.xrel.to/wiki/6435/api-errors.html), for all other
types and error codes see the `xrel/types/errors.go` file.

To use this in your code try to cast it:
```go
err := xrel.SomeMethod()
if eErr, ok := err.(*types.Error); ok {
	// Is of type types.Error, you can use the variables
} else {
	// Is normal error
}
```

If you use the OAuth authentication make sure to save the Config variable somewhere and set it again on your next run.
Here is an example how to use the OAuth2 authentication:

```go
xrel.ConfigureOAuth2("OAUTH2_CLIENT_KEY", "OAUTH2_CLIENT_SECRET", "", []string{"viewnfo", "addproof"})

fmt.Println("(1) Go to: " + xrel.GetOAuth2RequestURL())
fmt.Println("(2) Grant access, you should get back a verification code.")
fmt.Print("(3) Enter that verification code here: ")
verificationCode := ""
fmt.Scanln(&verificationCode)

err := xrel.InitiateOAuth2CodeExchange(verificationCode)
ok(err)

```
