    CREATE TABLE IF NOT EXISTS users
    (
        id SERIAL PRIMARY KEY,
        name varchar(255) not null,
        username varchar(255) not null unique,
        password_hash varchar(255) not null
    );

    CREATE TABLE IF NOT EXISTS todo_authors
    (
        id SERIAL PRIMARY KEY,
        fullname varchar(255) not null,
        pseudonym varchar(255),
        spec varchar(255)
    );

    CREATE TABLE IF NOT EXISTS users_authors
    (
        id SERIAL PRIMARY KEY,
        user_id integer not null,
        author_id integer not null,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (author_id) REFERENCES todo_authors(id) ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS todo_books
    (
        id SERIAL PRIMARY KEY,
        title varchar(255) not null,
        genre varchar(255),
        isbn int not null
    );

    CREATE TABLE IF NOT EXISTS authors_books
    (
        id SERIAL PRIMARY KEY,
        book_id integer not null,
        author_id integer not null,
        FOREIGN KEY (book_id) REFERENCES todo_books(id) ON DELETE CASCADE,
        FOREIGN KEY (author_id) REFERENCES todo_authors(id) ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS users_books
    (
        id SERIAL PRIMARY KEY,
        user_id integer not null,
        book_id integer not null,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (book_id) REFERENCES todo_books(id) ON DELETE CASCADE
    );
