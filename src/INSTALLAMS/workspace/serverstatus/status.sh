#!/bin/bash
AMSVersion=$1
AMSRevision=$2
FILE="/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server"
STATUSFILE="/ams_build/AMSwsp/log/$AMSVersion/$AMSRevision/stage/serverstatus.log"
LOGFILE="system.log"
chmod 755 *
if [ -e "$FILE" ] 
then
	/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server status > $STATUSFILE
	cat $STATUSFILE > $LOGFILE
else
	echo -e 'File is not NotExist!!!' > $LOGFILE
	echo "" > $STATUSFILE
fi

AMSstatus=$(cat $STATUSFILE | grep "AMS ser" | awk '{print $4}' )
#echo "$AMSstatus dsf"
if [ "$AMSstatus" = "" ]
then
	AMSstatus="NotExist"
fi
echo "$AMSstatus" >$STATUSFILE
cat $STATUSFILE