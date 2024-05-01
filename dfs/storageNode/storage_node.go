package main

import (
	"dfs/config"
	"dfs/handler/clientSNHandler"
	"dfs/handler/snHandler"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func calculateFreeSpace(storagePath string) (uint64, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs(storagePath, &stat)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return stat.Bavail * uint64(stat.Bsize), nil
}

func sendHeartbeat(handler *snHandler.StorageNodeHandler, storagePath string, snHostname string, snPortForClient string) {
	// defer snHandler.Close()

	// Calculate Free space
	freeSpace, stErr := calculateFreeSpace(storagePath)
	if stErr != nil {
		log.Fatalln(stErr.Error())
	}

	// Create a Heartbeat wrapper, for now you can create something temporary

	heartbeatMsg := &snHandler.Heartbeat{StorageNodeName: snHostname, SpaceAvailability: freeSpace, StoragePortNumber: snPortForClient}

	wrapperMsg := &snHandler.Wrapper{
		Task: &snHandler.Wrapper_HeartbeatTask{
			HeartbeatTask: heartbeatMsg,
		},
	}

	// Send Heartbeat message
	handler.Send(wrapperMsg)
}

func connectToController(controllerPort string, snHostname string, snPortForClient string, storagePath string) net.Conn {
	// defer conn.Close()
	conn, err := net.Dial("tcp", config.ControllerHostName+":"+controllerPort)

	if err != nil {
		log.Fatalln(err.Error())
	}
	// fmt.Println("Before conn")
	return conn

}

func sendRegReq(handler *snHandler.StorageNodeHandler, snHostname string, snPortForClient string) {

	registrationMsg := &snHandler.Registration{StorageNodeName: snHostname, StoragePortNumber: snPortForClient}

	wrapperMsg := &snHandler.Wrapper{
		Task: &snHandler.Wrapper_RegTask{
			RegTask: registrationMsg,
		},
	}

	handler.Send(wrapperMsg)
}

func handleHeartbeat(snHandler *snHandler.StorageNodeHandler, snName string, snPortForClient string, storagePath string, hbCounter *int) {

	// Send a heartbeat every 5 seconds
	// for {

	sendHeartbeat(snHandler, storagePath, snName, snPortForClient)
	if *hbCounter%10 == 0 {
		fmt.Println("Heartbeat sent")
	}
	*hbCounter = *hbCounter + 1

	time.Sleep(5 * time.Second)
	// }

}

func handleController() {
	snName, hnErr := os.Hostname()
	snName = strings.Split(snName, ".")[0]
	if hnErr != nil {
		log.Fatalln(hnErr.Error())
	}
	snPort := os.Args[1]
	snPortForClient := os.Args[2]
	_, err := net.Listen("tcp", ":"+snPort)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// /bigdata/students/skalse
	storagePath := config.StoragePath

	controllerPort := config.ControllerPortForSN
	// chunkNames := getChunkDets(config.StoragePath)
	// TODO : Calculate number of requests processed (storage/retrievals)
	// calculateRequestCount()
	flag := true
	hbCounter := 0
	for {
		conn := connectToController(controllerPort, snName, snPortForClient, storagePath)

		handler := snHandler.NewStorageNodeHandler(conn)

		if flag {
			sendRegReq(handler, snName, snPortForClient)

			wrapper, _ := handler.Receive()

			// Will receive an ok message if registration is successful
			if wrapper.GetRegTask().Status == "ok" {
				handleHeartbeat(handler, snName, snPortForClient, storagePath, &hbCounter)

			} else {
				fmt.Println("Register as new node")
			}
			flag = false
		} else {
			handleHeartbeat(handler, snName, snPortForClient, storagePath, &hbCounter)
		}

		conn.Close()
	}
}

func handleClientRequests(handler *clientSNHandler.ClientSNHandler, snPortForClient string) {
	defer handler.Close()
	chunkDetailsMsg, err := handler.Receive()

	if err != nil {
		log.Fatalln(err.Error())
	}
	fileName := chunkDetailsMsg.GetFileName()
	action := chunkDetailsMsg.GetAction()

	if action == "put" {
		chunkArrList := chunkDetailsMsg.GetChunkArray()
		noOfChunks := len(chunkArrList)
		for i := 0; i < noOfChunks; i++ {
			chunkName := fmt.Sprintf("%s/%s_chunk_%d", config.StoragePath, fileName, i)

			// file, createErr := os.Create(fileName)
			writeErr := os.WriteFile(chunkName, chunkArrList[i], 0644)

			if writeErr != nil {
				log.Fatalln(writeErr.Error())
			} else {
				fmt.Println("File Created: ", chunkName)
			}
		}
	} else if action == "get" {
		dirEntry, err := os.ReadDir(config.StoragePath)

		if err != nil {
			log.Fatalln(err.Error())
		}
		chunkNames := []string{}
		chunkArr := [][]byte{}
		// traverse through files get only those file related chunks and add it to slice chunkNames
		for _, entry := range dirEntry {
			entryStr := entry.Name()
			isChunkPresent := strings.Contains(entryStr, fileName)
			if isChunkPresent {
				chunkNames = append(chunkNames, entryStr)
			}
		}

		// sort the slice according to the digits which come at the end
		sort.Slice(chunkNames, func(i, j int) bool {
			splitStrOne := strings.Split(chunkNames[i], "_chunk_")
			numI, _ := strconv.Atoi(splitStrOne[1])

			splitStrTwo := strings.Split(chunkNames[j], "_chunk_")
			numJ, _ := strconv.Atoi(splitStrTwo[1])

			return numI < numJ
		})

		//read file into bytes
		for _, element := range chunkNames {
			fileName := fmt.Sprintf("%s/%s", config.StoragePath, element)
			byteArr, err := os.ReadFile(fileName)

			if err != nil {
				log.Fatalln(err.Error())
			}

			chunkArr = append(chunkArr, byteArr)

		}
		clientGetMsg := &clientSNHandler.ChunkDetails{
			ChunkArray: chunkArr,
		}

		// Send chunks for the filename requested by the client
		handler.Send(clientGetMsg)

	}

}
func handleClient() {
	snPortForClient := os.Args[2]

	clientListner, err := net.Listen("tcp", ":"+snPortForClient)

	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		fmt.Println("Started an infinite loop to handle Client")
		if conn, connErr := clientListner.Accept(); connErr == nil {
			fmt.Println("Accepted a client")
			handler := clientSNHandler.NewClientSNHandler(conn)
			handleClientRequests(handler, snPortForClient)
		}
	}

}
func main() {

	go handleController()
	handleClient()

}
