# Todo

RESTful HTTP API to manage todos in Go.

## Context

Users have difficulty remembering tasks, like chores that need to be done or groceries that have to be bought. They need a way to organize these type of tasks so they are not forgotten.

## Scope

A user needs to:

- Add a single task.
- List all tasks.
- Update a task to flag/unflag as done.
- Delete one task at a time.

Initially tasks will not be grouped by user, will all be public and not persisted.

## Goals

Create an RESTful HTTP API with which (web) clients can interact.

## Language

| Name | Description                                               |
| ---- | --------------------------------------------------------- |
| Task | Represent something that needs to be done, i.e. a "todo". |

## Design

Business logic and objects are modeled using domain types.

We create one package per dependency to make it easier to test and replace these dependecies. As well as preventing cyclic dependencies.

### Root package contains domain types

The root package `todo.go` _only_ contains **domain types** like `struct` and `interface`. It does _not_ contain any dependencies, i.e. it does not depend on any other package in the application.

### Packages are grouped by dependency

Domain types don't "do" anything. To interact with these types a **service** is required. Services are packages that operate on or with the domain types and are located in `pkg`. These services (or "subpackages") are adapters that live between domain and implementation.
In other words, a service is a specific implementation of a dependency that satisfies an `interface` of a domain type. For example `storage/memory.Service` satisfies `todo.TaskService`.

This approach makes it easier to change- and combine dependencies.

### Shared mock subpackage

Because domain interfaces isolate dependencies, these connection points can be used to inject mock implementations. This helps to isolate tests as well which improves testability.

### Main package hooks up all dependencies

Because an application may produce multiple binaries, the convention is used to place a main package as a sub dir of `cmd`. Every main package can use dependency injection to choose which packages are passed to which objects.

## Design Alternatives

### Monolithic Package

This means putting all code in a single package.

Works well for very small applications and is very simple to get started. It also prevents creating cyclic dependencies because there are no other packages. However, it will become more difficult to read, navigate and maintain the application as the amount of code grows--this also applies to testing. It also "doesn't scale" well when working with multiple people on the application.

### Group by Function

This means grouping entities like HTTP handlers, controllers and models in a package.

A danger here is that it's easy to create cyclic dependencies. In Go dependencies must be acyclic, otherwise code will not compile. Therefore dependencies _must_ be "one way" when using this approach.
Another downside is that as the project grows it will become more difficult to create a mental model of the application. The developer needs to make a greater effort to "piece all things together", since they are fragmented.

### Group by Entity

This means grouping by Domain entity. For example users or books.

Here it becomes easier to navigate the code, but suffers from the same downsides as grouping by function. Additionaly, here it's easier to introduce "naming stutter". For example `books.Book`, which conflicts with Go best practices.

## Resources

- [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)
- [Ashley McNamara + Brian Ketelsen - Go best practices](https://www.youtube.com/watch?v=MzTcsI6tn-0)
