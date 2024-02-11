# WC processor

yet another a unix like wc binary written in golang

# How to build

```
> go build -o ccwc
```

# How to use

`ccwc` can either have information piped to it or it have a file path passed via the command line.

```
Flags

-c      The number of bytes in each input file is written to the standard output.
-l      The number of lines in each input file is written to the standard output.
-m      The number of characters in each input file is written to the standard output.
-w      The number of words in each input file is written to the standard output.
```

# Example

```
./ccwc ./text/file.txt -lwc
7142 58164 335039 ./text/file.txt
```

## Notes

i'm still learning go, so expect there's some funny things you'll found here :D
