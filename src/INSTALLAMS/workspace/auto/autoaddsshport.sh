#!/bin/bash
#FILE="/opt/ams/software/ams-*/bin/ams_server"
STATUSFILE="status.log"
LOGFILE="system.log"
checkstaus()
{	
		AMSstatus=""
		if [ -d "/opt/ams/software/" ]
		then
			/opt/ams/software/ams-*/bin/ams_server status > $STATUSFILE
			cat $STATUSFILE > $LOGFILE
			AMSstatus=$(cat $STATUSFILE | grep "AMS ser" | awk '{print $4}' )
			rm -f $STATUSFILE
		fi	
}
start=$SECONDS
while true
do
	duration=$(( SECONDS - start ))
	if [ $duration -gt 300 ]
	then
		checkstaus
		if [ "$AMSstatus" = "Disabled" -o "$AMSstatus" = "Running" -o "$AMSstatus" = "" -o "$AMSstatus" = "WaitingStartup" ]
		then
			SSHPORT=""
			SSHPORT=$(iptables -L | awk '$1 == "ACCEPT" && $7 == "dpt:ssh" {print "yes"}')
			if [ "$SSHPORT" == "yes" ]
			then
				sudo iptables -I INPUT -p tcp -m tcp --dport 22 -j ACCEPT > $LOGFILE
				sudo service iptables save > $LOGFILE
			fi
		fi
		start=$SECONDS
	fi
	sleep 30
done
