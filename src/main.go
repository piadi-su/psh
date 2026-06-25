package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

)



type Host struct {
    Hostname string `json:"hostname"`
    User string `json:"user"`
    Ip string   `json:"ip"`
    Port string `json:"port"`
}

const FILE_PERM = 0644



//---

func getFileDir() (string, error) {

	var configPath string

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath = home + "/.config/psh/hosts.json"

	return configPath, nil
}


func getList () ([]Host, error){

	path , err := getFileDir()
	if err != nil {
		return nil, err
	}

	rawData, err := os.ReadFile(path)

	if(os.IsNotExist(err)){
		return []Host{}, nil
	}

	if err != nil {
		return nil, err
	}

	var hosts_list []Host

	err = json.Unmarshal(rawData, &hosts_list)
	if err != nil{
		return nil, err
	}

	return hosts_list, nil
}


func saveHost(hosts_list []Host) error {

	path , err := getFileDir()
	if err != nil {
		return err
	}

	rawJson, err := json.MarshalIndent(hosts_list, "", "  ")
	if err != nil {
		return err
	}
	
	//check if it gives an error
	return os.WriteFile(path, rawJson, FILE_PERM)

}

//----


//psh list
func listHost() error {
	
	hosts_list, err := getList()
	if err != nil {
		return err
	}
	
	
	fmt.Println("====== HOSTS LIST ======")
	if len(hosts_list) == 0 {
		fmt.Println("\n       Empty list")

	}else{

		for _, host := range(hosts_list){
			fmt.Printf("%v ->  %v@%v  -p%v\n", host.Hostname, host.User, host.Ip, host.Port)
		}
	}


	fmt.Println()

	return nil
}


// psh remove, rm <hostname>
func removeHost(hostname string ) error {

	hosts_list, err := getList()

	if err != nil{
		return err
	}

	var new_hosts_list [] Host

	for _, host := range(hosts_list){
		if host.Hostname != hostname{
			new_hosts_list = append(new_hosts_list, host)
		}
	}

	return saveHost(new_hosts_list)
}


// psh add, a
func addHost() error {
		
	var new_host Host

	// sta merda crea un lettore per avere tipo un input() di py
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("hostname: ")
	new_host.Hostname, _ = reader.ReadString('\n')
    new_host.Hostname = strings.TrimSpace(new_host.Hostname)

	fmt.Print("user: ")
	new_host.User, _ = reader.ReadString('\n')
    new_host.User = strings.TrimSpace(new_host.User)

	fmt.Print("ip: ")
	new_host.Ip, _ = reader.ReadString('\n')
    new_host.Ip = strings.TrimSpace(new_host.Ip)

	fmt.Print("port: ")
	new_host.Port, _ = reader.ReadString('\n')
    new_host.Port = strings.TrimSpace(new_host.Port)

	hosts_list, err := getList()
	if err != nil {
		return err
	}

	hosts_list = append(hosts_list, new_host)

	return saveHost(hosts_list)

}

//psh connect, c <hostname>
func connectToHost(hostname string) error {

	hosts_list, err := getList()
	if err != nil {
		return err
	}

	var host_to_connect *Host

	for _, host := range(hosts_list){
		if host.Hostname == hostname{
			host_to_connect = &host
			break
		}
	}

	if host_to_connect == nil {
		fmt.Println("Error: host not found")
	}

	cmd := exec.Command(
		"ssh",
		host_to_connect.User+"@"+host_to_connect.Ip,
		"-p",
		host_to_connect.Port,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin  = os.Stdin

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}



//TODO add list 

func print_args(){
	
	fmt.Println("")
	fmt.Println("usage: ")
	fmt.Println("	psh <args>.")
	fmt.Println("add, a")
	fmt.Println("	add a new host.")
	fmt.Println("list, l")
	fmt.Println("	list all the knowun host.")
	fmt.Println("conncect, c <hostname>")
	fmt.Println("	connect to the host.")
	fmt.Println("remove, rm <hostname>")
	fmt.Println("	remove the host form the list.")
	fmt.Println("")
	fmt.Println("The file that is going to be used for")
	fmt.Println("storage is in: ~/.config/psh/hosts.json")
	fmt.Println("")
	fmt.Println("[!] make sure to make this dir for")
	fmt.Println("    making psh work")
	fmt.Println("")
}


func main() {
    // fmt.Println(len(os.Args))

	if len(os.Args) < 2 {
		print_args()
		return
	} 

	switch os.Args[1]{

		case "add", "a":
			err := addHost()
			if err != nil{
				fmt.Println("error: ", err)
			}

		case "list", "l":
			err := listHost()
			if err != nil{
				fmt.Println("error: ", err)
			}

		case "remove", "rm":
			err := removeHost(os.Args[2])
			if err != nil{
				fmt.Println("error: ", err)
			}

		case "connect", "c":
			err := connectToHost(os.Args[2])
			if err != nil{
				fmt.Println("error: ", err)
			}
	}

}

