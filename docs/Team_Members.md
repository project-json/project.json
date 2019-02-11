## Team Members

Team members are included in a `project.json` file as a [List Property](List_Properties.md).

Team members support the following object properties:


Property|Description / Values | Example
--------|------------------|-------
`name`| Full name of the team member.| e.g. `"Mike Schinkel"`, `"The WPLib Team"` or `"NewClarity Consulting LLC"`.
`email`|Contact email for the team member.|`"mike@newclarity.net"`, `"team@wplib.org"` or `"team@newclarity.net"`.
`role`|Freeform description of role this team member plays on the team.|`"Architect"`, `Advocate` or `"Sponsor"`.
`type`|Indicates the type of team member from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"person"`, `"project"`, `"organization"` _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-documentor"`


### Example:

An example from the context of a `project.json` file:

    {
        ...,
        "team": {
            "mikeschinkel": {
                "name": "Mike Schinkel",
                "type": "person",
                "role": "Architect",
                "email": "mike@newclarity.net"
            },
            "newclarity": {
                "name":  "NewClarity Consulting LLC",
                "type": "organization",
                "role": "Sponsor",
                "email": "team@newclarity.net"
            },
            "wplib": {
                "name":  "The WPLib Team",
                "type": "project",
                "role": "Advocate",
                "email": "team@wplib.org"
            }
        },
        ...
    }




