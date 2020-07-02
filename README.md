Brazilian Cities API
====================

This API contains most of the brazilian cities and can be used as search database.

It contains the `cities` endpoint that will return all cities in the system and accept
a `q` parameter to do a `text` search based on the city name.



API
====

`GET /cities/`

Return all cities

    [{
        "id": "fdba782e-f421-4837-93f8-9f73fbd70da4",
        "name": "São José da Safira"
        "verbose": "São José da Safira, MG"
        "state: {
            "id": "d233ffc9-b681-4996-b65b-fa9f9393aa15",
            "name": "MG",
            "verbose": "Minas Gerais",
        },
    }]


`GET /cities/?q=São José`
   
Return all cities that contains `sao jose`, the api will normalize
the search parameters and validate it against a normalized city name.

    [{
        "id": "fdba782e-f421-4837-93f8-9f73fbd70da4",
        "name": "São José da Safira"
        "verbose": "São José da Safira, MG"
        "state: {
            "id": "d233ffc9-b681-4996-b65b-fa9f9393aa15",
            "name": "MG",
            "verbose": "Minas Gerais",
        },
    }]
