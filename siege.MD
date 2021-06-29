# Siege


````
siege 'https://jsonplaceholder.typicode.com/posts POST { "title": "foo", "body":"bar", "userId":1 }' -c25 -r10 -d1 -T 'application/json' --log=./data.log
````

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

