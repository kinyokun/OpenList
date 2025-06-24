appName="openlist"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor=$(git show -s --format='format:%aN <%ae>' HEAD)
gitCommit=$(git log --pretty=format:"%h" -1)
version=$(git describe --long --tags --dirty --always)
webVersion=$(curl -s --max-time 5 "https://api.github.com/repos/OpenListTeam/OpenList-Frontend/releases/latest" -L | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' | sed 's/^v//')
if [ -z "$webVersion" ]; then
    webVersion="0.0.0"
fi
ldflags="\
-w -s \
-X 'github.com/kinyokun/OpenList/internal/conf.BuiltAt=$builtAt' \
-X 'github.com/kinyokun/OpenList/internal/conf.GoVersion=$goVersion' \
-X 'github.com/kinyokun/OpenList/internal/conf.GitAuthor=$gitAuthor' \
-X 'github.com/kinyokun/OpenList/internal/conf.GitCommit=$gitCommit' \
-X 'github.com/kinyokun/OpenList/internal/conf.Version=$version' \
-X 'github.com/kinyokun/OpenList/internal/conf.WebVersion=$webVersion' \
"
go build -ldflags="$ldflags" .