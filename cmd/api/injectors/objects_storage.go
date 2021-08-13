package injectors

import (
	"sync"

	"github.com/user0608/mujeresapi/database"
	"github.com/user0608/mujeresapi/handlers"
	"github.com/user0608/mujeresapi/repository"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/storage"
	"gorm.io/gorm"
)

var ones sync.Once

var dbConfig = database.DBConfig{
	Host:     "localhost",
	UserDB:   "postgres",
	Password: "maira002",
	DBName:   "mujeres_db",
	Port:     "5432",
}

var ( //db connextion
	db *gorm.DB
	//torager
	usuarioStorage storage.UsuarioStorage
	//services
	usuarioService *services.UsuarioService

	//handlers
	usuarioHandler *handlers.UsuarioHandlear
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
}
func initServices() {
	usuarioService = services.NewUsuarioService(usuarioStorage)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
}

func GetNewUsuarioHandler() *handlers.UsuarioHandlear {
	return usuarioHandler
}
