package xclient

import (
	"context"
	"fmt"
	"net"
	. "rpcdemo"
	"testing"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func call(addr1, addr2 string) (int, error) {
	d := NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := NewXClient(d, RandomSelect, nil)
	defer func() { _ = xc.Close() }()
	var reply int
	ctx := context.Background()
	err := xc.Call(ctx, "Foo.Sum", &Args{1, 2}, &reply)
	return reply, err
}

func startServer(addr chan string) {
	var f Foo
	_ = Register(&f)
	l, _ := net.Listen("tcp", ":0")
	addr <- l.Addr().String()
	Accept(l)
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

func TestXClient_Call(t *testing.T) {
	t.Parallel()

	ch1 := make(chan string)
	ch2 := make(chan string)
	// start two servers
	go startServer(ch1)
	go startServer(ch2)

	addr1 := <-ch1
	addr2 := <-ch2

	t.Run("XClient call", func(t *testing.T) {
		d := NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
		xc := NewXClient(d, RandomSelect, nil)
		defer func() { _ = xc.Close() }()
		var reply int
		var args = &Args{1, 2}
		ctx := context.Background()
		_ = xc.Call(ctx, "Foo.Sum", args, &reply)
		sum := args.Num2 + args.Num1
		_assert(reply == sum, "expect Foo.Sum call success")
	})
}
