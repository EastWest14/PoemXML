#!/bin/bash

if ! [ $DB_USER ]
then
	echo 'DB_USER environment variable not set will not be able to connect to DB.'
	echo 'Set DB_USER environment variable and preferable put it in a script, for example set_env_var.sh.'
	exit 1
fi

if ! [ $DB_PASSWORD ]
then
        echo 'DB_PASSWORD environment variable not set will not be able to connect to DB.'
        echo 'Set DB_PASSWORD environment variable and preferable put it in a script, for example set_env_var.sh.'
        exit 1
fi

if [ $OSTYPE = 'linux-gnu' ]
then
	echo 'Running on Linux OS. Assuming EC2 environment.'
	if ! [ $DB_HOST ]
	then
        	echo 'DB_HOST environment variable not set will not be able to connect to DB.'
        	echo 'Set DB_HOST environment variable and preferable put it in a script, for example set_env_var.sh.'
        	exit 1
	fi
	mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD -e "show databases;"
else 
	echo 'Running on a non-Linux OS. Assuming local environment.'
	mysql -u$DB_USER -p$DB_PASSWORD -e "show databases;"
fi
