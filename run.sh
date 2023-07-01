#!/bin/bash

go_file="main.go"

if [ -r "$go_file" ] && [ -f "$go_file" ] && [ -s "$go_file" ]; then
    chmod +x "$go_file"
    
    if [ -x "$go_file" ]; then
        echo "Running $go_file"
        go run "$go_file"
    else
        echo "Go file '$go_file' is not executable."
    fi
else
    echo "Go file '$go_file' is either not readable, not a regular file, or empty."
fi

cd testapp && npm run dev
