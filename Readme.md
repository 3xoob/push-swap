# Push-Swap

**Push-Swap** is a project focused on implementing a non-comparative sorting algorithm using two stacks. The goal is to sort a stack of integers with the minimum number of operations. The project consists of two main programs:

1. **push-swap**: Generates and displays the smallest sequence of instructions needed to sort a given stack of integers.
2. **checker**: Verifies whether a sequence of instructions correctly sorts the stack.

## Objectives

The project aims to help you learn about:

- The use of basic and sorting algorithms.
- The manipulation of stacks.
- Efficient problem-solving and optimization techniques.

## Instructions

### push-swap

The `push-swap` program receives a stack of integers as input and outputs the smallest sequence of instructions to sort the stack in ascending order. The output must be optimized to use the minimum number of operations.

### checker

The `checker` program reads a sequence of instructions and applies them to the stack. It then checks if the stack is sorted and stack `b` is empty. If the stack is sorted, `checker` outputs "OK"; otherwise, it outputs "KO".

### Operations

The following instructions are used to manipulate the stacks:

- `pa`: Push the top element of stack `b` onto stack `a`.
- `pb`: Push the top element of stack `a` onto stack `b`.
- `sa`: Swap the top two elements of stack `a`.
- `sb`: Swap the top two elements of stack `b`.
- `ss`: Execute `sa` and `sb` simultaneously.
- `ra`: Rotate stack `a` upwards.
- `rb`: Rotate stack `b` upwards.
- `rr`: Execute `ra` and `rb` simultaneously.
- `rra`: Reverse rotate stack `a` downwards.
- `rrb`: Reverse rotate stack `b` downwards.
- `rrr`: Execute `rra` and `rrb` simultaneously.

### Error Handling

Both programs will handle errors gracefully:

- If invalid arguments are provided (non-integer values, duplicates, or invalid instructions), an `Error` message will be displayed.
- If no arguments are provided, the programs will output nothing.

## Usage

### Examples

```sh
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa

$ ./checker "3 2 1 0"
sa
rra
pb
KO

$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

### Performance Testing

```sh
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
8

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

## Installation

You have two options for building the executables:

### Option 1: Manual Build

1. Clone the repository:

```sh
git clone https://learn.reboot01.com/git/aabdulhu/push-swap
cd push-swap
```

2. Build the executables:

```sh
go build -o push-swap ./PushSwap
go build -o checker ./Check
```

### Option 2: Build Using a Bash Script

If you prefer an automated process, you can use the provided bash script to build the executables:

1. Clone the repository:

```sh
git clone https://learn.reboot01.com/git/aabdulhu/push-swap
cd push-swap
```

2. Run the bash script:

```sh
bash build.sh
```

This script will compile the `push-swap` and `checker` programs for you.

## Requirements

- **Go**: The project must be implemented using Go. Only standard Go packages are allowed.
- **Good Coding Practices**: The code should follow best practices, including proper error handling and clear, maintainable code.

## Authors
- Ali (aabdulhu)
- Sayed Ali (sayedalawi)
- Yusuf (yalsayya)
