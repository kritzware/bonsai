# bonsai
Key value store in Golang

### Setup
```bash
$ go build src/bonsai.go
```

How to run:
```bash
$ bonsai
```

### Commands

#### `store` 
Store a value in memory
```bash
bonsai> store [key] [value]
```

#### `get` 
Retrieve memory address and value by key
```bash
bonsai> get [key]
```

### `exit` 
Close bonsai instance
```bash
bonsai> exit
```
