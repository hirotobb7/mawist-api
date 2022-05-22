#!/bin/bash

# @args $1: directory path
function create_table() {
  aws dynamodb create-table --cli-input-yaml file://$1/table-definition.yml --endpoint ${DB_ENDPOINT}

  return 0
}

function create_tables() {
  for path in ./aws/dynamodb/*; do
    if [ -d $path ]; then
      create_table $path
    fi
  done
}

function main() {
 create_tables
}

main
