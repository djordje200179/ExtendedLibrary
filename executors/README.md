# Executors

Java-like Future tasks for Go, that can calculate and 
return values or be functions that return nothing.

Every task can be cancelled, and every panic is recovered and returned.
You can check status of the task whenever you want. And you can also
wait for the task to finish.

By default, there exists a global executor that has as many goroutines
as there are CPU cores. But, you can also create your own executors with
custom number of goroutines.

## Construction

### Tasks

Currently, two types of tasks are supported:
- Actions - tasks that return nothing
- Futures - generic tasks that return a value

Both types of tasks have two constructors that allow you to create new instances. 
You should supply a function that will be executed by the task and
optionally a context that will be used in the task.

```go
action := executors.NewDefaultAction(PrintText)
future := executors.NewDefaultFuture(CalcNumber)
```

### Executors
To create a new executor, you can use the `NewExecutor` function that takes
the number of goroutines and queue size as parameters.

```go
executor := executors.NewExecutor(10, 100)
```

Or, even better, you can use the `executors.DefaultExecutor` and don't
worry about custom executors.


## Invoking
To invoke a task, you can use the `Submit` method of the executor. It will
place the task in the queue and return immediately.

```go
executors.Submit(future)
```

## Task status
Because every task can panic or be cancelled, you can check the status of the
task whenever you want.

```go
task.IsStarted()
...
task.Wait()
if task.IsCancelled() {
	...
}
```

If the task failed and if you try to get the result from `Future`, the error
will be propagated to the caller (method will panic with same error).
