# Asigna un notificacion a un efectivo
## Create
POST: http://localhost:90/v1/asignacion
```json
    { 
        "efectivo_id":3,
        "notificacion_id":1
    }
```
Response:
```json
    {
    "code": "OK",
    "data": {
        "asignacion_id": 4,
        "created_at": "2021-09-24T03:51:30.6924927-05:00",
        "efectivo_id": 3,
        "notificacion_id": 1
        }
    }
```
## Update
PUT:http://localhost:90/v1/asignacion/<asignacion_id>
```json
    { 
        "efectivo_id":3,
        "notificacion_id":1
    }
```
```json
    {
        "code": "OK",
        "data": {
            "asignacion_id": 4,
            "created_at": "0001-01-01T00:00:00Z",
            "efectivo_id": 3,
            "notificacion_id": 1
        }
    }
```
## Delete
DELETE: http://localhost:90/v1/asignacion/<asignacion_id>
Response:
```json
    {
        "code": "OK",
        "message": "Eliminado correctamente"
    }
```
## Get ByID
GET: http://localhost:90/v1/asignacion/<asignacion_id>

## Get By Efectivo ID
GET: http://localhost:90/v1/asignacion/efectivo/<efectivo_id>

## Get By Notificacion ID
GET: http://localhost:90/v1/asignacion/notificacion/<notificacion_id>