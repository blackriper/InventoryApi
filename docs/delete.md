# Delete Producto

Borrar producto de la base de datos

**URL** : `/inventory/deleteproduct/:sko/`

**URL Parameters** : `sku=[string]` donde `sku` es el id del producto en la base de datos.

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : NONE

**Data** : `{}`

## Success Response

**Condition** : si el producto existe

**Code** : `200 OK`

**Content** 
```json
 {
  "body": "product ea071c49-6497-4dcb-9d74-7f5bec87c369 Delete sucesfully"
}

```

## Error Responses

**Condition** : Si el producto no existe.

**Code** : `404 NOT FOUND`

**Content** : `{}`
