#!/usr/bin/env bash

while ./butler status $1 | grep " • ";
do
  echo "Still deploying $1 ..."
  sleep 1
done
echo "Deployed $1 successfully!"
