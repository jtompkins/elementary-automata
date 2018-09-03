# elementary-automata

A simple program to generate images for Wolfram's [elementary automata](). The output looks like this (for [rule 126]()):

![rule 126 output](assets/one-twenty-six.png)

## Usage

`elementary-automata [<flags>] <rule> <file>`

### Flags

| Flag    | What does it do?                                                   | Default |
| ------- | ------------------------------------------------------------------ | ------- |
| `-s`    | How large is each generation?                                      | 1000    |
| `-g`    | How many generations should be simulated?                          | 1000    |
| `scale` | How large should each cell appear in the output image (in pixels)? | 1       |

## Testing

You need [Ginkgo](), [Gomega](), and [Counterfeiter]() installed to run the tests.

1. `go generate ./...`
1. `go test ./...` _or_ `ginkgo -r`
