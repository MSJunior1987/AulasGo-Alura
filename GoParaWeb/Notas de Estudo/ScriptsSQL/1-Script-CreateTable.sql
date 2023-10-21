-- Crie o banco de dados alura_loja
CREATE DATABASE alura_loja;

-- Use o banco de dados alura_loja
USE alura_loja;

-- Crie uma tabela para armazenar informações dos produtos
CREATE TABLE produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    preco DECIMAL(10, 2) NOT NULL,
    quantidade INT 
);
