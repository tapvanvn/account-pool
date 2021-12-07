DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

mkdir -p "$DIR/#temp/config/"

go build -o "$DIR/#temp/accpool" ../main.go 

rs=$?
if [ $rs -eq 0 ]; then 
    echo "SUCCESS"
    cp ../config/config_local.jsonc     "$DIR/#temp/config/config.jsonc"
    cp -r ../static                     "$DIR/#temp/"
    cp ../route.jsonc                   "$DIR/#temp/"
    accpool="$DIR/#temp/accpool"
    kill $(lsof -t -i:8080)
    PORT=8080 $accpool
else
    echo "FAIL"
fi