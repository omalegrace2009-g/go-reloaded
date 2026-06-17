# Go Reloaded

## Description

Go Reloaded is a text processing and auto-correction tool written in Go. The program reads text from an input file, applies a series of transformations and formatting rules, and writes the corrected result to an output file.

The tool supports:

* Converting hexadecimal numbers to decimal using `(hex)`
* Converting binary numbers to decimal using `(bin)`
* Converting words to uppercase using `(up)`
* Converting words to lowercase using `(low)`
* Capitalizing words using `(cap)`
* Applying transformations to multiple previous words using `(up, n)`, `(low, n)`, and `(cap, n)`
* Correcting punctuation spacing
* Formatting quotation marks correctly
* Replacing `a` with `an` when followed by a word beginning with a vowel or `h`

This project focuses on file handling, string manipulation, text formatting, and algorithmic processing.

---

## Authors

* Ooja Omale (@oomale)

---

## Usage

### Prerequisites

* Go installed on your machine

### Running the Program

```bash
go run . input.txt output.txt
```

### Example

Input file:

```text
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

Run:

```bash
go run . sample.txt result.txt
```

Output file:

```text
Simply add 66 and 2 and you will see the result is 68.
```

---

## Implementation Details (Algorithm)

1. Read the contents of the input file.
2. Tokenize the text into words, punctuation, and special commands.
3. Traverse the tokens sequentially.
4. Detect and process transformation commands:

   * `(hex)` converts the previous hexadecimal value to decimal.
   * `(bin)` converts the previous binary value to decimal.
   * `(up)` converts the previous word to uppercase.
   * `(low)` converts the previous word to lowercase.
   * `(cap)` capitalizes the previous word.
   * `(up, n)`, `(low, n)`, and `(cap, n)` apply transformations to the previous `n` words.
5. Correct punctuation placement by removing unnecessary spaces before punctuation and ensuring proper spacing afterward.
6. Process quotation marks by attaching opening and closing apostrophes directly to the enclosed text.
7. Replace occurrences of `a` with `an` when the following word begins with a vowel or the letter `h`.
8. Reconstruct the corrected text.
9. Write the final result to the output file.
10. Handle file and processing errors appropriately.

---


The application uses only standard Go packages and follows good coding practices.
