NAME = "ft_ality"

$(NAME) :
	go build -o $(NAME) src/main.go

clean :
	rm -f $(NAME)
