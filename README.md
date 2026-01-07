# EXFS: Extensible File System

A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.

## Definitions

On initiating an EXFS repository, a "views" folder (`exfs_views/`) is created which holds multiple codebases of your choice. The top level folder is nomenclatively referred to as a "view".

1. Any public file in a view is a recognized **source file**.
2. Any public YAML is a referenceable **configuration file**.

An exfs repository looks like this:

```text
my_project/
├── .git/                         ← (optional, exfs is git-agnostic)
├── exfs_views/                   ← (exfs jurisdiction)
│   ├── .exfs/                    ← (root config, applies to all views)
│   │
│   ├── engines/                  ← (engines view - free roaming configs)
│   │   ├── go_to_md.yml
│   │   └── ai_review.yml
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

## Schemas

### File Schema

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
    calls:
      type: core.go.calls
```

**Referencing** — captures are addressable via `exfs.<view>.<path>.<ext>.<file>.<capture>`:

```text
server/
├── .exfs/ts/router.yml        ← config applies to all router.ts below
├── router.ts
├── api/router.ts
├── auth/router.ts
└── auth/login/router.ts
```

**Sync** — links captures across views:

```yaml
# server/.exfs/ts/router.yml
funcs:
  type: core.ts.functions
  sync:
    # Bidirectional (core.transpiler, default)
    - with: exfs.docs.api_ref.md.endpoint.content

    # Unidirectional: this capture → peer
    - to: exfs.tests.stubs.ts.router.mocks
      engine: exfs.engines.func_to_stub

    # Unidirectional: peer → this capture
    - from: exfs.schemas.openapi.yaml.routes.endpoints
      engine: exfs.engines.openapi_to_ts
```

**Sync entry types:**

| Key    | Direction       | Engine                          |
| ------ | --------------- | ------------------------------- |
| `with` | bidirectional   | `core.transpiler` only          |
| `to`   | this → peer     | any (default: core.transpiler)  |
| `from` | peer → this     | any (default: core.transpiler)  |

Because subtrees match (`api/`, `auth/`, `auth/login/`), a single config syncs all corresponding files. This is the **repeating subproblem** — the same structural pattern appears at multiple levels.

### Engine Schema

Engines are **free-roaming** YAML files (not in `.exfs/`) that define transforms for captures. Reference them as `exfs.<view>.<path>.<engine_name>`.

**Core transpiler** (`core.transpiler`) is built-in and bidirectional:

- comment ↔ comment
- function signature ↔ function
- variable ↔ variable
- string ↔ string
- import/require/use statements

**Custom engines** are unidirectional (use `to` or `from` in sync):

```yaml
# engines/go_to_md.yml
engine:
  name: go_to_md
  version: "1.0"

  # Container (optional - omit for native execution)
  image: golang:1.22-alpine

  # Environment (optional)
  env:
    - GOPROXY=direct
    - exfs.config.secrets.env # can reference exfs files

  # Execution steps (required)
  # Last step: stdin = capture content, stdout = transformed content
  steps:
    - name: Install dependencies
      run: go mod download

    - name: Build transformer
      run: go build -o /tmp/transform ./cmd/transform

    - name: Transform
      run: /tmp/transform
```

**Usage in file schema:**

```yaml
# api/.exfs/go/handler.yml
funcs:
  type: core.go.functions
  sync:
    # Push to docs using custom engine
    - to: exfs.docs.api.md.handler.content
      engine: exfs.engines.go_to_md

    # Pull from schema using another engine
    - from: exfs.schemas.api.yaml.handler.spec
      engine: exfs.engines.yaml_to_go
```

**Use cases:**

- **Transpilation** — Go comments → Python docstrings → Markdown docs
- **AI in the loop** — pipe captures through LLMs for code review, generation, translation
- **CI/CD** — test, build, deploy pipelines triggered by capture changes
- **Schema generation** — JSON schema → TypeScript types → Swift structs
- **Documentation** — extract function signatures → generate API docs
- **Anything** — if it takes input and produces output, it's an engine

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
