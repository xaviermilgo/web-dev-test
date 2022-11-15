# Web development challenge
In the following challenge, you will be required to test and implement a processor that aggregates data efficiently.

## Introduction

The current implementation generates millions of trips across 100 hotels and 10000 drivers. You are required to calculate the top-rated driver and top-rated hotel. You are also required to test for this.

The test requires you to implement `processors/interface.go` and `processors/processor.go`.

You are also required to test the functioning of the processor by implementing `main_test.go`


### Requirements

1. Running main.go prints out the top rated driver and hotel. Calculate the average rating by referring to Rating fields of the Trip struct.
2. Running main_test.go (`go test`) validates that the processor calculates the above correctly.
3. Memory usage should be below 128mb. Extra points awarded for less usage.
4. Modifying other parts of the code is allowed but ensure the basics of the challenge remain the same.

## Submission or Questions

After completing the implementation, push it to github and share the link via email to `xavier@novek.io`

For any questions and clarifications, feel free to contact me via email.
