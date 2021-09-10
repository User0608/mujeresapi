## Consultar todo los registros de persona
para realizar esta operacion es necesarios 
contar con permisos de administrador de lo contrario no es posible,
pro tanto todas las operaciones relacionadas con persona, tiene que 
tener en la cabecera de la peticion un TOKEN, con rola admin.

GET:http://localhost:90/v1/persona
ejemplo de respuesta, en data, retorna un array con todos los registros de persona...
```json
    {
    "code": "OK",
    "data": [
        {
            "id": 1,
            "nombre": "Richard",
            "apellido_materno": "Escobal",
            "apellido_paterno": "Cotrina",
            "telefono": "98763784",
            "dni": "99324354",
            "direccion_id": 22,
            "direccion": {
                "direccion_id": 22,
                "provincia": "Pacasmayo",
                "distrito": "Guadalupe",
                "direccion": "Calle Junin 235",
                "referencia": "Frente a veterinaria salud"
            }
        }
    ]
}
```
## Consultar por un registro concreto
para realizar esta operacion al igual que la anterior, es necesario contar 
con un rola admin, o por otro lado, debera ser el propietario de la informacion
solicitada.
GET:http://localhost:90/v1/persona/:persona_id

```json
    // el ejemplo es el mismo que el anterior, a diferencia que ahora no 
    // se retornara un array, sino un objeto concreto de persona. 
```

## Create 
POST: 
El objeto que enviado tiene que tener la estructura siguiente,
el unico campo que se puede omitir, es direccion, que se puede insertar en
operaciones posteriores. En el caso de los demas campos, son obligatorios, 
dni contiene 8 caracteres y numero de celular, tiene que tener almeno caracteres,
todos ellos numeros.

Por otro lado, la respuesta que se devuelve, contiene los datos completos, despues de la operacion.

* Example
data send:
```json
    {  
        "nombre": "Richard",
        "apellido_materno": "Escobal",
        "apellido_paterno": "Cotrina",
        "telefono": "98763784",
        "dni": "99324354",
        "direccion": {
            "provincia": "Pacasmayo",
            "distrito": "Guadalupe",
            "direccion": "Calle Junin 235",
            "referencia": "Frente a veterinaria salud"
        }    
    }
```

server response:
```json
    {
    "code": "OK",
    "data": {
        "id": 2,
        "nombre": "Angel",
        "apellido_materno": "Burga",
        "apellido_paterno": "Coronel",
        "telefonoo": "878786762",
        "dni": "98762534",
        "direccion_id": 23,
        "direccion": {
            "direccion_id": 23,
            "provincia": "Chepen",
            "distrito": "Chepen",
            "direccion": "Calle cajamarca 235 jr. 4",
            "referencia": "Paradero de motos el apurado."
        }
    }
}
```

## Update
Para realizar actualizaciones, se puede hacer exceptuando el campo direccion
PUT: http://localhost:90/v1/persona
```json
    {
        "id": 2,
        "nombre": "Angel",
        "apellido_materno": "Escobal",
        "apellido_paterno": "Cotrina",
        "telefonoo": "878786762",
        "dni": "98762534"
    }
```
con lo cual la respuesta sera:

```json
    {
        
        "code": "OK",
        "data": {
            "id": 2,
            "nombre": "Angel",
            "apellido_materno": "Escobal",
            "apellido_paterno": "Cotrina",
            "telefono": "878786762",
            "dni": "98762534",
            "direccion_id": 0,
            "direccion": null
        }
    }  
```

Cabe resaltar tambien la posibilidad de tambien enviar la informacion de direccion,
pero esta debe estar completa de lo contrario seran rechazados, se recomienda hacer primero una consulta de los datos, por id, y editar los campos necesarios y enviarlos. 