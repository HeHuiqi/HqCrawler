#!/usr/bin/env bash

if [[ $1 == '' ]]; then
    echo "Usage: sh hq_commit your_commit_message"
    exit
fi
git pull
git add .
git commit -m $1
git push
