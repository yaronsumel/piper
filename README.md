# piper
pipe things over TLS


## serevr
```bash
$ tail -f log.txt | piper 0.0.0.0:8080
```

## client
```bash
piper 0.0.0.0:8080 > log.txt
```