-- +goose Up
CREATE TABLE IF NOT EXISTS chats(
    chat_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    chat_title TEXT NOT NULL,
    
    CONSTRAINT pk_chats_chat_id PRIMARY KEY(chat_id)
);

CREATE TABLE IF NOT EXISTS users(
    user_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    name TEXT NOT NULL,
    
    CONSTRAINT pk_users_user_id PRIMARY KEY(user_id)
);

CREATE TABLE IF NOT EXISTS users_chats(
    user_id INT NOT NULL,
    chat_id INT NOT NULL,

    CONSTRAINT fk_users_chats_user_id FOREIGN KEY(user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_users_chats_chat_id FOREIGN KEY(chat_id) REFERENCES chats (chat_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages(
    message_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    chat_id INT NOT NULL,
    user_id INT NOT NULL,
    text_message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_messages_message_id PRIMARY KEY(message_id),
    CONSTRAINT fk_chats_messages_chat_id FOREIGN KEY(chat_id) REFERENCES chats (chat_id) ON DELETE CASCADE,
    CONSTRAINT fk_chats_messages_user_id FOREIGN KEY(user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS users_chats;
DROP TABLE IF EXISTS chats_messages;