1BRC adapted for Go (no Java)
=============================

**Credits:** The original project can be found at https://github.com/gunnarmorling/1brc

Status
------
WIP


Get started
-----------

````
make all
./cmd/create_measurements [-file <outputfile>] 1000000000
````

will create a file containing 1 billion records. Allow several minutes and around 12GB of disk space. 

If no `-file` flag is provided, the file is written to `measurements.txt`.   






