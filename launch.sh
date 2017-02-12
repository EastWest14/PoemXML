#!/bin/sh

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
 
