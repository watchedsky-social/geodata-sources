#!/usr/bin/env bash

set -euo pipefail

# shellcheck disable=SC2207
ALL_URLS=( $(curl 'https://api.weather.gov/zones?limit=100000' | jq -r '.features[].id') )

# shellcheck disable=SC2068
for url in ${ALL_URLS[@]}; do
  file=$(sed -e 's|https://api.weather.gov/zones/||' -e 's|/|_|g' <<< "$url")
  curl "$url" > "${file}.json"
done

jq --slurp -c '{"type":"FeatureCollection", "features":.}' *.json > all.json
