# Elasticsearch

## Creating or deleting indices

```

DELETE /pages

PUT /products
{
    "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 2
    }
}
```

## Indexing documents

Una vez hemos creado nuestro índice, podemos indexar documentos al mismo. Para ello deberemos utilizar el **verbo HTTP** `POST` indicando el índice así como la palabra clave `_doc`, a continuación definimos el documento que queremos insertar. Un ejemplo:

```json
POST /<index_name>/_doc
{
    "name": "myName",
    ...
}
```

Obtendremos una respuesta del servidor que contendrá
* Identificador
* Resultado


```json
{
  "_index": "products",
  "_id": "0B9ZlIMBhdKKR9gSA6Ft",
  "_version": 1,
  "result": "created",
  "_shards": {
    "total": 2,
    "successful": 1,
    "failed": 0
  },
  "_seq_no": 0,
  "_primary_term": 1
}
```

## Documents underhood

Elasticsearch ha diseño el concepto de documentos inmutables. Es decir, los documentos no pueden ser modificados. Por tanto cada vez que hacemos una actualización de un documento lo que estamos haciendo realmente es reemplazarlo.

* The current document is retrieved
* The field values are changed
* The existing document is *replaced* with modified document.



## Scripted updates

Elasticsearch admite *scripting* dentro de las operaciones CRUD. Esto nos permite añadir lógica a estas operaciones. Veamos algunos ejemplos:

```json
POST /products/_update/100
{
    "script": {
        "source": "ctx._source.in_stock++"
    }
}
```

```json
POST /products/_update/100
{
    "script": {
        "source": "ctx._source.in_stock = params.quantity",
        "params": {
            "quantity": 4
        }
    }
}
```
El siguiente script actualiza el documento únicamente si el stock es diferente de cero. De serlo, no habría modificaciones.

```json
POST /products/_update/100
{
    "script": {
        "source": """
        if (ctx._source.in_stock == 0) {
            ctx.op = 'noop';
        }

        ctx._source.in_stock--;
        """
    }
}
```

## Upserts

Elasticsearch tiene un mecanismo para crear o actualizar documentos en función de si dichos documentos existen o no. Es decir, si el documento no existe lo crea y si el documento no existe lo actualiza.


```json
POST /products/_update/100
{
    "script": {
        "source": "ctx._source.in_stock++"
    },
    "upsert": {
        "name": "Blender",
        "price": 233,
        "in_stock": 5
    }
}
```

## Update by query

```json
POST /products/_update_by_query
{
    "script": {
        "source": "ctx._source.in_stock--"
    },
    "query": {
        "match_all": {}
    }
}
```

## Batch processing

Introduction to the builk api

* you learned how to index, update and delete documents
* let's see how we can perform these actions on many documents with a single query.
    * That's done 

