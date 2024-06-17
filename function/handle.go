package function

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Handle the connection page (for get the username and the channel who connectther user)
func HandleConnection(usrConn net.Conn) {
	username := make([]byte, 1024)
	channel := make([]byte, 1024)

	for {
		usrConn.Write([]byte("Write an username: "))
		usrConn.Read(username)

		usrConn.Write([]byte("Select your channel (1 to 10): "))
		usrConn.Read(channel)

		tmp := strings.Split(string(channel), "\n")[0]
		intChannel, err := strconv.Atoi(tmp)

		if err != nil || intChannel < 1 || intChannel > 10 {
			usrConn.Write([]byte(ClearTerminal + "The channel choose isn't a good number\n"))
		} else if !NameAlreadyUsedInChannel(strings.Split(string(username), "\n")[0], strings.Split(string(channel), "\n")[0]) {
			usrConn.Write([]byte(ClearTerminal + "The username is already used in this channel\n"))
		} else if ChannelFull(tmp) {
			usrConn.Write([]byte(ClearTerminal + "The channel selected is full\n"))
		} else {
			break
		}
	}

	usrConn.Write([]byte(ClearTerminal))

	var user = User{
		Username: strings.Split(string(username), "\n")[0],
		Channel:  strings.Split(string(channel), "\n")[0],
		UserConn: usrConn,
	}

	Users = append(Users, user)
	HandleChat(user)
}

// Look if a channel is full
func ChannelFull(channel string) bool {
	i := 0
	for _, u := range Users {
		if u.Channel == channel {
			i++
		}
	}

	return i >= 10
}

// Handle the connection between the users connected in the same server and the message send in the chat
func HandleChat(user User) {
	entryMsg := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + "Say hello to the newcomers " + user.Username + "\n"

	msg := make([]byte, 1024)

	for _, u := range Users {
		if u.Channel != user.Channel {
			continue
		}

		if u == user {
			intChannel, _ := strconv.Atoi(user.Channel)

			datas, err := os.ReadFile(Folder + "/" + strconv.Itoa(intChannel) + ".log")
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(datas) > 0 {
				formateData := strings.Split(string(datas), "\n")
				for i := range formateData {
					if !strings.Contains(formateData[i], "[") {
						continue
					}

					formateData[i] = strings.Split(formateData[i], "[")[1]
				}

				u.UserConn.Write([]byte("[" + strings.Join(formateData, "\n")))
			}
		}

		u.UserConn.Write([]byte("\033[38;5;46m" + entryMsg + "\033[38;5;255;255;255m"))
	}

	intChannel, _ := strconv.Atoi(user.Channel)
	LogFiles[intChannel-1].WriteString(user.UserConn.RemoteAddr().String() + entryMsg)

	for {
		_, err := user.UserConn.Read(msg)
		if err != nil {
			RemoveUserFromSlice(user)
			SayByeToOther(user)
			break
		}

		if strings.Split(string(msg), "\n")[0] == "" {
			user.UserConn.Write([]byte("\033[1A\033[K"))
			continue
		}

		if strings.Split(string(msg), "\n")[0] == "/switch" {
			user.UserConn.Write([]byte(ClearTerminal))

			RemoveUserFromSlice(user)
			SayByeToOther(user)
			HandleConnection(user.UserConn)
			return
		} else if strings.Split(string(msg), "\n")[0] == "/rename" {
			user = Rename(user)
			continue
		} else if strings.Split(string(msg), "\n")[0] == "/exit" {

			user.UserConn.Write([]byte(ClearTerminal))
			user.UserConn.Write([]byte("You are logout. Press Enter to quit the program\n"))
			user.UserConn.Close()

			RemoveUserFromSlice(user)
			SayByeToOther(user)
			return
		}

		message := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + user.Username + ": " + strings.Split(string(msg), "\n")[0] + "\n"

		intChannel, _ := strconv.Atoi(user.Channel)
		LogFiles[intChannel-1].WriteString(user.UserConn.RemoteAddr().String() + message)

		for _, u := range Users {
			if u.Channel != user.Channel {
				continue
			}

			if u.Username == user.Username {
				u.UserConn.Write([]byte("\033[1A\033[K" + "\033[38;5;14m" + message + "\033[38;5;255;255;255m"))
			} else {
				u.UserConn.Write([]byte(message))
			}
		}
	}
}
