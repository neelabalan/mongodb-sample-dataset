// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// type LogEntry struct {
// 	Timestamp string        `json:"timestamp"`
// 	Message   string        `json:"message"`
// 	Duplicate ObjectIdEntry `json:"duplicate"`
// }

// type ObjectIdEntry struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// func main() {
// 	// Open the MongoDB log file
// 	filePath := "mongo.log"
// 	file, err := os.Open("mongo.log")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Create a slice to store log entries
// 	var logEntries []LogEntry

// 	// Create a scanner to read the file line by line
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Split the line by tab to separate timestamp and message
// 		parts := strings.Split(line, "\t")
// 		if len(parts) != 2 {
// 			continue // Skip lines that don't match the expected format
// 		}

// 		entry := LogEntry{
// 			Timestamp: parts[0],
// 			Message:   parts[1],
// 		}
// 		logEntries = append(logEntries, entry)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(logEntries)

// 	regex := regexp.MustCompile(`([^:]+): ObjectId\('([^']+)'\)`)

// 	// Extract key-value pairs containing ObjectId from log entries
// 	var objectIdEntries []ObjectIdEntry
// 	for _, v := range logEntries {
// 		if strings.Contains(v.Message, "_id_") {
// 			matches := regex.FindAllStringSubmatch(v.Message, -1)
// 			for _, match := range matches {
// 				if len(match) >= 3 {
// 					entry := ObjectIdEntry{
// 						Key:   match[1],
// 						Value: match[2],
// 					}
// 					v.Duplicate.Key = entry.Key
// 					v.Duplicate.Value = entry.Value
// 					objectIdEntries = append(objectIdEntries, entry)
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(objectIdEntries)

// 	// Convert log entries to JSON
// 	jsonData, err := json.Marshal(logEntries)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = os.WriteFile(filePath, jsonData, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print the JSON data
// 	// fmt.Println(string(jsonData))
// }



// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// type LogEntry struct {
// 	Timestamp string        `json:"timestamp"`
// 	Message   string        `json:"message"`
// 	Duplicate []ObjectIdEntry `json:"duplicate,omitempty"`
// }

// type ObjectIdEntry struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// func main() {
// 	// Open the MongoDB log file
// 	filePath := "mongo.log"
// 	file, err := os.Open("mongo.log")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Create a slice to store log entries
// 	var logEntries []LogEntry

// 	// Create a scanner to read the file line by line
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Split the line by tab to separate timestamp and message
// 		parts := strings.Split(line, "\t")
// 		if len(parts) != 2 {
// 			continue // Skip lines that don't match the expected format
// 		}

// 		entry := LogEntry{
// 			Timestamp: parts[0],
// 			Message:   parts[1],
// 		}

// 		// Initialize the Duplicate field as an empty slice
// 		entry.Duplicate = []ObjectIdEntry{}

// 		logEntries = append(logEntries, entry)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	regex := regexp.MustCompile(`([^:]+): ObjectId\('([^']+)'\)`)

// 	// Extract key-value pairs containing ObjectId from log entries
// 	for i := range logEntries {
// 		if strings.Contains(logEntries[i].Message, "_id_") {
// 			matches := regex.FindAllStringSubmatch(logEntries[i].Message, -1)
// 			for _, match := range matches {
// 				if len(match) >= 3 {
// 					entry := ObjectIdEntry{
// 						Key:   match[1],
// 						Value: match[2],
// 					}
// 					logEntries[i].
// 					logEntries[i].Duplicate = append(logEntries[i].Duplicate, entry)
// 				}
// 			}
// 		}
// 	}
// 	// fmt.Println(logEntries)

// 	// Convert log entries to JSON
// 	jsonData, err := json.Marshal(logEntries)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = os.WriteFile(filePath, jsonData, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print the JSON data
// 	// fmt.Println(string(jsonData))
// }



package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
)

type LogEntry struct {
	Timestamp string          `json:"timestamp"`
	Message   string          `json:"message"`
	Duplicate []ObjectIdEntry `json:"duplicate,omitempty"`
}

type ObjectIdEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	// Open the MongoDB log file
	filePath := "mongo.log"
	file, err := os.Open("mongo.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a slice to store log entries
	var logEntries []LogEntry

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by tab to separate timestamp and message
		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			continue // Skip lines that don't match the expected format
		}

		entry := LogEntry{
			Timestamp: parts[0],
			Message:   parts[1],
		}

		// Initialize the Duplicate field as an empty slice
		entry.Duplicate = []ObjectIdEntry{}

		logEntries = append(logEntries, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	regex := regexp.MustCompile(`([^:]+): ObjectId\('([^']+)'\)`)
	// regex2 := regexp.MustCompile(`continuing through error: ([^:]+): ObjectId\('([^']+)'\)`)

	// Extract key-value pairs containing ObjectId from log entries
	for i := range logEntries {
		if strings.Contains(logEntries[i].Message, "_id_") {

			matches := regex.FindAllStringSubmatch(logEntries[i].Message, -1)
			for _, match := range matches {
				if len(match) >= 3 {
					entry := ObjectIdEntry{
						Key:   match[1],
						Value: match[2],
					}
					logEntries[i].Duplicate = append(logEntries[i].Duplicate, entry)
				}
			}
		}
	}
	// fmt.Println(logEntries)

	// Convert log entries to JSON
	jsonData, err := json.Marshal(logEntries)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	// fmt.Println(string(jsonData))
}
