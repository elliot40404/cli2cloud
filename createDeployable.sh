#!/usr/bin/env bash

mkdir -p deploy
cd client
cp -r dist ../deploy
cd ../server
rsync -a --exclude={'node_modules/','db/'} . ../deploy