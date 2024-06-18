#!/bin/bash

export VECIN_DB_PASSWORD=${VECIN_DB_PASSWORD}
export VECIN_DB_USER=${VECIN_DB_USER}
export VECIN_DB=${VECIN_DB}
export VECIN_DB_HOST=${VECIN_DB_HOST}
export PORT=8180
export PGPORT=5432

docker-compose -f docker-compose.postgres.yml up --build

exit
