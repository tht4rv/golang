#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision
CheckStatus
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $STARTAMS" >> $LOGFILE
if [ "$AMSstatus" = "Disabled" ]
then
	./updfirewall.sh $AMSVersion $AMSRevision
	echo $EMSstatus >> $LOGFILE
	if [ "$EMSstatus" = "(maintenance)" ]
	then
		./StopAMS.sh $AMSVersion $AMSRevision
	fi
	/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server start >> $LOGFILE
	./addsshporwithstatus.sh $AMSVersion $AMSRevision >> $LOGFILE
fi
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exist $STARTAMS" >> $LOGFILE
