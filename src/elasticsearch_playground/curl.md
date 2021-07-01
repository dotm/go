The `?pretty` makes the response returned to be prettified first

## Ping:
```
curl -X GET "localhost:9200/_cat/health?v&pretty"
```

## Add document (id = 1)
successful insert will return status 201
```
curl -X PUT "localhost:9200/index_name/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "field": "value"
}
'
```

## Add document without id
```
curl -X POST "localhost:9200/index_name/_doc/?pretty" -H 'Content-Type: application/json' -d'
{
    "field": "value"
}
'
```

## Get document by id (id = 1)
```
curl -X GET "localhost:9200/index_name/_doc/1?pretty"
```

## Get document by searching
```
curl -X GET "localhost:9200/index_name/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}
'
```

## Data aggregation
```
curl -X GET "localhost:9200/index_name/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      }
    }
  }
}
'
```
size=0 means that the response should only contains the aggregation results

## Create an index with an explicit mapping
```
curl -X PUT "localhost:9200/index_name?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "dynamic": "strict",
    "properties": {
      "age":    { "type": "integer" },  
      "email":  { "type": "keyword"  }, 
      "name":   { "type": "text"  }     
    }
  }
}
'
```

## Add a field to an existing mapping
```
curl -X PUT "localhost:9200/index_name/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "employee-id": {
      "type": "keyword",
      "index": false
    }
  }
}
'
```

## Reindex (move data from one index to another)
https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html

## View the mapping
of an index:
```
curl -X GET "localhost:9200/index_name/_mapping?pretty"
```

of a field:
```
curl -X GET "localhost:9200/index_name/_mapping/field/employee-id?pretty"
```

## Delete index
```
curl -X DELETE "localhost:9200/index_name?pretty"
```