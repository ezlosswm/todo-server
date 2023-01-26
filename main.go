package main

var listenAddr int = 3000

func main() {
	s := NewAPIServer(listenAddr)
	s.Run()
}
