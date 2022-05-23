echo "Executing gofmt -s -d ."
gofmt -s -d .
if gofmt -s -d . > /dev/null; then
    exit 1
fi
