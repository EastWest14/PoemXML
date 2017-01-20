#!/bin/bash
go test ./...
[ $? -eq 0 ] || exit $?;
echo Unit Tests Pass!

echo Building the server.
go build poemserver/main.go
[ $? -eq 0 ] || exit $?;

echo Launching the server
./main 'poemserver/index/test_indices/regression_test_index.xml' &
ID=$!;
[ $? -eq 0 ] || exit $?;

cd regression/Java_Regression/src
sleep 3
java Tester.Tester
[ $? -eq 0 ] || exit $?;
echo Regression Tests Pass!
cd -

pkill -P $$;
echo 'Run Complete!';
