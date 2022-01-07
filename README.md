# Temporal Go SDK samples

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B18405%2Fgithub.com%2Ftemporalio%2Fsamples-go.svg?type=shield)](https://app.fossa.com/projects/custom%2B18405%2Fgithub.com%2Ftemporalio%2Fsamples-go?ref=badge_shield)

This repository contains several sample Workflow applications that demonstrate the various capabilities of the Temporal Server via the Temporal Go SDK.

- Temporal Server repo: [https://github.com/temporalio/temporal](https://github.com/temporalio/temporal)
- Temporal Go SDK repo: [https://github.com/temporalio/sdk-go](https://github.com/temporalio/sdk-go)
- Go SDK docs: https://docs.temporal.io/docs/go/introduction

## How to use

Make sure the Temporal Server is running locally.
Follow the [Quick install guide](https://docs.temporal.io/docs/server/quick-install) to do that.

Then, clone this repository and follow the instructions in the README that is included with each sample.
The [helloworld](helloworld/README.md) sample is a good place to start.

You can learn more about running the Server locally in the [temporalio/docker-compose README](https://github.com/temporalio/docker-compose/blob/main/README.md).
And you can learn more about the Temporal Server technology in general via our [documentation](https://docs.temporal.io/).

## Samples directory

Each sample demonstrates one feature of the SDK, together with tests.

<!-- @@@SNIPSTART samples-go-readme-samples-directory -->

- [**Basic hello world**](https://github.com/temporalio/samples-go/tree/master/helloworld): Simple example of a Workflow Definition and an Activity Definition.

### API demonstrations

- **Async activity completion**: Example of an [Expense reporting](https://github.com/temporalio/samples-go/tree/master/expense) Workflow that communicates with a server API. Additional documentation: [How to complete an Activity Execution asynchronously in Go](https://docs.temporal.io/docs/go/activities/#asynchronous-activity-completion)

- [**Retry Activity Execution**](https://github.com/temporalio/samples-go/tree/master/retryactivity): This samples executes an unreliable Activity. The Activity is executed with a custom Retry Policy. If the Activity Execution fails, the Server will schedule a retry based on the Retry Policy. This Activity also includes a Heartbeat, which enables it to resume from the Activity Execution's last reported progress when it retries.

- [**Child Workflow**](https://github.com/temporalio/samples-go/tree/master/child-workflow): Demonstrates how to use execute a Child Workflow from a Parent Workflow Execution. A Child Workflow Execution only returns to the Parent Workflow Execution after completing its last Run.

- [**Child Workflow with ContinueAsNew**](https://github.com/temporalio/samples-go/tree/master/child-workflow-continue-as-new): Demonstrates that the call to Continue-As-New, by a Child Workflow Execution, is *not visible to the a parent*. The Parent Workflow Execution receives a notification only when a Child Workflow Execution completes, fails or times out. This is a useful feature when there is a need to **process a large set of data**. The child can iterate over the data set calling Continue-As-New periodically without polluting the parents' history.

- [**Cancellation**](https://github.com/temporalio/samples-go/tree/master/cancellation): Demonstrates how to cancel a Workflow Execution by calling `CancelWorkflow`, an how to defer an Activity Execution that "cleans up" after the Workflow Execution has been cancelled.

- **Coroutines**: Do not use native `go` routines in Workflows. Instead use Temporal coroutines (`workflow.Go()`) to maintain a [deterministic](https://docs.temporal.io/docs/go/workflows/#how-to-write-workflow-code) Workflow. Can be seen in the [Split/Merge](https://github.com/temporalio/samples-go/tree/master/splitmerge), [DSL](https://github.com/temporalio/samples-go/tree/master/dsl), [Recovery](https://github.com/temporalio/samples-go/tree/master/recovery), [PSO](https://github.com/temporalio/samples-go/tree/master/pso), and [Parallel](https://github.com/temporalio/samples-go/tree/master/parallel) Workflow examples.

- [**Cron Workflow**](https://github.com/temporalio/samples-go/tree/master/cron): Demonstrates a recurring Workflow Execution that occurs according to a cron schedule. This samples showcases the `HasLastCompletionResult` and `GetLastCompletionResult` APIs which are used to pass information between executions. Additional documentation: [What is a Temporal Cron Job?](https://docs.temporal.io/docs/content/what-is-a-temporal-cron-job).

- [**Encryption**](https://github.com/temporalio/samples-go/tree/master/encryption): How to use encryption for Workflow/Activity data with the DataConverter API. Also includes an example of stacking encoders (in this case encryption and compression). [Docs](https://docs.temporal.io/docs/go/workflows/#custom-serialization-and-workflow-security).

- [**Query Example**](https://github.com/temporalio/samples-go/tree/master/query): Demonstrates how to Query the state of a single Workflow Execution using the `QueryWorkflow` and `SetQueryHandler` APIs. Additional documentation: [How to send a Query to a Workflow Execution in Go](https://docs.temporal.io/docs/go/how-to-send-a-query-to-a-workflow-execution-in-go), [How to handle a Query in a Workflow in Go](https://docs.temporal.io/docs/go/how-to-handle-a-query-in-a-workflow-in-go).

- **Selectors**: Do not use the native Go `select` statement. Instead use [Go SDK Selectors](https://docs.temporal.io/docs/go/selectors) (`selector.Select(ctx)`) to maintain a [deterministic](https://docs.temporal.io/docs/go/workflows/#how-to-write-workflow-code) Workflow. Can be seen in the [Pick First](https://github.com/temporalio/samples-go/tree/master/pickfirst), [Mutex](https://github.com/temporalio/samples-go/tree/master/mutex), [DSL](https://github.com/temporalio/samples-go/tree/master/dsl), and [Timer](https://github.com/temporalio/samples-go/tree/master/timer) examples.

- **Sessions**: Demonstrates how to bind a set of Activity Executions to a specific Worker after the first Activity executes. This feature is showcased in the [File Processing example](https://github.com/temporalio/samples-go/tree/master/fileprocessing). Addition documentation: [How to use Sessions in Go](https://docs.temporal.io/docs/go/how-to-use-worker-sessions-in-go).

- **Signals**: Can be seen in the [Recovery](https://github.com/temporalio/samples-go/tree/master/recovery) and [Mutex](https://github.com/temporalio/samples-go/tree/master/mutex) examples. Additional documentation: [eCommerce application tutorial](https://docs.temporal.io/blog/tags/go-ecommerce-tutorial), [How to send a Signal to a Workflow Execution in Go](https://docs.temporal.io/docs/go/how-to-send-a-signal-to-a-workflow-execution-in-go), [How to handle a Signal in a Workflow Execution in Go](https://docs.temporal.io/docs/how-to-handle-a-signal-in-a-workflow-in-go).

- [**Search Attributes**](https://github.com/temporalio/samples-go/tree/master/searchattributes): Demonstrates how to use custom Search Attributes that can be used to find Workflow Executions using predicates (must use with [Elasticsearch](https://docs.temporal.io/docs/content/how-to-integrate-elasticsearch-into-a-temporal-cluster)).

- [**Timer Futures**](https://github.com/temporalio/samples-go/tree/master/timer): The sample starts a long running order processing operation and starts a Timer (`workflow.NewTimer()`). If the processing time is too long, a notification email is "sent" to the user regarding the delay (the execution does not cancel). If the operation finishes before the Timer fires, then the Timer is cancelled.

- [**Tracing and Context Propagation**](https://github.com/temporalio/samples-go/tree/master/ctxpropagation): Demonstrates  the client initialization with a context propagator, which propagates specific information in the `context.Context` object across the Workflow Execution. The `context.Context` object is populated with information prior to calling `StartWorkflow`. This example demonstrates that the information is available in the Workflow Execution and Activity Executions. Additional documentation: [How to use tracing in Go](https://docs.temporal.io/docs/go/tracing/).

- [**Updatable Timer**](https://github.com/temporalio/samples-go/tree/master/updatabletimer): Demonstrates timer cancellation and use of a Selector to wait on a Future and a Channel simultaneously.

- [**Greetings**](https://github.com/temporalio/samples-go/tree/master/greetings): Demonstrates how to pass dependencies to activities defined as struct methods.
-
- [**Greetings Local**](https://github.com/temporalio/samples-go/tree/master/greetingslocal): Demonstrates how to pass dependencies to local activities defined as struct methods.

### Dynamic Workflow logic examples

These samples demonstrate some common control flow patterns using Temporal's Go SDK API.

- [**Dynamic Execution**](https://github.com/temporalio/samples-go/tree/master/dynamic): Demonstrates how to execute Workflows and Activities using a name rather than a strongly typed function.

- [**Branching Acitivites**](https://github.com/temporalio/samples-go/blob/master/branch): Executes multiple Activities in parallel. The number of branches is controlled by a parameter that is passed in at the start of the Workflow Execution.

- [**Exclusive Choice**](https://github.com/temporalio/samples-go/tree/master/choice-exclusive): Demonstrates how to execute Activities based on a dynamic input.

- [**Multi-Choice**](https://github.com/temporalio/samples-go/tree/master/choice-multi): Demonstrates how to execute multiple Activities in parallel based on a dynamic input.

- [**Mutex Workflow**](https://github.com/temporalio/samples-go/tree/master/mutex): Demonstrates the ability to lock/unlock a particular resource within a particular Temporal Namespace. In this examples the other Workflow Executions within the same Namespace wait until a locked resource is unlocked. This shows how to avoid race conditions or parallel mutually exclusive operations on the same resource.

- [**Parallel Workflow**](https://github.com/temporalio/samples-go/tree/master/parallel): This sample executes multiple branches in parallel using the `workflow.Go()` API.

- [**Pick First**](https://github.com/temporalio/samples-go/tree/master/pickfirst): This sample executes Activities in parallel branches, picks the result of the branch that completes first, and then cancels other Activities that have not finished.

- [**Split/Merge**](https://github.com/temporalio/samples-go/tree/master/splitmerge): Demonstrates how to use multiple Temporal coroutines, instead of native goroutines. This samples to processes chunks of a large work item in parallel, and then merges the intermediate results to generate the final result. Do not use native goroutines in Temporal Workflows. Instead, the coroutine provided by the workflow.Go API.

- [**Synchronous Proxy Workflow pattern**](https://github.com/temporalio/samples-go/tree/master/synchronous-proxy): This sample demonstrates a synchronous interaction with a "main" Workflow Execution from a "proxy" Workflow Execution. The proxy Workflow Execution sends a Signal to the "main" Workflow Execution, then blocks, waiting for a Signal in response.

### Scenario based examples

- [**DSL Workflow**](https://github.com/temporalio/samples-go/tree/master/dsl): Demonstrates how to implement a DSL-based Workflow. This sample contains 2 yaml files that each define a custom "workflow" which instructs the Temporal Workflow. This is useful if you want to build in a "low code" layer.

- [**Expense Request**](https://github.com/temporalio/samples-go/tree/master/expense): This demonstrates how to process an expense request. This sample showcases how to complete an Activity Execution asynchronously.

- [**File Processing**](https://github.com/temporalio/samples-go/tree/master/fileprocessing): Demonstrates how to download and process a file using set of Activities that run on the same host. Activities are executed to download a file from the web, store it locally on the host, and then "process it". This samples showcases how to handle a scenario where all subsequent Activities need to execute on the same host as the first Activity in the sequence. In Go, this is achieved by using the Session APIs.

- [**Particle Swarm Optimization**](https://github.com/temporalio/samples-go/tree/master/pso): Demonstrates how to perform a long iterative math optimization process using particle swarm optimization (PSO). This sample showcases the use of parallel executions, `ContinueAsNew` for long histories, a Query API, and the use of a custom `DataConverter` for serialization.

- [**Prometheus Metrics**](https://github.com/temporalio/samples-go/tree/master/metrics): Demonstrates how to instrument Temporal with Prometheus and Uber's Tally library.

<!-- @@@SNIPEND -->

### Pending examples

Mostly examples we haven't yet ported from https://github.com/temporalio/samples-java/

  - Async activity calling: *Example to be completed*
  - Async lambda:  *Example to be completed*
  - Periodic Workflow: Workflow that executes some logic periodically. *Example to be completed*
  - Exception propagation and wrapping: *Example to be completed*
  - Polymorphic activity: *Example to be completed*
  - SAGA pattern:  *Example to be completed*
  - Side Effect:  *Example to be completed* - [Docs](https://docs.temporal.io/docs/go/side-effect)

### Fixtures

These are edge case examples useful for Temporal internal development and bug reporting. [See their readme for more details](https://github.com/temporalio/samples-go/tree/master/temporal-fixtures).
