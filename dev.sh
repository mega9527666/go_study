# go run ./webserver/webserver.go 9090

# ./mega_go_webserver 9090

env=1

portList=(
  9090
)
for(( i=0;i<${#portList[@]};i++)) do
    port=${portList[i]}
echo "port="$port
echo "env="$env
pm2 start ./build/go_webserver  --name go_webserver-$port   --output="./logs/go_webserver"-$port".log" --log-date-format="YYYY-MM-DD HH:mm:ss:SSS"   -- $port $env  
done;


# portList=(
#   9091
# )
# for(( i=0;i<${#portList[@]};i++)) do
#     port=${portList[i]}
# echo "port="$port
# pm2 start ./build/go_account_server --name go_account_server-$port   --output="./logs/go_account_server"-$port".log" --log-date-format="YYYY-MM-DD HH:mm:ss:SSS"   -- $port $env
# done;