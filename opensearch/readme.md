# Elasticsearch

## Mapping & Analysis

### Introduction to analysis

* Sometimes referred to as *text analysis*
* Applicable to text field/values
* Text values are analyzed when indexing documents
* The result is stored in data structures that are efficient for searching etc
* The `_source` object is **not** used when searching for documents.
    * It contains the exact values specified when indexing a document.



Cuando un texto o cadena es indexada, un proceso denominado *analizador* procesa el texto. Este analizador consiste en tres bloques: filtro de caracteres, un token y filtros de tokens. El resultado es guardado en un estructura de datos que favorece la busqueda.

#### Character filters
Adds, removes or changes characters
Analyzers contain zero or more character files

#### Tokenizers
Characters my be stripped as part of the tokenization

#### Token filters

Receive the output of the tokenizer as input
A token filter can add, remove or modifuy tokens



### Undestanding inverted indices

Mapping between terms and which documents contain them
Outside the context of analyzers, we use the terminology "terms"
Terms are sorted alphabetically
Inverted indices contain more than just terms and documents IDs