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
	usuarioStorage      storage.UsuarioStorage
	alertaStorage       storage.AlertaStorage
	multimediaStorage   storage.MultimediaStorage
	persoaStorage       storage.PersonaStorage
	roleStorage         storage.RoleStorage
	colaboradorStorage  storage.ColaboradorStorage
	institucionStorage  storage.InstitucionStorage
	notificacionStorage storage.NotificacionStorage
	efectivoStorage     storage.EfectivoStorage

	//services
	usuarioService      *services.UsuarioService
	alertaService       *services.AlertaService
	multimediaService   *services.MultimediaService
	personaService      *services.PersonaService
	roleService         *services.RoleService
	colaboradorService  *services.ColaboradorService
	institucionService  *services.InsitucionService
	notificacionService *services.NotificacionService
	efectivoService     *services.EfectivoService

	//handlers
	usuarioHandler      *handlers.UsuarioHandlear
	alertaHandler       *handlers.AlertaHandlear
	multimediaHandler   *handlers.MultimediaHandlear
	personaHandler      *handlers.PersonaHandlear
	rolesHandler        *handlers.RolesHandlear
	colaboradorHandler  *handlers.ColaboradorHandlear
	institucionHandler  *handlers.InstitucionHandlear
	notificacionHandler *handlers.NotificacionHandler
	efectivoHandler     *handlers.EfectivoHandler
)
