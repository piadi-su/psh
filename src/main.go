
package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"encoding/json"
	
)

type Host struct {
    Name string `json:"name"`
    Host string `json:"host"`
    User string `json:"user"`
    Port int    `json:"port"`
}


func print_args(){
	
	fmt.Println("")
	fmt.Println("usage: ")
	fmt.Println("	psh <args>.")
	fmt.Println("add")
	fmt.Println("	add a new host.")
	fmt.Println("list")
	fmt.Println("	list all the knowun host.")
	fmt.Println("conncect <hostname>")
	fmt.Println("	connect to the host.")
	fmt.Println("remove <hostname>")
	fmt.Println("	remove the host form the list.")
	fmt.Println("")
}	

func add_host(){
		
	var new_host Host

	// sta merda crea un lettore per avere tipo un input() di py
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("hostname: ")
	hostname, _ := reader.ReadString('\n')
    hostname = strings.TrimSpace(hostname)

	fmt.Print("user: ")
	user, _ := reader.ReadString('\n')
    user = strings.TrimSpace(user)

	fmt.Print("ip: ")
	ip, _ := reader.ReadString('\n')
    ip = strings.TrimSpace(ip)

	fmt.Print("port: ")
	port , _ := reader.ReadString('\n')
    port = strings.TrimSpace(port)

	
	fmt.Printf("ip: %s\n", ip)
	fmt.Printf("hostname: %s\n", hostname)
	fmt.Printf("user: %s\n", user)
	fmt.Printf("port: %s\n", port)
}


//TODO add list 

func main() {
    // fmt.Println(len(os.Args))

	if len(os.Args) < 2 {
		print_args()
		return
	} 

	if os.Args[1] == "add" {
		add_host()
	}
	

	if os.Args[1] == "list" {
		fmt.Println("connecting...")
	}
	


}

