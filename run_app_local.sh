#!/bin/bash

set -e

export VECIN_DB_PASSWORD=${VECIN_DB_PASSWORD}
export VECIN_DB_USER=${VECIN_DB_USER}
export VECIN_DB=${VECIN_DB}
export VECIN_DB_HOST=${VECIN_DB_HOST}
export PORT=8180
export PGPORT=5432
export MAILSENDER_API_KEY=${MAILSENDER_API_KEY}

readonly error_wrong_option=81
readonly binary_file="vecin"

case "${1}" in
    build|b)
        echo "Building project first..."
        make clean
        make
        echo "Done..."

        ;;

    *)
        ;;
esac

./"${binary_file}"

exit
