package v2

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Vioneta/VionetaOS/codegen"
	"github.com/Vioneta/VionetaOS/pkg/utils/file"
	"github.com/Vioneta/VionetaOS/service"
	"github.com/labstack/echo/v4"
	"github.com/mholt/archiver/v3"
)

func (s *VionetaOS) GetHealthServices(ctx echo.Context) error {
	services, err := service.MyService.Health().Services()
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: &message,
		})
	}

	return ctx.JSON(http.StatusOK, codegen.GetHealthServicesOK{
		Data: &codegen.HealthServices{
			Running:    services[true],
			NotRunning: services[false],
		},
	})
}

func (s *VionetaOS) GetHealthPorts(ctx echo.Context) error {
	tcpPorts, udpPorts, err := service.MyService.Health().Ports()
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: &message,
		})
	}

	return ctx.JSON(http.StatusOK, codegen.GetHealthPortsOK{
		Data: &codegen.HealthPorts{
			TCP: &tcpPorts,
			UDP: &udpPorts,
		},
	})
}
func (c *VionetaOS) GetHealthlogs(ctx echo.Context) error {
	var name, currentPath, commonDir, extension string
	var err error
	var ar archiver.Writer
	fileList, err := os.ReadDir("/var/log/vionetaos")
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: &message,
		})
	}
	extension, ar, err = file.GetCompressionAlgorithm("zip")
	if err != nil {
		ctx.Response().Header().Set("Content-Type", "application/json")
		message := err.Error()
		return ctx.JSON(http.StatusNotFound, codegen.ResponseInternalServerError{
			Message: &message,
		})
	}
	err = ar.Create(ctx.Response().Writer)
	if err != nil {
		ctx.Response().Header().Set("Content-Type", "application/json")
		message := err.Error()
		return ctx.JSON(http.StatusNotFound, codegen.ResponseInternalServerError{
			Message: &message,
		})
	}
	defer ar.Close()

	commonDir = "/var/log/vionetaos"

	currentPath = filepath.Base(commonDir)

	name = currentPath
	name += extension
	ctx.Response().Header().Add("Content-Type", "application/octet-stream")
	ctx.Response().Header().Add("Content-Transfer-Encoding", "binary")
	ctx.Response().Header().Add("Cache-Control", "no-cache")
	ctx.Response().Header().Add("Content-Disposition", "attachment; filename*=utf-8''"+url.PathEscape(name))

	for _, fname := range fileList {
		err := file.AddFile(ar, filepath.Join("/var/log/vionetaos", fname.Name()), commonDir)
		if err != nil {
			message := err.Error()
			return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
				Message: &message,
			})
		}

	}
	return nil
}
