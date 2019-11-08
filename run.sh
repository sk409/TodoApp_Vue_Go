#!/bin/sh
while getopts d:w:bs opt
do
    case "$opt" in
    "w" ) w="$OPTARG";;
    "d" ) d="$OPTARG";;
    "b" ) b="true";;
    "s" ) s="true";;
    esac
done

if [ -z $w ];
then
w=5000
fi

if [ -z $d ];
then
d=4000
fi

echo "WEB_SERVER_PORT=$w\nDATABASE_SERVER_PORT=$d" > .env

cat << EOS > web/app/src/constants.js
export const serverPortNumber = $d;
EOS

if [ "$s" = "true" ];
then
    docker-compose down
fi

if [ "$b" = "true" ];
then
    docker-compose up -d --build
else
    docker-compose up -d
fi

echo "$w => WebServer"
echo "$d => DatabaseServer"
echo "\033[0;32mYou can access TodoApp at localhost:$w\033[0;39m"
