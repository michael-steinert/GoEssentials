# Go Essentials

## Features

- **Conciseness, Simplicity, and Safety**: Go is designed to be simple to understand and easy to read, aiming to reduce the Complexity of Coding
- **Efficiency and Performance**: Comparable to C++ in terms of Efficiency and Performance, making it suitable for Systems Programming
- **High Scalability**: _Go Routines_ and _Channels_ allow for easy and efficient Concurrency, making it ideal for Cloud and Network Services
- **Garbage Collection**: Automated Memory Management helps prevent Memory Leaks
- **Rich Standard Library**: Go offers a broad Range of built-in Packages for various Functionalities like HTTP Servers, JSON Encoding, and more
- **Strongly Typed and Compilation**: Go is a statically typed Language that compiles to Machine Code, ensuring Type Safety and improved Performance
- **Cross-Platform Development**: Compiled Go Programs can be run on any Platform without Modifications, enhancing Portability
- **Tools and Documentation**: Go comes with a Suite of Tools for Testing, Formatting, and Documentation, making the Development Process smoother

## Structure

- **Modularity**: Each Program or Library is a _Module_ in Go. This enhances Maintainability and Versioning
- **Package System**: _Modules_ consist of _Packages_, and each _Package_ corresponds to a Directory. An Executable in Go is created from one (_Main Package_, i.e., Root Directory) or more _Packages_
- **Scope and Accessibility**:
  - Method and Variable Names that start with a lowercase Letter are private and only visible within their _Package_
  - Names that begin with an uppercase Letter are public and visible across different _Packages_
- **Importing _Packages_**: Go allows you to import other _Packages_, enabling Code Reuse and modular Design
- **Naming Conventions**: Go has strict Naming Conventions that encourage Readability and Maintainability of Code

## Testing

- **File Naming**: Test Files must have the Suffix `_test`. Go will exclude these Files during Compilation and Execution.
- **Package Naming for Tests**: _Test Packages_ should have the Suffix `_test`. This ensures that Tests are Black-Box, meaning they cannot interact with the Internal (i.e., private Variables and Functions) of the _Package under Test_. Tests are restricted to using the public Variables and Functions
- **Table-Driven Testing**: Go supports Table-driven Tests, allowing for clear and concise Testing of multiple Scenarios
- **Benchmarking**: Go includes support for Benchmark Tests, which are useful for Measuring and Improving Performance
- **Mocking and Interfaces**: Go's Interface System makes it easier to create Mocks for Testing
- **Coverage Tools**: Go has built-in Tools to analyze Test Coverage, helping to identify untested Parts of a Codebase

## Asynchronous

- **GoRoutines**: In Go, GoRoutines are lightweight Threads managed by the Go Runtime rather than the Operating System. They are fundamental for concurrent Execution
- **Parallel Execution with GoRoutines**: GoRoutines enable the parallel Execution of Tasks. This allows for efficient Handling of Operations like I/O-bound or CPU-bound Processes
- **Non-blocking Execution**: By default, Go does not wait for a GoRoutine to complete if the Main Program finishes Execution. It's important to manage GoRoutine Lifecycles to ensure they complete as intended
- **Synchronization with WaitGroups**: WaitGroups provide a Mechanism to synchronize multiple GoRoutines. They can be used to block the Main Program until all GoRoutines have finished Executing, ensuring proper Completion of concurrent Tasks
- **Channels for Communication**: Channels in Go provide a Way for GoRoutines to communicate with each other, enabling the Passing of Data or Signaling between concurrent Threads
- **Select Statement for Concurrency Control**: The Select Statement in Go allows a GoRoutine to wait on multiple Communication Operations, enabling more sophisticated Concurrency Patterns and Handling of Timeouts and Cancellation
- **Handling Goroutine Lifecycles**: It's crucial to manage the Lifecycle of GoRoutines to avoid issues like GoRoutine Leaks. Utilizing Context Package for Cancellation and Deadlines is a common Practice for Lifecycle Management in Go
- **Mutex for Resource Synchronization**: _Mutex_ (Mutual Exclusion) in Go provides a Mechanism to ensure that only one GoRoutine can access a particular Resource at a Time. When a GoRoutine acquires a _Mutex_ to access a Resource, it locks the Resource. This Lock prevents other GoRoutines from Accessing the same Resource until the Lock is released. This ensures safe and synchronized Access to shared Resources, preventing Race Conditions and ensuring Data Integrity during concurrent Execution

