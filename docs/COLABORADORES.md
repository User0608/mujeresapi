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
        "message": "Operacion realizada con exito!"        
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