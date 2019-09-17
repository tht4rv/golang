#!/bin/bash
AMSVersion=$1
AMSRevison=$2
source config.sh $AMSVersion $AMSRevison
CheckStatus
while [ "$AMSstatus" != "Running" ]
do
	CheckStatus
	sleep 60
done
./addsshport.sh >> $LOGFILE