package paths

import "fmt"

var (
	API      = "/v1"
	APP_USER = fmt.Sprintf("%s/usuario", API)
	ALERTA   = fmt.Sprintf("%s/alerta", API)
)
