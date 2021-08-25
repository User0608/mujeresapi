package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/cmd/api/router"
)

func main() {
	if err := authorization.LoadFiles("./cmd/certificates/app.rsa", "./cmd/certificates/app.rsa.pub"); err != nil {
		log.Fatalln("No se cargaron los certificados,", err.Error())
	}
	log.Println("Certificados cargados....")
	server := echo.New()
	router.Upgrade(server)
	log.Fatal(server.Start("0.0.0.0:91"))
}
