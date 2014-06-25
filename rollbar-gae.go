// +build appengine
package rollbar

import (
  "appengine"
  "appengine/urlfetch"
  "net/http"
)


func NewGaeClient(r *http.Request) *Client {
  c := appengine.NewContext(r)
  client := urlfetch.Client(c)
  return &Client{ client }
}
