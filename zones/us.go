package zones

import "embed"

//go:embed us/nws_zones.json
var USAWeatherServiceZones []byte

//go:embed us/zip_code_database.csv
var USAPostalCodeDatabase []byte

//go:embed all:us/zip-code-geojson
var USAPostalCodeGeoJSON embed.FS
