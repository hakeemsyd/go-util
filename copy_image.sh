#!/bin/sh
IMAGESOURCE=$1
DEST=$2

if test -z "$IMAGESOURCE" 
then
      echo "IMAGESOURCE is empty"
      exit 1
fi

if test -z "$DEST" 
then
      echo "DEST is empty"
      exit 1
fi

echo "IMAGESOURCE: " $IMAGESOURCE
echo "DEST: " $DEST

skopeo copy $IMAGESOURCE dir:$DEST