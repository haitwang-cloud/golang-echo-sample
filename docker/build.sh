set -e
image="golang-echo-sample"
tag=$image:latest
docker build -t $tag .