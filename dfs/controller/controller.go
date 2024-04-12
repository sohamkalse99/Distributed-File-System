package main

import (
	"dfs/handler"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func checkSNValidity(snTimeMap map[string]time.Time) {

	for {
		currTime := time.Now().Format("2006-01-02 15:04:05")
		currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)

		if formatErr != nil {
			log.Fatalln(formatErr)
		}

		for key, value := range snTimeMap {

			diff := currTimeFormatted.Sub(value)
			if diff > 15*time.Second {
				fmt.Println(key, "is Off", "Reinitalize yourself as new node")
			} else {
				fmt.Println(key, "is On")
			}

		}
		time.Sleep(5 * time.Second)
	}

}

func registerStorageNode(snHandler *handler.StorageNodeHandler, activeSNSet map[string]bool, snName string, snPortNo string) {

	// Add SNname to the set
	element := snName + ":" + snPortNo
	activeSNSet[element] = true

	// Send an ok message
	okMsg := &handler.Registration{Status: "ok"}

	wrapper := &handler.Wrapper{
		Task: &handler.Wrapper_RegTask{
			RegTask: okMsg,
		},
	}

	snHandler.Send(wrapper)
}

func handleHeartbeat(snHandler *handler.StorageNodeHandler, wrapper *handler.Wrapper, activeSNSet map[string]bool, snTimeMap map[string]time.Time) {

	snName := wrapper.GetHeartbeatTask().StorageNodeName
	snPort := wrapper.GetHeartbeatTask().StoragePortNumber
	key := snName + ":" + snPort
	if _, ok := activeSNSet[key]; ok {

		currTime := time.Now().Format("2006-01-02 15:04:05")
		currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)
		if formatErr != nil {
			log.Fatalln(formatErr)
		}
		if value, ok := snTimeMap[key]; ok {
			diff := currTimeFormatted.Sub(value)
			fmt.Println("Time diff", diff)
			if diff > 15*time.Second {
				fmt.Println("Failure. Reinitalize yourself as new node")
			} else {
				fmt.Println("Success")
				snTimeMap[key] = currTimeFormatted
			}
		} else {
			snTimeMap[key] = currTimeFormatted
		}

	}

}

func handleStorageNode(snHandler *handler.StorageNodeHandler, snTimeMap map[string]time.Time, activeSNSet map[string]bool) {

	// listen to registration/heartbeat from a storage node
	// defer snHandler.Close()

	wrapper, _ := snHandler.Receive()

	switch task := wrapper.Task.(type) {

	case *handler.Wrapper_RegTask:
		snName := task.RegTask.StorageNodeName
		snPortNo := task.RegTask.StoragePortNumber
		registerStorageNode(snHandler, activeSNSet, snName, snPortNo)
		fmt.Println(snName + ":" + snPortNo + " registered")
		// handleStorageNode(snHandler, snTimeMap, activeSNSet)
	case *handler.Wrapper_HeartbeatTask:
		handleHeartbeat(snHandler, wrapper, activeSNSet, snTimeMap)
	case nil:
		log.Println("Got a empty message. Disconnecting Storage Node")
		return
	default:
		log.Printf("Unexpected message type %T", task)
	}
}
func main() {
	listner, err := net.Listen("tcp", ":"+os.Args[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	// Used to store SN name and latest time of heartbeat
	snTimeMap := make(map[string]time.Time)

	go checkSNValidity(snTimeMap)

	// Keeps list of storage nodes which are active
	activeSNSet := make(map[string]bool)
	for {
		fmt.Println("Started an infinite loop")
		if conn, connErr := listner.Accept(); connErr == nil {
			snHandler := handler.NewStorageNodeHandler(conn)
			handleStorageNode(snHandler, snTimeMap, activeSNSet)

		}

	}
}
