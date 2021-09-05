package injectors

import (
	"github.com/user0608/mujeresapi/database"
	"github.com/user0608/mujeresapi/handlers"
	"github.com/user0608/mujeresapi/repository"
	"github.com/user0608/mujeresapi/services"
	"gorm.io/gorm"
)

func init() {
	ones.Do(func() {
		db = database.GetDBConnextion(&dbConfig)
		initRepository(db)
		initServices()
		initHandlers()
	})
}
func initRepository(gdb *gorm.DB) {
	usuarioStorage = repository.NewUsuarioRepository(gdb)
	alertaStorage = repository.NewAlertaRepository(gdb)
	multimediaStorage = repository.NewMultimediaRepository(gdb)
}
func initServices() {
	usuarioService = services.NewUsuarioService(usuarioStorage)
	alertaService = services.NewAlertaService(alertaStorage)
	multimediaService = services.NewMultimediaService(multimediaStorage)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
	alertaHandler = handlers.NewAlertaHandler(alertaService)
	multimediaHandler = handlers.NewMultimediaHandler(multimediaService)
}
