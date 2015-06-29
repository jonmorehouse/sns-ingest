package main

func main() {
	ParseFlags()

	server := NewServer()
	server.serve()
}

