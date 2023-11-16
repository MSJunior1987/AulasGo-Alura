document.addEventListener("DOMContentLoaded", function () {
    // Adicionar evento de envio do formulário
    const inserirForm = document.getElementById('inserirForm');
    inserirForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Obter os valores do formulário
        const nome = document.getElementById('nome').value;
        const historia = document.getElementById('historia').value;

        // Enviar os dados para a API usando o método POST
        inserirNovaPersonalidade(nome, historia);
    });

    // Adicionar evento ao botão de Voltar
    const btnVoltar = document.getElementById('btnVoltar');
    btnVoltar.addEventListener('click', function () {
        voltarParaPrincipal();
    });
});

function inserirNovaPersonalidade(nome, historia) {
    const url = 'http://localhost:8000/api/personalidades';
    const dadosInsercao = {
        nome: nome,
        historia: historia
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dadosInsercao)
    })
    .then(response => response.json())
    .then(data => {
        // Exibir pop-up de sucesso
        alert('Dados salvos com sucesso!');
        // Redirecionar de volta para a página principal após a inserção
        window.location.href = 'Principal.html';
    })
    .catch(error => console.error('Erro ao inserir nova personalidade:', error));
}

function voltarParaPrincipal() {
    window.location.href = 'Principal.html';
}
