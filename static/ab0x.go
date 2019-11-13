// Code generated by fileb0x at "2019-11-13 09:52:14.649183 -0600 CST m=+0.034907601" from config file "b0x.json" DO NOT EDIT.
// modification hash(275dda2afcd354de6384462093b6ce52.54f8d5f27683cbfb9b6b290cae5dfe8b)

package static


import (
  "bytes"
  
  "context"
  "io"
  "net/http"
  "os"
  "path"


  "golang.org/x/net/webdav"


)

var ( 
  // CTX is a context for webdav vfs
  CTX = context.Background()

  
  // FS is a virtual memory file system
  FS = webdav.NewMemFS()
  

  // Handler is used to server files through a http handler
  Handler *webdav.Handler

  // HTTP is the http file system
  HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}



// FileIndexHTML is "/index.html"
var FileIndexHTML = []byte("\x3c\x21\x44\x4f\x43\x54\x59\x50\x45\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x3e\x0a\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x20\x20\x20\x20\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x22\x73\x74\x79\x6c\x65\x73\x68\x65\x65\x74\x22\x20\x68\x72\x65\x66\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x63\x64\x6e\x6a\x73\x2e\x63\x6c\x6f\x75\x64\x66\x6c\x61\x72\x65\x2e\x63\x6f\x6d\x2f\x61\x6a\x61\x78\x2f\x6c\x69\x62\x73\x2f\x62\x75\x6c\x6d\x61\x2f\x30\x2e\x36\x2e\x32\x2f\x63\x73\x73\x2f\x62\x75\x6c\x6d\x61\x2e\x6d\x69\x6e\x2e\x63\x73\x73\x22\x3e\x0a\x20\x20\x20\x20\x3c\x73\x63\x72\x69\x70\x74\x20\x73\x72\x63\x3d\x22\x6a\x73\x2f\x66\x6f\x6f\x73\x2e\x6a\x73\x22\x3e\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x73\x65\x63\x74\x69\x6f\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x68\x65\x72\x6f\x20\x69\x73\x2d\x73\x75\x63\x63\x65\x73\x73\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x68\x65\x72\x6f\x2d\x62\x6f\x64\x79\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x68\x61\x73\x2d\x74\x65\x78\x74\x2d\x63\x65\x6e\x74\x65\x72\x65\x64\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x68\x31\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x74\x6c\x65\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x42\x6c\x75\x65\x70\x72\x69\x6e\x74\x73\x20\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x68\x31\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x68\x65\x72\x6f\x2d\x66\x6f\x6f\x74\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6e\x61\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x61\x62\x73\x20\x69\x73\x2d\x62\x6f\x78\x65\x64\x20\x69\x73\x2d\x66\x75\x6c\x6c\x77\x69\x64\x74\x68\x20\x69\x73\x2d\x6c\x61\x72\x67\x65\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x75\x6c\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6c\x69\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x61\x62\x20\x69\x73\x2d\x61\x63\x74\x69\x76\x65\x22\x20\x6f\x6e\x63\x6c\x69\x63\x6b\x3d\x22\x53\x65\x6e\x64\x54\x6f\x42\x61\x63\x6b\x65\x6e\x64\x28\x29\x22\x3e\x3c\x61\x3e\x43\x72\x65\x61\x74\x65\x20\x62\x6c\x75\x65\x70\x72\x69\x6e\x74\x3c\x2f\x61\x3e\x3c\x2f\x6c\x69\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x75\x6c\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x6e\x61\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x73\x65\x63\x74\x69\x6f\x6e\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x73\x65\x63\x74\x69\x6f\x6e\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x69\x64\x3d\x22\x57\x65\x62\x44\x65\x76\x22\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6e\x74\x65\x6e\x74\x2d\x74\x61\x62\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x61\x6e\x63\x65\x73\x74\x6f\x72\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x34\x20\x69\x73\x2d\x76\x65\x72\x74\x69\x63\x61\x6c\x20\x69\x73\x2d\x70\x61\x72\x65\x6e\x74\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x63\x68\x69\x6c\x64\x20\x62\x6f\x78\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x69\x64\x3d\x22\x66\x69\x6c\x65\x2d\x75\x70\x6c\x6f\x61\x64\x22\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x20\x68\x61\x73\x2d\x6e\x61\x6d\x65\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6c\x61\x62\x65\x6c\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x6c\x61\x62\x65\x6c\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x69\x6e\x70\x75\x74\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x69\x6e\x70\x75\x74\x22\x20\x74\x79\x70\x65\x3d\x22\x66\x69\x6c\x65\x22\x20\x6e\x61\x6d\x65\x3d\x22\x72\x65\x73\x75\x6d\x65\x22\x20\x6f\x6e\x63\x68\x61\x6e\x67\x65\x3d\x22\x55\x70\x6c\x6f\x61\x64\x46\x69\x6c\x65\x28\x29\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x63\x74\x61\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x69\x63\x6f\x6e\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x69\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x61\x73\x20\x66\x61\x2d\x75\x70\x6c\x6f\x61\x64\x22\x3e\x3c\x2f\x69\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x6c\x61\x62\x65\x6c\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x43\x68\x6f\x6f\x73\x65\x20\x61\x20\x66\x69\x6c\x65\xe2\x80\xa6\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x6c\x65\x2d\x6e\x61\x6d\x65\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x4e\x6f\x20\x66\x69\x6c\x65\x20\x75\x70\x6c\x6f\x61\x64\x65\x64\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x6c\x61\x62\x65\x6c\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x63\x68\x69\x6c\x64\x20\x62\x6f\x78\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x69\x65\x6c\x64\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6e\x74\x72\x6f\x6c\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x74\x65\x78\x74\x61\x72\x65\x61\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x65\x78\x74\x61\x72\x65\x61\x22\x20\x70\x6c\x61\x63\x65\x68\x6f\x6c\x64\x65\x72\x3d\x22\x43\x6f\x6e\x74\x65\x6e\x74\x20\x66\x69\x6c\x65\x22\x20\x69\x64\x3d\x22\x69\x6e\x70\x75\x74\x2d\x64\x6f\x63\x22\x3e\x3c\x2f\x74\x65\x78\x74\x61\x72\x65\x61\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x70\x61\x72\x65\x6e\x74\x20\x2d\x69\x73\x2d\x76\x65\x72\x74\x69\x63\x61\x6c\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x63\x68\x69\x6c\x64\x20\x62\x6f\x78\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x66\x69\x67\x75\x72\x65\x20\x63\x6c\x61\x73\x73\x3d\x22\x69\x6d\x61\x67\x65\x20\x69\x73\x2d\x31\x36\x62\x79\x39\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x69\x6d\x67\x20\x73\x72\x63\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x62\x75\x6c\x6d\x61\x2e\x69\x6f\x2f\x69\x6d\x61\x67\x65\x73\x2f\x70\x6c\x61\x63\x65\x68\x6f\x6c\x64\x65\x72\x73\x2f\x36\x34\x30\x78\x33\x36\x30\x2e\x70\x6e\x67\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x66\x69\x67\x75\x72\x65\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x74\x69\x6c\x65\x20\x69\x73\x2d\x63\x68\x69\x6c\x64\x20\x62\x6f\x78\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x66\x69\x67\x75\x72\x65\x20\x63\x6c\x61\x73\x73\x3d\x22\x69\x6d\x61\x67\x65\x20\x69\x73\x2d\x31\x36\x62\x79\x39\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x69\x6d\x67\x20\x73\x72\x63\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x62\x75\x6c\x6d\x61\x2e\x69\x6f\x2f\x69\x6d\x61\x67\x65\x73\x2f\x70\x6c\x61\x63\x65\x68\x6f\x6c\x64\x65\x72\x73\x2f\x36\x34\x30\x78\x33\x36\x30\x2e\x70\x6e\x67\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x66\x69\x67\x75\x72\x65\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e")

