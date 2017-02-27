#!/bin/bash

echo '' 
echo 'RUNNING THE UNIT + REGRESSION TESTS:'

echo '===================================='
echo 'Running the Unit Tests:'
go test ./...
[ $? -eq 0 ] || exit $?;
echo 'Unit Tests Pass'
echo '===================================='
echo ''

echo '===================================='
echo 'Building the Server:'
go build poemserver/main.go
[ $? -eq 0 ] || exit $?;
echo 'Server Built Succesfully'
echo '===================================='
echo ''

echo '===================================='
echo 'Recompiling Regression Suit:'
set -e
./recompile_regression.sh
set +e
echo 'Regression Recompiled Succesfully'
echo '===================================='
echo ''


echo '===================================='
echo 'Preparing the index directory copy:'
set -e
original_regression_index_path="./poemserver/index/test_indices/regression_test_index/*"

namebase="index_workcopy_"
datepart=$(date +%s%N)
new_dir_name="XXX$namebase$datepart"

temp_index_dir=$(mktemp -d $new_dir_name)
cp -R $original_regression_index_path $temp_index_dir
set +e
echo 'Index Directory Copy Created'
echo '===================================='
echo ''



echo 'Creating Regression DB:'
echo '===================================='
echo ''
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

export DB_NAME='Hello'
DB_STARTUP_SCRIPT="CREATE DATABASE $DB_NAME;"
DB_REMOVE_SCRIPT="DROP DATABASE $DB_NAME;"
if ! [ $DB_HOST ]
then
  	echo 'DB_HOST environment variable not set will not be able to connect to DB.'
    echo 'Set DB_HOST environment variable and preferable put it in a script, for example set_env_var.sh.'
    exit 1
fi

set -e
echo $DB_STARTUP_SCRIPT | mysql -h$DB_HOST -u $DB_USER -p$DB_PASSWORD
cat pet_db.sql | mysql -h$DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME
echo 'Regression DB Created'
echo '===================================='
echo ''



echo '===================================='
echo 'Launching the Server:'
set -e
index_name="regression_test_index.xml"
path_to_index_xml="$temp_index_dir/$index_name"
./main $path_to_index_xml $DB_HOST $DB_USER $DB_NAME $DB_PASSWORD &
ID=$!;
set +e
sleep 2
echo 'Server Launched'
echo '===================================='
echo ''


echo '===================================='
echo 'Launching Regression suit:'
cd regression/Java_Regression/src -> /dev/null
sleep 1
java Tester.Tester
cd - -> /dev/null
echo '===================================='
echo ''


echo '===================================='
echo 'Shutting down the server:'
set -e
echo 'Process ID: ' $ID
kill $ID
sleep 1
set +e
echo 'Server Shut Down'
echo '===================================='
echo ''


echo '===================================='
echo 'Removing index regression copy:'
set -e
cd $temp_index_dir -> /dev/null
rm *
cd - -> /dev/null
rmdir $temp_index_dir
set +e
echo 'Index regression copy removed'
echo '===================================='
echo ''


echo '===================================='
echo 'Removing regression DB:'
set -e
echo $DB_REMOVE_SCRIPT | mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME
set +e
echo 'Removed regression database'
echo '===================================='
echo ''


echo '###REGRESSION RUN COMPLETE###'
