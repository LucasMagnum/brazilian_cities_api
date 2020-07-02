Brazilian Cities API
====================

This API contains the 99% of the brazilian cities, with a search endpoint.

The search will normalize both the output as the input to return a wider number of results.


API
====

GET /cities/
    Return all cities

    [{
        "id": "fdba782e-f421-4837-93f8-9f73fbd70da4",
        "name": "Governador Valadares"
        "verbose": "Governador Valadares, MG"
        "state: {
            "id": "d233ffc9-b681-4996-b65b-fa9f9393aa15",
            "name": "MG",
            "verbose": "Minas Gerais",
        },
    }]


GET /cities/?q="gov"
    Return all cities

    [{
        "id": "fdba782e-f421-4837-93f8-9f73fbd70da4",
        "name": "Governador Valadares"
        "verbose": "Governador Valadares, MG"
        "state: {
            "id": "d233ffc9-b681-4996-b65b-fa9f9393aa15",
            "name": "MG",
            "verbose": "Minas Gerais",
        },
    }]