[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=yescrypt_crack&theme=gruvbox)](https://github.com/cyclone-github/yescrypt_crack/)

[![Go Report Card](https://goreportcard.com/badge/github.com/cyclone-github/yescrypt_crack)](https://goreportcard.com/report/github.com/cyclone-github/yescrypt_crack)
[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/yescrypt_crack.svg)](https://github.com/cyclone-github/yescrypt_crack/issues)
[![License](https://img.shields.io/github/license/cyclone-github/yescrypt_crack.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/yescrypt_crack.svg)](https://github.com/cyclone-github/yescrypt_crack/releases)
<!--
[![Go Reference](https://pkg.go.dev/badge/github.com/cyclone-github/yescrypt_crack.svg)](https://pkg.go.dev/github.com/cyclone-github/yescrypt_crack)
-->

```
./yescrypt_crack.bin -h hash.txt -w wordlist.txt

 ---------------------------------------------------------- 
|                Cyclone's Yescrypt Cracker                |
| https://https://github.com/cyclone-github/yescrypt_crack |
 ---------------------------------------------------------- 

Hash file:      hash.txt
Total Hashes:   1
CPU Threads:    56
Wordlist:       wordlist.txt
2025/03/04 10:23:08 Working...
$y$j9T$z7lNWyBfW4ZruGHCsFzDz/$Sz1GtrDDnsf0KfUE8mQHNJqGyG32TDWC287DdU97dz.:cyclone123
2025/03/04 10:23:20 Cracked: 1/1 431.78 h/s 00h:00m:11s

2025/03/04 10:23:20 Finished
```
### Info:
I wrote this tool since yescrypt has become the default /etc/shadow hash for many popular linux distros such as Debian, Ubuntu, RHEL, Fedora, Arch, etc, and due to the very limited hash cracking tools that support yescrypt. Hashcat currently lacks support for yescrypt, and while John the Ripper supports yescrypt in its bleeding-edge version, getting JtR compiled with yescrypt support can be challenging for newer users due to the required linux lib "libxcrypt" -- this is especially true for Windows users. 

Since `yescrypt_crack` is written in pure Go, it easily compiles and runs on just about any OS and architecture such as Intel/ARM, Linux, Windows, Mac.

It is worth noting that JtR is currently faster than `yescrypt_crack`, so using JtR is recommended when possible. 

### Example hash:plain:
```
$y$j9T$z7lNWyBfW4ZruGHCsFzDz/$Sz1GtrDDnsf0KfUE8mQHNJqGyG32TDWC287DdU97dz.:cyclone123
```

### Supported options:
```
-w {wordlist} (omit -w to read from stdin)
-h {yescrypt_hash}
-o {output} (omit -o to write to stdout)
-t {cpu threads}
-s {print status every nth sec}

-version (version info)
-help (usage instructions)

./yescrypt_crack.bin -h {yescrypt_hash} -w {wordlist} -o {output} -t {cpu threads} -s {print status every nth sec}

./yescrypt_crack.bin -h yescrypt.txt -w wordlist.txt -o cracked.txt -t 16 -s 10

cat wordlist | ./yescrypt_crack.bin -h yescrypt.txt

./yescrypt_crack.bin -h yescrypt.txt -w wordlist.txt -o output.txt
```

### Credits:
* `yescrypt_crack` tool was written by cyclone in pure Go
* `yescrypt_crack` uses Solar Designer's yescrypt-go implementation: https://github.com/openwall/yescrypt-go
* The yescrypt algo was written by Solar Designer: https://www.openwall.com/yescrypt/

### Changelog:
- https://github.com/cyclone-github/yescrypt_crack/blob/main/CHANGELOG.md

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/yescrypt_crack.git`
  - `cd yescrypt_crack`
  - `go mod init yescrypt_crack`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
  - `./yescrypt_crack -h {hash file} -w {wordlist file} -t {CPU threads to use (optional)}`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
