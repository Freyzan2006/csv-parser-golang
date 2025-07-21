#!/bin/bash

set -e

go run cmd/main.go -file ./test.csv -required "Name,Age" -range "Age:18-70" -validate-type "Age:int" -verbose
