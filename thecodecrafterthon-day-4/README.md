## What it does

Reads a text file, cleans it up, and writes a new file.

* Replace TODO with ACTION        
* Replace CLASSIFIED: with [REDACTED]:
* Fixes ALL CAPS lines
* Removes empty lines

## How to run

go run main.go input.txt output.txt

## Example

Input (`input.txt`):
TODO: check perimeter
CLASSIFIED: agent report
HELLO WORLD

Output (output.txt):
SENTINEL FIELD REPORT — PROCESSED
  1: Action: Check Perimeter
  2: [REDACTED]: Agent Report
  3: Hello World

## Author
MAKOLO ELEOJO GIFT