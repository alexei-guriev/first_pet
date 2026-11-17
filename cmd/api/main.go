package main

import "github.com/alexei-guriev/first_pet/iternal/httpapi"

func main() {
	srv := httpapi.New(httpapi.NewRouter())
	srv.Run()
}
