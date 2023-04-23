#!/bin/sh

if [ $# == 0 ]; then
    docker-compose up
elif [ "$1" == "-b" ]; then
    docker-compose up --build
elif [ "$1" == "-db" ]; then
    docker-compose up --build -d
elif [ "$1" == "-d" ]; then
    docker-compose up -d
else
    echo "Invalid argument."
fi
