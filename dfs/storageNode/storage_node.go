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
	"sync"
	"syscall"
	"time"
)

var chunkNameArr []string
var replicaNameArr []string
var chunkDataArr [][]byte
var replicaChunkNameArr []string
var fellowSNNamesList []string

// var SNPortMap = make(map[string]string)

func calculateFreeSpace(storagePath string) (uint64, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs(storagePath, &stat)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return stat.Bavail * uint64(stat.Bsize), nil
}

func sendHeartbeat(handler *snHandler.StorageNodeHandler, storagePath string, snHostname string, snPortForClient string, snPortForFellowSN string) {
	// defer snHandler.Close()

	// Calculate Free space
	freeSpace, stErr := calculateFreeSpace(storagePath)
	if stErr != nil {
		log.Fatalln(stErr.Error())
	}

	// Create a Heartbeat wrapper, for now you can create something temporary

	heartbeatMsg := &snHandler.Heartbeat{StorageNodeName: snHostname, SpaceAvailability: freeSpace, StoragePortNumber: snPortForClient, SnFellowPN: snPortForFellowSN}

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

func sendRegReq(handler *snHandler.StorageNodeHandler, snHostname string, snPortForClient string, snPortForFellowSN string) {

	registrationMsg := &snHandler.Registration{StorageNodeName: snHostname, StoragePortNumber: snPortForClient, SnFellowPN: snPortForFellowSN}

	wrapperMsg := &snHandler.Wrapper{
		Task: &snHandler.Wrapper_RegTask{
			RegTask: registrationMsg,
		},
	}

	handler.Send(wrapperMsg)
}

func handleHeartbeat(snHandler *snHandler.StorageNodeHandler, snName string, snPortForClient string, snPortForFellowSN string, storagePath string, hbCounter *int) {

	// Send a heartbeat every 5 seconds
	// for {

	sendHeartbeat(snHandler, storagePath, snName, snPortForClient, snPortForFellowSN)
	if *hbCounter%10 == 0 {
		fmt.Println("Heartbeat sent")
	}
	*hbCounter = *hbCounter + 1

	time.Sleep(5 * time.Second)
	// }

}

func handleController(wg *sync.WaitGroup) {
	defer wg.Done()
	snName, hnErr := os.Hostname()
	snName = strings.Split(snName, ".")[0]
	if hnErr != nil {
		log.Fatalln(hnErr.Error())
	}
	snPort := os.Args[1]
	snPortForClient := os.Args[2]
	snPortForFellowSN := os.Args[3]

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
			sendRegReq(handler, snName, snPortForClient, snPortForFellowSN)

			wrapper, _ := handler.Receive()

			// Will receive an ok message if registration is successful
			if wrapper.GetRegTask().Status == "ok" {
				handleHeartbeat(handler, snName, snPortForClient, snPortForFellowSN, storagePath, &hbCounter)

			} else {
				fmt.Println("Register as new node")
			}
			flag = false
		} else {
			handleHeartbeat(handler, snName, snPortForClient, snPortForFellowSN, storagePath, &hbCounter)
		}

		conn.Close()
	}
}

func sendDataToFellowSN(fellowSNPortMap map[string]string, clientSNPortMap map[string]string) {
	// replicaChunkNameArr, chunkDataArr, fellowSNNamesList

	for _, sn := range fellowSNNamesList {

		if value, ok := fellowSNPortMap[sn]; ok {

			snNameAndPort := sn + ":" + value

			conn, err := net.Dial("tcp", snNameAndPort)
			if err != nil {
				fmt.Println("Inside sendDataToFellowSN")
				log.Fatalln(err.Error())
			}
			handler := snHandler.NewStorageNodeHandler(conn)

			wrapper := &snHandler.Wrapper{
				Task: &snHandler.Wrapper_FellowSNTask{
					FellowSNTask: &snHandler.FellowSNMsg{
						Data:            chunkDataArr,
						ReplicaNames:    replicaNameArr,
						ClientsnPortMap: clientSNPortMap,
						SnName:          sn,
					},
				},
			}
			handler.Send(wrapper)
		}

	}

}

