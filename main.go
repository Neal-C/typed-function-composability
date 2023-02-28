//lint:file-ignore ST1006 because it's consistent

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


type Server struct{
	filenameTransformFn TransformFn
}

type TransformFn func(string) string;

func (self *Server) handleRequest(filename string) error {

	newFilename := self.filenameTransformFn(filename);

	fmt.Println("new filename :", newFilename);

	return nil;
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename));

	newFilename := hex.EncodeToString(hash[:]);

	return newFilename;
}

func PrefixFilename(prefix string) TransformFn {
	
	return func(filename string) string {
		return prefix + filename;
	}
	
}

//hashed filenames can be split, to do a content-addressable folder structure

func main(){
	server := &Server{
		filenameTransformFn: PrefixFilename("Peter_"),
	};

	server.handleRequest("cool_nude.jpg");

	server2 := &Server{
		filenameTransformFn: hashFilename,
	};

	server2.handleRequest("cool_nude.jpg");
}



//!bad
// func (self *Server) handleRequest(filename string) error {

// 	hash := sha256.Sum256([]byte(filename));

// 	newFilename := hex.EncodeToString(hash[:]);

// 	fmt.Println("new filename :", newFilename);

// 	return nil;
// }

// if the thing you want to do, does not take any state: an interface doesn't make sense