## Overview

While creating [WPLib Box](https://github.com/wplib/wplib-box) we came to realize that we could do so much more for the developer if
there were a central file that contained all the information about a project, or in the case
of secrets at least a reference to where the secrets could be find given the proper authorization.

_So we decided to create one:_ `Project.JSON`

## Core Principles
The following are core principles of `Project.JSON`:

1. **All properties are declarative**: Properties only contain values, not instructions. For example, properties must not contain [Bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) commands.
1. **Property values are context free**: Properties of Project.JSON must not require a given context to be interpretted correctly. For example, it they make sense for macOS they should also make sense for Windows and Linux.

## Reference
The following table provides the known top-level properties of a Project.JSON file.
Note that additional top-level properties are potential to document framework-specific values.

Property|Description / Values | Example
--------|------------------|-------
`name`| Proper name for project.| `"ACME Website"`
`identity`|Unique identity for project relative to its `source`.|`"acme/website"`
`description`|Prose description of the project.|`"Website for ACME Corporation at acme.example.com"`
`source`|The canonical source code host for the project. |A well-known host such as `"github"`, `"bitbucket"` or `"gitlab"`, or an Internet domain identified by dots in the value, e.g. `"git.wpengine.com"`.
`symbol`|Symbol name that can be used in code for namespaces, class name prefixes, etc. Typically this will be in either proper case or proper case with underscores unless the abbreviation comes at the end in which case the abbreviation can be fully capitalized.|e.g. `"AcmeWeb"` or  `"Acme_Web"`<br>Trailing abbreviation: `"ProjectJSON"`
`slug`|A version of the shortname that can be used in URLs. Dashes are used to separate words in slugs, not underscores. Periods can also be used, but generally dashes are preferred. |e.g. `acme-web` or `"acme.web"`,<br>never `acme_web`.
`prefix`|A lowercase prefix to be used to _"namespace"_ identifiers such as databases, table names, field names, etc. The prefix will have a trailing underscore.| e.g. `acme_`
`type`|Indicates the type of project from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"website"`, `"documentation"` , `"library"` _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-webservice"`
`version`|A [semantic version](http://semver.org/) compatible value labeling the current state of project.|e.g. `2.1.7`, `1.0.0-alpha+123`
`frameworks`|A command separated list of named frameworks for the project from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"wordpress"`, `"drupal"` , `"joomla"`, `"nodejs"`, `"angularjs"`, `"reactjs"`, `"vuejs"`, `"emberjs"`, `"lavarel"`, `"n/a"`  _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-rails"`
`languages`|A command separated list of named languages for the project from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"php"`, `"javascript"`, `"python"`, `"ruby"`, `"golang"`, `"java"`, `"bash"`, `"powershell"`, `"html"`, `"markdown"`, `"css"`, `"less"`, `"sass"`, `"n/a"` _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-obscure-lang"`
`team`|An object containing an object property of each of those involved on the team.|Team members are documented [here](Team_Members.md).
`repositories`|An object containing an object property of each of the repositories containing source for the project.|Repositories are documented [here](Repositories.md).


