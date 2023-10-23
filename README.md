# kz

A case converter for your terminal.

## Installation

You can grab a pre-compiled binary from the [releases](https://github.com/jultty/kz/releases) page.

To build from source:

```sh
go install https://github.com/jultty/kz@latest
```

This requires that you have Go [properly installed](https://go.dev/doc/install).

## Usage

You can invoke `kz` by piping something to it and passing arguments to specify the desired output case styles or by passing the text to be converted alongside the arguments.

So the following:

```sh
$ kz "this will be converted to camelCase" -c
thisWillBeConvertedToCamelCase
```

Is equivalent to:

```sh
$ echo "this will be converted to camelCase" | kz -c
thisWillBeConvertedToCamelCase
```

If you pipe anything to `kz`, text passed alongside arguments will be ignored.

Currently the following options are supported:

```sh
  --lower,  -l   outputs a lowercase sentence
  --upper,  -u   OUTPUTS AN UPPERCASE SENTENCE
  --title,  -t   Outputs A Title Case Sentence
  --camel,  -c   outputsACamelCaseSentence
  --pascal, -p   OutputsAPascalCaseSentence
  --snake,  -s   outputs_a_snake_case_sentence
  --Snake,  -S   OUTPUTS_AN_UPPERCASE_SNAKE_CASE_SENTENCE
  --kebab,  -k   outputs-a-kebab-case-sentence
  --Kebab,  -K   OUTPUTS-AN-UPPERCASE-KEBAB-CASE-SENTENCE
```

The order of the options does _not_ matter. Anything not matching the arguments above will be case-converted, while anything matching them will be stripped from the target text.

If your input may contain words starting with dashes that you do not want interpreted as arguments, make sure to quote them:

```sh
$ kz -u "the first -u in this line will be interpreted literally"
THE FIRST -U IN THIS LINE WILL BE INTERPRETED LITERALLY
```

Note that if you pass multiple case style arguments, all corresponding conversions will be printed:

```sh
$ kz "a linha de fronteira se rompeu" -u --title
A LINHA DE FRONTEIRA SE ROMPEU
A Linha De Fronteira Se Rompeu
```

  
