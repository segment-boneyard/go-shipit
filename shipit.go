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
import "io"

// To ships stdio to the given writer.
func To(w io.Writer) (err error) {
	if err = ship(1, "stdout", w); err != nil {
		return
	}

	if err = ship(2, "stderr", w); err != nil {
		return
	}

	return
}

// dup `fd` and read from the pipe to send log lines
// to loggly, then write back to the original stream.
func ship(fd int, name string, log io.Writer) error {
	r, w, err := Dup(fd, name)

	if err != nil {
		return err
	}

	go func() {
		buf := bufio.NewReader(r)

		for {
			line, err := buf.ReadBytes('\n')

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			_, err = log.Write(line)

			if err != nil {
				panic(err)
			}

			_, err = w.Write(line)

			if err != nil {
				panic(err)
			}
		}
	}()

	return nil
}
