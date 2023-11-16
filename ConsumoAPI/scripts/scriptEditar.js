document.addEventListener("DOMContentLoaded", function () {
    // Extrair parâmetros da URL
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    // Preencher o campo do formulário com o ID
    document.getElementById('id').value = id;

    // Carregar dados da API usando o ID
    carregarDadosDaAPI(id);

    // Adicionar evento de envio do formulário
    const editarForm = document.getElementById('editarForm');
    editarForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Obter os valores atualizados do formulário
        const novoNome = document.getElementById('nome').value;
        const novaHistoria = document.getElementById('historia').value;

        // Enviar os dados atualizados para a API usando o método PUT
        atualizarPersonalidade(id, novoNome, novaHistoria);
    });

    // Adicionar evento ao botão de Voltar
    const btnVoltar = document.getElementById('btnVoltar');
    btnVoltar.addEventListener('click', function () {
        voltarParaPrincipal();
    });
});

function carregarDadosDaAPI(id) {
    const url = `http://localhost:8000/api/personalidades/${id}`;

    fetch(url)
        .then(response => response.json())
        .then(data => {
            // Preencher os campos do formulário com os dados da personalidade
            document.getElementById('nome').value = data.nome;
            document.getElementById('historia').value = data.historia;
        })
        .catch(error => console.error('Erro ao obter dados da API:', error));
}

function atualizarPersonalidade(id, nome, historia) {
    const url = `http://localhost:8000/api/personalidades/${id}`;
    const dadosAtualizados = {
        nome: nome,
        historia: historia
    };

    fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dadosAtualizados)
    })
    .then(response => response.json())
    .then(data => {
        // Exibir pop-up de sucesso
        alert('Dados salvos com sucesso!');
        // Redirecionar de volta para a página principal após a atualização
        window.location.href = 'Principal.html';
    })
    .catch(error => console.error('Erro ao atualizar dados:', error));
}

function voltarParaPrincipal() {
    window.location.href = 'Principal.html';
}
