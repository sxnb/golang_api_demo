/**
 * File containing the route definitions.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

type Routes []Route

//--------------------------------------------------------------------------

var routes = Routes{
    Route {
        "Index",
        "GET",
        "/",
        index,
    },

    // TripService
    Route {
        "GetTrips",
        "GET",
        "/trip",
        GetTrips,
    },
    Route {
        "GetTrip",
        "GET",
        "/trip/{tripId}",
        GetTrip,
    },
    Route {
        "AddTrip",
        "PUT",
        "/trip",
        AddTrip,
    },

    // UserService
    Route {
        "GetUsers",
        "GET",
        "/user",
        GetUsers,
    },
    Route {
        "GetUser",
        "GET",
        "/user/{userId}",
        GetUser,
    },
    Route {
        "AddUser",
        "PUT",
        "/user",
        AddUser,
    },
    Route {
        "UpdateUser",
        "PUT",
        "/user/{userId}",
        UpdateUser,
    },
    Route {
        "AddUser",
        "DELETE",
        "/user/{userId}",
        DeleteUser,
    },
}

//--------------------------------------------------------------------------
