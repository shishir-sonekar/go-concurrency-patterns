# Go Concurrency Patterns

Collection of practical and idiomatic **Go concurrency examples**, covering core patterns like goroutines, channels, worker pools, fan-out/fan-in, context cancellation, and more.

> Ideal for learning, experimenting, and mastering concurrent programming in Go.

---

## 📦 What's Inside

| Pattern                     | Description                                                |
|----------------------------|------------------------------------------------------------|
| [Worker Pool](worker-pool/)        | Fixed number of goroutines processing jobs via channel     |
| [Per-Task Goroutines](per-task/)   | One goroutine per task, tracked via `sync.WaitGroup`       |
| [Semaphore (Channel-based)](semaphore/) | Limits max concurrent tasks using buffered channel         |
| [Rate Limiter](rate-limiter/)     | Basic token bucket implementation using time + channels    |
| [Fan-Out / Fan-In](fan-out-in/)   | Distribute work to many goroutines, collect results        |
| [Graceful Shutdown](graceful-shutdown/) | Cancel in-flight goroutines via `context.Context`        |
| [Timeouts with Context](timeout/) | Time-bound task execution using `context.WithTimeout`      |
| [Pipeline Pattern](pipeline/)     | Connect stages with channels for streaming transformation  |
| [Concurrent Map Access](sync-map/) | Safe read/write to shared map using `sync.Mutex` or `sync.Map` |

> More examples coming soon!

---

## Requirements

- Go 1.18 or later

---

## Getting Started

```bash
git clone https://github.com/yourusername/go-concurrency-patterns.git
cd go-concurrency-patterns
go run worker-pool/minimalWorkerPool.go
