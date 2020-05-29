# go-Mandelbrot

The Mandelbrot set programmed in Go

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.
Note: These instructions have been tested on Ubuntu 20.04.

### Prerequisites

You will have to install these things first to get the project to work.

#### Go

Install [Go](https://golang.org/doc/install) on you local machine.

#### C compiler

A C compiler is required since ebiten also uses C.

```
sudo apt install gcc
```

#### Dependencies

```
sudo apt install libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
```

#### Go Packages

Finally let's download the ebiten package

```
go get -u github.com/hajimehoshi/ebiten
```

### Install

To download the project simply just type:

```
go get -u github.com/eliashauksson/go-Mandelbrot
```
```
cd go-Mandelbrot && go build main.go
```

### Built With

* [Ebiten](https://github.com/hajimehoshi/ebiten) - The 2D game library used
