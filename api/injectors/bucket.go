package injectors

import (
	"sync"

	"github.com/user0608/mujeresapi/database"
	"github.com/user0608/mujeresapi/handlers"
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
	//storage
	usuarioStorage    storage.UsuarioStorage
	alertaStorage     storage.AlertaStorage
	multimediaStorage storage.MultimediaStorage
	//services
	usuarioService    *services.UsuarioService
	alertaService     *services.AlertaService
	multimediaService *services.MultimediaService

	//handlers
	usuarioHandler    *handlers.UsuarioHandlear
	alertaHandler     *handlers.AlertaHandlear
	multimediaHandler *handlers.MultimediaHandlear
)
