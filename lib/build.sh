#!/usr/bin/env bash
CGO_CFLAGS="-I/usr/lib/jvm/java-18-openjdk/include -I/usr/lib/jvm/java-18-openjdk/include/linux" go build -buildmode c-shared -o idblib.so main