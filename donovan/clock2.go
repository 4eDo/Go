package main

import(
	"io"
	"log"
	"net"
	"time"
	"fmt"
	"zzz/check"
)

func main() {
	
	var portCount int
	fmt.Println("\tВведите количество портов:")
	portCount = check.InputIntJNBE()
	
	var ports []string = make([]string, portCount)
	for i := 0; i < portCount; i++ {
		fmt.Printf("\tВведите номер порта %d:\n", i+1)
		ports[i] = check.InputStringJNBE()
	}
	
	fmt.Println(ports) 
	
	c := make(chan struct{})
	
	for i := 0; i < portCount; {
		var port string = ports[i]
		fmt.Printf("\tЗапуск порта %s...\n", port)
		go func() {
			addConn(port)
			c <- struct{}{}
		}()
		
		i++
	}
	
	for i := 0; i < portCount; i++ {
        <- c
    }
}

func addConn(port string) {
	localhost := "localhost:" + port
	listener, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err!= nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err!= nil {
			return
		}
		time.Sleep(1*time.Second)
	}
}