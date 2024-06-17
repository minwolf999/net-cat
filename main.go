package main

import (
	"fmt"
	"net"
	"net-cat2/function"
	"os"
	"strings"
	"time"
)

// This function create the logs files for all the servers
func init() {
	err := os.Mkdir("log/"+strings.ReplaceAll(time.Now().Format("2006-01-02 15:04:05"), " ", "_"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 10; i++ {
		namefile := fmt.Sprintf("log/%v/%v.log", strings.ReplaceAll(time.Now().Format("2006-01-02 15:04:05"), " ", "_"), i)
		file, err := os.Create(namefile)
		if err != nil {
			fmt.Println(err)
			return
		}

		function.LogFiles = append(function.LogFiles, file)
	}
}

func main() {
	if len(os.Args) == 2 {
		function.Port = ":" + os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	
	var ip string

	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	fmt.Printf("Started at ip: %s%s\n", ip, function.Port)

	ln, err := net.Listen("tcp", function.Port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		usrConn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go function.HandleConnection(usrConn)
	}
}
