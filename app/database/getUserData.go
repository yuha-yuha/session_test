package database

func GetUserData(Id string) (User, error) {
	db := DBConnect()

	user := User{}

	err := db.Debug().Where("user_id = ?", Id).First(&user).Error

	return user, err
}
