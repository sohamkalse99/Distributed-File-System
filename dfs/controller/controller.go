package main

import (
	"dfs/config"
	"dfs/handler/clientHandler"
	"dfs/handler/snHandler"
	"fmt"
	"log"
	"math"
	"net"
	"time"
)

// var activeSNSet = make(map[string]bool)
var snTimeMap = make(map[string]time.Time)

func checkSNValidity() {

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
				// activeSNSet[key] = false
			} else {
				fmt.Println(key, "is On")
			}

		}
		time.Sleep(5 * time.Second)

	}

}

func registerStorageNode(handler *snHandler.StorageNodeHandler, snName string, snPortNo string) {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)

	if formatErr != nil {
		log.Fatalln(formatErr)
	}
	// if element is present in snTimeMap, and the diff > 15sec then do not change the activSNSet
	element := snName + ":" + snPortNo

	if value, ok := snTimeMap[element]; ok {
		diff := currTimeFormatted.Sub(value)
		fmt.Println("Time diff", diff)

		if diff < 15*time.Second {
			// Add SNname to the set
			// activeSNSet[element] = true
			// Set current time to snTimeMap
			snTimeMap[element] = currTimeFormatted
			fmt.Println(element, " registered")
			// Send an ok message

			okMsg := &snHandler.Registration{Status: "ok"}

			wrapper := &snHandler.Wrapper{
				Task: &snHandler.Wrapper_RegTask{
					RegTask: okMsg,
				},
			}

			handler.Send(wrapper)
		} else {
			fmt.Println("Cannot register as time diff greater than 15 secs")
		}
	} else {
		// Registering for the first time, so key would not be there in the snTimeMap

		// Add SNname to the set
		// activeSNSet[element] = true

		// set current time in snTimeMap
		snTimeMap[element] = currTimeFormatted
		fmt.Println(element, " registered")
		// Send an ok message
		okMsg := &snHandler.Registration{Status: "ok"}

		wrapper := &snHandler.Wrapper{
			Task: &snHandler.Wrapper_RegTask{
				RegTask: okMsg,
			},
		}

		handler.Send(wrapper)
	}

}

func handleHeartbeat(handler *snHandler.StorageNodeHandler, wrapper *snHandler.Wrapper) {

	snName := wrapper.GetHeartbeatTask().StorageNodeName
	snPort := wrapper.GetHeartbeatTask().StoragePortNumber
	key := snName + ":" + snPort
	// if _, ok := activeSNSet[key]; ok {

	currTime := time.Now().Format("2006-01-02 15:04:05")
	currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)
	if formatErr != nil {
		log.Fatalln(formatErr)
	}
	if value, ok := snTimeMap[key]; ok {
		diff := currTimeFormatted.Sub(value)
		// fmt.Println("Time diff", diff)
		if diff > 15*time.Second {
			fmt.Println("Failure. Reinitalize yourself as new node")
		} else {
			fmt.Println("Success")
			snTimeMap[key] = currTimeFormatted
		}
	} else {
		fmt.Println("You need to register before sending the heartbeat")
	}

	// }

}

func handleStorageNodeRequests(handler *snHandler.StorageNodeHandler) {

	// listen to registration/heartbeat from a storage node
	// defer snHandler.Close()

	wrapper, _ := handler.Receive()

	switch task := wrapper.Task.(type) {

	case *snHandler.Wrapper_RegTask:
		snName := task.RegTask.StorageNodeName
		snPortNo := task.RegTask.StoragePortNumber
		registerStorageNode(handler, snName, snPortNo)
		// handleStorageNode(snHandler, snTimeMap, activeSNSet)
	case *snHandler.Wrapper_HeartbeatTask:
		handleHeartbeat(handler, wrapper)
	case nil:
		log.Println("Got a empty message. Disconnecting Storage Node")
		return
	default:
		log.Printf("Unexpected message type %T", task)
	}
}

func handleStorageNode() {
	snListner, snListnerErr := net.Listen("tcp", ":"+config.ControllerPortForSN)

	if snListnerErr != nil {
		log.Fatalln(snListnerErr.Error())
	}

	// Used to store SN name and latest time of heartbeat

	go checkSNValidity()

	for {
		fmt.Println("Started an infinite loop to handle SN")
		if conn, connErr := snListner.Accept(); connErr == nil {
			handler := snHandler.NewStorageNodeHandler(conn)
			handleStorageNodeRequests(handler)
		}

	}
}

func getDestinationSN(noOfChunks int64) []string {

	currTime := time.Now().Format("2006-01-02 15:04:05")
	currTimeFormatted, formatErr := time.Parse("2006-01-02 15:04:05", currTime)
	if formatErr != nil {
		log.Fatalln(formatErr)
	}
	// i:=0
	keyList := []string{}
	for key, value := range snTimeMap {
		diff := currTimeFormatted.Sub(value)
		if diff < 15*time.Second {
			// count++
			keyList = append(keyList, key)
		}

	}
	// keyList := make([]string, len(activeSNSet))
	dstSNList := make([]string, noOfChunks)
	/*i := 0
	for key, value := range activeSNSet {
		if value {
			keyList[i] = key
			i++
		}
	}*/

	for i := 0; i < int(noOfChunks); i++ {
		dstSNList[i] = keyList[i%len(keyList)]
	}

	return dstSNList
}

func handleClientPutReq(chunkSize int64, fileSize int64) []string {
	// fmt.Println("file size", fileSize)
	// fmt.Println("chunk size", chunkSize)

	noOfChunks := int64(math.Ceil(float64(fileSize) / float64(chunkSize)))
	// fmt.Println("no of chunks", noOfChunks)

	dstSNList := getDestinationSN(noOfChunks)

	return dstSNList
}

func handleClientRequests(handler *clientHandler.ClientHandler) {
	defer handler.Close()

	fileOpnsMsg, receiveErr := handler.Receive()
	if receiveErr != nil {
		log.Fatalln(receiveErr.Error())
	}

	action := fileOpnsMsg.GetAction()
	chunkSize := fileOpnsMsg.GetChunkSize()
	fileSize := fileOpnsMsg.GetFileSize()
	if action == "put" {
		dstSNList := handleClientPutReq(chunkSize, fileSize)

		dstSNListmsg := &clientHandler.FileOpns{
			DstSN: dstSNList,
		}
		fmt.Println("Dest SN list", dstSNList)
		handler.Send(dstSNListmsg)
	}
}
func handleClient() {
	clientListner, clientListnerErr := net.Listen("tcp", ":"+config.ControllerPortForClient)

	if clientListnerErr != nil {
		log.Fatalln(clientListnerErr.Error())
	}

	for {
		fmt.Println("Started an infinite loop to handle Client")

		if conn, connErr := clientListner.Accept(); connErr == nil {
			fmt.Println("Accepted client")
			handler := clientHandler.NewClientHandler(conn)
			handleClientRequests(handler)
		}
	}
}

func main() {

	// Keeps list of storage nodes which are active

	go handleStorageNode()
	handleClient()

}
