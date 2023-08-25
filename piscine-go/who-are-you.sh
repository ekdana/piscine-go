#! /bin/bash
subject_id=${1:-70}
#curl https://01.alem.school/assets/superhero/all.json | jq '.[] | select(.id == 70)'
name=$(curl -s https://01.alem.school/assets/superhero/all.json | jq --argjson sid "$subject_id" '.[] | select(.id == $sid) | .name')
echo "$name"