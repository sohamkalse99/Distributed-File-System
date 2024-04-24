package main

import (
	"bufio"
	"dfs/config"
	"dfs/handler/clientHandler"
	"dfs/handler/clientSNHandler"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func extractFileName(path string) string {
	fileName := filepath.Base(path)
	return fileName
}

func calcFileSize(fileName string) int64 {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err.Error())
	}

	info, _ := file.Stat()
	size := info.Size()
	defer file.Close()
	return size
}

func connectToController(controllerPort string) net.Conn {
	conn, connErr := net.Dial("tcp", config.ControllerHostName+":"+controllerPort)

	if connErr != nil {
		log.Fatalln(connErr.Error())
	}

	return conn
}

func sendFileDetails(handler *clientHandler.ClientHandler, action string, filePath string, chunkSize int64, fileSize int64) {
	clientMsg := &clientHandler.FileOpns{
		Action:    action,
		ChunkSize: chunkSize,
		FileSize:  fileSize,
	}
	handler.Send(clientMsg)
	fmt.Println("Sent details to controller")

}

func getStorageNodesDetails(handler *clientHandler.ClientHandler) []string {
	fileOpnsMsg, err := handler.Receive()

	if err != nil {
		log.Fatalln(err.Error())
	}

	dstSNList := fileOpnsMsg.GetDstSN()
	fmt.Println("received the active sn list", dstSNList)
	return dstSNList
}

func createChunks(filePath string, dstSNList []string, chunkSize int64, fileSize int64) map[string][][]byte {
	noOfChunks := len(dstSNList)

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		log.Fatalln(fileErr.Error())
	}

	defer file.Close()
	chunkSNMap := make(map[string][][]byte)
	// arr:= [][]byte{}
	for i := 0; i < noOfChunks; i++ {
		partSize := int(math.Min(float64(chunkSize), float64(fileSize-int64(i*int(chunkSize)))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		if arr, ok := chunkSNMap[dstSNList[i]]; !ok {
			arr = append(arr, partBuffer)
			chunkSNMap[dstSNList[i]] = arr
		} else {
			currArr := chunkSNMap[dstSNList[i]]
			currArr = append(currArr, partBuffer)
			chunkSNMap[dstSNList[i]] = currArr
		}
		/*file.Read(partBuffer)

		filename := "chunk_" + strconv.FormatUint(uint64(i), 10)

		_, createErr := os.Create(filename)

		if createErr != nil {
			log.Fatalln(createErr.Error())
			os.Exit(1)
		}

		os.WriteFile(filename, partBuffer, os.ModeAppend)
		fmt.Println("Split to : ", filename)*/
	}

	return chunkSNMap
}

func connectToSN(chunkSNMap map[string][][]byte, fileName string) {
	for key, value := range chunkSNMap {

		conn, err := net.Dial("tcp", key)

		if err != nil {
			log.Fatalln(err.Error())
		}

		handler := clientSNHandler.NewClientSNHandler(conn)
		msg := &clientSNHandler.ChunkDetails{
			FileName:   fileName,
			ChunkArray: value,
		}
		handler.Send(msg)

	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("message>")

		actionRes := scanner.Scan()

		if !actionRes {
			break
		}
		clientIp := scanner.Text()

		// Remove leading and trailing spaces as well as make it lower case
		if strings.ToLower(strings.Trim(clientIp, "")) == "put" {

			fmt.Println("Please enter chunk size")

			chRes := scanner.Scan()

			if !chRes {
				break
			}
			cS := scanner.Text()
			chunkSize, convErr := strconv.ParseInt(cS, 10, 64)
			if convErr != nil {
				log.Fatalln(convErr.Error())
			}
			controllerPort := config.ControllerPortForClient
			conn := connectToController(controllerPort)
			defer conn.Close()
			handler := clientHandler.NewClientHandler(conn)
			filePath := config.FilePath

			fileSize := calcFileSize(filePath)
			sendFileDetails(handler, "put", filePath, chunkSize, fileSize)
			dstSNList := getStorageNodesDetails(handler)
			chunkSNMap := createChunks(filePath, dstSNList, chunkSize, fileSize)
			// fmt.Println("Map of chunks", chunkSNMap)
			// uniqueSNList := getUniqueSN(dstSNList)
			fileName := extractFileName(filePath)
			connectToSN(chunkSNMap, fileName)
		}

	}

	/*_, err := net.Listen("tcp", ":"+os.Args[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	action := os.Args[2]
	// if put, will contain entire path else will just contain file name
	var chunkSize int64
	// var convErr error
	if len(os.Args) < 3 {
		chunkSize = config.ChunkSize
	} else {
		temp, convErr := strconv.Atoi(os.Args[3])
		if convErr != nil {
			log.Fatalln(convErr.Error())
		}
		chunkSize = int64(temp)
	}*/

}
