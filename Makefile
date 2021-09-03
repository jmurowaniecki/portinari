#
# 🜏
#
# For further information see `README.md`.
#

DESTDIR ?= /usr/local/bin

CMD = "\\e[1m%-10s\\e[0m%s\n"
STR = "\\e[0;2;3m%s\\e[0m\n"
HLP = sed -E 's/(`.*`)/\\e[1m\1\\e[0m/'

DEFAULT: help

install: dist/portinari # Installs application - see readme for details.
	cp -Rf ./man/* /usr/local/man
	install --mode 755 $< $(DESTDIR)

build: # Process book covers.
	@go build -o dist/portinari .

clean: # Remove temporary files.
	@rm -Rf dist

#
help: # Shows this help.
	@\
	echo -e """""""""""""""""""""""  \
	$$(awk 'BEGIN {   FS=":.*?#"   } \
	/^(\w+:.*|)#/ {                  \
	gsub("^( : |)#( |)", """""""" ); \
	LEN=length($$2); COND=(LEN < 1); \
	FORMAT=(COND ? $(STR) : $(CMD)); \
	printf(FORMAT, $$1, """"""$$2 ); \
	}' $(MAKEFILE_LIST) | ($(HLP)))"


#
%:
	@:
