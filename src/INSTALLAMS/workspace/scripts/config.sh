#!/bin/bash
AMSVersion=$1
AMSRevision=$2
FILE="/opt/ams/software/ams-$AMSVersion-$AMSRevision/bin/ams_server"
STATUSFILE="../log/$AMSVersion/$AMSRevision/stage/status.log"
AMSBINLOG="../log/$AMSVersion/$AMSRevision/stage/amsbin.log"
LOGFILE="../log/$AMSVersion/$AMSRevision/system/system.log"
ALOGFILE="/ams_build/AMSwsp/log/$AMSVersion/$AMSRevision/system/system.log"
AAMSBINLOG="/ams_build/AMSwsp/log/$AMSVersion/$AMSRevision/stage/amsbin.log"
INSTALLSTATUS="../log/$AMSVersion/$AMSRevision/stage/installstatus.log"
STARTAMS="StartAMS.sh"
STOPAMS="stopams.sh"
CRFUSER="CreateFirstUser.sh"
UPDATEFW="/opt/ams/software/ams-$AMSVersion-$AMSRevision*/bin/ams_updatefirewall"
UPDATEFWDB="/opt/ams/software/ams-$AMSVersion-$AMSRevision*/bin/ams_updatefirewall --debug"
UPDFIREWALL="updfirewall.sh"
STATUSSH="status.sh"
INSTALLFILE="installams.sh"
UNINSTALLAMS="uninstallams.sh"
UNINSTALL="uninstall.sh"
WORKSPACE="/ams_build/AMSwsp/scripts"
AMSBIN="/ams_build/build/$AMSVersion/$AMSRevision"
RUNAMSSH="runamsbin.sh"
LOGSTAGEFOLDER="/ams_build/AMSwsp/log/$AMSVersion/$AMSRevision/stage/"
LOGSYSTEMFOLDER="/ams_build/AMSwsp/log/$AMSVersion/$AMSRevision/system/"
CheckStatus()
{
	cd $WORKSPACE
	chmod 755 $WORKSPACE/*.sh
	./status.sh $AMSVersion $AMSRevision
	AMSstatus=$(cat $STATUSFILE | grep "AMS ser" | awk '{print $4}' )
	if [ "$AMSstatus" = "" ]
	then
		AMSstatus="NotExist"
	fi
	EMSstatus=$(cat $STATUSFILE | grep "EMS ser" | awk '{print $5}' )
	JBOSSstatus=$(cat $STATUSFILE | grep "JBoss ..." | awk '{print $3}' )
	rm -f $STATUSFILE
}
chmod 777 /tmp/
chmod +t /tmp/
MakeFolder(){
	FOLDER=$1
	if [ ! -d $FOLDER  ]
	then 
		mkdir -p $FOLDER 
	fi
}
MakeFolder $LOGSTAGEFOLDER
MakeFolder $LOGSYSTEMFOLDER