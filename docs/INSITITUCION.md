# Operaciones institucion
Todas las operaciones que se detallan a continuacion, requieren tener permisos, de administrador, el propietario de la informacion. el cual tiene permisos de ingresar y solicitar su informacion.
## Create
#POST http://localhost:90/v1/institucion
Estructurá de datos JSON a enviar
```json
{
        "nombre":"Comisaria Pacanga",
        "persona": "Jose Angel Gomez Diaz",
        "telefono": "995831649",
        "email": "josek@outlook.com",
        "tipo": "policia",
        "usuario_id": 53,
        "direccion": {
            "provincia": "Chepen chepen",
            "distrito": "pacanga",
            "direccion": "Calle chavin - san jose de moro",
            "referencia": "Frente a la panamericana norte"
        }    
}
```
La respuesta en casos todo salga de manera esperada sera:
```json
    {
        "code": "OK",
        "data": {
            "institucion_id": 1,
            "nombre":"Comisaria Pacanga",
            "persona": "Jose Angel Gomez Diaz",
            "telefono": "995831649",
            "email": "josek@outlook.com",
            "tipo": "policia",
            "usuario_id": 5,
            "direccion_id": 25,
            "direccion": {
                "direccion_id": 25,
                "provincia": "Chepen",
                "distrito": "pacanga",
                "direccion": "Calle chavin - san jose de moro",
                "referencia": "Frente a la panamericana norte"
            }
        }
    }
```

## Update
Para realizar la operacion de update, tenemos que enviar el objeto identico al la data que se obtiene como respuesta en el proceso de creado, mostrado anterior mente.
PUT: http://localhost:90/v1/institucion

# Consulta de información
Todos los objetos que se recuperen, seran identicos en estructura, al que se obtiene, despues de guardar los datos, a diferencia, de que si es una consulta multiple, esta vendra en array.


## Consulta un objeto, invitando su id asociado
GET: http://localhost:90/v1/institucion/:intitucion_id

## Consulta un Objeto, a traves de el username, asociado a su usuario del sistema
GET: http://localhost:90/v1/institucion/username/:username

## Consulta todo los datos
GET:  http://localhost:90/v1/institucion

## Consultar mi informacion, como usuario institucional
GET: http://localhost:90/v1/institucion/me
Ejemplo de respuesta:
```json
    {
        "code": "OK",
        "data": { // esta puede ser un arregle, segun se el caso.
            "institucion_id": 1,
            "nombre":"Comisaria Pacanga",
            "persona": "Jose Angel Gomez Diaz",
            "telefono": "995831649",
            "email": "josek@outlook.com",
            "tipo": "policia",
            "usuario_id": 5,
            "direccion_id": 25,
            "direccion": {
                "direccion_id": 25,
                "provincia": "Chepen",
                "distrito": "pacanga",
                "direccion": "Calle chavin - san jose de moro",
                "referencia": "Frente a la panamericana norte"
            }
        }
    }
```