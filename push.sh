#!/bin/bash

echo "commit name: "
read commit
git add .
git commit -m "$commit"
git push origin main