# mkr

**mkr** (marker) is marking a text tool.
It marks lines and a text matched pattern.
It is possible to customize marking style, charactor and bacground color and text format.

This tool helps **easy to read a text** by coloring and formatting.

## Demo

### Mark lines containing a text matched pattern

![line command](resources/demo001.gif)

### Mark a text matched pattern

![line command](resources/demo002.gif)

### Mark specified lines

![line command](resources/demo003.gif)

## Usage

#### Simple Usage

Mark lines containing a text matched given pattern in `file.txt`.

```bash
# Mark lines containing 'text'
$ cat file.txt | mkr mark line text
```

Mark a text matched given pattern in `file.txt`.

```bash
# Mark 'text'
$ cat file.txt | mkr mark word text
```

Mark specified lines in `file.txt`.

```bash
# Mark lines 5 to 9
$ cat file.txt | mkr mark number 5~9
```

#### Set marking style

```bash
# Make a marked text bold and italic
$ cat file.txt | mkr mark line --bold --italic text
# Color a marked text
$ cat file.txt | mkr mark word --charactor red text
```

#### Using Config File

Create new config file (default: `$HOME/.mkr.yaml`).

```bash
# Create new config file ($HOME/.mkr.yaml)
$ mkr init
# Create new config file (/path/to/mkr/.config.yaml)
$ mkr init /path/to/mkr/.config.yaml
```

Create new config file and set default marking style

```bash
# In this sample, set charactor color (green) and text formats (bold, italic)
$ mkr init --charactor green --bold --italic
```

Config file is YAML file. It is possible to customize marking a text style.

```yaml
# Config file
color:
    background: none
    charactor: green
format:
    bold: false
    hide: false
    italic: false
    strikethrough: false
    underline: false
```

## Flags

### Color Control

| flag | description |
|:-|:-|
| --background, -b | color background of a marked text |
| --charactor, -c | color charactor of a marked text |

Settable colors are

- none
- black
- blue
- cyan
- green
- magenta
- red
- yellow
- 0 ~ 255

### Format Control

| flag | description |
|:-|:-|
| --bold | make a marked text bold |
| --italic | make a marked text italic |
| --hide | hide a marked text |
| --underline | underline a marked text |
| --strikethrough | strikethrough a marked text |

### Context Control

This options can be set to `line` command only.

| option | description |
|:-|:-|
| --after | mark n lines after matching lines |
| --before | mark n lines before matching lines |

## Install

Prerequisite Tools

- Git
- Go (at least Go 1.11)

```bash
$ git clone https://github.com/x-color/mkr.git
$ cd mkr
$ go install
```

## License

Apache License 2.0
