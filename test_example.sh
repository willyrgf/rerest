#!/usr/bin/env bash

# general config
NULL='/dev/null'
REDIS_CLI=$(command -v redis-cli) || exit 1
CURL=$(command -v curl) || exit 1
WORKDIR=$(dirname $0)
TESTDATADIR="${WORKDIR}/testdata"
TESTDATAFILE="${WORKDIR}/testdata/example_doc.csv"
DELIMITER=";"
SUBDELIMITER=","
APIKEY=""
URLBASE="http://127.0.0.1:8088"

# redis config
HOST=127.0.0.1
PORT=63799
DB=0


_is_linux() {
    [[ "${OSTYPE}" == "linux-gnu" ]] 
}

_is_freebsd() {
    [[ "${OSTYPE}" == "FreeBSD" ]] 
}

_insert_redis_example_data() {
    [[ -n "$1" ]] && doc="$1" || exit 1
    [[ -n "$2" ]] && phone_number="$2" || exit 1

    ${REDIS_CLI} -h ${HOST} -p ${PORT} -n ${DB} sadd doc:${doc} ${phone_number}
}

_create_api_key() {
    APIKEY="$(${WORKDIR}/create_apikey.sh -d 'test_example.sh APIKEY' -t)"
    [[ -n ${APIKEY} ]]
}

_insert_example_data() {
    [[ -f ${TESTDATAFILE} ]] || exit 1

    for l in $(sed 1d ${TESTDATAFILE}); do
        doc="$(echo ${l} | cut -f1 -d"${DELIMITER}")"
        for n in $(echo ${l} | cut -f2 -d"${DELIMITER}" | sed "s/${SUBDELIMITER}/ /g"); do
            _insert_redis_example_data "${doc}" "${n}" || exit 1
        done
    done
}

_test_api_request() {
    [[ -f ${TESTDATAFILE} ]] || exit 1
    apikey="$(echo "${APIKEY}" | cut -f2 -d':')"

    for l in $(sed 1d ${TESTDATAFILE}); do
        doc="$(echo ${l} | cut -f1 -d"${DELIMITER}")"
        phone_numbers="$(echo ${l} | cut -f2 -d"${DELIMITER}" | sed "s/${SUBDELIMITER}/ /g" | xargs -n 1 | sort)"
        ${CURL} -s -k -L "${URLBASE}/doc/${doc}?api_key=${apikey}" | grep -q "${phone_numbers}" || exit 1
    done 
}


_create_api_key || exit 1
_insert_example_data || exit 1
_test_api_request || exit 1
