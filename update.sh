#!/bin/bash

INIT_PATH="init"
BUILD_NAME="server"

BRANCH = "main"

PID=$(sudo lsof -i | grep $BUILD_NAME | awk '{print $2}')

echo $PID

kill -9 $PID

git checkout $BRAANCH
git pull origin $BRANCH

CD $INIT_PATH
GO build -o $BUILD_NAME

echo $PWD

nohup ./$BUILD_NAME &