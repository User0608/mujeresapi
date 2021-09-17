## ASOCIAR UNA PERSONA CON UN USUARIO, EN MODO COLABORADOR

POST: http://localhost:90/v1/colaborador

pasamos la ID del la persona y el usuario a asociar.
```json
    {
        "persona_id":1,
        "usuario_id":1
    }
```
Respuesta:
```json
    {
        "code": "OK",
        "data": {
            "id": 3,
            "persona_id": 2,
            "usuario_id": 4
    }
}   
```

## Listar datos por id
GET: http://localhost:90/v1/colaborador/:id

```json
    {
        "code": "OK",
        "data": {
            "id": 1,
            "persona_id": 1,
            "persona": {
                "id": 1,
                "nombre": "Richard",
                "apellido_materno": "Escobal",
                "apellido_paterno": "Cotrina",
                "telefono": "98763784",
                "dni": "99324354",
                "direccion_id": 22
            },
            "usuario_id": 1,
            "usuario": {
                "usuario_id": 1,
                "username": "kevin002",
                "estado": true
            }
        }
    }
```

## Listar Todos los colaboradores
GET: http://localhost:90/v1/colaborador