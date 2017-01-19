#!/bin/bash
go test ./...
[ $? -eq 0 ] || exit $?;
echo Unit Tests Pass!

go run poemserver/main.go 'poemserver/index/test_indices/regression_test_index.xml' &
ID=$!;
#[ $? -eq 0 ] || echo Failed to launch the server!; exit $?;

cd regression/Java_Regression/src
sleep 3
java Tester.Tester
[ $? -eq 0 ] || exit $?;
echo Regression Tests Pass!
cd -

pkill -P $$;
echo 'Run Complete!';
