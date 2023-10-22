package main

import (
  "os"
  "io"
	"fmt"
  "strings"
	"github.com/chanced/caps"
)

const DEBUG = false

func main() {

  styles := set_styles()
  text, args := parse_options(styles)
  convert(text, args)

}

func parse_options(styles map[string]bool) (string, map[string]bool) {

  var text_args []string
  style_args := make(map[string]bool)

  for _, arg := range os.Args[1:] {

    log(fmt.Sprintf("\t[main]\t found argument: %s", arg))

    if arg == "--help" || arg == "-h" {
      print_help_text()
      os.Exit(0)
    }

    if is_style_arg(arg, styles) {
      style_args[arg] = true
    } else {
      text_args = append(text_args, arg)
    }

  }

  var text string

  stat, _ := os.Stdin.Stat()
  if (stat.Mode() & os.ModeCharDevice) == 0 {
    bytes, _ := io.ReadAll(os.Stdin)
    text = string(bytes)
  } else {
    text = strings.Join(text_args, " ")
  }

  return text, style_args
}

func convert(text string, style_args map[string]bool) {

  for i, _ := range style_args {

    if (i == "--camel" || i == "-c") {
      fmt.Println(caps.ToLowerCamel(text))
    }

    if (i == "--kebab" || i == "-k") {
      fmt.Println(caps.ToKebab(text))
    }

    if (i == "--Kebab" || i == "-K") {
      fmt.Println(caps.ToScreamingKebab(text))
    }

    if (i == "--lower" || i == "-l") {
      fmt.Println(caps.ToLower(text))
    }

    if (i == "--pascal" || i == "-p") {
      fmt.Println(caps.ToCamel(text))
    }

    if (i == "--snake" || i == "-s") {
      fmt.Println(caps.ToSnake(text))
    }

    if (i == "--Snake" || i == "-S") {
      fmt.Println(caps.ToScreamingSnake(text))
    }

    if (i == "--title" || i == "-t") {
      fmt.Println(caps.ToTitle(text))
    }

    if (i == "--upper" || i == "-u") {
      fmt.Println(caps.ToUpper(text))
    }
  }
}

func set_styles() map[string]bool {

  style_list := [...]string{
    "camel", 
    "kebab", 
    "Kebab",
    "lower", 
    "pascal", 
    "snake", 
    "Snake", 
    "title", 
    "upper", 
  }
  
  styles := make(map[string]bool)

  for _, v := range style_list {
    styles[v] = true
    styles[v[0:1]] = true
  }

  return styles
}

func is_style_arg(input string, styles map[string]bool) bool {
  var first_two string
  var first string
  var candidate string

  log(fmt.Sprintf("[is_style_arg]\t checking if argument is a style: %s", input))

  if len(input) > 2 {
    first_two = input[:2]
  } else if len(input) == 2 {
    first = input[0:1]
  } else { 
    log("[is_style_arg]\t returning false: too short for an argument")
    return false 
  }

  if first_two == "--" {
    candidate = input[2:]
  } else if first == "-" {
    candidate = input[1:2]
  } else {
    log("[is_style_arg]\t returning false: starts with neither one nor two dashes")
    return false
  }

  if DEBUG && styles[candidate] {
    log(fmt.Sprintf("[is_style_arg]\t returning true for arg candidate %s", candidate))
  } else {
    log(fmt.Sprintf("[is_style_arg]\t returning false for arg candidate %s", candidate))
  }

  return styles[candidate]
}

func print_help_text() {
  fmt.Println(`
    Usage: kz [args] [text]

      Arguments:

        --lower,  -l   outputs a lowercase sentence
        --upper,  -u   OUTPUTS AN UPPERCASE SENTENCE
        --title,  -t   Outputs A Title Case Sentence
        --camel,  -c   outputsACamelCaseSentence
        --pascal, -p   OutputsAPascalCaseSentence
        --snake,  -s   outputs_a_snake_case_sentence
        --Snake,  -S   OUTPUTS_AN_UPPERCASE_SNAKE_CASE_SENTENCE
        --kebab,  -k   outputs-a-kebab-case-sentence
        --Kebab,  -K   OUTPUTS-AN-UPPERCASE-KEBAB-CASE-SENTENCE

    Any argument not matching the arguments above will be case-converted.

    If your input may contain words starting with dashes that you
    do not want interpreted as arguments, make sure to quote them:

      "the -u in this sentece will be interpreted literally"

    If you pass multiple case style arguments, they will all be printed:

      kz "a linha de fronteira se rompeu" -u --title
      A LINHA DE FRONTEIRA SE ROMPEU
      A Linha De Fronteira Se Rompeu

    You can also pipe text into kz in order to convert it:

      echo "a sentence to convert to camelCase" | kz -c

    If you do so, text passed as an argument will be ignored.

    `)
}

func log(message string) {
  if DEBUG {
    fmt.Println(message)
  }
}

