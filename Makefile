include $(GOROOT)/src/Make.inc

TARG=pcomb
GOFILES=\
	reader.go\
	types.go\

include $(GOROOT)/src/Make.pkg
