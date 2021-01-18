## Laboratory cli
A cli to obtain a reaction and obtain smallest lenght removing a bridle element from a polymer. ::nerd_face:

## Subjet

A polymer is made up of the element by letters of the alphabet, lower case (negative polarization) or upper case (positive polarization))

When you start a chemical reaction on a polymer, the elements with reverse polarity cancel each other out.

## Motivation

Confirm my technical skills as a coder and integrate this awesome team!

## Tech/framework used

<b>Built with</b>
- [Cobra](https://github.com/spf13/cobra)

## How to use?
```
$> make help
Usage: 

  build   build the application
  run     runs go run main.go
  clean   cleans the binary
  test    runs go test with default values
  bench   runs go bench with default values
  setup   setup go modules
  help    Prints this help message
```

```
Currently, just write your polymere.

Usage:
  laboratory [command]

Available Commands:
  help        Help about any command
  minlen      minimal length reaction
  react       make polymere reaction

Flags:
  -h, --help   help for laboratory

```
```
./laboratory react aBbAScaleway
polymere: aBbAScaleway, reaction: Scaleway
```
```
./laboratory minlen dabAcCaCBAcCcaDA
polymere: dabAcCaCBAcCcaDA, minimal length: 4
```