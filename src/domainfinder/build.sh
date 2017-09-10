#!/bin/bash

echo "build domainfinder..."
go build -o domainfinder

libs=("synonyms" "available" "sprinkle" "coolify" "domainify")

for lib in "${libs[@]}"
do
echo "build ${lib} ..."
cd ../${lib}
go build -o ../domainfinder/lib/${lib}
done

echo "build finished!!"




