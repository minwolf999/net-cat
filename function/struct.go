package function

import (
	"net"
	"os"
	"strings"
	"time"
)

var (
	Port          string = ":8080"
	ClearTerminal string = "\x1bc"
	Users         []User
	LogFiles      []*os.File
	Folder        string = "log/" + strings.ReplaceAll(time.Now().Format("2006-01-02 15:04:05"), " ", "_")
)

type User struct {
	Username string
	Channel  string

	UserConn net.Conn
}
