/**
 * File containing the database actions that are related to users.
 * Author: _sb
 */

package main

//--------------------------------------------------------------------------

import (
     "database/sql"
)

import _ "github.com/go-sql-driver/mysql"

//--------------------------------------------------------------------------

type UserDAO struct {
    SqlConn     *SqlConnection
    TableName   string
}

func UserDAOInit(dao *UserDAO) bool {
    db, err := sql.Open("mysql", dao.SqlConn.Username + ":" + dao.SqlConn.Password + "@/" + dao.SqlConn.Database)
}

//--------------------------------------------------------------------------

func (dao UserDAO) GetUsers() (Users, error) {
    if dao.SqlConn == nil || dao.SqlConn._Connection == "" {
        // init
    }

    users := Users{}

    return users, nil
}

//--------------------------------------------------------------------------

func (dao UserDAO) GetUser(userId string) (string, error) {
    return string("User show:" + userId), nil
}

//--------------------------------------------------------------------------

func (dao UserDAO) AddUser() (string, error) {
    return GenerateSecureId()
}

//--------------------------------------------------------------------------

func (dao UserDAO) UpdateUser() (string, error) {
    return GenerateSecureId()
}

//--------------------------------------------------------------------------

func (dao UserDAO) DeleteUser() (bool, error) {
    return true, nil
}

//--------------------------------------------------------------------------
