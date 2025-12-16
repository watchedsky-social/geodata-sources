package zones

import "embed"

//go:embed us/nws_zones.json
var USAWeatherServiceZones []byte

//go:embed us/zip_code_database.csv
var USAPostalCodeDatabase []byte

//go:embed us/zip-code-geojson/*.json
var USAPostalCodeGeoJSON embed.FS
