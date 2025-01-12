# llmpatch

> Tool for modifying individual files using an LLM

### Install

``` sh
go install github.com/icholy/llmpatch/cmd/llmpatch@latest
```

### Usage:

``` sh
llmpatch -f main.go llm "use slog instead of log"
```