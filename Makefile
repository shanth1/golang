# c - chapter (1 -> 01, 11 -> 11)
# f - folder
# a - args
# Example: make run c=1 f=echo a=test

run:
	@if ! echo $(c) | grep -Eq '^[0-9]+$$'; then \
		echo "Error: $(c) is not a number"; \
     	exit 1; \
    fi; \
	if [ $(c) -lt 9 ]; then \
		NEWNUM=0$(c); \
	else \
    	NEWNUM=$(c); \
    fi; \
    go run chapter_$${NEWNUM}/$(f)/main.go $(a)
