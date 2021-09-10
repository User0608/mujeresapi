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
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9"
}
```
# Registro de Usuario, este metodo de registro, solo aplica para usuarios del tipo t1, t2 y t3 son gestionados por el admin.
POST: http://localhost:91/v1/registrar

```json
       {
        "username":"torress05",
        "password":"fortaleza12"
    }
```
Response:
```json
    {
    "code": "OK",
    "data": {
        "usuario_id": 6,
        "username": "torress05",
        "password": "fortaleza12",
        "estado": true,
        "roles": [{"role_id": 2,"nombre": "name"}
        ]
    }
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
