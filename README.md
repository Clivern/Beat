## Beat

Beat is a command line tool built with golang to do a fare estimation for a big dataset of rides.

[![Build Status](https://travis-ci.com/Clivern/Beat.svg?branch=master)](https://travis-ci.com/bitbucket/Clivern/beat)

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
$ cd src/bitbucket.org/clivern
$ git clone https://Clivern@bitbucket.org/Clivern/beat.git
$ cd beat
```

Here is the steps to build new binaries or run interactively

```
# Install goreleaser on Mac. Or download a binary from here https://github.com/goreleaser/goreleaser/releases
$ brew install goreleaser

$ goreleaser --snapshot --skip-publish --rm-dist --parallelism 1


# Build manually
$ make build
$ ./releases/beat_linux_amd64
$ ./releases/beat_darwin_amd64


# Run interactively
$ go run beat.go
```

Anytime you want to show the tool verbose logs, add `-v` to the command as explained in the tool help

```bash
$ ./releases/beat_darwin_amd64

A Command Line Tool to do Fare Estimation for a big set of Rides.

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://bitbucket.org/clivern/beat>

Usage:
  beat [command]

Available Commands:
  calculate   Calculate fare for a big set of rides
  help        Help about any command
  license     Print the license
  version     Print the version number

Flags:
  -h, --help      help for beat
  -v, --verbose   verbose output

Use "beat [command] --help" for more information about a command.
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
Also there is another github workflow included to create new binaries each time we create a new tag using goreleaser. Right now i use travis for build.


### Under the hood

I use a data pipeline to process the input CSV file, It works like the following:

- It will read the CSV file line by line and send the whole ride lines to a golang channel.

- Another function will take that channel as input and it will launch a concurrent goroutines (configurable and can change) to do the fare calculation. This function waits till all goroutines finish. once each goroutine finishes, it sends the result (rideid, fare) to another output channel.

- Finally there is a function listening to the output channel of the second function and store the data to output file (line by line too) in CSV format.

- It is worth mentioning that the number of goroutines used for processing can be increased or decreased from the config file, property `app.max_goroutines`. this can speed things if the dataset is huge.

The command line tool is organized as packages:

- `cmd`: Holding all commands.
- `core/model`: Holding the entites we have like `Ride`, `Coordinate` ... etc
- `core/modules`: Holding the domain logic like loading data from CSV, fare calculations, the data pipeline.
- `core/utils`: Small functions to do the data conversion and some for file manipulation.
- `pkg`: Contains a small reusable functions for unit and functional test cases purposes.

And Some other folders and files:

- `testdata`: Files uses as test data.
- `cache`: Used during build since some tests will create a test files.
- `.github`: Github workflow files. If this solution hosted on github, it will run build on push and create a release once we create a new tag. Bitbucket offer something but it is unfortunately not for free. Here i am using travis.com.
- `config.toml`: Used to configure golang revive linter. It is a nice linter https://github.com/mgechev/revive I use together with `golint`.



### Improvements

This tools can be improved by:

- More test cases and debug logs.

- Pricing & currency should be configured based on the ride country or city.

- It can be a REST service that store coordinates, calculate the fare on real time or at the end of the ride. The calculations should run on Async way (by goroutines or separate processes/workers).

- If the tool meant to work with a big dataset of rides that already happened in the past, we can improve the way it works. at least support multiple data sources as input not just CSV.
