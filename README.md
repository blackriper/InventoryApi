# Inventory Api 

Esta api  esta hecha en golang usando [Gin framework](https://gin-gonic.com/) la api consiste en una base de datos de productos la cual esta alojada en [Planet Scale](https://planetscale.com/) un servicio que ejecuta una base de datos MySQL.

## Objectivo del Proyecto
Aprender la configuracion basica de una api en golang siguiendo el proyecto MVC y utlizando el orm [Gorm](https://planetscale.com/) para conectar y hacer las operaciones CRUD basicas de la base de datos.


## Estructura de carpetas 
En el proyecto puedes encontrar las siguientes carpetas 

1. Config: paquete donde se cargan las validaciones de los datos enviados por gin y donde se cargan las configuraciones de variables de entorno.

2. Docs: archivos para la documentacion de markdown

3. Controller: paquete donde se guardan las funciones de las rutas de gin 

4. Modeles: paquete donde se encuentra la configuracion del modelo para gorm la conexion de la base de datos y las operaciones basicas del crud.

5. Routes: paquete donde se encuentra la configuracion del  router de gin

## Endpoints Inventory Api

rutas de la aplicación

* [New Product](docs/newprod.md) : `POST /inventory/newproduct/`
* [Products](docs/products.md) :   `GET /inventory/products/`
* [Update Product](docs/update.md) : `PUT /inventory/updateproduct/:sku/`
* [Rest Cant](docs/rest.md) : `PUT /Inventory/restcant/:sku/`
* [Delete Product](docs/delete.md) : `DELETE /Inventory/deleteproduct/:sku/`

## Ejecutar proyecto 
 ```bash
    go run main.go
 ```