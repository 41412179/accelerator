go build
kill -9 $(ps -ef | grep "./accelerator" | awk '{print $2}')
nohup ./accelerator >acc.log 2>&1 &
