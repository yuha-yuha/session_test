package database

func GetUserData(Id int) (User, error) {
	db := DBConnect()

	user := User{}

	err := db.Debug().Where("id = ?", Id).First(&user).Error

	return user, err
}
