## Beat

Beat is a command line tool built with golang to do a fare estimation for a big dataset of rides.

[![Build Status](https://travis-ci.com/Clivern/Beat.svg?branch=master)](https://travis-ci.com/Clivern/Beat)

## Documentation


### Solution

A final working binaries are included so you can use the command line tool directly with the `paths.csv` file. Here is the steps

```bash
$ git clone https://Clivern@bitbucket.org/Clivern/beat.git
$ cd beat/solution

# Command line tool help
$ ./beat_linux_amd64
$ ./beat_darwin_amd64

# The command to calculate the rides fare from a CSV file and output the result to another CSV file
$ ./beat_linux_amd64 calculate -c config.yml -i paths.csv -o output.csv
$ ./beat_darwin_amd64 calculate -c config.yml -i paths.csv -o output.csv
$ cat output.csv
```


### Development

After installing golang, You need to fork the repository like the following:

```bash
$ cd $GOPATH
$ mkdir -p src/bitbucket.org/clivern
$ git clone https://Clivern@bitbucket.org/Clivern/beat.git
$ cd beat
```

Here is the steps to build new binaries or run interactively

```
# Install & build with goreleaser
$ brew install goreleaser
$ goreleaser --snapshot --skip-publish --rm-dist --parallelism 1


# Build manually
$ make build
$ ./releases/beat_linux_amd64
$ ./releases/beat_darwin_amd64


# Run interactively
$ go run beat.go
```

In case you did some code changes, it is pretty easy to run the sanity check locally:

```bash
# Install linters locally
$ make install_revive
$ make install_golint

# Reset any changes made to go.mod since these are only for test env
$ git reset --hard

# Now you can run sanity checks
$ make ci
```

Ideally this should run each time we push changes to the git server. `.github/workflows` included for that purpose.
Also there is another github workflow included to create new binaries each time we create a new tag using goreleaser. Seems like bitbucket pipelines not for free anymore.


### Under the hood

I use a data pipeline to process the input CSV file, It works like the following:

- It will read the CSV file line by line `bufio` and send a the whole ride lines to a golang channel.
- Another function will take that channel as input and it will launch a concurrent goroutines to do the fare calculation. This function waits till all goroutines finish. once each goroutine finishes, it sends the result (rideid, fare) to another output channel.
- Finally there is a function listening to the output channel of the second function and store the data to output file (line by line too).


The command line tool is organized as packages:

- `cmd`: Holding all commands.
- `core/model`: Holding the entites we have like `Ride`, `Coordinate` ... etc
- `core/modules`: Holding the domain logic.
- `core/utils`: Small functions to do the data conversion and some for file manipulation.
- `pkg`: Contains a small reusable functions for unit and functional test cases.

And Some other folders and files:

- `testdata`: Files uses as test data.
- `cache`: Used during build since some tests will create a test files.
- `.github`: Github workflow files. If this solution hosted on github, it will run build on push and create a release once we create a new tag. Bitbucket offer something but it is unfortunately not for free.
- `config.toml`: Used to configure golang revive linter. It is a nice linter https://github.com/mgechev/revive


### Improvements

This tools can be improved by:

- More test cases and debug logs.
- It can be a REST service that calculate the fare on a real time or at the end of the ride. The calculations should run on Async way (by goroutines or separate processes/workers)
- If the tool meant to work with a big dataset of rides that already happened in the past, we can improve the way it works. at least support multiple data sources as input not just CSV.
- The tool assume that the data are in a good shape, like coordinates are always sorted based on timestamp.




