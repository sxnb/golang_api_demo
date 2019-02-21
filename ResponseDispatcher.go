/**
 * Method that handles the correct forming of an answer.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "fmt"
    "encoding/json"
    "net/http"
)

//--------------------------------------------------------------------------

func _getErrorMessage(errorCode int) string {
    switch errorCode {
    case 401:
        return "Not authorized"
    }

    return "This error does not have a custom error message."
}

//--------------------------------------------------------------------------

func _dispatchSuccessResponse(w http.ResponseWriter, response string) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintln(w, "{\"success\": true, \"response\": " + response + "}")
}

//--------------------------------------------------------------------------

func _dispatchErrorResponse(w http.ResponseWriter, errorMessage string, errorCode int) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(errorCode)

    fmt.Fprintln(w, "{\"error\": \"" + errorMessage + "\", \"success\": \"false\"}")
}

//--------------------------------------------------------------------------

func dispatchResponse(w http.ResponseWriter, response string, errorCode int) {
    if errorCode == 0 {
        _dispatchSuccessResponse(w, "\"" + response + "\"")
    } else {
        _dispatchErrorResponse(w, response, errorCode)
    }
}

//--------------------------------------------------------------------------

func dispatchResponseObject(w http.ResponseWriter, response interface{}, errorCode int) {
    serializedStruct, err := json.Marshal(response)

    if err != nil {
        _dispatchErrorResponse(w, "Internal Error", 500)
    }

    if errorCode == 0 {
        _dispatchSuccessResponse(w, string(serializedStruct))
    } else {
        _dispatchErrorResponse(w,  _getErrorMessage(errorCode), errorCode)
    }
}

//--------------------------------------------------------------------------
