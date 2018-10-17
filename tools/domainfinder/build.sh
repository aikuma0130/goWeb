#!/bin/bash
echo domainfinderをビルドします...
go build -o domainfinder

echo synonymsをビルドします...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms

echo availableをビルドします...
cd ../available
go build -o ../domainfinder/lib/available

echo sprinkleをビルドします...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle

echo coolifyをビルドします...
cd ../coolify
go build -o ../domainfinder/lib/coolify

echo domainifyをビルドします...
cd ../domainify
go build -o ../domainfinde