#!/bin/bash

mkdir -p cmd
    mkdir -p cmd/server
        touch cmd/server/server.go
    touch cmd/main.go
mkdir -p config
    touch config/config.go
mkdir -p pkg
    mkdir -p ./pkg/db
    mkdir -p ./pkg/node
    mkdir -p ./pkg/transaction
    mkdir -p ./pkg/util
    mkdir -p ./pkg/signer
touch sample.env.json
touch .gitignore
touch LICENSE
touch README.md