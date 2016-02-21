# Usage:
#   ./deploy.sh production
#   ./deploy.sh dev

if [ -n "$*" ]; then
  env=$*
else
  env=dev
fi

echo "Deploying \033[1;31m$env\033[0m from branch \033[1;33m$(git branch | sed -n '/\* /s///p')\033[0m..."

# build enterprise.go
echo "Building enterprise seeds..."
#GOOS=linux GOARCH=amd64 go build -o db/seeds/enterprise db/seeds/enterprise.go

# build seeds.go
echo "Building main seeds..."
#GOOS=linux GOARCH=amd64 go build -o db/seeds/main db/seeds/main.go

echo "Deploying..."
#harp -s $env -log deploy
