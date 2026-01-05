# EXFS: Extensible File System

A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.

## Definitions

On initiating an EXFS repository, a "views" folder (`exfs_views/`) is created which holds multiple codebases of your choice. The top level folder is nomenclatively referred to as a "view".

1. Any public file in a view is a recognized **source file**.
2. Any public YAML in `.exfs/` is a recognized **configuration file**.

An exfs repository looks like this:

```text
myproject/
├── .git/                         ← (optional, exfs is git-agnostic)
├── exfs_views/                   ← (exfs jurisdiction)
│   ├── .exfs/                    ← (root config, applies to all views)
│   │   └── exec/
│   │       └── go_to_md.yml
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
