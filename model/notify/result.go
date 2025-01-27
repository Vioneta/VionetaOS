/*
 * @Author: LinkLeong link@vioneta.com
 * @Date: 2022-05-26 14:21:11
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-05-27 11:15:59
 * @FilePath: /VionetaOS/model/notify/result.go
 * @Description:
 * @Website: https://www.vionetaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */

package notify

// Notify struct for Notify
type NotifyModel struct {
	Data  interface{} `json:"data"`
	State string      `json:"state"`
}
