# pwdc - Coilorized and Formatted `pwd` Command

`pwdc` は、現在のディレクトリのパスをカスタマイズされた形式で表示するコマンドです。標準の`pwd`コマンドの動作を拡張子、色分けや階段状の表示、スペース挿入などのオプションを提供します。

## インストール方法

### Go でのインストール

```bash
go install github.com/Morishita-mm/pwdc@latest
```

インストール後、`pwdc`コマンドが使用できるようになります。

### macOS へのバイナリインストール

1. 上記の Go インストールを行った後、`pwdc`バイナリが`$GOPATH/bin/`にインストールされます。
2. もしパスが通っていない場合は、`$GOPATH/bin/`にパスを通すか、シンボリックリンクをサック制してください。

### 使用方法

```bash
pwdc [オプション]
```

### オプション

以下のオプションを使用できます：

- `L`,`l`
  - **説明**：論理パスを表示します（デフォルト）。
  - **例**：`/var/www/project`
- `P`,`p`
  - **説明**：物理パスを表示します。
  - **例**：`/var/www/project`（シンボリックリンクが解決されたパス
- `-C`,`-c`
  - **説明**：各ディレクトリレベルを色分けして表示します。
  - **例**：`/Users/example/Documents/Projects/Go`が色分けされて表示されます。
- `-S`, `-s`
  - **説明**: 各ディレクトリ間にスペースを挿入して表示します。
  - **例**: /Users/example/Documents/Projects/Go → / Users / example / Documents / Projects / Go
- -T, -t
  - **説明**: 現在のディレクトリを階段状に表示します。ディレクトリごとにインデントされます。
  - **例**：
    ```bash
    /
    Users
        example
        Documents
            Projects
            Go
    ```

### 複数オプションの使用

オプションは複数同時に指定できます。例えば、`-T -C -S`を指定すると、階段状に出力され、さらにディレクトリ内にスペースが挿入されます。

### 動作環境

- Go 1.18 以降
- macOS, Linux, Windows に対応
