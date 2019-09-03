# httpull

Parses masscan XML file and pulls out HTTP and HTTPS servers found.

## Install

```
$ go get -u github.com/simonuvarov/httpull
```

## Basic Usage

httpull accepts XML report from `stdin`:

```
$ head recon/example/nmap.xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE nmaprun>
...
$ cat recon/example/nmap.xml | httpull
http://example.com
http://example.net
http://example.edu
https://example.com
https://example.edu
https://example.net
```
# Credit
Inspired by [httprobe](https://github.com/tomnomnom/httprobe)
