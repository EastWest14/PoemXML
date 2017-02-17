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
echo 'Preparing the index directory copy:'
set -e
original_regression_index_path="./poemserver/index/test_indices/regression_test_index/*"

namebase="index_workcopy_"
datepart=$(date +%s%N)
new_dir_name="$namebase$datepart"

temp_index_dir=$(mktemp -d $new_dir_name)
cp -R $original_regression_index_path $temp_index_dir
set +e
echo 'Index Directory Copy Created'
echo '===================================='
echo ''

echo '===================================='
echo 'Launching the Server:'
set -e
index_name="regression_test_index.xml"
path_to_index_xml="$temp_index_dir/$index_name"
./main $path_to_index_xml &
ID=$!;
set +e
sleep 2
echo 'Server Launched'
echo '===================================='
echo ''

echo '===================================='
echo 'Launching Regression suit:'
cd regression/Java_Regression/src
sleep 1
java Tester.Tester
cd -
echo 'Regression Suit Passes'
echo '===================================='
echo ''

echo '===================================='
echo 'Shutting down the server:'
set -e
echo 'Process ID: ' $ID
kill $ID
set +e
echo 'Server Shut Down'
echo '===================================='
echo ''

echo '===================================='
echo 'Removing index regression copy:'
set -e
cd $temp_index_dir
rm *
cd -
rmdir $temp_index_dir
set +e
echo '`index regression copy removed'
echo '===================================='
echo ''

echo '###ALL TESTS PASS###'
