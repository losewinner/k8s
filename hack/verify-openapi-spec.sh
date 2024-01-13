#!/usr/bin/env bash

# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script checks whether updating of OpenAPI specification is needed or not.
# We should run `hack/update-openapi-spec.sh` if OpenAPI specification is out of
# date.
# Usage: `hack/verify-openapi-spec.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env
kube::etcd::install

GOPROXY=off go install ./cmd/kube-apiserver

SPECROOT="${KUBE_ROOT}/api/openapi-spec"
SPECV3PATH="${SPECROOT}/v3"
DISCOVERYROOT="${KUBE_ROOT}/api/discovery"
TMP_SPECROOT="${KUBE_ROOT}/_tmp/openapi-spec"
TMP_DISCOVERYROOT="${KUBE_ROOT}/_tmp/discovery"
_tmp="${KUBE_ROOT}/_tmp"

mkdir -p "${_tmp}"
cp -a "${SPECROOT}" "${TMP_SPECROOT}"
cp -a "${DISCOVERYROOT}" "${TMP_DISCOVERYROOT}"
trap 'cp -a ${TMP_SPECROOT} ${SPECROOT}/..; cp -a ${TMP_DISCOVERYROOT} ${DISCOVERYROOT}/..; rm -rf ${_tmp}' EXIT SIGINT
rm -r "${SPECROOT:?}"/*
rm -r "${DISCOVERYROOT:?}"/*
mkdir -p "${SPECV3PATH}"
cp "${TMP_SPECROOT}/README.md" "${SPECROOT}/README.md"

"${KUBE_ROOT}/hack/update-openapi-spec.sh"
echo "diffing ${SPECROOT} against freshly generated openapi spec"
ret=0
diff -Naupr -I 'Auto generated by' "${SPECROOT}" "${TMP_SPECROOT}" || ret=$?
if [[ $ret -eq 0 ]]
then
  echo "${SPECROOT} up to date."
else
  echo "${SPECROOT} is out of date. Please run hack/update-openapi-spec.sh" >&2
  exit 1
fi

echo "diffing ${DISCOVERYROOT} against freshly generated discovery"
ret=0
diff -Naupr "${DISCOVERYROOT}" "${TMP_DISCOVERYROOT}" || ret=$?
if [[ $ret -eq 0 ]]
then
  echo "${DISCOVERYROOT} up to date."
else
  echo "${DISCOVERYROOT} is out of date. Please run hack/update-openapi-spec.sh" >&2
  exit 1
fi

# ex: ts=2 sw=2 et filetype=sh
