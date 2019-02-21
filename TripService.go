/**
 * File containing the TripService methods.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
    "net/http"
    "github.com/gorilla/mux"
)

//--------------------------------------------------------------------------

func GetTrips(w http.ResponseWriter, r *http.Request) {
    trips, err := GetDAO(DAO_TYPE_TRIP).(TripDAO).GetTrips()
    if err != nil {
    }

    dispatchResponseObject(w, trips, 0)
}

//--------------------------------------------------------------------------

func GetTrip(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    tripId := vars["tripId"]

    trip, err := GetDAO(DAO_TYPE_TRIP).(TripDAO).GetTrip(tripId)
    if err != nil {
    }

    dispatchResponse(w, trip, 0)
}

//--------------------------------------------------------------------------

func AddTrip(w http.ResponseWriter, r *http.Request) {
    tripId, err := GetDAO(DAO_TYPE_TRIP).(TripDAO).AddTrip()
    if err != nil {
    }

    dispatchResponse(w, tripId, 0)
}

//--------------------------------------------------------------------------
