CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  email VARCHAR(200) UNIQUE NOT NULL,
  phno VARCHAR(20),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


INSERT INTO users (name, email, phno) VALUES ('Pranav Reddy', 'pranav@example.com', '9876543210');
