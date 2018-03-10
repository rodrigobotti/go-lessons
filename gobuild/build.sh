# build directory: outputs same name
go build

# build with output name
go build -o <name>

# cross-compilation
GOOS=windows GOARCH=386 go build -o build/app-windows.exe

GOOS=linux GARCH=arm go build -o build/app-raspberry \
  -v \ # verbose
  $GOPATH/src/<path_to_application>
