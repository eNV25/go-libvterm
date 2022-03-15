#!/bin/sh -e

ver=0.2
curl -sSLO "https://www.leonerd.org.uk/code/libvterm/libvterm-$ver.tar.gz"
bsdtar xf "libvterm-$ver.tar.gz"

cd "libvterm-$ver"
	make
	cp ./include/*.h ./src/*.h ./src/*.c ./src/*.inc ./src/encoding/*.inc ../
	cp ./LICENSE ../libvterm-LICENSE
cd ..

rm -rf "libvterm-$ver.tar.gz" "libvterm-$ver"
sed -E 's|^#include "encoding/([^/]*)"$|#include "\1"|g' -i encoding.c
! grep '#include "[^/].*/.*"' ./*.c

