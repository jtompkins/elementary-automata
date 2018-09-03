# elementary-automata

A simple program to generate images for Wolfram's [elementary automata](). The output looks like this (for [rule 126]()):

![](assets/one-twenty-six.png)

## Usage

`elementary-automata -r 126 -s 1000 -g 1000 -f onetwentysix.png`

### Flags

| Flag | What does it do?                                                                      | Default  |
| ---- | ------------------------------------------------------------------------------------- | -------- |
| `-r` | The rule to use                                                                       | N/A      |
| `-s` | How large is each generation?                                                         | 1000     |
| `-g` | How many generations should be simulated?                                             | 1000     |
| `-f` | The output file name                                                                  | N/A      |
| `-l` | How should the initial generation be populated?<br>Valid values: `center` or `random` | `random` |

## Testing

You need [Ginkgo](), [Gomega](), and [Counterfeiter]() installed to run the tests.

1. `go generate ./...`
1. `go test ./...` _or_ `ginkgo -r`
