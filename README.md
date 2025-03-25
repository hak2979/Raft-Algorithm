# **Raft-Based Key-Value Store**  

This project implements a **distributed key-value store** using the **Raft consensus algorithm**. It ensures fault tolerance and consistency by handling **leader elections** and **log replication** across multiple nodes.  

## **📌 Features**  
- **Put(Key, Value)** – Stores a key-value pair in the system.  
- **Append(Key, Value)** – Appends a value to an existing key.  
- **Get(Key)** – Retrieves the value associated with a key.  

## **🔗 Raft Implementation**  
This project focuses on two core components of Raft:  
1. **Leader Election** – Ensures only one leader is active at a time and correctly handles elections.  
2. **Log Replication** – Maintains a consistent and fault-tolerant log across all nodes.  

## **⚡ Fault Tolerance**  
- Handles **leader failures** by electing a new leader automatically.  
- Recovers **worker nodes** by syncing logs upon reconnection.  
- Implements **timeout and retries** to ensure network reliability.  

## **📂 File Structure**  
```
📦 raft-kv-store
├── 📂 shared             # Shared utilities and helper functions
├── 📄 main.go            # Entry point of the application
├── 📄 raft_node.go       # Raft node implementation (leader & follower)
├── 📄 client.go          # Client implementation to interact with the system
├── 📄 go.mod             # Go module file
```

## **🚀 Running the Project**  
1. **Initialize Go module**  
   ```sh
   go mod init github.com/yourusername/raft-kv-store
   go mod tidy
   ```
2. **Run multiple nodes in separate terminals**  
   ```sh
   go run raft_node.go 5001  # Node 1
   go run raft_node.go 5002  # Node 2
   go run raft_node.go 5003  # Node 3
   ```
3. **Interact with the key-value store using the client**  
   ```sh
   go run client.go
   ```

## **📖 References**  
- [Raft Paper](https://raft.github.io/raft.pdf)  
- [Raft Consensus Algorithm Explained](https://thesecretlivesofdata.com/raft/)  

