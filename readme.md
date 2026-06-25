# psh

a simple ssh connection manager I made to practice Go.
 

### usage

```text
usage: 
        psh <args>.
add, a
        add a new host.
list, l
        list all the knowun host.
conncect, c <hostname>
        connect to the host.
remove, rm <hostname>
        remove the host form the list.

The file that is going to be used for
storage is in: ~/.config/psh/hosts.json

[!] make sure to make this dir for
    making psh work
```


### Install / Uninstall

```bash
git clone https://github.com/piadi-su/psh.git
cd psh
chmod +x installer.sh
./installer.sh
```





