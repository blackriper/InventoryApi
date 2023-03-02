# Update Producto

Actualizar los datos de un producto

**URL** : `/inventory/updateproduct/:sku/`

**Method** : `PUT`

**Auth required** : NO

**Permissions required** : NO

**Data constraints**

```json
{
    "name": "[unicode 64 chars max]",
    "cant": Integer
}
```

**Data example** Csmpos requeridos en el request

```json
{
    "name":"Mac Mini",
    "cant":12
}
```

## Success Responses

**Condition** : Actualizar alguno de los campos ya sea name o cant

**Code** : `200 OK`

**Content example** :no hay que olvidar agregar el sku del producto `/inventory/updateproduct/33de1301-09b3-47b5-a037-33dee8bbd08e/`.

```json
{
  "body": "product 067eb46e-00ec-470c-80b6-6aa5e06d2d91 updated sucesfully"
}
```

## Error Responses

**Condition** : si falta algun campo

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
  "error": "Key: 'InputProduct.Cant' Error:Field validation for 'Cant' failed on the 'required' tag"

}



