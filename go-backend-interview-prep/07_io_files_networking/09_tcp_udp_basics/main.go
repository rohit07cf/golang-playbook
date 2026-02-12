package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// echoServer accepts one TCP connection and echoes lines back.
func echoServer(ln net.Listener, ready chan<- struct{}) {
	ready <- struct{}{} // signal that server is listening

	conn, err := ln.Accept()
	if err != nil {
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "QUIT" {
			fmt.Fprintf(conn, "BYE\n")
			return
		}
		fmt.Fprintf(conn, "ECHO: %s\n", line)
	}
}

func main() {
	// --- TCP echo server + client ---
	fmt.Println("=== TCP Echo Server + Client ===")

	// Start server on random port
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer ln.Close()
	addr := ln.Addr().String()
	fmt.Println("  server listening on", addr)

	ready := make(chan struct{})
	go echoServer(ln, ready)
	<-ready // wait for server to be ready

	// Client connects
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		fmt.Println("  dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("  client connected")

	// Send messages and read echoes
	messages := []string{"hello", "world", "Go networking", "QUIT"}
	reader := bufio.NewReader(conn)

	for _, msg := range messages {
		fmt.Fprintf(conn, "%s\n", msg)
		reply, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("  sent: %q -> got: %s", msg, reply)
	}

	// --- UDP example (quick) ---
	fmt.Println("\n=== UDP (quick demo) ===")
	udpAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("  udp listen error:", err)
		return
	}
	defer udpConn.Close()

	udpServerAddr := udpConn.LocalAddr().String()
	fmt.Println("  udp server on", udpServerAddr)

	// Client sends to server
	go func() {
		clientAddr, _ := net.ResolveUDPAddr("udp", udpServerAddr)
		client, _ := net.DialUDP("udp", nil, clientAddr)
		defer client.Close()
		client.Write([]byte("udp hello"))
	}()

	buf := make([]byte, 1024)
	udpConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("  udp read error:", err)
	} else {
		fmt.Printf("  received %d bytes from %s: %q\n", n, remoteAddr, buf[:n])
	}
}
