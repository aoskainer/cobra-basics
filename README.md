# cobra-basics

[![Run Tests and Generate Coverage Report](https://github.com/aoskainer/cobra-basics/actions/workflows/test.yml/badge.svg)](https://github.com/aoskainer/cobra-basics/actions/workflows/test.yml)

[![codecov](https://codecov.io/gh/aoskainer/cobra-basics/graph/badge.svg?token=W0HSTF9ZWF)](https://codecov.io/gh/aoskainer/cobra-basics)

<br>

## 最初にやったこと

### cobra-cliのインストール

```sh
$ go install github.com/spf13/cobra-cli@latest
```

### 初期化

```sh
$ go mod init github.com/aoskainer/cobra-basics
$ cobra-cli init
```

### サブコマンドを追加するときにやったこと

```sh
$ cobra-cli add guessgame
```
