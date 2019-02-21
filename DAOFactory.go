/**
 * DAO factory, makes sure only maximum one instance of each exists.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
)

//--------------------------------------------------------------------------

const (
        DAO_TYPE_TRIP = iota
        DAO_TYPE_USER
        DAO_TYPE_PLACE
)

//--------------------------------------------------------------------------

var DAOList map[int]interface{}

//--------------------------------------------------------------------------

func initDAOFactory() bool {
    DAOList = make(map[int]interface{})
    return true
}

//--------------------------------------------------------------------------

func GetDAO(daoType int) interface{} {
    switch daoType {
    case DAO_TYPE_TRIP:
        if _, ok := DAOList[DAO_TYPE_TRIP]; ok {
            return DAOList[DAO_TYPE_TRIP]
        } else {
            DAOList[DAO_TYPE_TRIP] = TripDAO {
                nil,
                "trip",
            }

            return DAOList[DAO_TYPE_TRIP]
        }

    case DAO_TYPE_USER:
        if _, ok := DAOList[DAO_TYPE_USER]; ok {
            return DAOList[DAO_TYPE_USER]
        } else {
            DAOList[DAO_TYPE_USER] = UserDAO {
                nil,
                "user",
            }

            return DAOList[DAO_TYPE_USER]
        }
    default:
        return nil
    }
}

//--------------------------------------------------------------------------
