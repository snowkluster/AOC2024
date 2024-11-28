#!/bin/sh

date_total=$(date +%d/%Y)

day=$(echo "$date_total" | cut -d '/' -f1)
year=$(echo "$date_total" | cut -d '/' -f2)

if [ $day -gt 25 ];
then
	echo "it's so over, maybe next year"
	exit 0
fi

today_url="https://adventofcode.com/${year}/day/${day}/input"
wget -q $today_url -P "Day ${day}"
