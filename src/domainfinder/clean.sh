#!/bin/bash

echo "clean domainfinder..."
rm domainfinder

libs=("synonyms" "available" "sprinkle" "coolify" "domainify")

for lib in "${libs[@]}"
do
echo "clean ${lib} ..."
rm ../domainfinder/lib/${lib}
done

echo "clean finished!!"




