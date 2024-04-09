package main

import (
	"fmt"
	"os"
	"time"
)

func sender(message string) {
	file, err := os.OpenFile("storage.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, char := range message {
		file.WriteString(string(char))
		time.Sleep(500 * time.Millisecond)
	}
}

func receiver() {
	file, err := os.Open("storage.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Received message:")
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	go sender("new message")

	receiver()

	select {}
}
