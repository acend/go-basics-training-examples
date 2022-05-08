# tgz2zip

```bash
# create tar.gz
tar czvf files.tar.gz testdata/

# conver to zip
cat files.tar.gz | ./tgz2zip > files.zip

# verify
zipinfo files.zip
unzip -d tmp files.zip

cat tmp/testdata/file1.txt
cat tmp/testdata/file2.txt
```

# http handler

run convert server in main:

```golang
	err := convertServer()
```

```bash
curl --data-binary @files.tar.gz -o output.zip http://localhost:8080
zipinfo output.zip
```
