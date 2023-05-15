#!/bin/sh

set -e

REDIS_HOST=${REDIS_HOST:-"localhost"}
REDIS_PORT=${REDIS_PORT:-"6379"}
MAX_RETRIES=10

# Wait for Redis to be up and running
retry_counter=0
echo "Waiting for Redis to be available at ${REDIS_HOST}:${REDIS_PORT}..."
while ! nc -z "$REDIS_HOST" "$REDIS_PORT"; do
    sleep 1
    retry_counter=$((retry_counter + 1))
    if [ $retry_counter -ge $MAX_RETRIES ]; then
        echo "Redis is not available after ${MAX_RETRIES} retries. Exiting."
        exit 1
    fi
    echo "Retrying Redis connection (${retry_counter}/${MAX_RETRIES})..."
done

echo "Redis is available. Starting the Go microservice..."
exec ./main
