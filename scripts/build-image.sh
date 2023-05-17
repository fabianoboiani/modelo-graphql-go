#!/bin/bash

workdir="$(pwd)"
docker_dir=${workdir}/build/package/docker

name=$1
version=$2
arch=$3
suffix=$4
split=(${arch//\// })
OS_=${split[0]}
ARCH_=${split[1]}
tags="--tag  $name:$version --tag $name:latest"
if [ -n "$suffix" ]; then
  tags="$tags --tag $name:$version$suffix"
fi

docker build ${tags} -f "${docker_dir}/Dockerfile"  --build-arg OS_NAME=${OS_}  --build-arg ARCH=${ARCH_} "${workdir}"