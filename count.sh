totalfiles=$(find . |wc -l)
calc=$((totalfiles * 5))
printf "\t\vtotalfiles * 5: %calc\v\n"
