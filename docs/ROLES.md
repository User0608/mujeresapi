## Recuperar los reles, asignables a los usuarios
GET: http://localhost:90/v1/roles
respuesta json:

```json
    {
        "code": "OK",
        "data": [
            {
                "role_id": 3,
                "nombre": "user-t2",
                "descripcion": "usuario control, web"
            },
            {
                "role_id": 4,
                "nombre": "user-ii",
                "descripcion": "usuario institucion"
            }
        ]
    }
```