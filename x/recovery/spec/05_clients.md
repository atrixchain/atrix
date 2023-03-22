<!--
order: 5
-->

# Clients

A user can query the `x/recovery` module using the CLI, gRPC or REST.

## CLI

Find below a list of `Atrixd` commands added with the `x/recovery` module. You can obtain the full list by using the `Atrixd` -h command.

### Queries

The query commands allow users to query Recovery state.

**`params`**
Allows users to query the module parameters.

```bash
Atrixd query recovery params [flags]
```

## gRPC

### Queries

| Verb   |              Method              |           Description |
| :----- | :------------------------------- | :-------------------- |
| `gRPC` | `Atrix.recovery.v1.Query/Params` | `Get Recovery params` |
| `GET`  |   `/Atrix/recovery/v1/params`    | `Get Recovery params` |
