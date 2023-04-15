#!/bin/bash
# vim:sw=4:ts=4:et:ai:ci:sr:nu:syntax=sh
##############################################################
# Usage ( * = optional ):                                    #
# ./script.sh *<db-address> *<db-port> *<username> *<password> #
##############################################################

if [ ! -z "$3" ]; then
    if [ ! -z "$4" ]; then
        echo "Using password authentication!"
        auth="--authenticationDatabase admin -u $3 -p $4"
    fi
fi

HOST=${1:-localhost} # default server is the localhost
PORT=${2:-27017}     # default port for MongoDB is 27017

for directory in *; do
    if [ -d "${directory}" ] ; then
        echo "$directory"
        for data_file in $directory/*; do
            mongoimport --drop --host $HOST --port $PORT --db "$directory" --collection "$(basename $data_file .json)" --file $data_file $auth
        done
    fi
done
