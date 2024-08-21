package tcp_sever

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //思考题：这里不填写会有啥问题？
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
	}
}

func main() {
	//listen port
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	//defer listenner.Close()

	//建立socket connection
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}
		go process(conn)

	}
}
