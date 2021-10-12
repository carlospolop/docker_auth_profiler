#const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows "
#const goarchList = "386 amd64 amd64p32 arm arm64 ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 " # (new)

echo "Building for Linux 32-bit"
env GOOS=linux GOARCH=386 go build -o builds/docker_auth_profiler_386

echo "Building for Linux 64-bit"
env GOOS=linux GOARCH=amd64 go build -o builds/docker_auth_profiler_amd64