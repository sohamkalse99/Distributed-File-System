package main

import (
	"dfs/config"
	"dfs/handler/clientSNHandler"
	"dfs/handler/snHandler"
	"fmt"
	"log"
	"net"
	"os"
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

func handleHeartbeat(snHandler *snHandler.StorageNodeHandler, snName string, snPortForClient string, storagePath string) {

	// Send a heartbeat every 5 seconds
	// for {

	sendHeartbeat(snHandler, storagePath, snName, snPortForClient)
	fmt.Println("Heartbeat sent")
	time.Sleep(5 * time.Second)
	// }

}

func handleController() {
	snName, hnErr := os.Hostname()

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

	// TODO : Calculate number of requests processed (storage/retrievals)
	// calculateRequestCount()
	flag := true
	for {
		conn := connectToController(controllerPort, snName, snPortForClient, storagePath)

		handler := snHandler.NewStorageNodeHandler(conn)
		if flag {
			sendRegReq(handler, snName, snPortForClient)

			wrapper, _ := handler.Receive()

			// Will receive an ok message if registration is successful
			if wrapper.GetRegTask().Status == "ok" {
				handleHeartbeat(handler, snName, snPortForClient, storagePath)
			} else {
				fmt.Println("Register as new node")
			}
			flag = false
		} else {
			handleHeartbeat(handler, snName, snPortForClient, storagePath)
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
	fileArrList := chunkDetailsMsg.GetChunkArray()
	noOfChunks := len(fileArrList)

	for i := 0; i < noOfChunks; i++ {
		fileName := fmt.Sprintf("%s/chunk_%s", config.StoragePath, snPortForClient)

		// file, createErr := os.Create(fileName)
		writeErr := os.WriteFile(fileName, fileArrList[i], 0644)

		if writeErr != nil {
			log.Fatalln(writeErr.Error())
		} else {
			fmt.Println("File Created: ", fileName)
		}
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
