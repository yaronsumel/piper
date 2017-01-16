# piper 
###### piper is a small devtool designed to instatly share stdout from process over the secured connection (TLS). 

<p align="center">
<img src="http://i.imgur.com/xFHwmyF.jpg" width="500" >
</p>

#### Installation
```bash
$ go get github.com/yaronsumel/piper
```
#### Usage

##### Server
```bash
$ tail -f log.txt | piper -a 0.0.0.0:8080
```
##### Client
```bash
$ piper -a remotehost:8080 > log.txt
```
#### TBD
* testing

> ##### Written and Maintained by [@YaronSumel](https://twitter.com/yaronsumel) #####
