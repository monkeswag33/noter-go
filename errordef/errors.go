package errordef // Error definitions

import "errors"

var ErrNoteNameTooShort error = errors.New("note name must be at least 5 characters long")
var ErrNoteAlreadyExists error = errors.New("note already exists")
var ErrUserDoesntExist error = errors.New("user doesn't exist")
