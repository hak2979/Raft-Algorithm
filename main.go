package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"block/raft" // Import the local raft package
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <nodeID> <port> [peerPort1] [peerPort2] ...")
		os.Exit(1)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Parse node ID and port
	nodeID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid node ID: %v\n", err)
		os.Exit(1)
	}

	// Use a different seed for each node
	rand.Seed(time.Now().UnixNano() + int64(nodeID))

	port := os.Args[2]
	address := fmt.Sprintf("localhost:%s", port)

	// Set up peer addresses - FIXED: create the map with proper size
	peerAddresses := make(map[int]string)
	
	// Add self to peer list first
	peerAddresses[nodeID] = address
	
	// Add peers with sequential IDs
	peerID := 1
	for i := 3; i < len(os.Args); i++ {
		peerPort := os.Args[i]
		
		// Skip our own ID when assigning peers
		if peerID == nodeID {
			peerID++
		}
		
		peerAddresses[peerID] = fmt.Sprintf("localhost:%s", peerPort)
		peerID++
	}

	fmt.Printf("Starting node %d at %s with peers: %v\n", nodeID, address, peerAddresses)

	// Create key-value store
	kvStore := raft.NewKeyValueStore()

	// Create and start Raft node
	node := raft.NewRaftNode(nodeID, peerAddresses, kvStore)
	err = node.Start(address)
	if err != nil {
		fmt.Printf("Failed to start node: %v\n", err)
		os.Exit(1)
	}

	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown signal
	<-c
	fmt.Println("Shutting down...")
	node.Stop()
	fmt.Println("Node stopped")
}