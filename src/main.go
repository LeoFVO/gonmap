package main
import(
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func (portNumber int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", portNumber)
			con, err := net.Dial("tcp", address)
			if err != nil { 
				// Port closed or filtered
				return
			}
			con.Close()
			fmt.Printf("%v open\n", portNumber)
		}(i)
	}
	wg.Wait()
}