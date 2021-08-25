package injectors

import "github.com/user0608/mujeresapi/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandlear {
	return usuarioHandler
}
func GetAlertaHandler() *handlers.AlertaHandlear {
	return alertaHandler
}
