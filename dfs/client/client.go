package main

import (
	"bufio"
	"crypto/sha512"
	"dfs/config"
	"dfs/handler/clientHandler"
	"dfs/handler/clientSNHandler"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var SNChunkArrMap = make(map[string][][]byte)
var uniqueSNMap = make(map[string]bool)
var indexMap = make(map[string]int)

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

func calcCheckSum(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	h := sha512.New()

	if _, shaErr := io.Copy(h, file); err != nil {
		log.Fatalln(shaErr)
	}

	return h.Sum(nil)

}
func connectToController(controllerPort string) net.Conn {
	conn, connErr := net.Dial("tcp", config.ControllerHostName+":"+controllerPort)

	if connErr != nil {
		log.Fatalln(connErr.Error())
	}

	return conn
}

func sendFileDetails(handler *clientHandler.ClientHandler, action string, filePath string, chunkSize int64, fileSize int64, fileName string, checkSum []byte) {
	clientMsg := &clientHandler.FileOpns{
		FileName:  fileName,
		Action:    action,
		ChunkSize: chunkSize,
		FileSize:  fileSize,
		Checksum:  checkSum,
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
	action := "put"
	for key, value := range chunkSNMap {

		conn, err := net.Dial("tcp", key)

		if err != nil {
			log.Fatalln(err.Error())
		}

		handler := clientSNHandler.NewClientSNHandler(conn)
		msg := &clientSNHandler.ChunkDetails{
			FileName:   fileName,
			ChunkArray: value,
			Action:     action,
		}
		handler.Send(msg)

	}
}
func sendGetMsg(fileName string, clientContHandler *clientHandler.ClientHandler) {
	getMsg := &clientHandler.FileOpns{
		FileName: fileName,
		Action:   "get",
	}

	clientContHandler.Send(getMsg)
}

func getSNListAndCheckSum(fileName string, clientContHandler *clientHandler.ClientHandler) ([]string, []byte) {
	controllerMsg, err := clientContHandler.Receive()

	if err != nil {
		log.Fatalln(err.Error())
	}
	checkSum := controllerMsg.GetChecksum()
	dstSNList := controllerMsg.GetDstSN()
	return dstSNList, checkSum
}

func getChunkFromSN(handler *clientSNHandler.ClientSNHandler, sn string, wg *sync.WaitGroup) {
	defer wg.Done()
	chunkDetails, err := handler.Receive()
	chunkArr := chunkDetails.GetChunkArray()
	if err != nil {
		log.Fatalln(err.Error())
	}

	SNChunkArrMap[sn] = chunkArr
}
func sendMsgToSN(fileName string, dstSNList []string) {
	action := "get"
	var wg sync.WaitGroup
	for sn := range uniqueSNMap {
		conn, err := net.Dial("tcp", sn)
		if err != nil {
			log.Fatalln(err.Error())
		}

		handler := clientSNHandler.NewClientSNHandler(conn)

		snMsg := &clientSNHandler.ChunkDetails{
			FileName: fileName,
			Action:   action,
		}

		handler.Send(snMsg)
		wg.Add(1)
		go getChunkFromSN(handler, sn, &wg)
	}
	wg.Wait()
}
func createFileFromChunks(fileName string, dstSNList []string, checkSum []byte) {
	// Initialize indexes to zero in indexMap
	for key := range SNChunkArrMap {
		indexMap[key] = 0
	}

	fileArr := []byte{}
	for _, element := range dstSNList {

		if values, ok := SNChunkArrMap[element]; ok && indexMap[element] < len(values) {
			fileArr = append(fileArr, values[indexMap[element]]...)
			indexMap[element]++
		}
	}

	// write to a file

	filePath := fmt.Sprintf("%s/%s", config.GetFolderPath, config.GetFileName)

	writeErr := os.WriteFile(filePath, fileArr, 0644)

	if writeErr != nil {
		fmt.Println("Error creating the file")
		log.Fatalln(writeErr.Error())
	} else {
		fmt.Println("File Created: ", filePath)
		cs := calcCheckSum(filePath)
		if string(cs) == string(checkSum) {
			fmt.Println("Original File's checksum and assembled file's checksum match")
		}
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
		filePath := config.FilePath
		fileName := extractFileName(filePath)
		controllerPort := config.ControllerPortForClient
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

			conn := connectToController(controllerPort)
			defer conn.Close()
			handler := clientHandler.NewClientHandler(conn)

			fileSize := calcFileSize(filePath)
			checkSum := calcCheckSum(filePath)
			sendFileDetails(handler, "put", filePath, chunkSize, fileSize, fileName, checkSum)
			dstSNList := getStorageNodesDetails(handler)
			chunkSNMap := createChunks(filePath, dstSNList, chunkSize, fileSize)
			// fmt.Println("Map of chunks", chunkSNMap)
			// uniqueSNList := getUniqueSN(dstSNList)
			connectToSN(chunkSNMap, fileName)
		} else if strings.ToLower(strings.Trim(clientIp, "")) == "get" {
			conn := connectToController(controllerPort)
			defer conn.Close()

			clientContHandler := clientHandler.NewClientHandler(conn)
			sendGetMsg(fileName, clientContHandler)

			dstSNList, checkSum := getSNListAndCheckSum(fileName, clientContHandler)

			// Create a unique map of SN's as you want to send requests to unique SN's
			for _, element := range dstSNList {
				if _, ok := uniqueSNMap[element]; !ok {
					uniqueSNMap[element] = true
				}
			}
			sendMsgToSN(fileName, dstSNList)
			createFileFromChunks(fileName, dstSNList, checkSum)
		}

		SNChunkArrMap = make(map[string][][]byte)
		uniqueSNMap = make(map[string]bool)
		indexMap = make(map[string]int)
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
