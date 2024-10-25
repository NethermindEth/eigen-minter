#!/bin/sh
set -e

# Execute the application with the APP_NAME environment variable
exec /app/"$APP_NAME" "$@"
