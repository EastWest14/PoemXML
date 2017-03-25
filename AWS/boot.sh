#!/bin/bash

set -e
cd project/src/poemXML
source ./set_env_var.sh
export GOPATH=/home/ec2-user/project
git pull origin master
set +e
