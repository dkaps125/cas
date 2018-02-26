package cas

import (
	"net/http"
    "fmt"

	"github.com/golang/glog"
)

func (c *Client) Handler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if glog.V(2) {
			glog.Infof("cas: handling %v request for %v", r.Method, r.URL)
		}

		setClient(r, c)
        c.getSession(w, r)

		if !IsAuthenticated(r) {
			RedirectToLogin(w, r)
			return
		}

		if r.URL.Path == "/logout" {
			RedirectToLogout(w, r)
			return
		}

        next(w, r)
	}
}
