#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $STOPAMS" >> $LOGFILE
CheckStatus
if [ "$AMSstatus" = "Running" -o "$EMSstatus" = "(maintenance)" ]
then
	./updfirewall.sh $AMSVersion $AMSRevision
	/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server stop >> $LOGFILE
fi
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $STOPAMS" >> $LOGFILE