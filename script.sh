#!/bin/bash

function import_collection {
    local collections=("$@")
    # cd "$dir"
    echo "$LOGNAME"
    for ((i = 0; i < ${#collections[@]}; i++)); do
        # echo "${json_files[$i]}"
        collection="${collections[$i]}"
        json_file="${json_files[$i]}.json"
        docker cp "$dir/$json_file" mongodb:"/$LOGNAME/"
        docker exec -it mongodb mongoimport --host localhost --port 27017 --db vinay --collection "$collection" --file "/$LOGNAME/$json_file" --drop
            if [ $? -eq 0 ]; then
                echo "Import successful for collection '$collection'."
                docker exec -it mongodb rm -rf "/$LOGNAME/$json_file"
            else
                echo "Error importing data into collection '$collection'."
            fi
    done
}



echo "Hey Hi $LOGNAME, do you want to import a MongoDB collection?"
echo "ravemongodata: Enter 1"
echo "arcusmongodata: Enter 2"
echo "reszomongodata: Enter 3"
echo "exit: Enter 4"
while true; do
    read -rp "Enter your choice: " action
    case "$action" in
        1)
            # List of collection as well JSON files
            collections=(
                accounts
                customers
                transactions
            )
            json_files=(
                accounts
                customers
                transactions
            )
            dir="/workspaces/mongodb-sample-dataset/sample_analytics"
            import_collection "${collections[@]}"
            break
            ;;
        2)
            # List of collection names and corresponding JSON files
            collections=(
                shipwrecks
            )

            json_files=(
                shipwrecks
            )
            dir="/workspaces/mongodb-sample-dataset/sample_geospatial"
            import_collection "${collections[@]}"
            break
            ;;
        3)
            # List of collection names and corresponding JSON files
            collections=(
                templates.StorageMPFile
                templates.StorageMPFile.automationcode
                automationmap
            )
            json_files=(
                templates
                templates.automationcode
                casemap
            )
            dir="/home/$LOGNAME/go/src/nimble.com/fileStorageRavePlugins/testPlugins/inputs"
            import_collection "${collections[@]}"
            break
            ;;
        4)
        break
        ;;
        *)
            echo "Hi $LOGNAME You entered $action. Wrong input. Please enter 1, 2, or 3 || 4."
            ;;
    esac
done


# docker cp "/workspaces/mongodb-sample-dataset/sample_geospatial/shipwrecks.json" mongodb:/vinay/