package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"syscall"
	"time"
)

var (
	c       = flag.Int("c", 10, "concurrency")
	opsRate = metrics.NewRegisteredMeter("ops", nil)
)

var (
	epoller    *epoll
	workerPoll *pool
)

func main() {
	flag.Parse()

	setLimit()

	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}

	workerPoll = newPool(*c, 1000000)
	workerPoll.start()
	defer workerPoll.Close()

	epoller, err = MkEpoll()
	if err != nil {
		panic(err)
	}

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err : %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		if err := epoller.Add(conn); err != nil {
			log.Printf("failed to add connection %v", err)
			conn.Close()
		}
	}

}

func start() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			workerPoll.addTask(conn)
		}
	}
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	log.Printf("set cur limit: %d", rLimit.Cur)
}
