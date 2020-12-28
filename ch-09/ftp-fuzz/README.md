# Fuzzer

The file `main.go` in this folder is a fuzzer intended for identifying buffer overflows over a network. The program will repeatedly sent the character `A` in increments of one (e.g. A, AA, AAA, n), where 'n' is the maximum length specified (2500 by default) to an IP address and port specified by the user.

## Arguments

```
--ip	 :: Destination IP (default 127.0.0.1)
--port   :: Destination port (default 21)
--buffer :: Length of the buffer (default 2500)
```
