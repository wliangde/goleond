#!/bin/bash

_NORMAL="\033[0m"
_YELLOW="\033[0;33m"
_CYAN="\033[1;36m"
_GREEN="\033[1;32m"
_RED="\033[1;31m"
_PERPLE="\033[0;35m"

SERVER_PIDS=()

USER_NAME=`whoami`

check_server_exist()
{
	SERVER_PIDS=()
	pid=`ps -o ruser=$USER_NAME -A -o pid,ppid,cmd | grep -w "$USER_NAME" | grep "./gmweb none" | grep -v grep | grep -v $0 | awk '{print $2}'`
	if [ "$pid" != "" ]
	then
		printf "${_YELLOW}gmweb with config file \"${_PERPLE}none${_YELLOW}\": pid=${_GREEN}${pid}\t${_NORMAL}[${_GREEN}RUNNING${_NORMAL}]\n"
		SERVER_PIDS["gmweb"]=$pid
	else
		printf "${_YELLOW}gmweb with config file \"${_PERPLE}none${_YELLOW}\": \t\t${_NORMAL}[${_RED} STOPED${_NORMAL}]\n"
	fi
}

kill_server()
{
	check_server_exist
	if [ ${#SERVER_PIDS["gmweb"]} -gt 0 ]
	then
		printf "\t${_YELLOW}killing running server${_CYAN} $1 ${_YELLOW}with pid=${SERVER_PIDS["gmweb"]}...${_NORMAL}\t\t"
		kill ${SERVER_PIDS["gmweb"]} || exit 0
		printf "\t[${_RED} KILLED${_NORMAL}]\n"
	fi
}

kill_server
