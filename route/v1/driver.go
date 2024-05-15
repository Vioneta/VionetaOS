package v1

import (
	"github.com/Vioneta/VionetaOS-Common/utils/common_err"
	"github.com/Vioneta/VionetaOS/drivers/dropbox"
	"github.com/Vioneta/VionetaOS/drivers/google_drive"
	"github.com/Vioneta/VionetaOS/drivers/onedrive"
	"github.com/Vioneta/VionetaOS/model"
	"github.com/gin-gonic/gin"
)

func ListDriverInfo(c *gin.Context) {
	list := []model.Drive{}

	google := google_drive.GetConfig()
	list = append(list, model.Drive{
		Name:    "Google Drive",
		Icon:    google.Icon,
		AuthUrl: google.AuthUrl,
	})
	dp := dropbox.GetConfig()
	list = append(list, model.Drive{
		Name:    "Dropbox",
		Icon:    dp.Icon,
		AuthUrl: dp.AuthUrl,
	})
	od := onedrive.GetConfig()
	list = append(list, model.Drive{
		Name:    "OneDrive",
		Icon:    od.Icon,
		AuthUrl: od.AuthUrl,
	})
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS), Data: list})
}
