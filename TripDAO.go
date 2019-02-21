/**
 * File containing the database actions that are related to trips.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
)

//--------------------------------------------------------------------------

type TripDAO struct {
    SqlConn     *SqlConnection
    TableName   string
}

//--------------------------------------------------------------------------

func (dao TripDAO) GetTrips() (Trips, error) {
    if dao.SqlConn == nil || dao.SqlConn._Connection == "" {
        // init
    }

    trips := Trips{
        Trip{
            StartPosition: Position{Latitude: 45, Longitude: 26},
            PositionList: []Position{
                {Latitude: 45.01, Longitude: 26.02},
                {Latitude: 45.05, Longitude: 26.06},
            },
            Limits: "none",
        },
    }

    return trips, nil
}

//--------------------------------------------------------------------------

func (dao TripDAO) GetTrip(tripId string) (string, error) {
    return string("Trip show:" + tripId), nil
}

//--------------------------------------------------------------------------

func (dao TripDAO) AddTrip() (string, error) {
    return GenerateSecureId()
}

//--------------------------------------------------------------------------
