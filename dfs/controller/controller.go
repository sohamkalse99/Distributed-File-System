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
var fileSNMap = make(map[string][]string)                                      //Initialized during put operation, used during get operation. Used to map file name and orions
var fileSNChunkMap = make(map[string]map[string]*clientHandler.FileOpnsChunks) // {temp:{orion01:[temp_chunk_1, temp_chunk_2]}}
var snChunkMap = make(map[string]*clientHandler.FileOpnsChunks)                //Stores storage node names and list of chunks
var checkSum []byte

// Replication
var fellowsnSNMap = make(map[string]*clientHandler.FileOpnsFellowSNNames)
var replicaFileSNReplicaChunks = make(map[string]map[string]*clientHandler.FileOpnsReplicaList)
var replicaSNChunkMap = make(map[string]*clientHandler.FileOpnsReplicaList) // will map storage node name and the replica chunks
var snSNSet = make(map[string]bool)
var fellowSNPortMap = make(map[string]string)
var clientSNPortMap = make(map[string]string)
var replicaNameMap = make(map[string]*clientHandler.FileOpnsReplicaName) //Stores storage node names and list of names replica chunks, used for put

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
	element := snName
	// element2 := snName
	if value, ok := snTimeMap[element]; ok {
		diff := currTimeFormatted.Sub(value)
		fmt.Println("Time diff", diff)

		if diff < 15*time.Second {
			// Add SNname to the set
			// activeSNSet[element] = true
			// Set current time to snTimeMap
			snTimeMap[element] = currTimeFormatted
			snSNSet[element] = true
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
			snSNSet[element] = false
			fmt.Println("Cannot register as time diff greater than 15 secs")
		}
	} else {
		// Registering for the first time, so key would not be there in the snTimeMap

		// Add SNname to the set
		// activeSNSet[element] = true

		// set current time in snTimeMap
		snTimeMap[element] = currTimeFormatted
		snSNSet[element] = true
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
	snPortForFellowSN := wrapper.GetHeartbeatTask().SnFellowPN

	// Check if fellowSNport exist in the map for the orion name
	if _, ok := fellowSNPortMap[snName]; !ok {
		fellowSNPortMap[snName] = snPortForFellowSN
	}

	// Check if clientSNport exist in the map for the orion name
	if _, ok := clientSNPortMap[snName]; !ok {
		clientSNPortMap[snName] = snPort
	}

	// chunkNames := wrapper.GetHeartbeatTask().ChunkNames
	key := snName
	// key2 := snName
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
			snSNSet[key] = false
		} else {
			// fmt.Println("Success")
			snTimeMap[key] = currTimeFormatted
			snSNSet[key] = true
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
		// fmt.Println("Started an infinite loop to handle SN")
		if conn, connErr := snListner.Accept(); connErr == nil {
			handler := snHandler.NewStorageNodeHandler(conn)
			handleStorageNodeRequests(handler)
		}

	}
}
func getActiveSNList() []string {
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

	return keyList
}

func getDestinationSN(noOfChunks int64) []string {

	activeSNList := getActiveSNList()
	dstSNList := make([]string, noOfChunks)

	for i := 0; i < int(noOfChunks); i++ {
		dstSNList[i] = activeSNList[i%len(activeSNList)]
	}

	return dstSNList
}

func mapFileSNAndChunks(fileName string, noOfChunks int64, activeSNList []string, chunkSlice []string) {
	if _, ok := fileSNChunkMap[fileName]; !ok {
		fileSNChunkMap[fileName] = snChunkMap
	}
}

func getChunkSlice(fileName string, noOfChunks int64) []string {
	chunkSlice := make([]string, noOfChunks)

	for i := 0; i < int(noOfChunks); i++ {
		chunkName := fmt.Sprintf("%s_chunk_%d", fileName, (i + 1))
		chunkSlice[i] = chunkName
	}
	return chunkSlice
}

func createSNChunkMap(dstSNList []string, chunkSlice []string, handler *clientHandler.ClientHandler) {

	for i := 0; i < len(chunkSlice); i++ {
		if values, ok := snChunkMap[dstSNList[i]]; ok {
			values.ChunkList = append(values.ChunkList, chunkSlice[i])

			snChunkMap[dstSNList[i]] = values
		} else {
			arr := []string{}
			arr = append(arr, chunkSlice[i])

			//add list which of type fileopnschunk to the map
			snChunkMap[dstSNList[i]] = &clientHandler.FileOpnsChunks{ChunkList: arr}
		}

	}

}

