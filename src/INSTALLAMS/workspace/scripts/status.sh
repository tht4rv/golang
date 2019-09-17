#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter status.sh" >> $LOGFILE
if [ -e "$FILE" ] 
then
	/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server status > $STATUSFILE
	cat $STATUSFILE >> $LOGFILE
else
	echo -e 'File is not NotExist!!!' >> $LOGFILE
	echo "" > $STATUSFILE
fi
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit status.sh" >> $LOGFILE
