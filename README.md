# PMapper-go

## Description

PMapper-go is a Go implementation of the [PMapper](https://github.com/nccgroup/PMapper) Python project

As of 11/29/2023, it is currently at parity with the original project's main branch.

## Installation

### Requirements
PMapper-go was built with go 1.24.4, and the project includes a [Hermit](https://github.com/cashapp/hermit) `bin` folder to 
mirror my dev environment if you want to tinker and run locally.

After cloning, you should be able to fetch all of the necessary Go dependencies with
```zsh
go mod tidy
```

### Installing from `Releases`
- Download the latest release from the [Releases](https://github.com/sheikhrachel/PMapper-go/releases) page
- [Docker](https://docs.docker.com/get-docker/)

## Usage

### Quick examples
_referenced from the [PMapper](https://github.com/nccgroup/PMapper) project to help bridge the knowledge gap_

```
# Create a graph for the account, accessed through AWS CLI profile "skywalker"
pmapper --profile skywalker graph create
# [... graph-creation output goes here ...]

# Run a query to see who can make IAM Users
$ pmapper --profile skywalker query 'who can do iam:CreateUser'
# [... query output goes here ...]

# Run a query to see who can launch a big expensive EC2 instance, aside from "admin" users
$ pmapper --account 000000000000 argquery -s --action 'ec2:RunInstances' --condition 'ec2:InstanceType=c6gd.16xlarge'
# [... query output goes here ...]

# Run the privilege escalation preset query, skip reporting current "admin" users
$ pmapper --account 000000000000 query -s 'preset privesc *'
# [... privesc report goes here ...]

# Create an SVG representation of the admins/privescs/inter-principal access
$ pmapper --account 000000000000 visualize --filetype svg
# [... information output goes here, file created ...]
```

### Contributions

### License
```text
Copyright 2023 Rachel Sheikh

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```