func createReplicaChunkNames(chunkSlice []string, dstSNList []string) {
	replicaNameSlice := make([]string, len(chunkSlice))
	for i := range chunkSlice {
		replicaNameSlice[i] = fmt.Sprintf("replica_%s", chunkSlice[i])
		// fmt.Println(replicaNameSlice[i])
	}

	for i := 0; i < len(replicaNameSlice); i++ {
		if values, ok := replicaNameMap[dstSNList[i]]; ok {
			values.ReplicaNameList = append(values.ReplicaNameList, replicaNameSlice[i])

			replicaNameMap[dstSNList[i]] = values
		} else {
			arr := []string{}
			arr = append(arr, replicaNameSlice[i])
			//add list which of type fileopnschunk to the map
			replicaNameMap[dstSNList[i]] = &clientHandler.FileOpnsReplicaName{ReplicaNameList: arr}
		}

	}

}
func createSNMapping() {
	list := []string{}

	for val := range snSNSet {
		list = append(list, val)
	}
	for i, value := range list {
		neighbors := []string{}

		neighbors = append(neighbors, list[(i+1)%len(list)])

		if len(list) > 2 {
			neighbors = append(neighbors, list[(i+2)%len(list)])
		}

		// snChunkMap[dstSNList[i]] = &clientHandler.FileOpnsChunks{ChunkList: arr}

		fellowsnSNMap[value] = &clientHandler.FileOpnsFellowSNNames{Fellow_SNNamesList: neighbors}
		// replicasnSNMap[value].Fellow_SNNamesList = neighbors
	}
}

func createSNReplicaMapping(fileName string) {

	for key, snList := range fellowsnSNMap {
		if innerMap, ok := fileSNChunkMap[fileName]; ok {
			if innerList, ok := innerMap[key]; ok {
				arr := []string{}
				if len(snList.Fellow_SNNamesList) == 2 {
					for _, element := range innerList.ChunkList {
						replicaChunkName := fmt.Sprintf("replica_%s", element)
						arr = append(arr, replicaChunkName)
					}

					if val1, ok := replicaSNChunkMap[snList.Fellow_SNNamesList[0]]; ok {
						val1.ReplicaChunkList = append(val1.ReplicaChunkList, arr...)
						replicaSNChunkMap[snList.Fellow_SNNamesList[0]] = val1
					} else {
						replicaSNChunkMap[snList.Fellow_SNNamesList[0]] = &clientHandler.FileOpnsReplicaList{ReplicaChunkList: arr}
					}

					if val2, ok := replicaSNChunkMap[snList.Fellow_SNNamesList[1]]; ok {
						val2.ReplicaChunkList = append(val2.ReplicaChunkList, arr...)
						replicaSNChunkMap[snList.Fellow_SNNamesList[1]] = val2
					} else {
						replicaSNChunkMap[snList.Fellow_SNNamesList[1]] = &clientHandler.FileOpnsReplicaList{ReplicaChunkList: arr}
					}

				}
			}
		}
	}

	if _, ok := replicaFileSNReplicaChunks[fileName]; !ok {
		replicaFileSNReplicaChunks[fileName] = replicaSNChunkMap
	}
}
func handleClientPutReq(chunkSize int64, fileSize int64, fileName string, handler *clientHandler.ClientHandler) []string {
	// fmt.Println("file size", fileSize)
	// fmt.Println("chunk size", chunkSize)

	noOfChunks := int64(math.Ceil(float64(fileSize) / float64(chunkSize)))
	// fmt.Println("no of chunks", noOfChunks)
	activeSNList := getActiveSNList()

	dstSNList := getDestinationSN(noOfChunks)

	chunkSlice := getChunkSlice(fileName, noOfChunks)
	createSNChunkMap(dstSNList, chunkSlice, handler)
	mapFileSNAndChunks(fileName, noOfChunks, activeSNList, chunkSlice)

	//Get Replication Data
	createSNMapping()
	createSNReplicaMapping(fileName)
	createReplicaChunkNames(chunkSlice, dstSNList)

	return dstSNList
}

func handleClientRequests(handler *clientHandler.ClientHandler) {
	defer handler.Close()

	fileOpnsMsg, receiveErr := handler.Receive()
	if receiveErr != nil {
		log.Fatalln(receiveErr.Error())
	}
	fileName := fileOpnsMsg.GetFileName()
	action := fileOpnsMsg.GetAction()

	if action == "put" {
		chunkSize := fileOpnsMsg.GetChunkSize()
		fileSize := fileOpnsMsg.GetFileSize()
		checkSum = fileOpnsMsg.GetChecksum()
		dstSNList := handleClientPutReq(chunkSize, fileSize, fileName, handler)

		// add to the map
		fileSNMap[fileName] = dstSNList

		handler.Send(&clientHandler.FileOpns{DstSNList: dstSNList, SnChunkMap: snChunkMap, ReplicaSNChunkMap: replicaSNChunkMap, FellowSNNamesMap: fellowsnSNMap, FellowsnPortMap: fellowSNPortMap, ClientsnPortMap: clientSNPortMap, ReplicaNameMap: replicaNameMap})
	} else if action == "get" {
		if SNChunkMapValue, ok := fileSNChunkMap[fileName]; ok {
			if dstSNListValue, ok2 := fileSNMap[fileName]; ok2 {
				msg := &clientHandler.FileOpns{
					Checksum:        checkSum,
					SnChunkMap:      SNChunkMapValue,
					DstSNList:       dstSNListValue,
					ClientsnPortMap: clientSNPortMap,
				}
				handler.Send(msg)
			}

		} else {
			// TODO : Probabily need to send a message to the client with the below message
			fmt.Println("File not present, please insert a file")
		}
	}
}
func handleClient() {
	clientListner, clientListnerErr := net.Listen("tcp", ":"+config.ControllerPortForClient)

	if clientListnerErr != nil {
		log.Fatalln(clientListnerErr.Error())
	}

	for {
		// fmt.Println("Started an infinite loop to handle Client")

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
