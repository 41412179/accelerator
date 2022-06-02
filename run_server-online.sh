set -x
git pull
# 启动正式环境
export GIN_MODE=release
go mod tidy
go build
pid=$(ps -ef | grep "./accelerator" | grep -v 'grep' | awk '{print $2}')
if [ $pid -gt 0 ]; then
    /bin/kill -9 $pid
fi

nohup ./accelerator >>acc.log 2>&1 &
tail -200f acc.log