func handleClientRequests(handler *clientSNHandler.ClientSNHandler, snPortForClient string) {
	defer handler.Close()
	chunkDetailsMsg, err := handler.Receive()

	if err != nil {
		log.Fatalln(err.Error())
	}
	// fileName := chunkDetailsMsg.GetFileName()
	action := chunkDetailsMsg.GetAction()
	chunkNameArr = chunkDetailsMsg.GetChunkNameArray()
	replicaNameArr = chunkDetailsMsg.GetReplicaNameArray()
	if action == "put" {
		chunkDataArr = chunkDetailsMsg.GetChunkDataArray()
		replicaChunkNameArr = chunkDetailsMsg.GetReplicaChunkNameArray()
		fellowSNNamesList = chunkDetailsMsg.GetFellow_SNNamesList()
		fellowSNPortMap := chunkDetailsMsg.GetFellowsnPortMap()
		clientSNPortMap := chunkDetailsMsg.GetClientsnPortMap()
		noOfChunks := len(chunkDataArr)

		// Send replicas to fellow SN on sepearte threads
		go sendDataToFellowSN(fellowSNPortMap, clientSNPortMap)

		for i := 0; i < noOfChunks; i++ {
			chunkFullName := fmt.Sprintf("%s/%s", config.StoragePath, chunkNameArr[i])

			// file, createErr := os.Create(fileName)
			writeErr := os.WriteFile(chunkFullName, chunkDataArr[i], 0644)

			if writeErr != nil {
				log.Fatalln(writeErr.Error())
			} else {
				fmt.Println("File Created: ", chunkNameArr[i])
			}
		}
	} else if action == "get" {
		/*dirEntry, err := os.ReadDir(config.StoragePath)

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
		}*/

		// sort the slice according to the digits which come at the end
		sort.Slice(chunkNameArr, func(i, j int) bool {
			splitStrOne := strings.Split(chunkNameArr[i], "_chunk_")
			numI, _ := strconv.Atoi(splitStrOne[1])

			splitStrTwo := strings.Split(chunkNameArr[j], "_chunk_")
			numJ, _ := strconv.Atoi(splitStrTwo[1])

			return numI < numJ
		})

		//read file into bytes
		chunkArr := [][]byte{}
		for _, element := range chunkNameArr {
			chunkName := fmt.Sprintf("%s/%s", config.StoragePath, element)
			byteArr, err := os.ReadFile(chunkName)

			if err != nil {
				log.Fatalln(err.Error())
			}

			chunkArr = append(chunkArr, byteArr)

		}
		clientGetMsg := &clientSNHandler.ChunkDetails{
			ChunkDataArray: chunkArr,
		}

		// Send chunks for the filename requested by the client
		handler.Send(clientGetMsg)

	}

}
func handleClient(wg *sync.WaitGroup) {
	defer wg.Done()
	snPortForClient := os.Args[2]

	clientListner, err := net.Listen("tcp", ":"+snPortForClient)
	fmt.Println("Listening for client on Port: ", snPortForClient)

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

func handleFellowSNRequests(handler *snHandler.StorageNodeHandler) {
	var snDataMap = make(map[string][][]byte)

	defer handler.Close()
	fellowSNMsg, err := handler.Receive()
	if err != nil {
		fmt.Println("Inside handle fellow err 2")

		log.Fatalln(err)
	}

	//Replica names and data
	replicaNameArr := fellowSNMsg.GetFellowSNTask().GetReplicaNames()
	chunkArr := fellowSNMsg.GetFellowSNTask().GetData()
	sn := fellowSNMsg.GetFellowSNTask().GetSnName()

	if _, ok := snDataMap[sn]; !ok {
		snDataMap[sn] = chunkArr
	}

	// fmt.Println("chunkNameslice", chunkNameSlice)
	// fmt.Println("length of chunk arr", len(chunkArr))

	for i := 0; i < len(replicaNameArr); i++ {

		chunkFullName := fmt.Sprintf("%s/%s", config.StoragePath, replicaNameArr[i])

		// file, createErr := os.Create(fileName)
		writeErr := os.WriteFile(chunkFullName, chunkArr[i], 0644)

		if writeErr != nil {
			fmt.Println("Inside handle fellow err 3")

			log.Fatalln(writeErr.Error())
		} else {
			fmt.Println("Replica Chunks Created: ", replicaNameArr[i])
		}
	}
}

func handleFellowSN(wg *sync.WaitGroup) {
	defer wg.Done()
	snPortForFellowSN := os.Args[3]

	felowSNListner, err := net.Listen("tcp", ":"+snPortForFellowSN)
	fmt.Println("Listening for fellow storage nodes on Port: ", snPortForFellowSN)

	if err != nil {
		fmt.Println("Inside handle fellow err")
		log.Fatalln(err.Error())
	}

	for {
		fmt.Println("Started an infinite loop to handle Fellow SN")
		if conn, connErr := felowSNListner.Accept(); connErr == nil {
			fmt.Println("Accepted a SN")
			handler := snHandler.NewStorageNodeHandler(conn)
			handleFellowSNRequests(handler)
		} else {
			fmt.Println("Inside err")
		}
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go handleController(&wg)
	go handleClient(&wg)

	//When you get the list/Map of storage nodes, then only you startlistening to fellow storage nodes

	go handleFellowSN(&wg)
	wg.Wait()
}
