# Pal Split

The app for pals spliting the bill...

### Usage

#### Retrieve all articles

```
curl http://localhost:10000/articles
```

#### Retrieve a single article

```
curl http://localhost:10000/articles/1
```

#### Create a single article

````
curl -H "Content-Type: application/json"
      --request POST
      --data {
                "id":"3",
                "title":"Posted Article",
                "desc":"New Article",
                "content":"This is it"
             }'
      http://localhost:10000/articles ```
````

#### Delete an article

```
curl --request DELETE http://localhost:10000/article/1
```

#### Update an article

```
curl -H "Content-Type: application/json"
      --request PUT
      --data {
                  "id" : "3",
                  "title" : "Posted Article",
                  "desc"  : "New Article",
                  "content" : "This is updated"
             }
```
