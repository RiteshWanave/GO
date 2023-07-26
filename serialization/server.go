package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const BUFFERSIZE = 1024   // Bytes

func fillString (retunString string, toLength int) string {
	for {
		lengthString := len(retunString)
		if (lengthString < toLength) {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func sendFileToClient (connection net.Conn) {
	println("A client has connected!")
	defer connection.Close()
	file, err := os.Open("test.txt")
	if err != nil {
		println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	println("Sending filename and filesize!")
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	println("Start sending file!")
	for {
		data, err := file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		println(data)
		connection.Write(sendBuffer)
	}
	println("File has been sent, closing connection!")
	return
}

func main () {
	server, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error While Listening: ", err)
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Server started! Waiting for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		println("Client Connected")
		go sendFileToClient (connection)
	}
}