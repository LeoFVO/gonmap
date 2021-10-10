package main
import(
	"fmt"
	"net"
)

func main() {

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		con, err := net.Dial("tcp", address)
		if err != nil { 
			// Port closed or filtered
			continue
		}
		con.Close()
		fmt.Println("%d open\n", i)
	}
}