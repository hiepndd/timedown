# Countdown

## Usage

Specify duration in go format `2h3m4s`.

```bash
run `go build .`
and `./timedown 10s`
```

Add command with `&&` to run after countdown.

```bash
./timedown 1m30s && say "Hello, world"
```

## Key binding

- `x` or `X`: To pause the countdown.
- `c` or `C`: To resume the countdown.
- `Esc` or `Ctrl+C`: To stop the countdown without running next command.
