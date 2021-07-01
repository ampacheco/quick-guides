# Siege
## Install Siege

````
sudo apt-get update
sudo apt-get install siege -y
````


## Basic Execution
````
siege 'https://jsonplaceholder.typicode.com/posts POST { "title": "foo", "body":"bar", "userId":1 }' -c25 -r10 -d1 -T 'application/json' --log=./data.log
````
|Switch| | 
|:----:| ------|
|-c25|25 usuarios concurrentes|
|-r10|10 ejecuciones por cada usuario (25x10 = 250 calls)|
|-d1|delay between calls|
|--log=./data.log|execution log in a file in the local dir| 

## Advanced Executions using a file as service input and a json file as service input data
````
siege -f siege.input -c25 -r10 -d1 -T 'application/json' --log=./data.log
````

````
cat << EOF >> siege.input
https://jsonplaceholder.typicode.com/posts POST < data.json 
EOF
````
````
cat << EOF >> data.json
{
    "title": "foo",
    "body": "bar",
    "userId": 1
}
EOF
````

