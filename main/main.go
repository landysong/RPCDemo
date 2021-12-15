package main

import (
	"log"
	"net"
	"rpcdemo"
	"sync"
	"time"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	// 注册方法
	var foo Foo
	if err := rpcdemo.Register(&foo); err != nil {
		log.Fatal("register error: ", err)
	}

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}

	log.Println("start rpc server on:", l.Addr())
	addr <- l.Addr().String()
	rpcdemo.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	// client
	client, _ := rpcdemo.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{i, i * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
