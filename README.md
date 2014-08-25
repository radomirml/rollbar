rollbar
-------

`rollbar` is a Golang Rollbar client that makes it easy to report errors to
Rollbar with full stacktraces. Errors are sent to Rollbar asynchronously in a
background goroutine.

Because Go's `error` type doesn't include stack information from when it was set
or allocated, we use the stack information from where the error was reported.

NOTE: This is a clone of [http://github.com/stvp/rollbar]() extended to support
use on Google AppEngine.

Documentation
=============

[API docs on godoc.org](http://godoc.org/github.com/radomirml/rollbar)

Usage
=====

```go
package main

import (
  "github.com/stvp/rollbar"
)

func main() {
  rollbar.Token = "MY_TOKEN"
  rollbar.Environment = "production" // defaults to "development"

  result, err := DoSomething()
  if err != nil {
    rollbar.Error(rollbar.ERR, err)
  }

  rollbar.Message("info", "Message body goes here")

  rollbar.Wait()
}
```

On Google AppEngine, instead of invoking rollbar package functions directly, create a
client passing *http.Request:

    rollbar.NewGaeClient(r).Error("error", err)


Running Tests
=============

Set up a dummy project in Rollbar and pass the access token as an environment
variable to `go test`:

    TOKEN=f0df01587b8f76b2c217af34c479f9ea go test

And verify the reported errors manually.

Contributors
============

A big thank you to everyone who has contributed pull requests and bug reports:

* @kjk
* @Soulou

