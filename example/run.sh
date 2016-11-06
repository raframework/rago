#!/usr/bin/env bash

LOG_FILE=/tmp/rago_example.log
PID_FILE=/tmp/rago_example.pid

kill `cat $PID_FILE`

go install github.com/raframework/rago/example
$GOPATH/bin/example >> $LOG_FILE 2>&1 &
echo $! > $PID_FILE