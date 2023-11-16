NAME = "ft_ality"
UNAME := $(shell uname -s)

$(NAME):
	@NAME=$(NAME) bash ./scripts/build.sh

clean:
	rm -f $(NAME)
