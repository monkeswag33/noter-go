package db

import (
	"errors"
)

func CreateUser(username string, password string) (User, error) {
	var user User = User{
		Username: username,
		Password: password,
	}
	if err := db.Create(&user).Error; err != nil {
		return User{}, errors.New("User already exists")
	}
	return user, nil
}

func CreateNote(name string, body string, username string) (Note, error) {
	var user User
	db.First(&user, "username = ?", username)
	var note Note = Note{
		Name: name,
		Body: body,
		User: user,
	}
	db.Create(&note)
	return note, nil
}

func CheckUserExists(username string) bool {
	var users []User
	db.Find(&users, "username = ?", username)
	if len(users) == 1 {
		return true
	} else {
		return false
	}
}

func CheckNoteExists(noteName string) bool {
	var notes []Note
	db.Find(&notes, "name = ?", noteName)
	if len(notes) == 1 {
		return true
	} else {
		return false
	}
}
