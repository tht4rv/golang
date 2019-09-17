#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $UNINSTALL" >> $LOGFILE
CheckStatus
if [ "$AMSstatus" = "Disabled" ]
then
	if [ "$EMSstatus" = "(maintenance)"]
	then
		./stopams.sh $AMSVersion $AMSRevision
	fi
	./updfirewall.sh $AMSVersion $AMSRevision
	./expuninstall.sh /opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_uninstall >> $LOGFILE
	./addsshport.sh >> $LOGFILE
fi
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $UNINSTALL" >> $LOGFILE