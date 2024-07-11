.PHONY: run gen check_args

# c - chapter (1 -> 01, 11 -> 11)
# f - folder
# a - args

# Example: make run c=1 f=echo a=test
run:
	@if [ $(c) -lt 9 ]; then \
		NEWNUM=0$(c); \
	else \
    	NEWNUM=$(c); \
    fi; \
    go run chapter_$${NEWNUM}/$(f)/main.go $(a)

# Example: make get c=1 f=template
gen: check_args
	@if [ $(c) -lt 9 ]; then \
		NEWNUM=0$(c); \
	else \
    	NEWNUM=$(c); \
    fi; \
	mkdir -p chapter_$${NEWNUM}/$(f); \
	cp template.go chapter_$${NEWNUM}/$(f)/main.go; \

check_args:
	@if [ -z "$(f)" ]; then \
		echo "Error: f is a mandatory argument"; \
		exit 1; \
	fi
	@if [ -z "$(c)" ]; then \
		echo "Error: c is a mandatory argument"; \
		exit 1; \
	fi
	@if ! echo $(c) | grep -Eq '^[0-9]+$$'; then \
		echo "Error: $(c) is not a number"; \
    	exit 1; \
    fi
