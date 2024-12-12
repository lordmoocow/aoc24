# Advent of Code 2024
Another attempt this year using Go. Not particularly proud of anything in here.

## Usage
Using the binary, you just need to provide your puzzle input and specify which day and part to run:
```sh
aoc run <day> [--part n] <input_file>
```

e.g. `aoc run 1 --part 1 puzzle1.txt` will run day one, part one using the contents of the `puzzle1.txt` file as input.

### Docker
Docker image has been created mostly as an excercise rather than any real utility - but here it is:

```sh
docker run -v <input_file>:/input --rm lordmoocow/aoc run <day> [--part n] /input
```

`-v` option is needed in order to mount `input_file` to use as puzzle input; absolute path is required. The `--rm` option is recommended to automatically clean up the container after execution. 