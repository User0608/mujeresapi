# CREAR ALERTA
POST:  http://localhost:91/v1/alerta
Los datos de usuario se toman de la session, solo los usuarios 
en la application movil pueden emitir alertas. 
usuarios con permisos t1
```json
    {
        "latitude":"-7.1816743",
        "longitude":"-74.4322543"
    }
```
Respuesta:
```json
    {
    "code": "OK",
    "data": {
        "alerta_id": 7,
        "latitude": "-7.1816743",
        "longitude": "-74.4322543",
        "usuario_id": 3,
        "estado": "enviado",
        "multimedias": [],
        "created": "2021-08-25T14:15:04.8082571-05:00",
        "updated": "2021-08-25T14:15:04.8082571-05:00",
        "deleted": null
    }
}
```

#Listar alertas
GET: http://localhost:91/v1/alerta
Un usuario con permisos de admin o t2, serán capases de ver la lista total de alertas emitidas por todo los usuarios, mientras que un usuario tipo t1 solo residirá la lista de alertas emitidas por si mismo.
Si la alerta, tupiese archivos multimedia asociados, estos también serán entregados.
Response
```json
{
    "code": "OK",
    "data": [
        {
            "alerta_id": 1,
            "latitude": "-7.1816743",
            "longitude": "-79.4322543",
            "usuario_id": 2,
            "estado": "enviado",
            "multimedias": [
                {
                    "multimedia_id": 1,
                    "url": "https://pkg.go.dev/testing",
                    "tipo": "video",
                    "alerta_id": 1
                }
            ],
            "created": "2021-08-25T10:04:54.54052Z",
            "updated": "2021-08-25T10:04:54.54052Z",
            "deleted": null
        }
    ]
}
```

# Alerta por ID
GET: http://localhost:91/v1/alerta/:alerta_id
Al igual que antes, los administradores y usuarios t2 pueden acceder a cualquier alerta, por el contrario los usuarios t1 están limitados a sus propias alertas, 
si dotara de archivos multimedia,estos serian mostrados

```json
{
    "code": "OK",
    "data": {
        "alerta_id": 2,
        "latitude": "-7.1816743",
        "longitude": "-74.4322543",
        "usuario_id": 2,
        "estado": "enviado",
        "multimedias": null,
        "created": "2021-08-25T10:17:08.14379Z",
        "updated": "2021-08-25T10:17:08.14379Z",
        "deleted": null
    }
}
```