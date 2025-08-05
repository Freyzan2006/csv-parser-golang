#!/bin/bash

set -e



# go run cmd/main.go -file ./test.csv \
#    -header=true -required "Name,Age" -range "Age:18-70" -validate-type "Age:int" -filter "Age>=30" \
#    -out-format "json" -sort "Age:asc" -aggregate "col_3:sum,col_3:avg,col_4:max" -verbose 


go run cmd/main.go -file ./test.csv \
  -header=true \
  -aggregate=Age:sum,Price:avg 
