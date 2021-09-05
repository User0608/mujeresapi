package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type MultimediaHandlear struct {
	service *services.MultimediaService
}

func NewMultimediaHandler(s *services.MultimediaService) *MultimediaHandlear {
	return &MultimediaHandlear{
		service: s,
	}
}

func (h *MultimediaHandlear) UploadMultimedia(c echo.Context) error {
	log.Println("Leg!")
	data := c.FormValue("data")
	log.Println("Leg>")
	alertaID := 0
	medias := []application.Multimedia{}
	if data == "" {
		log.Println("Alerta ID, null")
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Datos invalidos o faltantes"))
	} else {
		var err error
		if alertaID, err = strconv.Atoi(data); err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Datos invalidos o faltantes"))
		}

	}
	if err := h.service.VerificarAlertaID(alertaID); err != nil {
		log.Println(err.Error())
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("No se encontró la alerta asociada con el contenido multimedia"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}

LOOP:
	for prefix := 1; true; prefix++ {
		file, err := c.FormFile(fmt.Sprintf("file%d", prefix))
		if err != nil {
			break LOOP
		}
		f, err := file.Open()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		nf, err := os.OpenFile(fmt.Sprintf("private/%s", file.Filename), os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		if _, err = io.Copy(nf, f); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}

		f.Close()
		nf.Close()
		medias = append(medias, application.Multimedia{
			Url:      file.Filename,
			Tipo:     h.TypeFile(file.Filename),
			AlertaID: alertaID,
		})

	}
	err := h.service.RegistrarMultimedia(medias)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkMesage("Archivos guradados"))
}
func (*MultimediaHandlear) TypeFile(fileName string) string {
	if fileName == "" {
		return "otros"
	}
	tipos := map[string]string{
		"mp4": "video", "mp3": "audio",
		"png": "imagen", "jpg": "imagen",
	}
	r := strings.Split(fileName, ".")
	ext := r[len(r)-1]

	if tipo, ok := tipos[ext]; ok {
		return tipo
	}
	return "otros"
}
