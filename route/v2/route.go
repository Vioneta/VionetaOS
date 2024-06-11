package v2

import (
	"github.com/Vioneta/VionetaOS/codegen"
	"github.com/Vioneta/VionetaOS/service"
)

type VionetaOS struct {
	fileUploadService *service.FileUploadService
}

func NewVionetaOS() codegen.ServerInterface {
	return &VionetaOS{
		fileUploadService: service.NewFileUploadService(),
	}
}
