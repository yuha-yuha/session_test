package model

import (
	"app/database"
	"log"
)

func IsLoggedIn(Id interface{}) (bool, string) {

	if Id != nil {
		idString, _ := Id.(string)
		_, err := database.GetUserData(idString)

		if err != nil {
			log.Println(err)
			return false, ""
		}

	} else {
		log.Println("Id:", Id)
		return false, ""
	}

	return true, Id.(string)

}
