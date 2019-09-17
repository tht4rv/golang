#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $CRFUSER" >> $LOGFILE
/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_createfirstuser.sh admin 0.0.0.0/0 >> $LOGFILE
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $CRFUSER" >> $LOGFILE