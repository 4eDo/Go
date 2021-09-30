package main

import(
	"zzz/check"
	"io"
	"log"
	"net"
	"fmt"
	"os"
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
	
	c := make(chan struct{})
	
	for i := 0; i < portCount; {
		var port string = ports[i]
		fmt.Printf("\tНачало чтения порта %s...\n", port)
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
	
	conn, err := net.Dial("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}