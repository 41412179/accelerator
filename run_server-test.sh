set -x
git pull
go build
pid=$(ps -ef | grep "./accelerator" | grep -v 'grep' | awk '{print $2}')
if [ $pid -gt 0 ]; then
    /bin/kill -9 $pid
fi

nohup ./accelerator >>acc.log 2>&1 &
tail -200f acc.log