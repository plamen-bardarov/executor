# Executor

Let me run that for you.

Executor is a logical process running inside the
[Rep](https://github.com/cloudfoundry/rep) that: \* manages container
allocations against resource constraints on the Cell, such as memory and
disk space \* implements the actions detailed in the API documentation
\* streams stdout and stderr from container processes to the
metron-agent running on the Cell, which in turn forwards to the
Loggregator system \* periodically collects container metrics and emits
them to Loggregator

> \[!NOTE\]
>
> This repository should be imported as
> `code.cloudfoundry.org/executor`.

# Contributing

See the [Contributing.md](./.github/CONTRIBUTING.md) for more
information on how to contribute.

# Working Group Charter

This repository is maintained by [App Runtime
Platform](https://github.com/cloudfoundry/community/blob/main/toc/working-groups/app-runtime-platform.md)
under `Diego` area.

> \[!IMPORTANT\]
>
> Content in this file is managed by the [CI task
> `sync-readme`](https://github.com/cloudfoundry/wg-app-platform-runtime-ci/blob/c83c224ad06515ed52f51bdadf6075f56300ec93/shared/tasks/sync-readme/metadata.yml)
> and is generated by CI following a convention.
