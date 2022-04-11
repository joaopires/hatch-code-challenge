# Hatch code challenge

CLI program that reads two json files and compares them for equality, printing out the result on stdout.

## Build

In order to build the project you must have the [go sdk installed](https://go.dev/doc/install) on your system. The project was developed with the latest v1.18.

## Run

The program has two arguments for the filepaths of both json files to compare.

```bash
❯ executable filepath1 filepath2
```

There's a makefile to make compilation easier and several json file samples in **./resources** folder for testing.

You have three options:

Local build of the executable

```bash
projects/go-stuff/hatch-code-challenge on  main [!] via 🐹 v1.18
❯ make build
go build -o hatch ./cmd/hatch

projects/go-stuff/hatch-code-challenge on  main [!?] via 🐹 v1.18
❯ ./hatch resources/equal/unordered_small_file_before_backup.json resources/equal/unordered_small_file_after_backup.json
```

System wide installation

```bash
projects/go-stuff/hatch-code-challenge on  main [!] via 🐹 v1.18
❯ make install
go install ./cmd/hatch

projects/go-stuff/hatch-code-challenge on  main [!?] via 🐹 v1.18
❯ hatch resources/equal/unordered_small_file_before_backup.json resources/equal/unordered_small_file_after_backup.json
```

Run directly with temp generated executable

```bash
projects/go-stuff/hatch-code-challenge on  main [!?] via 🐹 v1.18
❯ go run ./cmd/hatch/main.go resources/equal/unordered_small_file_before_backup.json resources/equal/unordered_small_file_after_backup.json
```

## Testing

Along with the source code, there are unit test for all the main components. To run them use the makefile recipe

```bash
❯ make test
go test -v ./...
?       github.com/joaopires/hatch/cmd/hatch    [no test files]
=== RUN   TestProcess
0be62b92f5a97b2a48b98631c9909545ef341fd555646d645a939e89b35e95a1=== RUN   TestProcess/file_not_found
=== RUN   TestProcess/syntax_error_in_file_beginning
=== RUN   TestProcess/cannot_decode_object
=== RUN   TestProcess/syntax_error_in_file_ending
=== RUN   TestProcess/file_processed_successfully
--- PASS: TestProcess (0.00s)
    --- PASS: TestProcess/file_not_found (0.00s)
    --- PASS: TestProcess/syntax_error_in_file_beginning (0.00s)
    --- PASS: TestProcess/cannot_decode_object (0.00s)
    --- PASS: TestProcess/syntax_error_in_file_ending (0.00s)
    --- PASS: TestProcess/file_processed_successfully (0.00s)
=== RUN   TestEqualSize
=== RUN   TestEqualSize/files_have_the_same_size
=== RUN   TestEqualSize/files_don't_have_the_same_size
=== RUN   TestEqualSize/file_A_not_found
=== RUN   TestEqualSize/file_B_not_found
--- PASS: TestEqualSize (0.00s)
    --- PASS: TestEqualSize/files_have_the_same_size (0.00s)
    --- PASS: TestEqualSize/files_don't_have_the_same_size (0.00s)
    --- PASS: TestEqualSize/file_A_not_found (0.00s)
    --- PASS: TestEqualSize/file_B_not_found (0.00s)
=== RUN   TestGetFileSize
=== RUN   TestGetFileSize/empty_filepath
=== RUN   TestGetFileSize/file_not_found
=== RUN   TestGetFileSize/file_exists
--- PASS: TestGetFileSize (0.00s)
    --- PASS: TestGetFileSize/empty_filepath (0.00s)
    --- PASS: TestGetFileSize/file_not_found (0.00s)
    --- PASS: TestGetFileSize/file_exists (0.00s)
PASS
ok      github.com/joaopires/hatch/internal/file        (cached)
?       github.com/joaopires/hatch/internal/utils       [no test files
```
