CREATE DATABASE IF NOT EXISTS root;

USE root;

CREATE TABLE IF NOT EXISTS alunos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    rg VARCHAR(20) NOT NULL
);



-- Inserir registros na tabela "alunos"
INSERT INTO alunos (nome, cpf, rg) VALUES
    ('João Silva', '123.456.789-01', '1234567-A'),
    ('Maria Oliveira', '987.654.321-09', '7654321-B'),
    ('Pedro Santos', '111.222.333-44', '9876543-C'),
    ('Ana Pereira', '555.666.777-88', '4567890-D'),
    ('Carlos Souza', '999.888.777-66', '3456789-E'),
    ('Lucia Costa', '333.444.555-66', '2345678-F'),
    ('Gabriel Lima', '777.888.999-11', '1234567-G'),
    ('Juliana Mendes', '111.222.333-44', '8765432-H'),
    ('Ricardo Silva', '555.666.777-88', '7654321-I'),
    ('Fernanda Oliveira', '999.888.777-66', '6543210-J');
