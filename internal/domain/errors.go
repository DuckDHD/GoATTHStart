package domain

import "errors"

var (

	/***

	****    User Creadentials Errors   ****

	***/

	ErrUserCredAalreadyExists = errors.New("user already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrWrongCredentials       = errors.New("email or password is wrong")
	ErrFailedTocreateCreds    = errors.New("failed to create user credentials")

	/***

	****    User Errors   ****

	***/

)
