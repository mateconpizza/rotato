<div>
<img align="right" width="320px" src="assets/rotato-logo.png" alt="Rotato Logo">
<h1><b><span style="font-size: 1.2em">🌀</span> Rotato</b></h1>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mateconpizza/rotato)
![Linux](https://img.shields.io/badge/-Linux-grey?logo=linux)
[![Go Reference](https://pkg.go.dev/badge/github.com/mateconpizza/rotato.svg)](https://pkg.go.dev/github.com/mateconpizza/rotato)

Lightweight and highly customizable spinner library for Go, designed to enhance your command-line applications with visually appealing progress indicators. With **+70** spinners to choose from, you can easily integrate dynamic loading animations into your projects.

<br>
</div>

<div align="center">

![Demo](assets/rotato-demo.gif)

</div>

## More demos

You can check out the spinners with the following commands:

- Simple demo

```sh
go run github.com/mateconpizza/rotato/example@latest -demo
```

- All spinners

```sh
go run github.com/mateconpizza/rotato/example@latest -all
```

- List by groups

```sh
go run github.com/mateconpizza/rotato/example@latest -list
```

- Show by group

```sh
go run github.com/mateconpizza/rotato/example@latest -group <name>
```

- Show by name

```sh
go run github.com/mateconpizza/rotato/example@latest -show <name>
```

## Installation

```sh
go get github.com/mateconpizza/rotato@latest
```

## Example

- Simple example:

```go
r := rotato.New(
    rotato.WithPrefix("Repo"),
    rotato.WithSpinnerColor(rotato.FgBrightGreen),
    rotato.WithDoneColorMesg(rotato.FgBrightGreen.With(rotato.StyleItalic)),
    rotato.WithFailColorMesg(rotato.FgRed.With(rotato.StyleBlink)),
)
r.Start()
// do some stuff
repo := git.New("https://github.com/mateconpizza/rotato")
r.UpdateMesg("Syncing Repo...")
if err := repo.Sync(); err != nil {
    r.Fail("Sync Failed!")
    return err
}
r.Done("Sync Completed!")
```

> [!NOTE]
> To see more, go to [example](https://github.com/mateconpizza/rotato/blob/master/example/main.go)

## Credits

This package uses `symbols/spinners` from this libraries, and of course ideas!

Thanks to:

- [@briandowns](https://github.com/briandowns/spinner)
- [@theckman](https://github.com/theckman/yacspin)

## Others

Visible elements

- Terminal: [st-terminal](https://st.suckless.org/)
- Font: [maple-font](https://github.com/subframe7536/maple-font)
