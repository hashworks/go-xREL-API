# xREL API Package

[![GoDoc](https://godoc.org/github.com/hashworks/go-xREL-API/xrel?status.svg)](https://godoc.org/github.com/hashworks/go-xREL-API/xrel)

A golang package to authorize with and access the complete xrel.to API.

Import this way:
```go
import github.com/hashworks/go-xREL-API/xrel
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
