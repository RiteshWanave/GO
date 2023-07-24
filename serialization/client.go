package main
import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const CBUFFERSIZE = 1024

func main () {
	connection, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Println("Connected to server, start receiving the file name and file size")
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64
	for {
		if (fileSize - receivedBytes) < CBUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+CBUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, CBUFFERSIZE)
		receivedBytes += CBUFFERSIZE
	}
}