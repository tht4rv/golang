#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $UNINSTALLAMS" >> $LOGFILE
CheckStatus
if [ "$AMSstatus" = "Disabled" ]
then
	echo "UNINSTALLING" > $INSTALLSTATUS
	./uninstall.sh $AMSVersion $AMSRevision
	echo "UNINSTALLED" > $INSTALLSTATUS
elif [ "$AMSstatus" = "Running" ]
then
	echo "STOPPING" > $INSTALLSTATUS
	./stopams.sh $AMSVersion $AMSRevision
	echo "STOPPED" > $INSTALLSTATUS
elif [ "$AMSstatus" = "Starting" ]
then
	echo "" > $LOGFILE
elif [ "$AMSstatus" = "NotExist" ]
then
	echo "NOTEXIST" > $INSTALLSTATUS
fi

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $UNINSTALLAMS" >> $LOGFILE