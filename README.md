# GOROUTE

## Introduction
 - Goroute is a proxy server implementation similar to Nginx. It is implemented in GoLang. It uses memcached as caching provider.

## Steps to build gorouter
 
 -  Install pre-requisites
    ```
    go get -u github.com/bradfitz/gomemcache/memcache
    go get -u github.com/gorilla/mux
    go get -u github.com/NYTimes/gziphandler
    ```
 -  Build Executable
    ```
    go build -o goroute .
    ```

## Running Goroute 
   ```
   goroute
    -config <Required. Config File Path For Router>
    -env <Optional. Env File Path For Dynamic Router Values>
    -help <Optional. Help>
   ```

## Important Points

### Providing dynamic environment values
  - Provide values enwrapped with $ at both start and end. Ex - $namespace$
  - Provide environment variable name for value in environment config file. Ex - "namespace" : "DATA_STACK_NAMESPACE"
  - Sample config and enviroment files are present in repo. 

