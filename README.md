# piper
pipe things over TLS

<img src="http://i.imgur.com/xFHwmyF.jpg" width="80%"/>

## server
```bash
$ tail -f log.txt | piper -a 0.0.0.0:8080
```

## client
```bash
piper -a 0.0.0.0:8080 > log.txt
```
