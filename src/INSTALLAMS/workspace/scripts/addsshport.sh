#!/bin/bash
SSHPORT=""
SSHPORT=$(iptables -L | awk '$1 == "ACCEPT" && $7 == "dpt:ssh" {print "yes"}')
if [ "$SSHPORT" == "yes" ]
then
	sudo iptables -I INPUT -p tcp -m tcp --dport 22 -j ACCEPT
	sudo service iptables save
fi
