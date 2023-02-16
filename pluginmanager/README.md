# GRPC / protobuf

## Install protoc

### RHEL

Reference: <https://protobuf.dev/getting-started/gotutorial/>

```bash
sudo yum install protobuf
sudo yum install protobuf-compiler
```

### Ubuntu

Reference: https://gist.github.com/1duo/38ac328fce58c3f14a2f4ac75bb4484d

```bash
sudo apt-get install build-essential
wget https://github.com/google/protobuf/releases/download/v2.6.1/protobuf-2.6.1.tar.gz
tar -zxvf protobuf-2.6.1.tar.gz && cd protobuf-2.6.1/
./configure
make -j$(nproc) && make check

sudo make install
export LD_LIBRARY_PATH=/usr/local/lib
protoc --version
```


## Install protoc plugin for go

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```bash
bash; 
export PATH=$PATH:/home/abhijith/go/bin/;
protoc --proto_path=pm --go_out=pm --go_opt=paths=source_relative pm/pm.proto
```

```bash
bash-4.4$ ls pm
pm.pb.go  pm.proto  README.md
```