## Data Types

### Slices

- **Definition**: Slices in Go are dynamic Arrays that provide a flexible and powerful Way to work with Sequences of typed Data. They are built on Top of Arrays and provide a more convenient and efficient Interface for various Operations
- **Properties**
  - **Length and Capacity**: Slices have two Properties, Length and Capacity. The Length is the Number of Elements it contains, and the Capacity is the Number of Elements it can hold before needing to allocate more Memory
  - **Zero Value**: The Zero Value of a Slice is _nil_. A _nil_ Slice has a Length and Capacity of 0 and has no underlying Array
  - **Slicing Operations**: Slices can be re-sliced, creating a new Slice that points to the same Array but with different Start and End Positions
  - **Appending Elements**: The _append_ Function is used to add new Elements to a Slice. This Operation might result in Allocating a new underlying Array if the Capacity of the existing Array is exceeded
  - **Pass By Reference**: When a Slice is passed to a Function, it is done by Reference, meaning Changes made to the Elements of the Slice in the Function are visible outside the Function

### Maps

- **Definition**: Maps in Go are Hash Tables or Dictionaries that store Key-Value Pairs. They are used to lookup Values quickly by a Key
- **Characteristics**
  - **Unordered Collection**: The Elements in a Map are stored in an unordered Way. The Order of Insertion does not guarantee the Order of Iteration
  - **Key-Value Pairs**: Each Element in a Map is a Key-Value Pair. The Keys are unique within a Map while the Values can be duplicated
  - **Dynamic Size**: Maps can grow in Size as more Elements are added. They are not of a fixed Capacity
  - **Zero Value**: The Zero Value of a Map is _nil_. A _nil_ Map has no Keys and can not have Keys added to it
  - **Retrieving Elements**: When Retrieving an Element from a Map, Developer can also check if the Key was present in the Map using a two-value Assignment
  - **Deleting Elements**: Elements can be removed from a Map using the _delete_ Function

### Strings

- **Definition**: Strings in Go are Sequences of Bytes that are conventionally used to represent Text
- **Immutable**: Once created, the Content of a String can not be changed. Any Modification Operations on a String result in a new String being created
- **UTF-8 Encoding**: Strings in Go are UTF-8 encoded by default. This means they can efficiently represent Unicode Text
- **String Operations**
  - **Concatenation**: Strings can be concatenated using the `+` Operator or the `fmt.Sprintf` Function for more complex Scenarios
  - **Length**: The _len_ Function returns the Number of Bytes in the String, not necessarily the Number of Characters (i.e. _Runes_)
  - **Indexing and Slicing**: Developer can access individual Bytes of a String or slice a String, but this operates on Bytes, not Characters. For Unicode Characters, Developer need to decode the String into Runes
  - **Rune Operations**: To handle multi-byte Characters, Developer can convert a String to a Slice of Runes (i.e. `[]rune`), allowing them to iterate over and manipulate individual Unicode Characters

## Commands

| Command                                                           | Description                                                                        |
| ----------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| go mod init                                                       | Initializes a new _Module_, creating a `go.mod` File                               |
| go run .                                                          | Compiles (in-memory) and executes the _Main Package_ in the current Directory      |
| go build .                                                        | Compiles the current _Package_ into a binary Executable                            |
| GOOS=windows GOARCH=amd64 go build -o build/main.exe .            | Cross-compiles the current _Package_ for Windows, outputting an Executable         |
| go test ./...                                                     | Executes all Tests in the current _Module_ and its _Sub Packages_                  |
| go test -coverprofile=./coverage/cover.out ./...                  | Runs Tests for all _Packages_ and generates a Coverage Report                      |
| go tool cover -html=./coverage/cover.out -o ./coverage/cover.html | Generates an HTML Coverage Report from the `cover.out` File                        |
| go test -bench . ./...                                            | Executes Benchmark Tests for all _Packages_ in the current _Module_                |
| go mod tidy                                                       | Cleans the _Module_ by Adding missing Module Dependencies and Removing unused ones |
