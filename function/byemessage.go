package function

import (
	"strconv"
	"time"
)


// Write when a user quit a channel for all the users present in the channel and in the log file
func SayByeToOther(user User) {
	message := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + user.Username + " has left the channel" + "\n"

	intChannel, _ := strconv.Atoi(user.Channel)
	LogFiles[intChannel-1].WriteString(user.UserConn.LocalAddr().String() + message)

	for _, u := range Users {
		if u.Channel == user.Channel {
			u.UserConn.Write([]byte("\033[38;5;88m" + message + "\033[38;5;255;255;255m"))
		}
	}
}
