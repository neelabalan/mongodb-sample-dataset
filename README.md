# [MongoDB Sample Dataset](https://docs.atlas.mongodb.com/sample-data/available-sample-datasets/)

> this is a duplicate [repo](https://github.com/mcampo2/mongodb-sample-databases).

> the original repo has BSON dump. I've removed it and I added a bash [script](https://github.com/neelabalan/mongodb-sample-dataset/blob/main/script.sh) to import the JSON to respective db 

Atlas provides sample data you can load into your Atlas clusters. You can use this data to quickly get started experimenting with data in MongoDB and using tools such as the Atlas Perform CRUD Operations in Atlas and MongoDB Charts.

MongoDB does not provide any sample databases on their website, However, they do provide sample databases for their cloud service Atlas.  These databases have been dumped from a MongoDB Atlas cluster using the MongoDB Compass GUI.  There are 7 databases, with each database collection (table) stored in a seperate JSON file.  These files will accelerate learning of MongoDB's features by allowing new developers to quickly experiment with prepopulated datasets.


## Sample Datasets

| Dataset Name                                                                                | Description                                                       | Collections                                                                |
| ------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------------------------------------------------- |
| [Sample AirBnB Listings Dataset](https://docs.atlas.mongodb.com/sample-data/sample-airbnb/) | Contains details on AirBnB listings.                              | listingsAndReviews                                                         |
| [Sample Analytics Dataset](https://docs.atlas.mongodb.com/sample-data/sample-analytics/)    | Contains training data for a mock financial services application. | accounts, customers, transactions                                          |
| [Sample Geospatial Dataset](https://docs.atlas.mongodb.com/sample-data/sample-geospatial/)  | Contains shipwreck data.                                          | shipwrecks                                                                 |
| [Sample Mflix Dataset](https://docs.atlas.mongodb.com/sample-data/sample-mflix/)            | Contains movie data.                                              | comments, movies, theaters, users                                          |
| [Sample Supply Store Dataset](https://docs.atlas.mongodb.com/sample-data/sample-supplies/)  | Contains data from a mock office supply store.                    | sales                                                                      |
| [Sample Training Dataset](https://docs.atlas.mongodb.com/sample-data/sample-training/)      | Contains MongoDB training services dataset.                       | companies, grades, inspection, posts, routes, stories, trips, tweets, zips |
| [Sample Weather Dataset](https://docs.atlas.mongodb.com/sample-data/sample-weather/)        | Contains detailed weather reports.                                | data                                                                       |

## Running in docker

```bash
docker pull mvertes/alpine-mongo

docker run -d --name mongo -p 2717:27017 -v ~/mongodb:/data/db  mvertes/alpine-mongo

# args
#   hostname   
#   port
./script localhost 2717

# start mongo shell
docker exec -ti mongo mongo
```

