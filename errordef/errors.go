package errordef // Error definitions

import "errors"

var ErrUsernameTooShort error = errors.New("username must be at least 4 characters")
var ErrNoteNameTooShort error = errors.New("note name must be at least 5 characters long")
var ErrPasswordTooShort error = errors.New("password must be at least 8 characters")

var ErrNoteAlreadyExists error = errors.New("note already exists")
var ErrUserAlreadyExists error = errors.New("user already exists")

var ErrUserDoesntExist error = errors.New("user doesn't exist")
var ErrNoteDoesntExist error = errors.New("note doesn't exist")

var ErrUsernameMustContainAlphaNumeric error = errors.New("username can only contain alphanumeric characters")

var ErrArgon2InvalidHash error = errors.New("hash is invalid")
var ErrArgon2IncompatibleVersion error = errors.New("argon2 version is incompatible")

var ErrCouldNotFindPostgresURI error = errors.New("could not find postgres uri environment variable")
var ErrFailedToConnect error = errors.New("failed to connect to database")
