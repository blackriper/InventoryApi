# Mostrar Productos 

Muestra todos los productos disponibles en la base de datos

**URL** : `/inventory/products/`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : NONE


## Success Response

**Condition** : Si no hay ningun error

**Code** : `200 OK`

**Content example**

```json
{
  "body": [
    {
      "sku": "067eb46e-00ec-470c-80b6-6aa5e06d2d91",
      "name": "Gigabyte laptop",
      "cant": 50
    },
    {
      "sku": "559c6344-3f09-4e8f-b96c-5a4d0039c5cf",
      "name": "Teclados Logictech",
      "cant": 160
    },
    {
      "sku": "8a82eec6-952f-4257-b60e-117c114a68b9",
      "name": "Gigabyte SSD 240GB",
      "cant": 50
    },
    {
      "sku": "df9689e3-93a2-4a55-b314-bc04b97469bd",
      "name": "Ryzen 5 3600",
      "cant": 13
    }
  ]
}
```

