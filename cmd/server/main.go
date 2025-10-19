package main

import "github.com/tariq-ventura/Proyecto-go/internal/router"

func main() {
	r := &router.Routes{}
	r.Routes = r.SetupRouter()
	r.Run()
}
