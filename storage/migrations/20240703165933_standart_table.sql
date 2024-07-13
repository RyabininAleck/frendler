-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       username TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       first_name TEXT,
                       last_name TEXT,
                       role TEXT CHECK(role IN ('admin', 'user', 'tester')) NOT NULL DEFAULT 'user',
                       status TEXT CHECK(status IN ('active', 'block', 'deleted')) NOT NULL DEFAULT 'active',
                       avatar_url TEXT,
                       phone_number TEXT,
                       gender TEXT,
                       birthdate DATE,
                       UNIQUE (username, email)
);

CREATE TABLE social_profiles (
                                 id INTEGER PRIMARY KEY AUTOINCREMENT,
                                 user_id INTEGER NOT NULL,
                                 platform TEXT CHECK(platform IN ('google', 'vk', 'telegram', 'contact')) NOT NULL,
                                 profile_url TEXT NOT NULL,
                                 external_id TEXT NOT NULL,
                                 params TEXT,
                                 token TEXT NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 last_contact_updated_at TIMESTAMP,
                                 FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE settings (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          user_id INTEGER,
                          theme TEXT CHECK(theme IN ('light', 'dark', 'system')) NOT NULL DEFAULT 'light',
                          language TEXT NOT NULL DEFAULT 'en',
                          auto_update BOOLEAN NOT NULL DEFAULT 1,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE friends (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         ownerID INTEGER NOT NULL,
                         given_name TEXT NOT NULL,
                         family_name TEXT NOT NULL,
                         display_name TEXT NOT NULL,
                         birthdate DATE,
                         Organizations TEXT,
                         platform TEXT CHECK(platform IN ('vk', 'telegram', 'contact', 'google')) NOT NULL,
                         avatar_url TEXT,
                         FOREIGN KEY (ownerID) REFERENCES users(id)
);

CREATE TABLE friend_phone_numbers (
                                      id INTEGER PRIMARY KEY AUTOINCREMENT,
                                      friend_id INTEGER NOT NULL,
                                      phone_number TEXT NOT NULL,
                                      is_primary BOOLEAN NOT NULL DEFAULT FALSE,
                                      number_type TEXT CHECK(number_type IN ('mobile', 'work', 'home', 'other')) NOT NULL,
                                      FOREIGN KEY (friend_id) REFERENCES friends(id),
                                      CONSTRAINT unique_primary_number UNIQUE (friend_id, is_primary)
);

CREATE TABLE friend_emails (
                               id INTEGER PRIMARY KEY AUTOINCREMENT,
                               friend_id INTEGER NOT NULL,
                               email TEXT NOT NULL,
                               email_type TEXT CHECK(email_type IN ('work', 'home', 'other')) NOT NULL,
                               FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE friend_urls (
                             id INTEGER PRIMARY KEY AUTOINCREMENT,
                             friend_id INTEGER NOT NULL,
                             url TEXT NOT NULL,
                             url_description TEXT NOT NULL,
                             url_type TEXT CHECK(url_type IN ('site', 'other')) NOT NULL,
                             FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE friend_addresses (
                                  id INTEGER PRIMARY KEY AUTOINCREMENT,
                                  friend_id INTEGER NOT NULL,
                                  addresses TEXT NOT NULL,
                                  is_primary BOOLEAN NOT NULL DEFAULT FALSE,
                                  address_type TEXT CHECK(address_type IN ('work', 'home', 'other')) NOT NULL,
                                  country TEXT,
                                  country_code TEXT,
                                  FOREIGN KEY (friend_id) REFERENCES friends(id),
                                  CONSTRAINT unique_primary_number UNIQUE (friend_id, is_primary)
);

CREATE TABLE friend_notes (
                              note_id INTEGER PRIMARY KEY AUTOINCREMENT,
                              friend_id INTEGER NOT NULL,
                              title TEXT,
                              category TEXT,
                              content TEXT,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE friend_events (
                               note_id INTEGER PRIMARY KEY AUTOINCREMENT,
                               friend_id INTEGER NOT NULL,
                               title TEXT,
                               description TEXT,
                               category TEXT,
                               date TIMESTAMP NOT NULL,
                               FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE friend_tags (
                             id INTEGER PRIMARY KEY AUTOINCREMENT,
                             friend_id INTEGER NOT NULL,
                             tag TEXT NOT NULL,
                             platform TEXT CHECK(platform IN ('vk', 'telegram', 'contact', 'google')) NOT NULL,
                             FOREIGN KEY (friend_id) REFERENCES friends(id)
);

CREATE TABLE tokens (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        user_id INTEGER,
                        token TEXT NOT NULL,
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
