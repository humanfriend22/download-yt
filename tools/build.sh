#/bin/bash

cd ..

go install github.com/machinebox/appify
go build
appify -name "Download YT" -icon ./assets/youtube-logo-greyscale.png ./main

clear
echo "Packaged!"