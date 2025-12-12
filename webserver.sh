# go run ./webserver/webserver.go 9090

# ./mega_go_webserver 9090


portList=(
  9090
)
for(( i=0;i<${#portList[@]};i++)) do
    port=${portList[i]}
echo "port="$port
pm2 start ./build/go_webserver --name go_webserver-$port   --output="./logs/go_webserver"-$port".log" --log-date-format="YYYY-MM-DD HH:mm:ss:SSS"   -- $port $env
done;