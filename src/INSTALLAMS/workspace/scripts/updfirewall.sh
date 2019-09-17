#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $UPDFIREWALL" >> $LOGFILE
./expupdfirewall.sh $UPDATEFW >> $LOGFILE
./expupdfirewall_debug.sh $UPDATEFWDB >> $LOGFILE
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $UPDFIREWALL" >> $LOGFILE