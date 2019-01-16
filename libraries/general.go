package libraries

import "log"

func CheckError(err error) bool {
	if err != nil {
		log.Println("Error on page : ", err.Error())

		return true
	}

	return false
}
