#!/bin/bash
ulimit -c unlimited

_NORMAL="\033[0m"
_YELLOW="\033[0;33m"
_CYAN="\033[1;36m"
_GREEN="\033[1;32m"
_RED="\033[1;31m"
_PERPLE="\033[0;35m"

LAUNCHER_SVR="gmweb"
LAUNCHER_SVR_CFG="none"
LAUNCHER_ERR="../../../sggamelog/gm/gmweb_errors.log"

SERVER_ARRAY=([1]=${LAUNCHER_SVR})
SERVER_CFG_ARRAY=([1]=${LAUNCHER_SVR_CFG})
SERVER_ERR_ARRAY=([1]=${LAUNCHER_ERR})
SERVER_PIDS=()



USER_NAME=`whoami`

print_success()
{
	printf "${_NORMAL}[${_GREEN}SUCCESS${_NORMAL}]\n"
}

print_failed()
{
	printf "${_NORMAL}[${_RED} FAILED${_NORMAL}]\n"
}

check_server_exist()
{
	for index in ${!SERVER_ARRAY[@]}; do
		if [ ${SERVER_ARRAY[${index}]} == ${1} ]; then
			local __server_cfg_arr=(${SERVER_CFG_ARRAY[${index}]})
			local __server_cfg_num=${#__server_cfg_arr[@]}
			SERVER_PIDS=() 
			printf "${_YELLOW}check if server${_CYAN} $1 ${_YELLOW}running... \n"
			for cfg_file in ${__server_cfg_arr[@]}; do
				pid=`ps -A -o ruser=$USER_NAME -o pid,ppid,cmd | grep -w "$USER_NAME" | grep "./${1} ${cfg_file}" | grep -v grep | grep -v $0 | awk '{print $2}'`
				if [ "$pid" != "" ]
				then
					printf "\t${_YELLOW}with config file \"${_PERPLE}${cfg_file}${_YELLOW}\": pid=${_GREEN}${pid}\t${_NORMAL}[${_GREEN}RUNNING${_NORMAL}]\n"
					SERVER_PIDS[${#SERVER_PIDS[@]}]=$pid
				else
					printf "\t${_YELLOW}with config file \"${_PERPLE}${cfg_file}${_YELLOW}\": \t\t${_NORMAL}[${_RED} STOPED${_NORMAL}]\n"
				fi
			done
			
			break
		fi
	done
}

kill_server()
{
    check_server_exist $1
    if [ ${#SERVER_PIDS[@]} -gt 0 ]
    then
        printf "${_YELLOW}killing running server${_CYAN} $1 ${_YELLOW}with pid=${SERVER_PIDS[@]}...${_NORMAL}\t\t"
        kill -9 ${SERVER_PIDS[@]} || exit 0
        printf "[${_RED} KILLED${_NORMAL}]\n"
    fi
}

kill_all_server()
{
    printf "${_YELLOW}################           stop${_CYAN} all ${_YELLOW}server           ################${_NORMAL}\n"

    kill_server &LAUNCHER_SVR

    echo "all server killed"
}

start_server_fail()
{
    printf "\n${_RED}start server failed, will stop all started server${_NORMAL}\n"
    kill_all_server
    exit 0
}

start_server()
{
    for index in ${!SERVER_ARRAY[@]}; do
        if [ ${SERVER_ARRAY[${index}]} == ${1} ]; then
            local __server_cfg_arr=(${SERVER_CFG_ARRAY[${index}]})
            local __server_cfg_num=${#__server_cfg_arr[@]}
            
            printf "${_YELLOW}starting server ${_CYAN}$1${_YELLOW} ...${_NORMAL}\n"
            for cfg_file in ${__server_cfg_arr[@]}; do
				printf "\t${_YELLOW} executing \"./${1} ${cfg_file}\"...${_NORMAL}\t\t"
				echo "executing \"./${1} ${cfg_file}\"...`date`" >> ${SERVER_ERR_ARRAY[${index}]}
				nohup ./${1} ${cfg_file} >/dev/null 2>> ${SERVER_ERR_ARRAY[${index}]} & > /dev/null  && print_success || start_server_fail
            done
    
            break
        fi
    done
}

start_all_server()
{
    printf "${_YELLOW}################          start${_CYAN} all ${_YELLOW}server           ################${_NORMAL}\n" 
	
	check_server_exist $LAUNCHER_SVR
    if [ ${#SERVER_PIDS[@]} -lt 1 ]; then start_server $LAUNCHER_SVR ; fi
    sleep 1
	
	check_server_exist $LAUNCHER_SVR
    if [ ${#SERVER_PIDS[@]} -lt 1 ]
    then 
      printf "${_RED} $LAUNCHER_SVR server start fail${_NORMAL}\n"
    fi
}

start_all_server
