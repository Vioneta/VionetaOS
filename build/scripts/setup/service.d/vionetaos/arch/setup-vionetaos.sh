#!/bin/bash
###
# @Author: LinkLeong link@vioneta.org
# @Date: 2022-08-25 11:41:22
 # @LastEditors: LinkLeong
 # @LastEditTime: 2022-08-31 17:54:17
 # @FilePath: /CasaOS/build/scripts/setup/service.d/vionetaos/debian/setup-vionetaos.sh
# @Description:

# @Website: https://www.vionetaos.io
# Copyright (c) 2022 by icewhale, All Rights Reserved.
###

set -e

APP_NAME="vionetaos"

# copy config files
CONF_PATH=/etc/vionetaos
OLD_CONF_PATH=/etc/vionetaos.conf
CONF_FILE=${CONF_PATH}/${APP_NAME}.conf
CONF_FILE_SAMPLE=${CONF_PATH}/${APP_NAME}.conf.sample


if [ -f "${OLD_CONF_PATH}" ]; then
    echo "copy old conf"
    cp "${OLD_CONF_PATH}" "${CONF_FILE}"
fi
if [ ! -f "${CONF_FILE}" ]; then
    echo "Initializing config file..."
    cp -v "${CONF_FILE_SAMPLE}" "${CONF_FILE}"
fi

rm -rf /etc/systemd/system/vionetaos.service # remove old service file

systemctl daemon-reload

# enable service (without starting)
echo "Enabling service..."
systemctl enable --force --no-ask-password "${APP_NAME}.service"
