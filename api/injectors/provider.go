package injectors

import "github.com/user0608/mujeresapi/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandlear {
	return usuarioHandler
}
func GetAlertaHandler() *handlers.AlertaHandlear {
	return alertaHandler
}
func GetMultimediaHandler() *handlers.MultimediaHandlear {
	return multimediaHandler
}
func GetPersonaHandler() *handlers.PersonaHandlear {
	return personaHandler
}
func GetRolesHandler() *handlers.RolesHandlear {
	return rolesHandler
}
func GetColaboradorHandler() *handlers.ColaboradorHandlear {
	return colaboradorHandler
}
func GetInstitucionHandler() *handlers.InstitucionHandlear {
	return institucionHandler
}
func GetNotificationHandler() *handlers.NotificacionHandler {
	return notificacionHandler
}
func GetEfectivoHandler() *handlers.EfectivoHandler {
	return efectivoHandler
}
