package main

import (
  "os"
	"fmt"
  "strings"
	"github.com/chanced/caps"
)

const DEBUG = false

func main() {

  styles := set_styles()
  style_args, text := parse_args(styles)
  convert_case(text, style_args)
}

func log(message string) {
  if DEBUG {
    fmt.Println(message)
  }
}

func parse_args(styles map[string]bool) (map[string]bool, string) {

  var text_args []string
  style_args := make(map[string]bool)

  for _, arg := range os.Args[1:] {

    log(fmt.Sprintf("\t[main]\t found argument: %s", arg))

    if arg == "--help" || arg == "-h" {
      fmt.Println("Help text here")
      os.Exit(0)
    }

    if is_style_arg(arg, styles) {
      style_args[arg] = true
    } else {
      text_args = append(text_args, arg)
    }

  }

  text := strings.Join(text_args, " ")
  return style_args, text
}

func convert_case(text string, style_args map[string]bool) {

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
