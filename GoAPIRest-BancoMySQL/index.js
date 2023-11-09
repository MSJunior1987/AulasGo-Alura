// index.js

const express = require('express');
const bodyParser = require('body-parser'); // Exemplo de uma biblioteca adicional para processar dados do corpo das requisições
const app = express();
const port = process.env.PORT || 3000;

// Middleware para processar JSON e formulários
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// Rota de exemplo
app.get('/', (req, res) => {
  res.send('Bem-vindo ao seu servidor Node.js!');
});

// Rota com parâmetros da requisição
app.get('/api/:id', (req, res) => {
  const id = req.params.id;
  res.send(`Recebido o parâmetro ID: ${id}`);
});

// Rota de exemplo para lidar com POST
app.post('/api/enviar', (req, res) => {
  const data = req.body;
  res.json({ mensagem: 'Dados recebidos com sucesso!', dados: data });
});

// Inicialização do servidor
app.listen(port, () => {
  console.log(`Servidor Node.js está rodando na porta ${port}`);
});