1BRC adapted for Go (no Java)
=============================

**Credits:** The original project can be found at https://github.com/gunnarmorling/1brc

Status
------
WIP


Project layout
--------------
`/cmd/`  --- Go programs

`/cmd/baseline` --- Basic implementation of the 1BRC

`/cmd/create_measurements/` --- Tool for creating input files 

`/datasets/` --- some smaller input files ready to use, with corresponding result files

`/Makefile` -- Format, test, build  


Get set-up
----------

Install Go: https://go.dev/dl/ 

Build the tool for making input files:  
````
make all
./cmd/create-measurements [-file <outputfile>] 1000000000
````

This will create a file containing 1 billion records as per the 1BRC. Allow several minutes and around 12GB of disk space. 

If no `-file` flag is provided, the file is written to `measurements.txt`.

Some smaller files are already provided in the `/datasets/` directory.

Get coding!
-----------
Take a look at my naive 1BRC solution implemented in `/cmd/baseline`. Use the shell script to build and run it.

As a guide, this implementation takes around 18s on my Macbook Pro (M1 Pro)

Create a new directory under `/cmd/` for your own implementation. 

See here for ideas and direction: https://mrkaran.dev/posts/1brc/

**Remember:** Stick to the Standard Library! https://pkg.go.dev/std

Where is your program spending its time? https://go-language.org/go-docs/runtime/pprof/

_Good luck!_







