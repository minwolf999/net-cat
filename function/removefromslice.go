package function

// Remove an element for the Users array
func RemoveUserFromSlice(user User) {
	for i := range Users {
		if Users[i] == user {
			if i+1 < len(Users) {
				Users = append(Users[:i], Users[i+1:]...)
			} else {
				Users = Users[:i]
			}
		}
	}
}