## Create new Notificacion
Permisos: ControlUser
POST:http://localhost:90/v1/notificacion
```json
    {
        "titulo":"Mujer es golpeada",
        "alerta_id":1,
        "institucion_id":1,
        "nivel":2,
        "descripcion":"Se reporto un incidente...",
        "colaborador_id":3
    }
```
response
```json
    {
        "code": "OK",
        "data": {
            "notificacion_id": 1,
            "titulo": "Mujer es golpeada",
            "alerta_id": 1,
            "institucion_id": 1,
            "nivel": 2,
            "descripcion": "Se reporto un incidente...",
            "colaborador_id": 3
        }
    }
```
## Update
Permisos: Controla user
PUT: http://localhost:90/v1/notificacion
```json
    {
        "notificacion_id": 1,
        "titulo": "Mujer es golpeada",
        "alerta_id": 1,
        "institucion_id": 1,
        "nivel": 2,
        "descripcion": "Se reporto un incidente...",
        "colaborador_id": 3
    }
```
Response
```json
    {
        "code": "OK",
        "data": {
            "notificacion_id": 1,
            "titulo": "Mujer es golpeada",
            "alerta_id": 1,
            "institucion_id": 1,
            "nivel": 2,
            "descripcion": "Se reporto un incidente...",
            "colaborador_id": 3
        }
    }
```
## Find All
Permisos: Control user o admin
GET: http://localhost:90/v1/notificacion

## Find By ID
Permisos: Admin, Usuario de control, o usuario Institucional, con la limitacion de este ultimo, que solo puede ver las notificaciones que an sido generadas para esa institucion, de no ser el caso, se devolvera null
GET: http://localhost:90/v1/notificacion/:notificacion_id

## Lista todas las notificaciones por institucion
Permisos: Control user, Admin
GET: http://localhost:90/v1/notificacion/institucion/:insitucion_id

## Lista todas las notificaciones por colaborador
Permisos: Control user, Admin
GET: http://localhost:90/v1/notificacion/colaborador/:colaborador_id

## Lista todas la notificaciones de la institucion.
Este endpoint solo se puede consultar por un usuario Institucional,
ya que usarla el id de la instruccion que realiza la peticion para entregar las notificaciones disponibles para esa institucion.
Permisos: Institucion User
GET: http://localhost:90/v1/notificacion/institucion