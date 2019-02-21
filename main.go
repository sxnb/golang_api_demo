/**
 * A basic http server.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "os"        // for cmdline arguments and for exit with error
    "fmt"
    "html"      // for string escapation
    "log"       // any direct output should be done using the logger
    "net/http"  // for
    "strconv"   // for converting port to string
    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

const DefaultPort = 8080

//--------------------------------------------------------------------------

func getPort() int {
    var port int = DefaultPort
    var err error

    if len(os.Args) > 1 {
        port, err = strconv.Atoi(os.Args[1])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
    }

    return port
}

//--------------------------------------------------------------------------

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//--------------------------------------------------------------------------

func registerRoutes(r *mux.Router) {
    r.HandleFunc("/", index)
}

//--------------------------------------------------------------------------

func main() {
    router := NewRouter()

    registerRoutes(router)
    initDAOFactory()

    port := getPort()

    log.Printf("Listening on port %d\n", port)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), router))
}

//--------------------------------------------------------------------------
