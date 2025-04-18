# nmap工具的使用
```bash
# nmap -sS  192.168.100.2 -T4 -A -O
Starting Nmap 7.91 ( https://nmap.org ) at 2025-04-18 17:52 CST
Nmap scan report for 192.168.100.2
Host is up (0.00010s latency).
Not shown: 989 closed ports
PORT      STATE    SERVICE          VERSION
22/tcp    open     ssh              OpenSSH 8.0 (protocol 2.0)
| ssh-hostkey: 
|   3072 ad:f8:20:56:c3:93:58:0c:77:df:e7:dc:6a:c1:0e:12 (RSA)
|   256 dd:17:26:24:ac:26:34:b4:64:59:7d:5e:0e:22:91:23 (ECDSA)
|_  256 99:f5:13:8c:15:39:6a:d1:36:a8:da:b9:f9:55:d4:1e (ED25519)
23/tcp    open     telnet           Linux telnetd
3001/tcp  filtered nessus
3306/tcp  filtered mysql
7001/tcp  filtered afs3-callback
8080/tcp  filtered http-proxy
8090/tcp  filtered opsmessaging
9001/tcp  filtered tor-orport
9091/tcp  filtered xmltec-xmlmail
9200/tcp  filtered wap-wsp
10000/tcp filtered snet-sensor-mgmt
Device type: general purpose
Running: Linux 3.X
OS CPE: cpe:/o:linux:linux_kernel:3
OS details: Linux 3.7 - 3.10
Network Distance: 0 hops
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

OS and Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 10.70 seconds

# nmap 192.168.50.1/24 -sL 
```