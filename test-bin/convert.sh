#!/bin/sh
SRC_PATH="../http__e4ftl01.cr.usgs.gov__SRTM__SRTMGL3.003__2000.02.11"

for F in $(ls -tr $SRC_PATH/*.hgt.zip)
do
	SRC_FILE=$(basename $F | sed  '/[.]zip/s///')
	unzip -p "$SRC_PATH/$SRC_FILE.zip" | lz4 -9 - $SRC_FILE.lz4
done
