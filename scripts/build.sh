#!/bin/sh

echo "Building Dexer to default location..."

# Cria uma nova versão recente
go build -o ./build/dexer

# Cria uma cópia da versão recente com timestamp
cp ./build/dexer ./build/dexer-$(date -d "today" +"%Y-%m-%d-%H-%M-%S")

echo "Build finished."

if [ "$1" == "--exec" ]; then
    echo "Executing latest build binary."
    exec ./build/dexer
elif [ "$1" == "--clear" ]; then
    echo "Cleaning build folder."
    find ./build -type f -not -name "dexer" -print0 | xargs -0 rm --
    echo "Done."
fi

exit 0
