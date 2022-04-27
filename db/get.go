package db

func GetUsers(username string, id int) []User {
	var whereClause User
	if len(username) > 0 {
		whereClause.Username = username
	}
	if id > 0 {
		whereClause.ID = id
	}
	var users []User
	db.Where(&whereClause).Find(&users)
	return users
}

func GetNotes(owner string, id int, name string) []Note {
	var whereClause Note
	var user User
	if len(owner) > 0 {
		db.Find(&user, "username = ?", owner)
		whereClause.UserID = user.ID
	}
	if id > 0 {
		whereClause.ID = id
	}
	if len(name) > 0 {
		whereClause.Name = name
	}
	var notes []Note
	db.Preload("User").Where(&whereClause).Find(&notes)
	return notes
}
