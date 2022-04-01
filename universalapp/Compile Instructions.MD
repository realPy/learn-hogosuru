I'm using linux for my env development.
If you are using another OS please refer the doc of https://github.com/webview/webview


LINUX:
Be sure to have the webkit2gtk-4.0 package 

#sudo apt-get install webkit2gtk-4.0

Compile you universal app with:

#go build -o universalapp/bin/linux_app universalapp/server/main.go


WINDOWS via cross compilation:

Be sure to have the mingwin package installed :
#sudo apt-get install mingw-w64

Compile you universal app with:

#CC=/usr/bin/x86_64-w64-mingw32-gcc GOOS=windows CGO_ENABLED=1  go build -ldflags="-H windowsgui" -o universalapp/bin/app.exe universalapp/server/main.go 


Compile via cross compilation for mac:
Follow the https://github.com/tpoechtrager/osxcross instruction

#OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.13 CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o universalapp/bin/mac_app universalapp/server/main.go 