// FileJsFoosJs is "/js/foos.js"
var FileJsFoosJs = []byte("\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x55\x70\x6c\x6f\x61\x64\x46\x69\x6c\x65\x28\x29\x7b\x0a\x20\x20\x20\x20\x6c\x65\x74\x20\x69\x6e\x70\x75\x74\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x71\x75\x65\x72\x79\x53\x65\x6c\x65\x63\x74\x6f\x72\x28\x27\x23\x66\x69\x6c\x65\x2d\x75\x70\x6c\x6f\x61\x64\x20\x69\x6e\x70\x75\x74\x5b\x74\x79\x70\x65\x3d\x66\x69\x6c\x65\x5d\x27\x29\x3b\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x69\x6e\x70\x75\x74\x2e\x66\x69\x6c\x65\x73\x2e\x6c\x65\x6e\x67\x74\x68\x20\x3e\x20\x30\x29\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x66\x69\x6c\x65\x4e\x61\x6d\x65\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x71\x75\x65\x72\x79\x53\x65\x6c\x65\x63\x74\x6f\x72\x28\x27\x23\x66\x69\x6c\x65\x2d\x75\x70\x6c\x6f\x61\x64\x20\x2e\x66\x69\x6c\x65\x2d\x6e\x61\x6d\x65\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x66\x69\x6c\x65\x4e\x61\x6d\x65\x20\x2e\x74\x65\x78\x74\x43\x6f\x6e\x74\x65\x6e\x74\x20\x3d\x20\x69\x6e\x70\x75\x74\x2e\x66\x69\x6c\x65\x73\x5b\x30\x5d\x2e\x6e\x61\x6d\x65\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x73\x65\x74\x20\x63\x6f\x6e\x74\x65\x6e\x74\x20\x66\x6f\x72\x20\x7a\x6f\x6e\x65\x20\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x64\x69\x73\x70\x6c\x61\x79\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x67\x65\x74\x45\x6c\x65\x6d\x65\x6e\x74\x42\x79\x49\x64\x28\x22\x69\x6e\x70\x75\x74\x2d\x64\x6f\x63\x22\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x69\x6e\x70\x75\x74\x2e\x66\x69\x6c\x65\x73\x5b\x30\x5d\x2e\x6e\x61\x6d\x65\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x22\x68\x65\x72\x65\x20\x77\x65\x20\x61\x72\x65\x22\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x72\x65\x61\x64\x65\x72\x20\x3d\x20\x6e\x65\x77\x20\x46\x69\x6c\x65\x52\x65\x61\x64\x65\x72\x28\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x61\x64\x65\x72\x2e\x6f\x6e\x6c\x6f\x61\x64\x20\x3d\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x66\x69\x6c\x65\x52\x65\x61\x64\x43\x6f\x6d\x70\x6c\x65\x74\x65\x64\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x77\x68\x65\x6e\x20\x74\x68\x65\x20\x72\x65\x61\x64\x65\x72\x20\x69\x73\x20\x64\x6f\x6e\x65\x2c\x20\x74\x68\x65\x20\x63\x6f\x6e\x74\x65\x6e\x74\x20\x69\x73\x20\x69\x6e\x20\x72\x65\x61\x64\x65\x72\x2e\x72\x65\x73\x75\x6c\x74\x2e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x72\x65\x61\x64\x65\x72\x2e\x72\x65\x73\x75\x6c\x74\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x2e\x76\x61\x6c\x75\x65\x20\x3d\x72\x65\x61\x64\x65\x72\x2e\x72\x65\x73\x75\x6c\x74\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x61\x64\x65\x72\x2e\x72\x65\x61\x64\x41\x73\x54\x65\x78\x74\x28\x69\x6e\x70\x75\x74\x2e\x66\x69\x6c\x65\x73\x5b\x30\x5d\x29\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x53\x65\x6e\x64\x54\x6f\x42\x61\x63\x6b\x65\x6e\x64\x28\x29\x7b\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x22\x74\x68\x65\x72\x65\x20\x79\x6f\x75\x20\x61\x72\x65\x22\x29\x0a\x20\x20\x20\x20\x2f\x2f\x20\x76\x61\x6c\x69\x64\x61\x74\x69\x6e\x67\x20\x74\x79\x70\x65\x20\x66\x69\x6c\x65\x20\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x66\x69\x6c\x65\x4e\x61\x6d\x65\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x71\x75\x65\x72\x79\x53\x65\x6c\x65\x63\x74\x6f\x72\x28\x27\x23\x66\x69\x6c\x65\x2d\x75\x70\x6c\x6f\x61\x64\x20\x2e\x66\x69\x6c\x65\x2d\x6e\x61\x6d\x65\x27\x29\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x20\x3d\x20\x66\x69\x6c\x65\x4e\x61\x6d\x65\x2e\x74\x65\x78\x74\x43\x6f\x6e\x74\x65\x6e\x74\x2e\x6d\x61\x74\x63\x68\x28\x2f\x5c\x2e\x5b\x30\x2d\x39\x61\x2d\x7a\x5d\x2b\x24\x2f\x69\x29\x3b\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x64\x69\x73\x70\x6c\x61\x79\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x67\x65\x74\x45\x6c\x65\x6d\x65\x6e\x74\x42\x79\x49\x64\x28\x22\x69\x6e\x70\x75\x74\x2d\x64\x6f\x63\x22\x29\x3b\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x75\x72\x69\x20\x3d\x20\x22\x2f\x61\x70\x69\x2f\x63\x6f\x6e\x73\x75\x6c\x69\x7a\x65\x22\x3b\x0a\x0a\x20\x20\x20\x20\x2f\x2f\x20\x6c\x6f\x6f\x6b\x73\x20\x61\x77\x66\x75\x6c\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x29\x7b\x20\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x28\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x5b\x30\x5d\x20\x3d\x3d\x3d\x20\x27\x2e\x6a\x73\x6f\x6e\x27\x20\x7c\x7c\x20\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x5b\x30\x5d\x20\x3d\x3d\x3d\x20\x27\x2e\x68\x63\x6c\x27\x29\x20\x3f\x20\x50\x6f\x73\x74\x28\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x5b\x30\x5d\x2c\x20\x75\x72\x69\x2c\x20\x64\x69\x73\x70\x6c\x61\x79\x2e\x76\x61\x6c\x75\x65\x29\x3a\x20\x77\x69\x6e\x64\x6f\x77\x2e\x61\x6c\x65\x72\x74\x28\x22\x6e\x6f\x74\x20\x73\x75\x70\x70\x6f\x72\x74\x65\x64\x20\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x22\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x0a\x0a\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x61\x6c\x65\x72\x74\x28\x22\x53\x75\x70\x70\x6f\x72\x74\x65\x64\x20\x66\x69\x6c\x65\x20\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x73\x20\x61\x72\x65\x20\x6a\x73\x6f\x6e\x2f\x68\x63\x6c\x22\x29\x3b\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x50\x6f\x73\x74\x28\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x2c\x20\x75\x72\x69\x2c\x20\x70\x61\x79\x6c\x6f\x61\x64\x29\x7b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x78\x68\x72\x20\x3d\x20\x6e\x65\x77\x20\x58\x4d\x4c\x48\x74\x74\x70\x52\x65\x71\x75\x65\x73\x74\x28\x29\x3b\x0a\x20\x20\x20\x20\x78\x68\x72\x2e\x6f\x70\x65\x6e\x28\x22\x50\x4f\x53\x54\x22\x2c\x75\x72\x69\x2c\x20\x74\x72\x75\x65\x29\x3b\x0a\x20\x20\x20\x20\x78\x68\x72\x2e\x73\x65\x74\x52\x65\x71\x75\x65\x73\x74\x48\x65\x61\x64\x65\x72\x28\x22\x43\x6f\x6e\x74\x65\x6e\x74\x2d\x54\x79\x70\x65\x22\x2c\x20\x22\x61\x70\x70\x6c\x69\x63\x61\x74\x69\x6f\x6e\x2f\x6a\x73\x6f\x6e\x22\x29\x3b\x0a\x0a\x20\x20\x20\x20\x78\x68\x72\x2e\x6f\x6e\x72\x65\x61\x64\x79\x73\x74\x61\x74\x65\x63\x68\x61\x6e\x67\x65\x20\x3d\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x28\x29\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x78\x68\x72\x2e\x72\x65\x61\x64\x79\x53\x74\x61\x74\x65\x20\x3d\x3d\x3d\x20\x34\x29\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x22\x2e\x2e\x2e\x2e\x2e\x2e\x2e\x2e\x2e\x20\x73\x65\x6e\x74\x20\x5a\x5a\x5a\x22\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x78\x68\x72\x2e\x72\x65\x73\x70\x6f\x6e\x73\x65\x54\x65\x78\x74\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x64\x61\x74\x61\x20\x3d\x20\x4a\x53\x4f\x4e\x2e\x73\x74\x72\x69\x6e\x67\x69\x66\x79\x28\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x22\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x22\x3a\x20\x65\x78\x74\x65\x6e\x73\x69\x6f\x6e\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x22\x70\x61\x79\x6c\x6f\x61\x64\x22\x20\x3a\x20\x77\x69\x6e\x64\x6f\x77\x2e\x62\x74\x6f\x61\x28\x70\x61\x79\x6c\x6f\x61\x64\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x29\x3b\x0a\x0a\x20\x20\x20\x20\x78\x68\x72\x2e\x73\x65\x6e\x64\x28\x64\x61\x74\x61\x29\x3b\x0a\x7d")



func init() {
  err := CTX.Err()
  if err != nil {
		panic(err)
	}






  
  err = FS.Mkdir(CTX, "/js/", 0777)
  if err != nil && err != os.ErrExist {
    panic(err)
  }
  




  
  var f webdav.File
  

  

  
  

  f, err = FS.OpenFile(CTX, "/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileIndexHTML)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "/js/foos.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileJsFoosJs)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  


  Handler = &webdav.Handler{
    FileSystem: FS,
    LockSystem: webdav.NewMemLS(),
  }


}



// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
  path = hfs.Prefix + path


  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

  // If the buffer overflows, we will get bytes.ErrTooLarge.
  // Return that as an error. Any other panic remains.
  defer func() {
    e := recover()
    if e == nil {
      return
    }
    if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
      err = panicErr
    } else {
      panic(e)
    }
  }()
  _, err = buf.ReadFrom(f)
  return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
  if err != nil {
    return err
  }
  n, err := f.Write(data)
  if err == nil && n < len(data) {
    err = io.ErrShortWrite
  }
  if err1 := f.Close(); err == nil {
    err = err1
  }
  return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
    return nil, err
  }
  
  err = f.Close()
  if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}

