#!/usr/bin/env bash

# general config
NULL='/dev/null'
REDIS_CLI=$(command -v redis-cli) || exit 1
UUIDGEN=$(command -v uuidgen) || exit 1
MIN_ARGS=2

# redis config
HOST=127.0.0.1
PORT=63799
DB=1

_help() {
    echo -e "
Usage: $0 [OPTION]
Example: $0 -d ExemplaCustomer -t

OPTIONS:
  -d \t\t Description of the new key
  -t \t\t To create a api key to test
"
}

_is_linux() {
    [[ "${OSTYPE}" == "linux-gnu" ]] 
}

_is_freebsd() {
    [[ "${OSTYPE}" == "FreeBSD" ]] 
}

_create_api_key_test() {
    _create_api_key $@ &&
        ${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} sadd api_keys:enabled_test ${KEY} &> ${NULL}
}

_create_api_key() {
    ${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} sadd api_keys:enabled ${KEY} &> ${NULL} &&
        ${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} hset ${KEY} desc "${DESC}" enable true count 0 count_auth 0 net_allowed 0.0.0.0/0 &> ${NULL}
}

if [[ ${#@} -lt ${MIN_ARGS} ]]; then
    _help
    exit 1
fi

while getopts "d:t" OPT; do
  case "$OPT" in
    "d")  DESC=${OPTARG};; 
    "t") TEST=true;;
  esac
done

if _is_linux; then
    api_key="$(${UUIDGEN} -r)"
elif _is_freebsd; then
    api_key="$(${UUIDGEN})"
fi

[[ -n ${DESC} ]] || exit 1
[[ -n ${api_key} ]] || exit 1

KEY="api_key:${api_key}"

if [[ -n ${TEST} ]]; then
    _create_api_key_test && echo "${KEY}"
else
    _create_api_key && echo "${KEY}"
fi
