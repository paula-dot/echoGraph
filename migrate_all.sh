#!/bin/bash

# Set your database URL here
DB_URL="postgres://echograph_user:securepassword@localhost:5432/echograph?sslmode=disable"
MIGRATIONS_DIR="db/migrations"

# Find all .up.sql files, sort them, and apply each with psql
for file in $(ls "$MIGRATIONS_DIR"/*.up.sql 2>/dev/null | sort); do
    echo "Applying migration: $file"
    psql "$DB_URL" -f "$file"
    if [ $? -ne 0 ]; then
        echo "Error applying $file. Stopping."
        exit 1
    fi
done

echo "All migrations applied successfully."

