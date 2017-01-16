# piper
pipe things over TLS

![piper](http://i.imgur.com/9Mttcxg.jpg)

## server
```bash
$ tail -f log.txt | piper -a 0.0.0.0:8080
```

## client
```bash
piper -a 0.0.0.0:8080 > log.txt
```
