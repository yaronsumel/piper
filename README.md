# piper 
###### piper is a small devtool designed to instatly share stdout from process over the secured connection (TLS). 

![](http://i.imgur.com/xFHwmyF.jpg )


#### Installation
```bash
$ go get github.com/yaronsumel/piper
```
#### Usage

##### Server
when piper used with named piper it will run as server
```bash
$ tail -f log.txt | piper -a 0.0.0.0:8080
```
##### Client
```bash
$ piper -a remotehost:8080 > log.txt
```
#### TBD
* testing
