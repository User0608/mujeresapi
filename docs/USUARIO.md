# Mujeres Api

# Iniciar Session
POST:  http://localhost:91/v1/login
```json
    {
        "username":"username",
        "password":"password"
    }
```
Respuesta:
```json
    {
    "usuario": {
        "usuario_id": 1,
        "username": "kevin002",
        "password": "maira002",
        "estado": true,
        "roles": [
            {
                "role_id": 1,
                "nombre": "admin",
                "descripcion": "todos los permisos"
            }
        ]
    },
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6ImtldmluMDAyIiwiUm9sZXMiOlsiYWRtaW4iXX0.fXXKGNxhnrYuK5Pxx3uhqudA0tpaEUmFpWXipv-QJ_GHcoZ0e5PcC7Gm7NbGX-CO4EjiTMa15H3Xy1HbBi-oLP0dLUjQ-x25qPFI71FfbbxQyA5Mxo3oLfTkYmd9JDwN7RPzs6gnaTiek1zFLWz4V16ViNtcAtpxbf7oIByjaS8"
}
```
# Registro de Usuario
POST: http://localhost:91/v1/registrar

```json
    {
        "username":"usuario2020",
        "password":"mipassword",
        "roles":[
            {
                "role_id":3
            }
        ]

    }
```
Response:
```json
    {
        "code": "X000K",
        "message": "El username ya esta en uso"
    }
```
# Registro de datos usuario App
POST: PUT: http://localhost:91/v1/registrar
El campo `appuser_id` es referente al usuario.
```json
    {
        "appuser_id": 1,
        "nombre": "Kevin",
        "apellido_materno": "huaman",
        "apellido_paterno": "saucedo",
        "telefono": "983298327",
        "dni": "88787656",
        "direccion": {
            "provincia": "chepen",
            "distrito": "pacanga",
            "direccion": "Panamericana norte San jose de moro",
            "referencia": "costado de restaurante, don pepino"
        }
    }
```
Response:
```json
    {
    "code": "OK",
    "data": {
        "appuser_id": 2,
        "nombre": "Kevin",
        "apellido_materno": "huaman",
        "apellido_paterno": "saucedo",
        "telefono": "983298327",
        "dni": "88787656",
        "direccion_id": 19,
        "direccion": {
            "direccion_id": 19,
            "provincia": "chepen",
            "distrito": "pacanga",
            "direccion": "Panamericana norte San jose de moro",
            "referencia": "costado de restaurante, don pepino"
            }   
        }
    }
```

# Update de datos usuario App
PUT: http://localhost:91/v1/registrar

```json
    {
        "appuser_id": 1,
        "nombre": "Kevin",
        "apellido_materno": "Huaman",
        "apellido_paterno": "Saucedo",
        "telefono": "983298327",
        "dni": "88787656",
        "direccion_id": 12,
        "direccion": {
            "direccion_id": 12,
            "provincia": "chepen",
            "distrito": "pacanga",
            "direccion": "Panamericana norte San jose de moro",
            "referencia": "costado de restaurante, don pepino"
        }
    }
```
Response:
```json
    {
    "code": "OK",
    "data": {
        "appuser_id": 1,
        "nombre": "Kevin",
        "apellido_materno": "huaman",
        "apellido_paterno": "saucedo",
        "telefono": "983298327",
        "dni": "88787656",
        "direccion_id": 12,
        "direccion": {
            "direccion_id": 12,
            "provincia": "chepen",
            "distrito": "pacanga",
            "direccion": "Panamericana norte San jose de moro",
            "referencia": "costado de restaurante, don pepino"
        }
    }
}
```
