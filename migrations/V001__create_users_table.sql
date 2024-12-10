-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,                
    username VARCHAR(128) NOT NULL,       
    email VARCHAR(128) UNIQUE NOT NULL,  
    created_at TIMESTAMP DEFAULT NOW(),   
    updated_at TIMESTAMP                  
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);