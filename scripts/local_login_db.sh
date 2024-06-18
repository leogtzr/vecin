#!/bin/bash

set -u

psql -U "${VECIN_DB_USER}" -h localhost -p "${VECIN_DB_PORT}" -d "${VECIN_DB_HOST}"

exit 0