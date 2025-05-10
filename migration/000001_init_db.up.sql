CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     identity_number VARCHAR,
                                     first_name VARCHAR,
                                     last_name VARCHAR,
                                     email VARCHAR UNIQUE NOT NULL,
                                     password_hash VARCHAR NOT NULL,
                                     role VARCHAR NOT NULL DEFAULT 'user',
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS refresh_tokens (
                                              id SERIAL PRIMARY KEY,
                                              user_id INT NOT NULL,
                                              token TEXT NOT NULL,
                                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                              FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );