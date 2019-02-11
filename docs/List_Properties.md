## List Properties

Project.JSON typically uses JSON objects for properties _(though not always)_ where others might use arrays for properties.
For example, team members might otherwise be represented like this in another specification:

    {
        ..., 
        "team": [
            {
                "name": "Mike Schinkel",
                "type": "person"
            },
            {
                "name": "George P. Burdell",
                "type": "person"
            },
            {
                "name": "NewClarity Consulting LLC",
                "type": "organization"
            }
        ],  
        ...
    }        

For Project.JSON we instead use object properties rather than array elements:

    {
        ..., 
        "team": {
            "mikeschinkel": {
                "name": "Mike Schinkel",
                "type": "person"
            },
            "sarutole": {
                "name": "George P. Burdell",
                "type": "person"
            },
            "newclarity": {
                "name": "NewClarity Consulting LLC",
                "type": "organization"
            }
        ],  
        ...
    }        

### @Defaults
In addition, Project JSON supports defaults, so the above can be rewritten:

    {
        ..., 
        "team": {
            "@defaults": {
                "type": "person"
            },
            "mikeschinkel": {
                "name": "Mike Schinkel",
            },
            "georgepburdell": {
                "name": "George P. Burdell",
            },
            "newclarity": {
                "name": "NewClarity Consulting LLC",
                "type": "organization"
            }
        ],  
        ...
    }        

