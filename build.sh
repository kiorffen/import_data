#!/bin/bash


BIN_NAME=import_data

go build -o ${BIN_NAME} main.go
mv ${BIN_NAME} output/bin
