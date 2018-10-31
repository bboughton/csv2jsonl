# csv2jsonl

This tool converts CSV files to Jsonl (new line separated json
objects) output.

Note: that all values are interpreted as strings when reading the CSV
file so all output will also be strings.

## Install

    go get github.com/bboughton/csv2jsonl

## Usage

    csv2jsonl data.csv > data.jsonl

