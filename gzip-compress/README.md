# gzip-compress

```bash
hexdump -C testfile.txt

# compress
./gzip-compress testfile.txt > out.gz
hexdump -C out.gz

# decompress
./gzip-compress -d out.gz > out
cat out

# decompress from stdin
cat out.gz | ./gzip-compress -d -
```
