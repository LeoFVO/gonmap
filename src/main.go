package main
import(
	"fmt"
	"net"
	"sort"
)

/** 
 * Define worker
 * Get non treated port from channel and scan them
 * Send number of port in results channel if opened, else 0
*/
func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		con, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		con.Close()
		results <- p
	}
}

func main() {
	// Initialize variables
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// Launch worker
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	// send port needing to be treated in channel
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <- results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
