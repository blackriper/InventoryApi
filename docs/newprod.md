# Registrar nuevos productos

Registrar nuevo producto a la base de datos 

**URL** : `/inventory/newproduct/`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Se necesita el nombre del producto y la cantidad.

```json
{
    "name": "[unicode 64 chars max]",
    "cant": integer
}
```

**Data example** Ejemplo de los campos que son requeridos.

```json
{
    "name":"Monitor Asus",
    "cant":123
}
```

## Success Response

**Condition** : Si el producto no existe en la base datos y se enviaron los campos correctamente.

**Code** : `201 CREATED`

**Content example**

```json
{
  "body": "product 58bfd8c2-652a-4e91-b93f-86ab609f6a25 created sucesfully"
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