#!/bin/bash
YEAR_DIRS=$(find . -maxdepth 1 -type d -not -path '*\/.*' -not -path . -exec basename {} \; | grep -P '\d{4}' | sort -r)
LANG_DIRS=$(find . -maxdepth 1 -type d -not -path '*\/.*' -not -path . -exec basename {} \; | grep -vP '\d{4}')
README="README.md"
cat readme.in > $README

for year in $YEAR_DIRS
do
    echo "## $year" >> $README
    header="|Day|Data|"
    delim="|:-:|:-:|"
    for lang in $LANG_DIRS
    do
        header="$header${lang^}|"
        delim="$delim:-:|"
    done
    echo "$header" >> $README
    echo "$delim" >> $README
    
    for i in {1..25}
    do
        day="|$i|"
        data=$(git ls-files | grep -E "^${year}/day${i}/")
        if [ -z "$data" ]; then
            day="${day}|"
        else
            day="${day}[data]($year/day$i/)|"
        fi
        for lang in $LANG_DIRS
        do
            src=$(git ls-files | grep -E "$lang.*$year.*day${i}(\..*)?$")
            if [ -z "$src" ]; then
                day="${day}|"
            else
                day="${day}[src]($src)|"
            fi
        done
        echo "$day" >> $README
    done
done

exit 0