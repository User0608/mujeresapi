package injectors

import (
	"github.com/user0608/mujeresapi/handlers"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/storage"
	"gorm.io/gorm"
)

var ( //db connextion
	db *gorm.DB
	//storage
	usuarioStorage    storage.UsuarioStorage
	alertaStorage     storage.AlertaStorage
	multimediaStorage storage.MultimediaStorage
	persoaStorage     storage.PersonaStorage
	//services
	usuarioService    *services.UsuarioService
	alertaService     *services.AlertaService
	multimediaService *services.MultimediaService
	personaService    *services.PersonaService

	//handlers
	usuarioHandler    *handlers.UsuarioHandlear
	alertaHandler     *handlers.AlertaHandlear
	multimediaHandler *handlers.MultimediaHandlear
	personaHandler    *handlers.PersonaHandlear
)
