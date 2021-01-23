for coll in *; do
    if [ -d "${coll}" ] ; then
        echo "$coll"
        for file in $coll/*; do
            mongoimport --drop --host $1 --port $2 --db "$coll" --collection "$(basename $file .json)" --file $file
            #echo "$(basename $file .json)"
            #echo "$file"
        done
    fi
done

