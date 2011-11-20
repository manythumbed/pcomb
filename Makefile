include $(GOROOT)/src/Make.inc

TARG=pcomb
GOFILES=\
	state.go\
	error.go\
	combinators.go\
	types.go\

include $(GOROOT)/src/Make.pkg
