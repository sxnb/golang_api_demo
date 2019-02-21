/**
 * File containing the router object.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "net/http"
    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {
        var handler http.Handler

        //
        handler = route.HandlerFunc
        handler = LoggerInterceptor(handler, route.Name)

        handler = Middleware(handler)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

//--------------------------------------------------------------------------
