/*
 * @Author: LinkLeong link@vioneta.org
 * @Date: 2022-08-24 17:37:36
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-08-24 17:38:48
 * @FilePath: /VionetaOS/interfaces/migrationTool.go
 * @Description:
 * @Website: https://www.vionetaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package interfaces

type MigrationTool interface {
	IsMigrationNeeded() (bool, error)
	PostMigrate() error
	Migrate() error
	PreMigrate() error
}
