package main

import (
	"fmt"
	"net"
	"os"
	// "os"
)

// var (
// 	ipaddr = flag.String("ip", "", "ip address")
// 	port   = flag.String("p", "", "port")

// )

func main() {

	// s := flag.String("s", "s", "server")
	// c := flag.String("c", "c", "client")
	// flag.Var()

	// flag.Parse()
	// //	fmt.Println(*ipaddr, *port, *s, *c)
	// for k, v := range flag.Args() {
	// 	fmt.Println(k, v)

	// }

	//fmt.Println("test")
	workMode := 0
	ipaddr := ""
	port := ""

	for _, v := range os.Args[1:] {
		switch v {
		case "-s":
			workMode = 1
		case "-c":
			workMode = 2
		default:
			continue
		}
	}

	for k, v := range os.Args[1:] {
		switch v {
		case "-ip":
			ipaddr = os.Args[k+2]
		case "-p":
			port = os.Args[k+2]
		case "-h":
			help()
		default:
			continue
		}
	}

	fmt.Println(workMode, ipaddr, port)

	switch workMode {
	case 1:
		server(port)
	case 2:
		client(ipaddr, port)
	default:
		fmt.Println("workMode error")
		os.Exit(-1)
	}

}

func help() {
	fmt.Print(`
help:
	-s : server
	-c : client
	-ip : ip address
	-p : port 

	Exp:
		Server
			NetworkTest -s -p 500
		Client
			NetworkTest -c -ip 111.111.111.111 -p 500
	
	War!!!:
		no default value
	`)

	os.Exit(0)
}

func server(port string) {
	fmt.Println("server")

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Listening error", err)
		os.Exit(-1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accpet error")
			os.Exit(-1)
		}

		fmt.Println(c.RemoteAddr())
		c.Close()
	}

}

func client(ipaddr, port string) {
	fmt.Println("client")

	conn, err := net.Dial("tcp", ipaddr+":"+port)
	if err != nil {
		fmt.Println("dial error")
		os.Exit(-1)
	}
	defer conn.Close()

	fmt.Println("dial ok", conn.RemoteAddr())
}
