package main

import (
	"dfs/handler"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func checkSNValidity(snTimeMap map[string]time.Time, activeSNSet map[string]bool) {

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
				activeSNSet[key] = false
			} else {
				fmt.Println(key, "is On")
			}

		}

		time.Sleep(5 * time.Second)

	}

}

func registerStorageNode(snHandler *handler.StorageNodeHandler, activeSNSet map[string]bool, snTimeMap map[string]time.Time, snName string, snPortNo string) {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)

	if formatErr != nil {
		log.Fatalln(formatErr)
	}
	// if element is present in snTimeMap, and the diff > 15sec then do not change the activSNSet
	element := snName + ":" + snPortNo

	if value, ok := snTimeMap[element]; ok {
		diff := currTimeFormatted.Sub(value)
		if diff < 15*time.Second {
			// Add SNname to the set
			activeSNSet[element] = true

			fmt.Println(element, " registered")
			// Send an ok message
			okMsg := &handler.Registration{Status: "ok"}

			wrapper := &handler.Wrapper{
				Task: &handler.Wrapper_RegTask{
					RegTask: okMsg,
				},
			}

			snHandler.Send(wrapper)
		} else {
			fmt.Println("Cannot register as time diff greater than 15 secs")
		}
	} else {
		// Registering for the first time, so key would not be there in the snTimeMap

		// Add SNname to the set
		activeSNSet[element] = true

		fmt.Println(element, " registered")
		// Send an ok message
		okMsg := &handler.Registration{Status: "ok"}

		wrapper := &handler.Wrapper{
			Task: &handler.Wrapper_RegTask{
				RegTask: okMsg,
			},
		}

		snHandler.Send(wrapper)
	}

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
		registerStorageNode(snHandler, activeSNSet, snTimeMap, snName, snPortNo)
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

	// Keeps list of storage nodes which are active
	activeSNSet := make(map[string]bool)

	go checkSNValidity(snTimeMap, activeSNSet)

	for {
		fmt.Println("Started an infinite loop")
		if conn, connErr := listner.Accept(); connErr == nil {
			snHandler := handler.NewStorageNodeHandler(conn)
			handleStorageNode(snHandler, snTimeMap, activeSNSet)

		}

	}
}
