package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/user0608/mujeresapi/api/router"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/configs"
)

func main() {
	conf, err := configs.LoadConfigs("config.json")
	if err != nil {

	}
	if err := authorization.LoadFiles(conf.Certificates.Private, conf.Certificates.Public); err != nil {
		log.Fatalln("No se cargaron los certificados,", err.Error())
	}
	if _, err := os.Stat(conf.MediaRootDir); os.IsNotExist(err) {
		if err := os.Mkdir(conf.MediaRootDir, 0755); err != nil {
			log.Fatal(err.Error())
		}
	}
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.Cors.AllowOrigins,
		AllowMethods: conf.Cors.AllowMethods,
	}))
	server.HideBanner = true
	router.Upgrade(server)
	server.Static(conf.PublicMediaDirPath, conf.MediaRootDir)
	log.Fatal(server.Start(conf.Address))
}
