# docker pull image 
### A small golang tool to pull docker images use http 
 <br>
<br>



### 1)&emsp;start up
 - build
```
  sh build.sh
```
&emsp;&emsp; will generate an executable ./build/gopull,run "./build/gopull help" 

 - or run
```
  go run main.go help
```



<br>

### 2)&emsp;Pull docker images and generate a tar archive on a machine without docker
```
  ./gopull download redis
```
  
### 3)&emsp;Pull docker images whith Digest
```
  ./gopull download redis@sha256:31120dcdd310e9a65cbcadd504f4fe60a185bd634ab7c6a35e3e44a941904d97
```

### 4)&emsp;Pull amd64 images by default, user -p platform select the desired image
```
  ./gopull download -l redis 
  ./gopull download -p arm64 redis
```

### 4)&emsp;Compatible with docker pull
```
  ./gopull pull redis 
```

# Reference  https://github.com/NotGlop/docker-drag.git

