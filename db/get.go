package db

func (db *DB) GetUsers(conditions User) ([]User, error) {
	var users []User
	if err := db.DB.Where(&conditions).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) GetNotes(conditions Note) ([]Note, error) {
	var notes []Note
	if err := db.DB.Preload("User").Where(&conditions).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
