-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       username VARCHAR(100) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       first_name VARCHAR(100),
                       last_name VARCHAR(100),
                       role VARCHAR(10) CHECK(role IN ('admin', 'user', 'tester')) NOT NULL DEFAULT 'user',
                       status VARCHAR(10) CHECK(status IN ('active', 'block', 'deleted')) NOT NULL DEFAULT 'active',
                       avatar_url VARCHAR(255),
                       phone_number VARCHAR(20),
                       gender VARCHAR(10),
                       birthdate DATE,
                       UNIQUE (username, email)
);
-- // todo  возможно надо определить platform в отдельной таблице чтобы не происходило дублирование platform VARCHAR(50) CHECK(platform IN ('google', 'vk', 'telegram', 'contact')) NOT NULL, --вот это
--
CREATE TABLE social_profiles (
                                 id INTEGER PRIMARY KEY AUTOINCREMENT,
                                 user_id INTEGER NOT NULL,
                                 platform VARCHAR(50) CHECK(platform IN ('google', 'vk', 'telegram', 'contact')) NOT NULL, --вот это
                                 profile_url VARCHAR(255) NOT NULL,
                                 external_id VARCHAR(50) NOT NULL,
                                 params TEXT,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 last_contact_updated_at TIMESTAMP,
                                 FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE settings (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          user_id INTEGER,
                          theme VARCHAR(10) CHECK(theme IN ('light', 'dark', 'system'))NOT NULL DEFAULT 'light', -- Поле для хранения темы (например, 'light', 'dark')
                          language VARCHAR(10) NOT NULL DEFAULT 'en', -- Поле для хранения языка (например, 'en', 'ru')
                          auto_update BOOLEAN NOT NULL DEFAULT 1, -- Поле для хранения флага автообновления (1 - включено, 0 - выключено)
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE friends (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         ownerID INTEGER NOT NULL,
                         name VARCHAR(50) NOT NULL,
                         alternate_names TEXT,
                         birthdate DATE,
                         platform VARCHAR(50) CHECK(platform IN ('vk', 'telegram', 'contact', 'google')) NOT NULL,
                         phone_number VARCHAR(20),
                         alternate_phone_numbers TEXT,
                         avatar_url VARCHAR(255),
                         FOREIGN KEY (ownerID) REFERENCES users(id)
);

CREATE TABLE friend_notes (
                              note_id INTEGER PRIMARY KEY AUTOINCREMENT,
                              friend_id INTEGER NOT NULL,
                              title VARCHAR(100),
                              category VARCHAR(50),
                              content TEXT,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE friend_tags (
                             id INTEGER PRIMARY KEY AUTOINCREMENT,
                             friend_id INTEGER NOT NULL,
                             tag VARCHAR(30) NOT NULL,
                             platform VARCHAR(50) CHECK(platform IN ('vk', 'telegram', 'contact', 'google')) NOT NULL,
                             FOREIGN KEY (friend_id) REFERENCES friends(id)
);
CREATE TABLE tokens (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        user_id INTEGER,
                        token VARCHAR(50) NOT NULL,
                        created_at DATETIME NOT NULL,
                        is_active BOOLEAN NOT NULL DEFAULT TRUE,
                        FOREIGN KEY (user_id) REFERENCES users(id),
                        UNIQUE (user_id, token)
);

CREATE TABLE conflicts (
                           id INTEGER PRIMARY KEY AUTOINCREMENT,
                           user_id INTEGER NOT NULL,
                           left_friend_id INTEGER NOT NULL,
                           right_friend_id INTEGER NOT NULL,
                           is_active BOOLEAN NOT NULL DEFAULT TRUE,
                           FOREIGN KEY (user_id) REFERENCES users(id),
                           FOREIGN KEY (left_friend_id) REFERENCES friends(id),
                           FOREIGN KEY (right_friend_id) REFERENCES friends(id)
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friend_tags;
DROP TABLE IF EXISTS friend_notes;
DROP TABLE IF EXISTS friends;
DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS social_profiles;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
