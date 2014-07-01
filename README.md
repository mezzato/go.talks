# Viewing the slides and finding the source code #

Install Go and then the Go present tool:    

    go get code.google.com/p/go.tools/cmd/present
    
## Slides ##

Make sure $GOPATH\bin is added to your PATH environment variable.

Open a command prompt in the directory where the repository has been checked out and launch the present tool:

    cd $GOPATH/src/github.com/mezzato/go.talks

    present

Open a browser at the specified address and choose the slide file. Normally:

    http://127.0.0.1:3999

## Source code ##

The source code to the slide resides in a subfolder named after the slide.
