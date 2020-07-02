Brazilian Cities API
====================

This API contains the 99% of the brazilian cities, with a search endpoint.

The search will normalize both the output, and the input to return a wider number of results.


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