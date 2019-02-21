/**
 * File containing the UserService methods.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "net/http"
    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

func GetUsers(w http.ResponseWriter, r *http.Request) {
    trips, err := GetDAO(DAO_TYPE_USER).(UserDAO).GetUsers()
    if err != nil {
    }

    dispatchResponseObject(w, trips, 0)
}

//--------------------------------------------------------------------------

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]

    user, err := GetDAO(DAO_TYPE_USER).(UserDAO).GetUser(userId)
    if err != nil {
    }

    dispatchResponse(w, user, 0)
}

//--------------------------------------------------------------------------

func AddUser(w http.ResponseWriter, r *http.Request) {
    userId, err := GetDAO(DAO_TYPE_USER).(UserDAO).AddUser()
    if err != nil {
    }

    dispatchResponse(w, userId, 0)
}

//--------------------------------------------------------------------------

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    userId, err := GetDAO(DAO_TYPE_USER).(UserDAO).UpdateUser()
    if err != nil {
    }

    dispatchResponse(w, userId, 0)
}

//--------------------------------------------------------------------------

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    _, err := GetDAO(DAO_TYPE_USER).(UserDAO).DeleteUser()
    if err != nil {
    }

    dispatchResponse(w, "true", 0)
}

//--------------------------------------------------------------------------
