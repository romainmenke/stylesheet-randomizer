package main

import "github.com/romainmenke/stylesheet-randomizer/server"

func main() {
	s := server.Serve(":4567")
	defer s.Close()

	s.ListenAndServe()
}
