/**
 * File containing the TodoService methods.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "net/http"
    "time"
    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

type Todo struct {
    Name      string        `json:"name"`
    Completed bool          `json:"completed"`
    Due       time.Time     `json:"due"`
}

//--------------------------------------------------------------------------

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    dispatchResponseObject(w, todos, 0)
}

//--------------------------------------------------------------------------

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]

    dispatchResponse(w, "Todo show:" + todoId, 0)
}

//--------------------------------------------------------------------------
