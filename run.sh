#!/bin/bash

go_file="main.go"


if [ -r "$go_file" ]; then
    

    
    if [ -f "$go_file" ]; then
        

        
        if [ -s "$go_file" ]; then
           

            
            chmod +x "$go_file"

            
            if [ -x "$go_file" ]; then
                
                echo "Running $go_file"
                go run "$go_file"
            else
                echo "Go file '$go_file' is not executable."
            fi
        else
            echo "Go file '$go_file' is empty."
        fi
    else
        echo "Go file '$go_file' is not a regular file."
    fi
else
    echo "Go file '$go_file' is not readable."
fi
