#!/bin/sh
SRC_PATH="../http__e4ftl01.cr.usgs.gov__SRTM__SRTMGL3.003__2000.02.11"

for F in $(ls -tr $SRC_PATH/*.hgt.zip.xml)
do
	SRC_FILE=$(basename $F | sed  '/[.]zip/s///')
	cp "$F" "$SRC_FILE"
done
