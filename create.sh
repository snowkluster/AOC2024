#!/bin/sh

echo "Creating folders all the Day's"
for i in $(seq 1 25);
do
	mkdir "Day ${i}"
done

