/*
 * @Author: LinkLeong link@vioneta.com
 * @Date: 2021-09-30 18:18:14
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-08-31 17:04:02
 * @FilePath: /CasaOS/pkg/config/config.go
 * @Description:
 * @Website: https://www.vionetaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package config

import (
	"path/filepath"

	"github.com/Vioneta/VionetaOS-Common/utils/constants"
)

var CasaOSConfigFilePath = filepath.Join(constants.DefaultConfigPath, "vionetaos.conf")
