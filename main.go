package MP1

import (
	"bufio"
)

func read_config {
	//use the space delimiter to grab the values

	//put those values in an arrays?
	//arrays can't be resized though...slices?





}

func build_TCP_channels {

	//make the TCP servers here using refactored code from MP0

	//for loop to create the number of servers/nodes specified by the config file

}

func start_unicast {
	//collect input from the user
	//unicast_send built on TCP Client and unicast_recieve built on TCP server??
}

func main() {
	read_config()
	build_TCP_channels()
	start_unicast()
}