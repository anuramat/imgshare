#!/usr/bin/env sh

DBSTRING="host=$DBHOST port=$DBPORT user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=$DBSSL"

./bin/goose -dir migrations postgres "$DBSTRING" up