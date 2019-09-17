#!/usr/bin/expect -d
set timeout 10
set shellscript [lindex $argv 0]
set flag [lrange $argv 1 end]
spawn $shellscript {*}$flag
expect {
	"Are you sure you want to proceed" { send "yes\r";exp_continue  }
	"The user amssys has some running shells, Do you want to kill them" { send "yes\r";exp_continue  }
}
set timeout 200
expect EOF