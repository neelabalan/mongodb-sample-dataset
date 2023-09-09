#!/bin/bash

function import_collection {
    local collections=("$@")
    # cd "$dir"
    # echo "$LOGNAME"
    > /workspaces/mongodb-sample-dataset/work/mongo.log
    docker exec -it mongodb mkdir -p vinay
    for ((i = 0; i < ${#collections[@]}; i++)); do
        # echo "${json_files[$i]}"
        collection="${collections[$i]}"
        json_file="${json_files[$i]}.json"
        # docker cp "$dir/$json_file" mongodb:"/$LOGNAME/"
        # docker exec -it mongodb mongoimport --host localhost --port 27017 --db vinay --collection "$collection" --file "/$LOGNAME/$json_file" --drop

        docker cp "$dir/$json_file" mongodb:"/vinay/" >> /workspaces/mongodb-sample-dataset/work/mongo.log 2>&1
        docker exec -it mongodb mongoimport --host localhost --port 27017 --db vinay --collection "$collection" --file "/vinay/$json_file" --drop >> /workspaces/mongodb-sample-dataset/work/mongo.log 2>&1
        echo $import_log >> /workspaces/mongodb-sample-dataset/work/mongo.log
            if [ $? -eq 0 ]; then
                echo "Import successful for collection '$collection'."
                # docker exec -it mongodb rm -rf "/$LOGNAME/$json_file"
                # docker exec -it mongodb rm -rf "/vinay/$json_file"
            else
                echo "Error importing data into collection '$collection'."
            fi
        sleep 1
    done
    docker exec -it mongodb rm -rf /vinay/
}



echo "Hey Hi $LOGNAME, do you want to import a MongoDB collection?"
echo "sample_analyticsmongodata: Enter 1"
echo "sample_geospatialmongodata: Enter 2"
echo "sample_mflixmongodata: Enter 3"
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
                comments
                movies                
                sessions
            )
            json_files=(
                comments
                movies                
                sessions
            )
            dir="/workspaces/mongodb-sample-dataset/sample_mflix"
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
go build /workspaces/mongodb-sample-dataset/work/main.go
./main

# docker cp "/workspaces/mongodb-sample-dataset/sample_geospatial/shipwrecks.json" mongodb:/vinay/