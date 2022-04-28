package db

func DeleteUser(username string) User {
	var user User
	db.First(&user, "username = ?", username)
	db.Delete(&user)
	return user
}

func DeleteNote(noteName string) Note {
	var note Note
	db.First(&note, "name = ?", noteName)
	db.Delete(&note)
	return note
}
