# 06 -- Concurrency

The core strength of Go. Most interview questions touch this.

---

## Planned Topics

- [ ] Goroutines -- lightweight threads, `go` keyword
- [ ] Channels -- unbuffered vs buffered, send/receive
- [ ] Channel directions -- send-only, receive-only
- [ ] Select -- multiplexing channels, default case
- [ ] Timeouts -- `time.After`, context deadlines
- [ ] Context -- `context.WithCancel`, `WithTimeout`, `WithValue`
- [ ] WaitGroups -- `sync.WaitGroup`, coordinating goroutines
- [ ] Mutexes -- `sync.Mutex`, `sync.RWMutex`
- [ ] Once -- `sync.Once`, safe lazy initialization
- [ ] Worker pools -- bounded concurrency pattern
- [ ] Rate limiting -- `time.Tick`, token bucket
- [ ] Atomic operations -- `sync/atomic`, when to use
- [ ] Errgroup -- `golang.org/x/sync/errgroup`
