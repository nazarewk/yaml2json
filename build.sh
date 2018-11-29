#!/usr/bin/env bash
set -Eeuo pipefail
if [[ -n "${DEBUG:-}" ]] ; then export DEBUG; set -x; fi

build_file () {
  local out cmd

  >&2 echo "Building ${GOOS}/${GOARCH}..."
  for cmd in "${@}"; do
    out="dist/${cmd}-${GOOS}-${GOARCH}"
    [[ "${GOOS}" == "windows" ]] && out="${out}.exe"
    if ! go build -o "${out}" "cmd/${cmd}/main.go" ; then
      >&2 echo "    failed"
      return 1
    fi
    chmod +x "${out}"
  done
  >&2 echo "    success"
}

main () {
  mkdir -p dist

  for os_arch in $(go tool dist list); do
    export GOOS=${os_arch%/*}
    export GOARCH=${os_arch#*/}
    build_file yaml2json json2yaml || true
  done

  pushd dist
  sha256sum json2yaml* yaml2json* > sha256sums.txt
  popd
}

main