## Repositories

Repositories are included in a `project.json` file as a [List_Property](List_Properties.md).

Repositories support the following object properties:

Property|Description / Values | Example
--------|------------------|-------
`identity`|Unique identity for repository relative to its `source`.|`"project-json/project.json"`
`source`|The canonical source code host for the repository. |A well-known host such as `"github"`, `"bitbucket"` or `"gitlab"`, or an Internet domain identified by dots in the value, e.g. `"git.wpengine.com"`.
`type`|Indicates the type of repository from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"git"`, `"mecurial"`, `"svn"` _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-bzr"`
`description`|Prose description of the repository for benefit of future project maintainers.|`"The canonical source repository for the Project.JSON project."`


### Example:

An example from the context of a `project.json` file:

    {
        ...,
        "repos": {
            "canonical": {
                "identity":  "project-json/project.json",
                "source": "github",
                "description": "The canonical source repository for the Project.JSON project"
            }
        },
        ...
    }
