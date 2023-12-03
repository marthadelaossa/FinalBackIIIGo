#!/bin/sh

wait-for "${DB_HOST}:${DB_PORT}" -- "$@"

exec "./app"
