for file in $(git diff main --raw | grep packages/.*.go | awk "{ print \$6 }" | xargs) ; do
		staticcheck $file ;
	done