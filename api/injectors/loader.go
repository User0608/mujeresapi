package injectors

import (
	"log"
	"sync"

	"github.com/user0608/mujeresapi/configs"
	"github.com/user0608/mujeresapi/database"
	"github.com/user0608/mujeresapi/handlers"
	"github.com/user0608/mujeresapi/repository"
	"github.com/user0608/mujeresapi/services"
	"gorm.io/gorm"
)

var ones sync.Once

func init() {
	ones.Do(func() {
		conf, err := configs.LoadDBConfigs("db_config.json")
		if err != nil {
			log.Fatalln("Err db configs,", err.Error())
		}
		log.Println("Configuraciones de base de datos cargadas!")
		db = database.GetDBConnextion(conf)
		initRepository(db)
		initServices()
		initHandlers()
	})
}
func initRepository(gdb *gorm.DB) {
	usuarioStorage = repository.NewUsuarioRepository(gdb)
	alertaStorage = repository.NewAlertaRepository(gdb)
	multimediaStorage = repository.NewMultimediaRepository(gdb)
	persoaStorage = repository.NewPersonaRepository(gdb)
	roleStorage = repository.NewRolesRepository(gdb)
	colaboradorStorage = repository.NewColaboradorRepository(gdb)
}
func initServices() {
	usuarioService = services.NewUsuarioService(usuarioStorage)
	alertaService = services.NewAlertaService(alertaStorage)
	multimediaService = services.NewMultimediaService(multimediaStorage)
	personaService = services.NewPersonaService(persoaStorage)
	roleService = services.NewRoleService(roleStorage)
	colaboradorService = services.NewColaboradorService(colaboradorStorage)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
	alertaHandler = handlers.NewAlertaHandler(alertaService)
	multimediaHandler = handlers.NewMultimediaHandler(multimediaService)
	personaHandler = handlers.NewPersonaHandler(personaService)
	rolesHandler = handlers.NewRolesHandler(roleService)
	colaboradorHandler = handlers.NewColaboradorHandler(colaboradorService)
}
