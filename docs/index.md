## Overview

While working on client and open-source projects we came to realize that we could do so much more for the developer if
there were a central file that contained all the information about a project, or in the case
of secrets at least a reference to where the secrets could be find given the proper authorization.

**So we decided to create one.**

## Core Principles
The following are core principles of `Project Config`:

1. **All properties are declarative**: Properties only contain values, not instructions. For example, properties must not contain [Bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell) commands.
1. **Property values are context free**: Properties of Project Config must not require a given context to be interpretted correctly. For example, it they make sense for macOS they should also make sense for Windows and Linux.

## Reference
The following table provides the known top-level properties of a Project Config file.
Note that additional top-level properties are potential to document framework-specific values.

Property|Description / Values | Example
--------|------------------|-------
`schema`| Semantic version number of the "schema" used for JSON but  without a patch version — e.g. `1.0` but not `1.0.0` — of the schema used. We will start maintaining this once the spec reaches `1.0`| `"1.0"`
`name`| Proper name for project.| `"ACME Website"`
`identity`|Unique identity for project relative to its `source`.|`"github.com/acme/website"`
`hostname`|Domain where the production project will be hosted, if applicable. This will typically be without `"www"` which wcan be handled via `aliases` when applicable.|`"acme.com"`
`description`|Prose description of the project.|`"Website for ACME Corporation at acme.example.com"`
`namespace`|Symbol name that can be used in code for namespaces, class name prefixes, etc. Typically this will be in either proper case or proper case with underscores unless the abbreviation comes at the end.|e.g. `"AcmeWeb"` or  `"Acme_Web"`
`slug`|A lowercase version of the namespace that can be used in URLs. Dashes are used to separate words in slugs, not underscores. Periods can also be used, but generally dashes are preferred. |e.g. `acme-web` or `"acme.web"`,<br>never `acme_web`.
`prefix`|A lowercase prefix to be used to _"namespace"_ identifiers such as databases, table names, field names, etc. The prefix should have a trailing underscore.| e.g. `acme_`
`type`|Indicates the type of project from one of the well-known types listed here. Prefix with `"x-"` if not well-known.|One of: `"website"`, `"docs"` , `"library"` _(with more types to come; pull requests encouraged.)_ <br>Also: `"x-webservice"`
`aliases`|Array of string aliases for the `hostname`.  Can include the template var `{hostname}` so that you won't need to duplicate the hostname in many different places.
`dev`|A structure containing information about **development** tools and configurations for this project. Click to learn about the [**Dev**](Dev.md) structure.
`stack`|A structure containing information about the **stack** required. Click to learn about the [**Stack**](Stack.md) structure.
`source`|A structure containing information about the **source code**. Click to learn about the [**Source**](Source.md) structure.
`deploy`|A structure containing information about **deployment** processes and destinations. Click to read about [**Deploy**](Deploy.md) structure.
`hosts`|A map of structures containing information **host servers**. Click to read about [**Hosts**](Hosts.md) map.
`extra`|Freeform information to allow for any additional information to be maintain in the project.
