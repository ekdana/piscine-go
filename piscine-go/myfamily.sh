#!/bin/bash


HERO_ID=${HERO_ID:-1}  

API_ENDPOINT="https://01.alem.school/assets/superhero/all.json"

FAMILY=$(curl -s "$API_ENDPOINT" | jq --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives|  sub("; "; "; ")')
FAMILY=${FAMILY//\"/}
echo "$FAMILY"