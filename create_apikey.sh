#!/usr/bin/env bash

HOST=127.0.0.1
PORT=63799
DB=1

REDIS_CLI=$(command -v redis-cli) || exit 1

api_key="$(uuidgen -r)" || exit 1

key="api_key:${api_key}"


${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} sadd api_keys:enabled ${key}
${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} hset ${key} enable true count 0 count_auth 0
