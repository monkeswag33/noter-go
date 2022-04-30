package db

func (db *DB) CreateUser(user *User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateNote(note *Note) error {
	var user User
	if err := db.DB.First(&user, "username = ?", note.User.Username).Error; err != nil {
		return err
	}
	note.User = user
	if err := db.DB.Create(note).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) CheckUserExists(conditions User) (bool, error) {
	var count int64
	if err := db.DB.Model(&User{}).Where(&conditions).Count(&count).Error; err != nil {
		return false, err
	}
	if count == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (db *DB) CheckNoteExists(conditions Note) (bool, error) {
	var count int64
	if err := db.DB.Model(&Note{}).Where(&conditions).Count(&count).Error; err != nil {
		return false, err
	}
	if count == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
