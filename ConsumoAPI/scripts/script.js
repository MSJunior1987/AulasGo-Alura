document.addEventListener("DOMContentLoaded", function () {
    carregarDadosDaAPI();

    // Adicionar evento ao botão de Inserir
    const btnInserir = document.getElementById('btnInserir');
    btnInserir.addEventListener('click', function () {
        // Redirecionar para a página de inserção
        window.location.href = 'Inserir.html';
    });
});

function carregarDadosDaAPI() {
    fetch('http://localhost:8000/api/personalidades')
        .then(response => response.json())
        .then(data => exibirDadosNaTabela(data))
        .catch(error => console.error('Erro ao obter dados da API:', error));
}

function exibirDadosNaTabela(data) {
    const tbody = document.querySelector('#personalidades-table tbody');
    tbody.innerHTML = ''; // Limpar o conteúdo da tabela antes de adicionar novos dados

    data.forEach(personalidade => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td style="display: none;">${personalidade.id}</td>
            <td>${personalidade.nome}</td>
            <td>${personalidade.historia}</td>
            <td>
                <div class="btn-group">
                    <button class="btn btn-primary editar-btn mr-1" data-id="${personalidade.id}">Editar</button>
                    <button class="btn btn-danger excluir-btn" data-id="${personalidade.id}">Excluir</button>
                </div>
            </td>
        `;
        tbody.appendChild(row);
    });

    adicionarEventos();
}

function adicionarEventos() {
    adicionarEventoEditar();
    adicionarEventoExcluir();
}

function adicionarEventoEditar() {
    const editarButtons = document.querySelectorAll('.editar-btn');
    editarButtons.forEach(button => {
        button.addEventListener('click', function () {
            const id = this.getAttribute('data-id');
            // Redirecionar para a página de edição apenas com o ID
            window.location.href = `editar.html?id=${id}`;
        });
    });
}

function adicionarEventoExcluir() {
    const excluirButtons = document.querySelectorAll('.excluir-btn');
    excluirButtons.forEach(button => {
        button.addEventListener('click', function () {
            const id = this.getAttribute('data-id');
            // Confirmar exclusão com uma popup
            if (confirm('Tem certeza de que deseja excluir este registro?')) {
                // Enviar a solicitação de exclusão para a API
                excluirPersonalidade(id);
            }
        });
    });
}

function excluirPersonalidade(id) {
    const url = `http://localhost:8000/api/personalidades/${id}`;

    fetch(url, {
        method: 'DELETE'
    })
    .then(response => {
        if (response.ok) {
            // Exibir mensagem de exclusão bem-sucedida
            alert('Registro excluído com sucesso!');
            // Recarregar a página após a exclusão
            window.location.reload();
        } else {
            console.error('Erro ao excluir personalidade:', response.status);
        }
    })
    .catch(error => console.error('Erro ao excluir personalidade:', error));
}
