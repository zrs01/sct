#!/bin/bash
cd $(dirname $0)
go run main.go print -n $1
