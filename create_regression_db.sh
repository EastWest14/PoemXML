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

DB_NAME='Hello'
DB_STARTUP_SCRIPT='CREATE DATABASE Hello;'
DB_REMOVE_SCRIPT='DROP DATABASE Hello;'

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

	set -e
        echo $DB_STARTUP_SCRIPT | mysql -h$DB_HOST -u $DB_USER -p$DB_PASSWORD
        cat pet_db.sql | mysql -h$DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME
        echo 'Database exists!'
        echo $DB_REMOVE_SCRIPT | mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME
        set +e
else 
	echo 'Running on a non-Linux OS. Assuming local environment.'
	set -e
	echo $DB_STARTUP_SCRIPT | mysql -u $DB_USER -p$DB_PASSWORD
	cat pet_db.sql | mysql -u $DB_USER -p$DB_PASSWORD $DB_NAME
	echo 'Database exists!'
	echo $DB_REMOVE_SCRIPT | mysql -u $DB_USER -p$DB_PASSWORD $DB_NAME
	set +e
fi
