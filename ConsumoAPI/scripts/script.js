document.addEventListener("DOMContentLoaded", function () {
    // Função para carregar os dados da API
    function carregarDadosDaAPI() {
        fetch('http://localhost:8000/api/personalidades')
            .then(response => response.json())
            .then(data => exibirDadosNaTabela(data))
            .catch(error => console.error('Erro ao obter dados da API:', error));
    }

    // Função para exibir os dados na tabela
    function exibirDadosNaTabela(data) {
        const tbody = document.querySelector('#personalidades-table tbody');

        data.forEach(personalidade => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${personalidade.id}</td>
                <td>${personalidade.nome}</td>
                <td>${personalidade.historia}</td>
            `;
            tbody.appendChild(row);
        });
    }

    // Carregar os dados da API assim que a página for carregada
    carregarDadosDaAPI();
});
