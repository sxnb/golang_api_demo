/**
 * File containing the middleware logic.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "log"
    "net/http"
//    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

func _validateToken(token string) bool {
    if (token == string("abc123")) {
        return true
    }

    return false
}

//--------------------------------------------------------------------------

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("=== middleware")
        log.Println("Host: " + r.Host)

        var token string
        token = ""

        if val, ok := r.Header["Token"]; ok {
            if len(val) > 0 {
                token = val[0]
            }
        }

        if _validateToken(token) {
            h.ServeHTTP(w, r)
        } else {
            dispatchResponse(w, "Not authorized", 401)
        }
    })
}

//--------------------------------------------------------------------------
