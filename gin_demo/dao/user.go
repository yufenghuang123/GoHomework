package dao

// data access object

var database = map[string]string{
	"gobee":"123456",
	"gobee2":"456789",
}

func AddUser(username string,password string)  {
	database[username]=password
}

func SelectUser(username string) bool {
	if database[username]==""{
		return false
	}
	return true
}

func SelectUserPasswordUsername(username string) string {
	return database[username]
}







