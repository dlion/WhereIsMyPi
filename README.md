# WhereIsMyPi
A simple tool to find my raspberry pi on my LAN

## Usage
```
> ./whereismypi
Your Raspberry Pi is up on: 192.168.1.116
```

## How it works
Scans my subnet, when it finds an host up the software searchs in the ARP table on your machine about it, if the MAC Address contains `b8:27:eb` it's own by your Raspberry Pi and returns the IP address.

## Author
* Domenico Luciani
* http://dlion.it
* domenicoleoneluciani@gmail.com

## License
MIT © [Domenico Luciani](https://github.com/DLion)
