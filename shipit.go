//
// Shipit proxies stdio and logs to the given writer.
//
//   shipit.To(loggly.New("your-token-here"))
//
//   for {
//     time.Sleep(100 * time.Millisecond)
//     fmt.Fprintf(os.Stdout, "testing stdout\n")
//     fmt.Fprintf(os.Stderr, "testing stderr\n")
//   }
//
package shipit

import . "github.com/segmentio/go-dup"
import "bufio"
import "fmt"
import "os"
import "io"

// To ships stdio to the given writer.
func To(w io.Writer) (err error) {
	stderr, err := os.Open("/dev/stderr")

	if err != nil {
		return
	}

	if err = ship(1, "stdout", w, stderr); err != nil {
		return
	}

	if err = ship(2, "stderr", w, stderr); err != nil {
		return
	}

	return
}

// dup `fd` and read from the pipe to send log lines
// to loggly, then write back to the original stream.
func ship(fd int, name string, log io.Writer, stderr *os.File) error {
	r, w, err := Dup(fd, name)

	if err != nil {
		return err
	}

	go func() {
		buf := bufio.NewReader(r)

		for {
			line, err := buf.ReadBytes('\n')

			_, err = log.Write(line)

			if err != nil {
				fmt.Fprintf(stderr, "ERROR: shipit failed to write to writer: %s", err)
				break
			}

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Fprintf(stderr, "ERROR: shipit failed to read from %s: %s", name, err)
				break
			}

			_, err = w.Write(line)

			if err != nil {
				fmt.Fprintf(stderr, "ERROR: shipit failed to write to %s: %s", name, err)
				break
			}
		}
	}()

	return nil
}
