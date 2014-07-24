package main

import "github.com/segmentio/go-shipit"
import "github.com/segmentio/go-loggly"
import "time"
import "fmt"
import "os"

func main() {
	log := loggly.New("your-token-here")

	err := shipit.To(log)

	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(os.Stdout, "testing stdout\n")
		fmt.Fprintf(os.Stderr, "testing stderr\n")
	}
}
