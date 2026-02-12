# 06 -- Concurrency

Go's concurrency model is its **killer feature**. Goroutines are
cheap, channels are typed, and the runtime handles scheduling.
This is the most interview-heavy module in the entire repo.

Each example includes `main.go` + Python equivalent `main.py`.
Python equivalents use `threading` + `queue.Queue` or `asyncio`
depending on which maps best. Key differences (GIL, async model)
are noted in every topic.

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `01_goroutines_vs_threads` | `go` keyword, goroutine scheduling, cost comparison |
| 02 | `02_channels_basics` | Unbuffered channels, send/receive, close semantics |
| 03 | `03_buffered_channels` | Buffered channels, capacity, when to buffer |
| 04 | `04_channel_directions` | Send-only `chan<-` and receive-only `<-chan` types |
| 05 | `05_select` | Multiplexing channels, default case, non-blocking ops |
| 06 | `06_timeouts_and_timers` | `time.After`, `time.NewTimer`, `time.NewTicker` |
| 07 | `07_waitgroups` | `sync.WaitGroup` for joining goroutines |
| 08 | `08_mutexes` | `sync.Mutex`, `sync.RWMutex`, protecting shared state |
| 09 | `09_worker_pools` | Bounded concurrency with channel-based worker pool |
| 10 | `10_rate_limiting` | Ticker-based and token-bucket rate limiting |
| 11 | `11_context_cancellation` | `context.WithCancel`, `WithTimeout`, propagation |
| 12 | `12_fanout_fanin_pipeline` | Fan-out, fan-in, multi-stage pipeline |

---

## 10-Min Revision Path

1. Skim `01_goroutines_vs_threads` -- recall goroutine cost and scheduling
2. Skim `02_channels_basics` -- send, receive, close, range
3. Skim `05_select` -- multiplexing, non-blocking, timeouts
4. Skim `07_waitgroups` -- joining goroutines
5. Skim `09_worker_pools` -- the pattern interviewers love to ask
6. Skim `11_context_cancellation` -- propagation through call chains
7. Skim `_quick_revision/` -- one-screen cheat sheet

---

## Common Concurrency Mistakes

- Launching goroutines without a way to stop them (goroutine leak)
- Sending on a closed channel (panics)
- Forgetting to close channels (range blocks forever)
- Using unbuffered channels when buffered is needed (deadlock)
- Sharing memory without a mutex or channel (data race)
- Not using `go test -race` to detect races
- Forgetting `wg.Add` before launching the goroutine
- Ignoring context cancellation in long-running work
- Using `time.Sleep` instead of proper synchronization

---

## TL;DR

- Goroutines are ~2 KB each; OS threads are ~1 MB -- Go can run millions
- Channels are typed pipes: unbuffered = synchronous, buffered = async up to capacity
- `select` multiplexes channels like `epoll` for goroutines
- `sync.WaitGroup` waits for a group; `sync.Mutex` protects shared state
- Worker pools = fixed goroutines reading from a shared job channel
- `context.Context` propagates cancellation and deadlines through call chains
- Always run `go test -race` to catch data races
- Python: threads + GIL (no true parallelism), asyncio (cooperative concurrency)
