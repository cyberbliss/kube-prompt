#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $(dirname $0)); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ./bin/option-gen ./_tools/option-gen/main.go

kubectlsubcmds=(
"annotate"
"api-resources"
"apply"
"attach"
"auth can-i"
"auth reconcile"
"autoscale"
"certificate approve"
"certificate deny"
"cluster-info dump"
"cluster-info"
"config get-contexts"
"config set-cluster"
"config set-credentials"
"config set"
"config view"
"convert"
"cordon"
"cp"
"create"
"delete"
"describe"
"drain"
"edit"
"exec"
"explain"
"expose"
"get"
"label"
"logs"
"patch"
"port-forward"
"proxy"
"replace"
"rolling-update"
"rollout history"
"rollout pause"
"rollout resume"
"rollout status"
"rollout undo"
"run"
"scale"
"set env"
"set image"
"set resources"
"set selector"
"set serviceaccount"
"set subject"
"taint"
"top node"
"top pod"
"uncordon"
"version"
)

for cmd in "${kubectlsubcmds[@]}"; do
  echo "Generating code for ${cmd}"
  camelized=`echo ${cmd} | gsed -r 's/[- ](.)/\U\1\E/g'`
  snaked=`echo ${cmd} | gsed -r 's/[- ]/_/g'`
  
  kubectl ${cmd} --help | ./bin/option-gen -o ${KUBE_DIR}/option_${snaked}.gen.go -var ${camelized}Options
  goimports -w ${KUBE_DIR}/option_${snaked}.gen.go
done
