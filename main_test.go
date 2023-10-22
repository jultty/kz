package main

import (
  "os"
  "io"
  "bytes"
  "testing"
)

func TestCamelCase(t *testing.T) {
  sentence := "this is a no-stress sentence for testing"
  expected := "thisIsANoStressSentenceForTesting\n"

  args := make(map[string]bool)
  args["--camel"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestSnakeCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "this_is_a_no_stress_sentence\n"

  args := make(map[string]bool)
  args["--snake"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestUpperSnakeCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "THIS_IS_A_NO_STRESS_SENTENCE\n"

  args := make(map[string]bool)
  args["--Snake"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestToKebabCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "this-is-a-no-stress-sentence\n"

  args := make(map[string]bool)
  args["--kebab"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestToUpperKebabCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "THIS-IS-A-NO-STRESS-SENTENCE\n"

  args := make(map[string]bool)
  args["--Kebab"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestToPascalCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "ThisIsANoStressSentence\n"

  args := make(map[string]bool)
  args["--pascal"] = true
  
  runCaseTest(t, sentence, expected, args)
}

func TestToTitleCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "This Is A No Stress Sentence\n"

  args := make(map[string]bool)
  args["--title"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestToUpperCase(t *testing.T) {
  sentence := "this is a no-stress sentence"
  expected := "THIS IS A NO-STRESS SENTENCE\n"

  args := make(map[string]bool)
  args["--upper"] = true

  runCaseTest(t, sentence, expected, args)
}

func TestToLowerCase(t *testing.T) {
  sentence := "THIS is a NO-STRESS SENTENCE"
  expected := "this is a no-stress sentence\n"

  args := make(map[string]bool)
  args["--lower"] = true

  runCaseTest(t, sentence, expected, args)
}

func runCaseTest(t *testing.T, sentence string, expected string, args map[string]bool) {

  old := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w

  convert(sentence, args)

  w.Close()
  os.Stdout = old

  var buf bytes.Buffer
  io.Copy(&buf, r)
  output := buf.String()

  if output != expected {
    t.Errorf("%q = %q, want %q", sentence, output, expected)
  }
}
