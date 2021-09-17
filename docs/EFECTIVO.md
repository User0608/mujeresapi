## Crear un Efectivo
Permisos: Institucio
```json
    {
        "institucion_id":1,
        "nombre":"Jorge",
        "apellido_paterno":"Linares",
        "apellido_materno":"Castro",
        "telefono":"987367878",
        "dni":"98767897",
        "direccion":"San Jose De Moro",
        "provincia":"Chepen",
        "distrito":"Pacanga",
        "referencia":"Panamericanan norte"
    }
```
response
```json
{
    "code": "OK",
    "data": {
        "efectivo_id": 3,
        "institucion_id": 1,
        "persona_id": 6,
        "persona": {
            "id": 6,
            "nombre": "Jorge",
            "apellido_materno": "Castro",
            "apellido_paterno": "Linares",
            "telefono": "987367878",
            "dni": "98767897",
            "direccion_id": 30,
            "direccion": {
                "direccion_id": 30,
                "provincia": "Chepen",
                "distrito": "Pacanga",
                "direccion": "San Jose De Moro",
                "referencia": "Panamericanan norte"
            }
        }
    }
}
```
# Lista todo los efectivos
Permisos: Admin, Usuario control
GET: http://localhost:90/v1/efectivo

# Lista de efectivos de una institucion
Permisos: Admin, Usuario control
GET http://localhost:90/v1/efectivo/institucion/:insitucion_id

# Lista de efectivos de la institucion con session actual 
Permisos: Institucion User
GET : http://localhost:90/v1/efectivo/institucion