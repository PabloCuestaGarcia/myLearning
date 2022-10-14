# cluster health
GET /_cluster/health

# List of nodes
GET /_cat/nodes?v

# List of indices
GET /_cat/indices?v

# Shards
GET /_cat/shards?v

# Creating a new index
PUT /products

# Deleting an index
DELETE /products

## Creating a new index
PUT /products
{
  "settings": {
    "number_of_shards": 2,
    "number_of_replicas": 2
  }
}

# Create a new document
POST /products/_doc 
{
  "name": "Toaster",
  "price": 43,
  "in_stock": 10
}

PUT /products/_doc/100
{
  "name": "Toaster",
  "price": 43,
  "in_stock": 10
}

POST /products/_doc 
{
  "name": "Coffee",
  "price": 43,
  "in_stock": 1
}

## Get documents
GET /products/_doc/4am_WYIBp7Eiintmf8v2
GET /products/_doc/100

# Update a document
POST /products/_update/100
{
  "doc": {
    "in_stock": 12
  }
}

POST /products/_update/100
{
  "doc": {
    "tags": ["electronics"]
  }
}

## Scripted updates

## Decresing the value of in_stock -1
POST /products/_update/100
{
  "script": {
    "source": "ctx._source.in_stock--"
  }
}

POST /products/_update/100
{
  "script": {
    "source": "ctx._source.in_stock = 112"
  }
}
POST /products/_update/100
{
  "script": {
    "source": "ctx._source.in_stock -= params.quantity",
    "params": {
      "quantity": 4
    }
  }
}
GET /products/_doc/100





















