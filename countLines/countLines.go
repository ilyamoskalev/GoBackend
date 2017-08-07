package main

import (
    "fmt"
    "bytes"
    "os"
    "io/ioutil"
)

func countLines( fileNames []string ) {

	totalSize := 0

	for _, value := range fileNames{

		buffer, err := ioutil.ReadFile( value )

    		if err != nil {

    			fmt.Println( "No such file or directory." )

    		} else {

    			currentSize := bytes.Count( buffer, []byte{ '\n' } )
    			totalSize += currentSize
    			fmt.Println( value, currentSize )
    	
    		}
	}

	if len( fileNames ) > 1 {

		fmt.Println( "total",  totalSize )

	}
}

func main() {

	fileNames := os.Args[ 1: ]

	if len( fileNames ) == 0 {

		fmt.Println( "Enter the filename" )

	} else {

		countLines( fileNames )

	}

}
