# Data structures

Java-like Future tasks for Go, that can calculate and 
return values or be functions that return nothing.

Every task can be cancelled, and every panic is recovered and returned.
You can check status of the task whenever you want. And you can also
wait for the task to finish.

By default, there exists a global executor that has as much goroutines
as there are CPU cores. But, you can also create your own executors with
custom number of goroutines.

## Types
Currently, two types of tasks are supported:
- Actions - tasks that return nothing
- Futures - generic tasks that return a value

## Construction

### Constructors
Every type has two constructors that allow you to create new instances of
the type. You should supply a function that will be executed by the task and
optionally a context that will be used in the task.

```go
action := executors.NewDefaultAction(PrintText)
future := executors.NewDefaultFuture(CalcNumber)
```

### Invoking
To invoke a task, you can use the `Submit` method of the executor. It will
place the task in the queue and return immediately.

```go
executors.Submit(future)
```