package router

import "github.com/gin-contrib/cors"

func (r *Routes) SetupCors() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Routes.Use(cors.New(config))
}
