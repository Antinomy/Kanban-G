#!/bin/sh

current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

echo "Git working on path : $current_path"

cd $current_path

git add -A

git commit  -m "$1"

git push origin master