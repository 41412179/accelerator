set -x
git pull
go build
/bin/kill $(ps -ef | grep "./accelerator" | awk '{print $2}')
nohup ./accelerator >acc.log 2>&1 &
