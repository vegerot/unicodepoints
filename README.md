# convert unicode code point to UTF-8 character

## Usage

```sh
❯ go run cmd/cli/main.go --string 1F60A
😊
```

```sh
❯ go run cmd/cli/main.go --bytes 1F973 # 🥳
[240 159 165 179]
```

## Building
```sh
❯ go build -o unicodepoint cmd/cli/main.go

❯ ./unicodepoint --string 1F347
🍇
```

## Installing

```sh
# on POSIX
❯ go build -o "${GOBIN:-$HOME/go/bin}/unicodepoint" cmd/cli/main.go

❯ unicodepoint --string 1F495
💕
```

