#!/bin/bash
AMSVersion=$1
AMSRevision=$2
source config.sh $AMSVersion $AMSRevision
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Enter $INSTALLFILE" >> $LOGFILE
CheckStatus
if [ "$AMSstatus" = "Disabled" ]
then
	echo "Disabled" >> $LOGFILE
	echo "STARTING" > $INSTALLSTATUS
	./startams.sh $AMSVersion $AMSRevision	
	echo "STARTED" > $INSTALLSTATUS
elif [ "$AMSstatus" = "Running" ]
then	
	echo "CREATING" > $INSTALLSTATUS
	./createfirstuser.sh $AMSVersion $AMSRevision
	./addsshport.sh >> $LOGFILE
	echo "CREATED" > $INSTALLSTATUS
elif [ "$AMSstatus" = "Starting" ]
then
	echo "Starting" >> $LOGFILE
elif [ "$AMSstatus" = "WaitingStartup" ]
then
	if [ "$JBOSSstatus" != "WaitingStartup" ]
	then
		echo "WaitingStartup" > $INSTALLSTATUS
	fi
elif [ "$AMSstatus" = "NotExist" ]
then
	echo "NOT EXIST FILE" >> $LOGFILE
	echo "BINRUNNING" > $INSTALLSTATUS
	./runamsbin.sh $AMSVersion $AMSRevision >> $LOGFILE
	./addsshport.sh >> $LOGFILE
	SPROCESS2=""
	SPROCESS1=$(cat $WORKSPACE/$AMSBINLOG | grep "Activation finished"|awk '$2 == "finished" {print "BINRAN"}')
	SPROCESS2=$(cat $WORKSPACE/$AMSBINLOG | grep "Installation failed."|awk '$2 == "failed." {print "BINFAIL"}')
	if [ "$SPROCESS2" = "BINFAIL" ]
	then
		./uninstallams.sh $AMSVersion $AMSRevision
		echo "$SPROCESS1" > $INSTALLSTATUS
	else
		echo "$SPROCESS1" > $INSTALLSTATUS
	fi
	
fi
NOW=$(date '+%d/%m/%Y %H:%M:%S')
echo "$NOW :: Exit $INSTALLFILE" >> $LOGFILE
