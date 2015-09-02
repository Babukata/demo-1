#!/bin/bash

echo "Installing Go...";
yum install -y go

echo "Compiling application...";
cd /tmp/build && go build -o /opt/app 

