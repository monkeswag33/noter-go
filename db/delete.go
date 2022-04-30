package db

func (db *DB) DeleteUser(conditions User) error {
	if err := db.DB.Where(&conditions).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteNote(conditions Note) error {
	if err := db.DB.Where(&conditions).Delete(&Note{}).Error; err != nil {
		return err
	}
	return nil
}
