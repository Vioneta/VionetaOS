#!/bin/bash
###
 # @Author:  LinkLeong link@vioneta.com
 # @Date: 2022-06-30 10:08:33
 # @LastEditors: LinkLeong
 # @LastEditTime: 2022-07-01 11:17:54
 # @FilePath: /VionetaOS/shell/delete-old-service.sh
 # @Description:
### 

((EUID)) && sudo_cmd="sudo"

# SYSTEM INFO
readonly UNAME_M="$(uname -m)"

# VionetaOS PATHS
readonly CASA_REPO=Vioneta/VionetaOS
readonly CASA_UNZIP_TEMP_FOLDER=/tmp/vionetaos
readonly CASA_BIN=vionetaos
readonly CASA_BIN_PATH=/usr/bin/vionetaos
readonly CASA_CONF_PATH=/etc/vionetaos.conf
readonly CASA_SERVICE_PATH=/etc/systemd/system/vionetaos.service
readonly CASA_HELPER_PATH=/usr/share/vionetaos/shell/
readonly CASA_USER_CONF_PATH=/var/lib/vionetaos/conf/
readonly CASA_DB_PATH=/var/lib/vionetaos/db/
readonly CASA_TEMP_PATH=/var/lib/vionetaos/temp/
readonly CASA_LOGS_PATH=/var/log/vionetaos/
readonly CASA_PACKAGE_EXT=".tar.gz"
readonly CASA_RELEASE_API="https://api.github.com/repos/${CASA_REPO}/releases"
readonly CASA_OPENWRT_DOCS="https://github.com/Vioneta/VionetaOS-OpenWrt"

readonly COLOUR_RESET='\e[0m'
readonly aCOLOUR=(
    '\e[38;5;154m' # green  	| Lines, bullets and separators
    '\e[1m'        # Bold white	| Main descriptions
    '\e[90m'       # Grey		| Credits
    '\e[91m'       # Red		| Update notifications Alert
    '\e[33m'       # Yellow		| Emphasis
)

Target_Arch=""
Target_Distro="debian"
Target_OS="linux"
Casa_Tag=""


#######################################
# Custom printing function
# Globals:
#   None
# Arguments:
#   $1 0:OK   1:FAILED  2:INFO  3:NOTICE
#   message
# Returns:
#   None
#######################################

Show() {
    # OK
    if (($1 == 0)); then
        echo -e "${aCOLOUR[2]}[$COLOUR_RESET${aCOLOUR[0]}  OK  $COLOUR_RESET${aCOLOUR[2]}]$COLOUR_RESET $2"
    # FAILED
    elif (($1 == 1)); then
        echo -e "${aCOLOUR[2]}[$COLOUR_RESET${aCOLOUR[3]}FAILED$COLOUR_RESET${aCOLOUR[2]}]$COLOUR_RESET $2"
    # INFO
    elif (($1 == 2)); then
        echo -e "${aCOLOUR[2]}[$COLOUR_RESET${aCOLOUR[0]} INFO $COLOUR_RESET${aCOLOUR[2]}]$COLOUR_RESET $2"
    # NOTICE
    elif (($1 == 3)); then
        echo -e "${aCOLOUR[2]}[$COLOUR_RESET${aCOLOUR[4]}NOTICE$COLOUR_RESET${aCOLOUR[2]}]$COLOUR_RESET $2"
    fi
}

Warn() {
    echo -e "${aCOLOUR[3]}$1$COLOUR_RESET"
}

# 0 Check_exist
Check_Exist() {
    #Create Dir
    Show 2 "Create Folders."
    ${sudo_cmd} mkdir -p ${CASA_HELPER_PATH}
    ${sudo_cmd} mkdir -p ${CASA_LOGS_PATH}
    ${sudo_cmd} mkdir -p ${CASA_USER_CONF_PATH}
    ${sudo_cmd} mkdir -p ${CASA_DB_PATH}
    ${sudo_cmd} mkdir -p ${CASA_TEMP_PATH}

   
    Show 2 "Start cleaning up the old version."
    
    ${sudo_cmd} rm -rf /usr/lib/systemd/system/vionetaos.service
    
    ${sudo_cmd} rm -rf /lib/systemd/system/vionetaos.service
    
    ${sudo_cmd} rm -rf /usr/local/bin/${CASA_BIN}
    
    #Clean
    if [[ -d "/casaOS" ]]; then
        ${sudo_cmd} rm -rf /casaOS
    fi
    Show 0 "Clearance completed."    

    $sudo_cmd systemctl restart ${CASA_BIN}
}
Check_Exist
