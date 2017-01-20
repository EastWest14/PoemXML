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
echo 'Launching the Server:'
./main 'poemserver/index/test_indices/regression_test_index.xml' &
ID=$!;
sleep 2
echo 'Server Launched'
echo '===================================='
echo ''

echo '===================================='
echo 'Launching Regression suit:'
cd regression/Java_Regression/src
sleep 1
java Tester.Tester
[ $? -eq 0 ] || exit $?;
cd -
echo 'Regression Suit Passes'
echo '===================================='
echo ''

echo '===================================='
echo 'Shutting down the server:'
pkill -P $$;
[ $? -eq 0 ] || exit $?;
echo 'Server Shut Down'
echo '===================================='
echo ''

echo '###ALL TESTS PASS###'
