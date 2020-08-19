package simple_rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWrite(t *testing.T) {
	addr := "0.0.0.0:2333"
	cont := "yephahhago"
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		l, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := l.Accept()
		s := Session{conn: conn}
		err = s.Write([]byte(cont))
		if err != nil {
			t.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{conn: conn}
		data, err := s.Read()
		fmt.Print(string(data))
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != cont {
			t.FailNow()
		}
	}()
	wg.Wait()
}
