/**
 * File containing the basic structures.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "net/http"
    "encoding/json"
    "database/sql"
)

//--------------------------------------------------------------------------

type SqlConnection struct {
    Host            string
    Port            int
    Database        string
    Username        string
    Password        string
    _Connection     DB
}

//--------------------------------------------------------------------------

type Route struct {
    Name            string
    Method          string
    Pattern         string
    HandlerFunc     http.HandlerFunc
}

//--------------------------------------------------------------------------

type Error struct {
    Code        int64            `json:"code"`
    Message     string           `json:"message"`
    Data        *json.RawMessage `json:"data"`
}

//--------------------------------------------------------------------------

type Position struct {
    Latitude    float32     `json:"latitude"`
    Longitude   float32     `json:"longitude"`
}

//--------------------------------------------------------------------------

type Positions []Position

//--------------------------------------------------------------------------

type Place struct {
    Position    Position    `json:"position"`
    Name        string      `json:"name"`
    Type        string      `json:"type"`
}

//--------------------------------------------------------------------------

type Trip struct {
    StartPosition   Position    `json:"startPosition"`
    PositionList    []Position  `json:"positionList"`
    Limits          string      `json:"limits"`
}

//--------------------------------------------------------------------------

type Trips []Trip

//--------------------------------------------------------------------------

type User struct {
    Sid             string      `json:"sid"`
    Username        string      `json:"username"`
    Email           string      `json:"email"`
    Name            string      `json:"name"`
    Location        Position    `json:"location"`
    DateCreated     string      `json:"dateCreated"`
}

//--------------------------------------------------------------------------

type Users []User

//--------------------------------------------------------------------------
