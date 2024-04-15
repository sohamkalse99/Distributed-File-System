package main

import (
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

func sendHeartbeat(handler *snHandler.StorageNodeHandler, storagePath string, snHostname string, snPort string) {
	// defer snHandler.Close()

	// Calculate Free space
	freeSpace, stErr := calculateFreeSpace(storagePath)
	if stErr != nil {
		log.Fatalln(stErr.Error())
	}

	// Create a Heartbeat wrapper, for now you can create something temporary

	heartbeatMsg := &snHandler.Heartbeat{StorageNodeName: snHostname, SpaceAvailability: freeSpace, StoragePortNumber: snPort}

	wrapperMsg := &snHandler.Wrapper{
		Task: &snHandler.Wrapper_HeartbeatTask{
			HeartbeatTask: heartbeatMsg,
		},
	}

	// Send Heartbeat message
	handler.Send(wrapperMsg)
}

func connectToController(controllerPort string, snHostname string, snPort string, storagePath string) net.Conn {
	// defer conn.Close()

	conn, err := net.Dial("tcp", controllerPort)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return conn

}

func sendRegReq(handler *snHandler.StorageNodeHandler, snHostname string, snPort string) {

	registrationMsg := &snHandler.Registration{StorageNodeName: snHostname, StoragePortNumber: snPort}

	wrapperMsg := &snHandler.Wrapper{
		Task: &snHandler.Wrapper_RegTask{
			RegTask: registrationMsg,
		},
	}

	handler.Send(wrapperMsg)
}

func handleHeartbeat(snHandler *snHandler.StorageNodeHandler, snName string, snPort string, storagePath string) {

	// Send a heartbeat every 5 seconds
	// for {

	sendHeartbeat(snHandler, storagePath, snName, snPort)
	fmt.Println("Heartbeat sent")
	time.Sleep(5 * time.Second)
	// }

}
func main() {
	snName, hnErr := os.Hostname()

	if hnErr != nil {
		log.Fatalln(hnErr.Error())
	}
	snPort := os.Args[1]
	_, err := net.Listen("tcp", ":"+snPort)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// /bigdata/students/skalse
	storagePath := os.Args[2]

	controllerPort := os.Args[3]

	// TODO : Calculate number of requests processed (storage/retrievals)
	// calculateRequestCount()
	flag := true
	for {
		conn := connectToController(controllerPort, snName, snPort, storagePath)

		handler := snHandler.NewStorageNodeHandler(conn)
		if flag {
			sendRegReq(handler, snName, snPort)

			wrapper, _ := handler.Receive()

			// Will receive an ok message if registration is successful
			if wrapper.GetRegTask().Status == "ok" {
				handleHeartbeat(handler, snName, snPort, storagePath)
			}
			flag = false
		} else {
			handleHeartbeat(handler, snName, snPort, storagePath)
		}

		conn.Close()
	}
}
