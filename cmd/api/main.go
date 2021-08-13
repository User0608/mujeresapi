package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/cmd/api/router"
)

func main() {
	if err := authorization.LoadFiles("./certificates/app.rsa", "./certificates/app.rsa.pub"); err != nil {
		log.Fatalln("No se cargaron los certificados,", err.Error())
	}
	log.Println("Certificados cargados....")
	e := echo.New()
	router.RoutesUsuario(e)
	log.Fatal(e.Start("localhost:91"))
}
