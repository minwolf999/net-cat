package function

// Look if the username a user want to use ise already used in the channel
func NameAlreadyUsedInChannel(username string, channel string) bool {
	for _, u := range Users {
		if u.Channel != channel {
			continue
		}

		if u.Username == username {
			return false
		}
	}
	return true
}