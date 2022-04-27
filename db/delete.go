package db

func DeleteUser(username string) User {
	var user User
	db.First(&user, "username = ?", username)
	db.Delete(&user)
	return user
}
