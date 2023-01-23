# Go json token parser

this repository is for me to test out parsing token
in json using go

Example `data`
```json
{
    "string": "some string"
    "number": 123
    "object": {
        "key": "value",
        "key2": {
            "key3": "value3"
        }
    },
    "array": [
        "value1",
        "value2"
    ]
}
```

Example `input`
```json
{
    "name": "${string}",
    "age": "${number}",
    "object": "${object}",
    "array": "${array}",
    "arrayElement": "${array.[0]}",
    "Authorization": "Bearer ${string}",
    "nested": "Bearer ${object.key}",
    "nested2": "Bearer ${object.key2.key3}",
    "all": "${.}",
    "nested_object": "${object.key2}",
    "object_with_message": "Message: ${object.key2}",
}
```

Example `result`
```json
{
    "Authorization": "Bearer some string",
    "age": 123,
    "all": {
        "array": [
            "value1",
            "value2"
        ],
        "number": 123,
        "object": {
            "key": "value",
            "key2": {
                "key3": "value3"
            }
        },
        "string": "some string"
    },
    "array": [
        "value1",
        "value2"
    ],
    "arrayElement": "value1",
    "name": "some string",
    "nested": "Bearer value",
    "nested2": "Bearer value3",
    "nested_object": {
        "key3": "value3"
    },
    "object": {
        "key": "value",
        "key2": {
            "key3": "value3"
        }
    },
    "object_with_message": "Message: {\"key3\":\"value3\"}"
}
```

