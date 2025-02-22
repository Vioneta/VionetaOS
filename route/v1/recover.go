package v1

import (
	"strconv"
	"strings"
	"time"

	"github.com/Vioneta/VionetaOS-Common/utils/logger"
	"github.com/Vioneta/VionetaOS/drivers/dropbox"
	"github.com/Vioneta/VionetaOS/drivers/google_drive"
	"github.com/Vioneta/VionetaOS/drivers/onedrive"
	"github.com/Vioneta/VionetaOS/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRecoverStorage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	t := c.Param("type")
	currentTime := time.Now().UTC()
	currentDate := time.Now().UTC().Format("2006-01-02")
	notify := make(map[string]interface{})
	if t == "GoogleDrive" {
		google_drive := google_drive.GetConfig()
		google_drive.Code = c.Query("code")
		if len(google_drive.Code) == 0 {
			c.String(200, `<p>Code cannot be empty</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Code cannot be empty"
			logger.Error("Then code is empty: ", zap.String("code", google_drive.Code), zap.Any("name", "google_drive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}

		err := google_drive.Init(c)
		if err != nil {
			c.String(200, `<p>Initialization failure:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Initialization failure"
			logger.Error("Then init error: ", zap.Error(err), zap.Any("name", "google_drive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}

		username, err := google_drive.GetUserInfo(c)
		if err != nil {
			c.String(200, `<p>Failed to get user information:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get user information"
			logger.Error("Then get user info error: ", zap.Error(err), zap.Any("name", "google_drive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		dmap := make(map[string]string)
		dmap["username"] = username
		configs, err := service.MyService.Storage().GetConfig()
		if err != nil {
			c.String(200, `<p>Failed to get rclone config:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get rclone config"
			logger.Error("Then get config error: ", zap.Error(err), zap.Any("name", "google_drive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		for _, v := range configs.Remotes {
			cf, err := service.MyService.Storage().GetConfigByName(v)
			if err != nil {
				logger.Error("then get config by name error: ", zap.Error(err), zap.Any("name", v))
				continue
			}
			if cf["type"] == "drive" && cf["username"] == dmap["username"] {
				c.String(200, `<p>The same configuration has been added</p><script>window.close()</script>`)
				err := service.MyService.Storage().CheckAndMountByName(v)
				if err != nil {
					logger.Error("check and mount by name error: ", zap.Error(err), zap.Any("name", cf["username"]))
				}
				notify["status"] = "warn"
				notify["message"] = "The same configuration has been added"
				service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
				return
			}
		}
		if len(username) > 0 {
			a := strings.Split(username, "@")
			username = a[0]
		}

		//username = fileutil.NameAccumulation(username, "/mnt")
		username += "_google_drive_" + strconv.FormatInt(time.Now().Unix(), 10)

		dmap["client_id"] = google_drive.ClientID
		dmap["client_secret"] = google_drive.ClientSecret
		dmap["scope"] = "drive"
		dmap["mount_point"] = "/mnt/" + username
		dmap["token"] = `{"access_token":"` + google_drive.AccessToken + `","token_type":"Bearer","refresh_token":"` + google_drive.RefreshToken + `","expiry":"` + currentDate + `T` + currentTime.Add(time.Hour*1).Add(time.Minute*50).Format("15:04:05") + `Z"}`
		service.MyService.Storage().CreateConfig(dmap, username, "drive")
		service.MyService.Storage().MountStorage("/mnt/"+username, username+":")
		notify := make(map[string]interface{})
		notify["status"] = "success"
		notify["message"] = "Success"
		notify["driver"] = "GoogleDrive"
		service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
	} else if t == "Dropbox" {
		dropbox := dropbox.GetConfig()
		dropbox.Code = c.Query("code")
		if len(dropbox.Code) == 0 {
			c.String(200, `<p>Code cannot be empty</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Code cannot be empty"
			logger.Error("Then code is empty error: ", zap.String("code", dropbox.Code), zap.Any("name", "dropbox"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}

		err := dropbox.Init(c)
		if err != nil {
			c.String(200, `<p>Initialization failure:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Initialization failure"
			logger.Error("Then init error: ", zap.Error(err), zap.Any("name", "dropbox"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		username, err := dropbox.GetUserInfo(c)
		if err != nil {
			c.String(200, `<p>Failed to get user information:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get user information"
			logger.Error("Then get user information: ", zap.Error(err), zap.Any("name", "dropbox"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		dmap := make(map[string]string)
		dmap["username"] = username

		configs, err := service.MyService.Storage().GetConfig()
		if err != nil {
			c.String(200, `<p>Failed to get rclone config:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get rclone config"
			logger.Error("Then get config error: ", zap.Error(err), zap.Any("name", "dropbox"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		for _, v := range configs.Remotes {
			cf, err := service.MyService.Storage().GetConfigByName(v)
			if err != nil {
				logger.Error("then get config by name error: ", zap.Error(err), zap.Any("name", v))
				continue
			}
			if cf["type"] == "dropbox" && cf["username"] == dmap["username"] {
				c.String(200, `<p>The same configuration has been added</p><script>window.close()</script>`)
				err := service.MyService.Storage().CheckAndMountByName(v)
				if err != nil {
					logger.Error("check and mount by name error: ", zap.Error(err), zap.Any("name", cf["username"]))
				}

				notify["status"] = "warn"
				notify["message"] = "The same configuration has been added"
				service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
				return
			}
		}
		if len(username) > 0 {
			a := strings.Split(username, "@")
			username = a[0]
		}
		username += "_dropbox_" + strconv.FormatInt(time.Now().Unix(), 10)

		dmap["client_id"] = dropbox.AppKey
		dmap["client_secret"] = dropbox.AppSecret
		dmap["token"] = `{"access_token":"` + dropbox.AccessToken + `","token_type":"bearer","refresh_token":"` + dropbox.Addition.RefreshToken + `","expiry":"` + currentDate + `T` + currentTime.Add(time.Hour*3).Add(time.Minute*50).Format("15:04:05") + `.780385354Z"}`
		dmap["mount_point"] = "/mnt/" + username
		// data.SetValue(username, "type", "dropbox")
		// data.SetValue(username, "client_id", add.AppKey)
		// data.SetValue(username, "client_secret", add.AppSecret)
		// data.SetValue(username, "mount_point", "/mnt/"+username)

		// data.SetValue(username, "token", `{"access_token":"`+dropbox.AccessToken+`","token_type":"bearer","refresh_token":"`+dropbox.Addition.RefreshToken+`","expiry":"`+currentDate+`T`+currentTime.Add(time.Hour*3).Format("15:04:05")+`.780385354Z"}`)
		// e = data.Save()
		// if e != nil {
		// 	c.String(200, `<p>保存配置失败:`+e.Error()+`</p>`)

		// 	return
		// }
		service.MyService.Storage().CreateConfig(dmap, username, "dropbox")
		service.MyService.Storage().MountStorage("/mnt/"+username, username+":")

		notify["status"] = "success"
		notify["message"] = "Success"
		notify["driver"] = "Dropbox"
		service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
	} else if t == "Onedrive" {
		onedrive := onedrive.GetConfig()
		onedrive.Code = c.Query("code")
		if len(onedrive.Code) == 0 {
			c.String(200, `<p>Code cannot be empty</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Code cannot be empty"
			logger.Error("Then code is empty error: ", zap.String("code", onedrive.Code), zap.Any("name", "onedrive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}

		err := onedrive.Init(c)
		if err != nil {
			c.String(200, `<p>Initialization failure:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Initialization failure"
			logger.Error("Then init error: ", zap.Error(err), zap.Any("name", "onedrive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		username, driveId, driveType, err := onedrive.GetInfo(c)
		if err != nil {
			c.String(200, `<p>Failed to get user information:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get user information"
			logger.Error("Then get user information: ", zap.Error(err), zap.Any("name", "onedrive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		dmap := make(map[string]string)
		dmap["username"] = username

		configs, err := service.MyService.Storage().GetConfig()
		if err != nil {
			c.String(200, `<p>Failed to get rclone config:`+err.Error()+`</p><script>window.close()</script>`)
			notify["status"] = "fail"
			notify["message"] = "Failed to get rclone config"
			logger.Error("Then get config error: ", zap.Error(err), zap.Any("name", "onedrive"))
			service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
			return
		}
		for _, v := range configs.Remotes {
			cf, err := service.MyService.Storage().GetConfigByName(v)
			if err != nil {
				logger.Error("then get config by name error: ", zap.Error(err), zap.Any("name", v))
				continue
			}
			if cf["type"] == "onedrive" && cf["username"] == dmap["username"] {
				c.String(200, `<p>The same configuration has been added</p><script>window.close()</script>`)
				err := service.MyService.Storage().CheckAndMountByName(v)
				if err != nil {
					logger.Error("check and mount by name error: ", zap.Error(err), zap.Any("name", cf["username"]))
				}

				notify["status"] = "warn"
				notify["message"] = "The same configuration has been added"
				service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
				return
			}
		}
		if len(username) > 0 {
			a := strings.Split(username, "@")
			username = a[0]
		}
		username += "_onedrive_" + strconv.FormatInt(time.Now().Unix(), 10)

		dmap["client_id"] = onedrive.ClientID
		dmap["client_secret"] = onedrive.ClientSecret
		dmap["token"] = `{"access_token":"` + onedrive.AccessToken + `","token_type":"bearer","refresh_token":"` + onedrive.RefreshToken + `","expiry":"` + currentDate + `T` + currentTime.Add(time.Hour*3).Add(time.Minute*50).Format("15:04:05") + `.780385354Z"}`
		dmap["mount_point"] = "/mnt/" + username
		dmap["drive_id"] = driveId
		dmap["drive_type"] = driveType
		// data.SetValue(username, "type", "dropbox")
		// data.SetValue(username, "client_id", add.AppKey)
		// data.SetValue(username, "client_secret", add.AppSecret)
		// data.SetValue(username, "mount_point", "/mnt/"+username)

		// data.SetValue(username, "token", `{"access_token":"`+dropbox.AccessToken+`","token_type":"bearer","refresh_token":"`+dropbox.Addition.RefreshToken+`","expiry":"`+currentDate+`T`+currentTime.Add(time.Hour*3).Format("15:04:05")+`.780385354Z"}`)
		// e = data.Save()
		// if e != nil {
		// 	c.String(200, `<p>保存配置失败:`+e.Error()+`</p>`)

		// 	return
		// }
		service.MyService.Storage().CreateConfig(dmap, username, "onedrive")
		service.MyService.Storage().MountStorage("/mnt/"+username, username+":")

		notify["status"] = "success"
		notify["message"] = "Success"
		notify["driver"] = "Onedrive"
		service.MyService.Notify().SendNotify("vionetaos:file:recover", notify)
	}

	c.String(200, `<p>Just close the page</p><script>window.close()</script>`)
}
