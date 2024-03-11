package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ChatService -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i MessageService -o ./mocks/ -s "_minimock.go"
