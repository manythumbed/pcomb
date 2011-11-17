include $(GOROOT)/src/Make.inc

TARG=pcomb
GOFILES=\
	reader.go\
	combinators.go\
	types.go\

include $(GOROOT)/src/Make.pkg
