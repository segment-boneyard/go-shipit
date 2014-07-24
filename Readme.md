
# go-shipit

 Ship it sends stdio logs to the given writer and re-writes to
 the original stdio streams for seamless local/remote logging
 without introducing logging agents or tailers.

 ![](http://1.bp.blogspot.com/_v0neUj-VDa4/TFBEbqFQcII/AAAAAAAAFBU/E8kPNmF1h1E/s640/squirrelbacca-thumb.jpg)

## Example

  With the following example you'll see "testing {stdout,stderr}" in the terminal as well as loggly.

```go
shipit.To(loggly.New("your-token-here"))

for {
  time.Sleep(100 * time.Millisecond)
  fmt.Fprintf(os.Stdout, "testing stdout\n")
  fmt.Fprintf(os.Stderr, "testing stderr\n")
}
```

  March on Chewbacca.

# License

 MIT