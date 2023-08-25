#! /bin/bash

cat all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
