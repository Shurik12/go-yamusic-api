# yamusic api

## Description

Unofficial Go client library for [Yandex.Music API](https://music.yandex.ru).

Golang fork of [Node.js library](https://github.com/itsmepetrov/yandex-music-api).
Client style based on [google/go-github](https://github.com/google/go-github).

## Usage

```go
import "github.com/Shurik12/go-yamusic-api/yamusic"
```

Construct a new Yandex.Music client, then use the various services on the client to access different parts of the Yandex.Music API.

## 👷 Build

Build package

```bash
make build
```

## 🧪 Testing

Run unit tests:

```bash
make test
```

Run integration tests:

Note that you should set `YANDEX_USER_ID` and `YANDEX_ACCESS_TOKEN` environment variables.

```bash
make test_integration
```

## 🖍 Lint

Run linters

```bash
make lint
```
