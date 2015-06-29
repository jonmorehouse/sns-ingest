package main

func main() {
	ParseFlags()

	server := InitServer()
	server.serve()
}

