#!/usr/bin/env bash

set -e

readonly inputFile="${1}"
if [[ -z "${inputFile}" ]]; then
  echo "input file required" > /dev/stderr
  exit 1
elif [[ ! -f "${inputFile}" ]]; then
  echo "file not found: ${inputFile}" > /dev/stderr
  exit 2
fi

declare props=false
declare line
declare text
declare count=0
while  read -r line; do
  [[ "${line}" =~ ^\<\!\-\- ]] && continue
  if [[ "${line}" == "<properties>" ]]; then
    props=true
    (( count++ )) || true
  elif [[ "${line}" == "</properties>" ]]; then
    props=false
  elif [[ "${props}" == "true" ]] && [[ -n "${line}" ]]; then
    text="$(echo "${line}" | sed 's/<\//" /' | cut -d ' ' -f1 | tr '<' '"' | sed -e 's/>/":"/' | sed -e 's/$/,/')"
    echo "${count} ${text}"
  fi
done < "${inputFile}"
