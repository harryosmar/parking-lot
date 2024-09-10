# Usage

```shell
git clone git@github.com:harryosmar/parking-lot.git

cd parkinglot

# run using go
go run main.go input example_input.txt

# this command will create parking-lot binary file in to go bin dir, $(go env GOPATH)/bin
# export PATH="$PATH:$(go env GOPATH)/bin"
go install && parking-lot

# run using build : mac m1 
./build/m1/parkinglot input example_input.txt

# run using build : mac intel 
./build/intel/mac/parkinglot input example_input.txt

# run using build : windows 64bit 
./build/intel/windows-64bit/parkinglot input example_input.txt

# run using build : windows 32bit 
./build/intel/windows-32bit/parkinglot input example_input.txt
```

## ENVIRONMENT VARIABLES
[config/parkinglot.config.go](https://github.com/harryosmar/parking-lot/blob/master/config/parkinglot.config.go)
