package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ChatUserRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ChatRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i MessageRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i LogRepository -o ./mocks/ -s "_minimock.go"
