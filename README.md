# EXFS: Extensible File System

A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.

## Definitions

On initiating an EXFS repository, a "views" folder (`exfs_views/`) is created which holds multiple codebases of your choice. The top level folder is nomenclatively referred to as a "view".

1. Any public file in a view is a recognized **source file**.
2. Any public YAML in `.exfs/` is a recognized **configuration file**.

An exfs repository looks like this:

```text
my_project/
├── .git/                         ← (optional, exfs is git-agnostic)
├── exfs_views/                   ← (exfs jurisdiction)
│   ├── .exfs/                    ← (root config, applies to all views)
│   │
│   ├── my_execs/                 ← (my_execs view)
│   │   └── go_to_md.yml
│   │
│   ├── api/                      ← (api view)
│   │   ├── .exfs/
│   │   │   └── go/
│   │   │       ├── handler.yml
│   │   │       └── routes.yml
│   │   ├── handler.go
│   │   └── routes.go
│   │
│   ├── docs/                     ← (docs view)
│   │   ├── .exfs/
│   │   │   └── md/
│   │   │       └── API.yml
│   │   └── API.md
│   │
│   └── sdk/                      ← (sdk view)
│       ├── .exfs/
│       │   ├── go/
│       │   │   └── server.yml
│       │   └── py/
│       │       └── client.yml
│       ├── go/
│       │   └── server.go
│       └── py/
│           └── client.py
│
├── src/                          ← (outside jurisdiction)
└── README.md
```

### Source Files

All source files can be referenced inside configuration files using `exfs.view.<path_from_view>.<your_file.ext>`. Which essentially makes all exfs files first class citizens of an exfs codebase.

### Configuration Files

All recognized configurations are YAML files, and these files have rigid schemas as defined in [exfs.org/schemas/](https://exfs.org/schemas/).

### Schemas

#### File Schema

These are configs in `.exfs/<ext>/<filename>.yml` that define **captures** for `<filename>.<ext>`. A capture is a named block of code extracted from a source file.

```yaml
# api/.exfs/go/handler.yml → captures for handler.go
comments:
  type: core.go.comments # what to capture
  limit: 10 # how many (optional)
```

**Nesting** — captures can contain captures:

```yaml
# api/.exfs/go/handler.yml
my_funcs:
  type: core.go.functions
  contains:
    doc_strings:
      type: core.go.comments
    _:
      type: core.go.calls
```

**Referencing** — captures are addressable via `exfs.<path_to_.exfs>.<ext>.<file_name>.<capture_name>`:

```text
server/
├── .exfs/ts/router.yml        ← config applies to all router.ts below
├── router.ts
├── api/router.ts
├── auth/router.ts
└── auth/login/router.ts
```

**Sync** — links captures across views when folder structures match:

```text
server/                              docs/api_ref/
├── .exfs/ts/router.yml              ├── .exfs/md/endpoint.yml
├── router.ts                        ├── endpoint.md
├── api/router.ts                    ├── api/endpoint.md
├── auth/router.ts                   ├── auth/endpoint.md
└── auth/login/router.ts             └── auth/login/endpoint.md
```

```yaml
# server/.exfs/ts/router.yml         # docs/api_ref/.exfs/md/endpoint.yml
funcs:                                content:
  type: core.ts.functions               type: core.md.content
  sync:                                 sync:
    - path: exfs.docs.api_ref.md.endpoint.content
```

Because subtrees match (`api/`, `auth/`, `auth/login/`), a single config syncs all corresponding files. This is the **repeating subproblem** — the same structural pattern appears at multiple levels. EXFS hinges on this natural property to provide self-referencing single source of truth.

## CLI's Expected Behavior

```text
"exfs" | "xfs"?
├─ no  → OUT OF SCOPE
└─ yes →
    ├─ <enter> → BEHAVIOR: print usage, exit 1
    └─ <space> →
        ├─ "init"?
        │   └─ yes →
        │       ├─ <enter> → BEHAVIOR: init in current dir with defaults
        │       └─ <space> →
        │           ├─ <path>? → BEHAVIOR: init in <path>
        │           └─ "--name <name>"? → BEHAVIOR: use <name> instead of "exfs"
        │
        ├─ "sync"?
        │   └─ yes →
        │       ├─ <enter> → BEHAVIOR: sync all views
        │       └─ <space> →
        │           └─ <view>? → BEHAVIOR: sync only <view>
        │
        ├─ "status"?
        │   └─ yes →
        │       ├─ <enter> → BEHAVIOR: show pending changes integration with last synced DAG
        │       └─ <space> →
        │           └─ <view>? → BEHAVIOR: status for specific view
        │
        ├─ "purge"?
        │   └─ yes →
        │       ├─ <enter> → BEHAVIOR: error, missing linkage
        │       └─ <space> →
        │           └─ <linkage>? → BEHAVIOR: remove sync references on all sides
        │               ├─ Precondition: state must be synced (clean)
        │               ├─ Removes sync entry from all connected captures
        │               └─ Does NOT delete captures, only severs linkage
        │
        ├─ "help"?
        │   └─ yes → BEHAVIOR: print usage, exit 0
        │
        └─ "version"?
            └─ yes → BEHAVIOR: print version, exit 0
```
