## Registro de reporte... por parte de la institucion

## Crear
POST: http://localhost:90/v1/reporte
```json
   {
        "notificacion_id":1,
        "descripcion":"New descripcion"
    } 
```
```json
    {
        "code": "OK",
        "data": {
            "id": 3,
            "descripcion": "new descripcion",
            "notificacion_id": 1,
            "created_at": "2021-09-24T11:58:26.5296906-05:00",
            "create_at": "2021-09-24T11:58:26.5296906-05:00"
        }
    }
```
## Update
PUT: http://localhost:90/v1/reporte/<reporte_id>
POST: http://localhost:90/v1/reporte
```json
   {
        "notificacion_id":1,
        "descripcion":"update descripcion"
    } 
```
Response
```json
    {
        "code": "OK",
        "data": {
            "id": 3,
            "descripcion": "update descripcion",
            "notificacion_id": 1,
            "created_at": "0001-01-01T00:00:00Z",
            "create_at": "2021-09-24T11:59:55.1008489-05:00"
        }
    }
```
## Delete
DELETE: http://localhost:90/v1/reporte/<reporete_id>
Response:
```json
    {
        "code": "OK",
        "message": "Registro eliminado"
    }
```
### Get All
GET: http://localhost:90/v1/reporte

### Get By ID
GET: http://localhost:90/v1/reporte/<reporete_id>

### Get By Notificacion ID
GET: http://localhost:90/v1/reporte/notificacion/1
