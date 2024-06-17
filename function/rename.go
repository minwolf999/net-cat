package function

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Change the username of a user and alert all user in the same channel
func Rename(user User) User {
retry:

	newname := make([]byte, 1024)

	user.UserConn.Write([]byte(ClearTerminal + "Write your new name: "))
	user.UserConn.Read(newname)
	user.UserConn.Write([]byte(ClearTerminal))

	message := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + user.Username + " became " + strings.Split(string(newname), "\n")[0] + "\n"

	if !NameAlreadyUsedInChannel(strings.Split(string(newname), "\n")[0], user.Channel) {
		goto retry
	}

	intChannel, _ := strconv.Atoi(user.Channel)

	datas, _ := os.ReadFile(Folder + "/" + strconv.Itoa(intChannel) + ".log")

	if len(datas) != 0 {
		FormateData := strings.Split(string(datas), "[")[1]

		user.UserConn.Write([]byte("[" + FormateData))
	}

	LogFiles[intChannel-1].WriteString(user.UserConn.RemoteAddr().String() + message)

	for i := range Users {
		if Users[i] == user {
			Users[i].UserConn.Write([]byte(message))

			user.Username = strings.Split(string(newname), "\n")[0]
			Users[i].Username = strings.Split(string(newname), "\n")[0]
		} else if Users[i].Channel == user.Channel {
			Users[i].UserConn.Write([]byte(message))
		}
	}

	return user
}
