# sampleProgram1

このリポジトリは Go + Ebiten によるブロック崩しゲームのサンプルです。

## GitHub Pages
[公開ページはこちら](https://maniax-jp.github.io/sampleProgram1/)

- WASM版は `docs/` ディレクトリで公開されています。
- 最新のビルド成果物は自動で反映されます。

## ビルド・実行方法
- デスクトップ: `go build -o bin/blockbreaker.exe ./src/main_desktop.go`
- WASM: `GOOS=js GOARCH=wasm go build -o docs/main.wasm ./src/main_wasm.go`

## 開発環境
- Go 1.21 以上
- Ebiten v2
- VSCode タスク/Actions対応
