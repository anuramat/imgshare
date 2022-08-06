#!/usr/bin/env bash

DBSTRING="host=$DBHOST port=$DBPORT user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=$DBSSL"

goose postgres "$DBSTRING" up