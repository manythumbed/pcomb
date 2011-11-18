include $(GOROOT)/src/Make.inc

TARG=pcomb
GOFILES=\
	state.go\
	reader.go\
	combinators.go\
	types.go\

include $(GOROOT)/src/Make.pkg
