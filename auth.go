package main

func checkCredentials(username, password string) bool {
	return checkLocal(username, password) || checkLdap(username, password)
}

func checkLocal(username, password string) bool {
	return username == "steve" && password == "istesting"
}

func checkLdap(username, password string) bool {
	return ldapLogin(username, password)
}
