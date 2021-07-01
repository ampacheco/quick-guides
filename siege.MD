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

