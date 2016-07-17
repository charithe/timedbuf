Timed Buffer
============

Implementation of a buffer that flushes its' contents immediately after it's full or after a specified amount of time has passed.

Please note that in some cases it might be more efficient to use the actual data type (and buffer pools) instead of the generic `interface{}` type used in this implementation.

Usage
-----
`go get github.com/charithe/timedbuf`


```go
// define a flush function that will be called whenever the buffer is full or the time period has elapsed
func flushFunc(items []interface{}){
    // flush logic
}

// Initialize and use the buffer
tb := timedbuf.New(100, 10 * time.Second, flushFunc)
defer tb.Close()
tb.Put("foo", "bar")
tb.Put("baz")
...
```

