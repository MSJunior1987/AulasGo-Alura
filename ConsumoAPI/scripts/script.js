document.addEventListener("DOMContentLoaded", function () {
    carregarDadosDaAPI();
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
            <td>${personalidade.id}</td>
            <td>${personalidade.nome}</td>
            <td>${personalidade.historia}</td>
            <td>
                <button class="btn btn-primary editar-btn" data-id="${personalidade.id}" data-nome="${personalidade.nome}" data-historia="${personalidade.historia}">Editar</button>
            </td>
        `;
        tbody.appendChild(row);
    });

    adicionarEventoEditar();
}

function adicionarEventoEditar() {
    const editarButtons = document.querySelectorAll('.editar-btn');
    editarButtons.forEach(button => {
        button.addEventListener('click', function () {
            const id = this.getAttribute('data-id');
            const nome = this.getAttribute('data-nome');
            const historia = this.getAttribute('data-historia');
            
            // Redirecionar para a página de edição
            window.location.href = `editar.html?id=${id}&nome=${nome}&historia=${historia}`;
        });
    });
}
