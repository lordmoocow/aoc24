# Advent of Code 2024
Another attempt this year using Go. Not particularly proud of anything in here.

## Usage
Using the binary, you just need to specify the day and part to run:
```sh
aoc run <day> <part>
```

e.g. `aoc run 1 1` will run day one, part one.

### Docker
Docker image has been created mostly as an excercise rather than any real utility - but here it is:

```sh
docker run --rm lordmoocow/aoc <day> <part>
```

The `--rm` option is recommended to automatically clean up the container after execution.