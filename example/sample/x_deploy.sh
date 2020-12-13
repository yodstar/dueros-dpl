#! /bin/bash
cd `dirname $0`

echo Build
make rebuild

echo Deploy
systemctl stop sample
make deploy

echo Complete
systemctl daemon-reload
systemctl start sample
