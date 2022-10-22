package main

import "assignment-3/routes"

func main() {
	var PORT = ":3000"
	err := routes.StartServer().Run(PORT)

	if err != nil {
		panic(err)
	}
}
