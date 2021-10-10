package main
import(
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Allow use to synchronize main function with go routines
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // Add 1 routine to group
		go func (portNumber int) {
			defer wg.Done() // remove 1 routine to group
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
	wg.Wait() // wait all routines ended before close.
}