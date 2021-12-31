set -e
image="haitwang/golang-echo-sample"
tag=$image:echo
docker build -t $tag .