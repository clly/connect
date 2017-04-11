// Author: Connor Kelly

package main

import (
	"fmt"
	"net"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	host := kingpin.Arg("host", "Hostname to try to connect to").String()
	port := kingpin.Arg("port", "Port to try to connect to").Int()
	timeout := kingpin.Flag("timeout", "Timeout in seconds").Default("5s").Short('t').Duration()
	kingpin.Parse()
	ips, err := net.LookupIP(*host)
	if err != nil {
		fmt.Errorf("Unable to lookup host %s, %s", *host, err)
		os.Exit(-1)
	}
	connString := fmt.Sprintf("%s:%v", ips[0].String(), *port)

	conn, err := net.DialTimeout("tcp", connString, *timeout)
	if err != nil {
		fmt.Printf("Failed to connect to %s at ip %s at port %v within %s\n", *host, ips[0].String(), *port, timeout)
	} else {
		fmt.Printf("Successfully connected to %s at ip %s at port %v\n", *host, ips[0].String(), *port)
		conn.Close()
	}
}
