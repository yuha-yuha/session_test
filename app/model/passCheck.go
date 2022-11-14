package model

import "app/database"

func PassCheck(Uid string, UPass string) bool {
	user, err := database.GetUserData(Uid)

	if err != nil {
		return false
	}

	if user.UPass != UPass {
		return false
	}

	return true
}
