echo "start build"
GOOS=linux go build
echo "build success"
scp sgin rdserver:~/