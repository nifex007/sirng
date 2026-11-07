#!/bin/sh
set -eu
if ! awslocal s3api head-bucket --bucket sirng-evidence >/dev/null 2>&1; then
  awslocal s3api create-bucket --bucket sirng-evidence --region us-east-1
fi
