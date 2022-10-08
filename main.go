package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/reiver/go-telnet"
)

const TelnetAddress = "localhost:2121"
const TriggerText1 = "Player: "
const TriggerText2 = " - Damage Taken"
const ArgumentError = 1
const ConnectionError = 2
const ReadingError = 3

func TelnetReadUntilSequence(conn *telnet.Conn, sequence string) (out string) {
	var buffer [1]byte
	recvData := buffer[:]
	var n int
	var err error

	for {
		n, err = conn.Read(recvData)
		if err != nil {
			fmt.Println("[!] Disconnected")
			os.Exit(ReadingError)
		}

		if n > 0 {
			out += string(recvData)
		}

		if strings.Contains(out, sequence) {
			break
		}
	}

	return out
}

func main() {
	args := os.Args[1:]

	// checks
	if len(args) < 3 {
		fmt.Println("csgo-death-trigger")
		fmt.Println("=================")
		fmt.Println("Sends a POST to {url} with {json} when {nickname} dies in CS:GO.")
		fmt.Println("\nUsage: ./csgo-death-trigger.exe {nickname} {url} {json}")
		fmt.Println("Example: ./csgo-death-trigger.exe \"s1mple\" \"http://127.0.0.1:3000\" \"{ \"a_value\": 123 }\"")
		fmt.Println("\nMake sure you run the game with: -condebug -netconport 2121")
		os.Exit(ArgumentError)
	}

	// parameters
	nickname := args[0]
	url := args[1]
	body := args[2]

	// connection
	conn, err := telnet.DialTo(TelnetAddress)
	if err != nil {
		fmt.Println("[!] Telnet connection failed: " + TelnetAddress)
		os.Exit(ConnectionError)
	}
	fmt.Println("[!] Connected")

	// trigger
	for {
		line := TelnetReadUntilSequence(conn, "\n")
		triggerText := TriggerText1 + nickname + TriggerText2

		if strings.Contains(line, triggerText) {
			_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
			if err == nil {
				fmt.Println("[!!!] " + triggerText)
			} else {
				fmt.Println("[ERROR] " + triggerText)
			}
		}
	}
}
