#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision

NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $RUNAMSSH" >> $ALOGFILE
cd $AMSBIN
chmod 755 $AMSBIN/*
./ams-$AMSVersion-$AMSRevision-redhat-x86_64.bin -c ../simplex.conf --activate-force > $AAMSBINLOG
cat $AAMSBINLOG >> $ALOGFILE
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $RUNAMSSH" >> $ALOGFILE
