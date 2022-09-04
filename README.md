# ChibiGB <!-- omit in toc -->

[![Godoc Reference](https://pkg.go.dev/badge/github.com/kaishuu0123/chibigb)](https://pkg.go.dev/github.com/kaishuu0123/chibigb)
[![GitHub Release](https://img.shields.io/github/v/release/kaishuu0123/chibigb)](https://github.com/kaishuu0123/chibigb/releases)
[![Github Actions Release Workflow](https://github.com/kaishuu0123/chibigb/actions/workflows/release.yml/badge.svg)](https://github.com/kaishuu0123/chibigb/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/kaishuu0123/chibigb)](https://goreportcard.com/report/kaishuu0123/chibigb)

ChibiGB is GameBoy emulator written by Go. This is my favorite hobby project!

Based on [Gearboy](https://github.com/drhelius/Gearboy).

- [Spec](#spec)
- [Key binding](#key-binding)
- [Documents](#documents)
- [License](#license)

## Spec

- [X] CPU
- [X] PPU
- [X] APU
- [X] Controller
- [ ] Cartridge
  - [X] No MBC
  - [X] MBC1
  - [X] MBC2
  - [ ] MBC3
  - [ ] MBC5
  - [ ] MBC6
  - [ ] MBC7
  - [ ] MMM01
  - [ ] M161
  - [ ] HuC1

## Key binding

Player 1

|GB|Key|
|---|---|
| UP, DOWN, LEFT, RIGHT | Arrow Keys |
| Start | Enter |
| Select | Right Shift |
| A | Z |
| B | X |

## Documents

- [Projects | gbdev](https://gbdev.io/)
  - [Foreword - Pan Docs](https://gbdev.io/pandocs/)
- [pokemium/gb-docs-ja: WIP: GameBoyの日本語リファレンスです](https://github.com/pokemium/gb-docs-ja)

## License

- [drhelius/Gearboy](https://github.com/drhelius/Gearboy)
  - [GPL-3.0 license](https://github.com/drhelius/Gearboy/blob/master/LICENSE)
- [itouhiro/PixelMplus](https://github.com/itouhiro/PixelMplus)
  - [M+ FONT LICENSE](https://github.com/itouhiro/PixelMplus/blob/master/misc/mplus_bitmap_fonts/LICENSE_E)
