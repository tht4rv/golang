#!/usr/bin/expect -d
set timeout 10
set shellscript [lindex $argv 0]
set flag [lrange $argv 1 end]
spawn $shellscript {*}$flag
expect {
	"Do you want to add these ports to the iptables" { send "yes\r";exp_continue  }
}
set timeout 60
expect EOF
