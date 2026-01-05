# EXFS: Extensible File System

A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.

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
