#!/bin/sh

# Esperar a SurrealDB (puerto 8000)
while ! nc -z surrealdb 8000; do
  echo "Esperando a SurrealDB..."
  sleep 1
done

# Esperar a Minio (puerto 9000)
while ! nc -z minio 9000; do
  echo "Esperando a Minio..."
  sleep 1
done

echo "Todos los servicios están listos!"
exec "$@"