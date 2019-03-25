# go-poker-hand-identifier


This is a command-line application written in Go that when given a txt file containing poker hands will try to name each hand.

## Runtime
```
go version go1.9.7 linux/amd64
```

## Usage
```
go run main.go <txt_file>
```

## Example

```
Example input:
3H JS 3C 7C 5D
JH 2C JD 2H 4C
9H 9D 3S 9S 9C
9C 3H 9S 9H 3S
```

```
Example output:
3H JS 3C 7C 5D => One pair
JH 2C JD 2H 4C => Two pair
9H 9D 3S 9S 9C => Four of a kind
9C 3H 9S 9H 3S => Full house
```

Some test files are supplied in the `tests` folder including some invalid ones to test the boundaries of error handling.