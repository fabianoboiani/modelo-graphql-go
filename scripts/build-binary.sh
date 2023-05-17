#!/bin/bash

version=$1
commit_sha=$2
workdir="$(pwd)"
binary_dir=${workdir}/target/bin
archs=`cat ${workdir}/scripts/plataforms`

CGO_ENABLED=0

for arch in ${archs}
do
  split=(${arch//\// })
  GOOS=${split[0]}
  GOARCH=${split[1]}
  if [ -f "/etc/alpine-release" -a "${GOOS}" = "linux" ]; then
    TAGS="-tags musl"
  fi

  go build -mod=vendor ${TAGS} -ldflags "-X 'main.Version=${version}' -X 'main.CommitSha=${commit_sha}' -s -w"  -trimpath -o ${binary_dir}/${GOOS}_${GOARCH} ${workdir}/...
done

for arch in ${archs}
do
  split=(${arch//\// })
  GOOS=${split[0]}
  GOARCH=${split[1]}
  binaries=`ls ${binary_dir}/${GOOS}_${GOARCH}`

  for binary in ${binaries}
  do
    output_name=${binary_dir}/${GOOS}_${GOARCH}/${binary}_${version}_${GOOS}_${GOARCH}
    if [ ${GOOS} = "windows" ]
    then
      output_name+=.exe
  fi
    mv  ${binary_dir}/${GOOS}_${GOARCH}/${binary} ${output_name}
 done
done

