package v2

import (
	"github.com/Vioneta/VionetaOS/codegen"
	"github.com/Vioneta/VionetaOS/service"
)

type CasaOS struct {
	fileUploadService *service.FileUploadService
}

func NewCasaOS() codegen.ServerInterface {
	return &CasaOS{
		fileUploadService: service.NewFileUploadService(),
	}
}
