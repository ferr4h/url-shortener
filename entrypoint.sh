#!/bin/sh

./migrations
if [ $? -ne 0 ]; then
  echo "Migrations failed"
  exit 1
fi

exec ./app