document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById('inserirForm');

    form.addEventListener('submit', function (event) {
        event.preventDefault();

        const nome = document.getElementById('nome').value;
        const historia = document.getElementById('historia').value;

        // Enviar dados para a API via método POST
        fetch('http://localhost:8000/api/personalidades', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                nome: nome,
                historia: historia,
            }),
        })
        .then(response => response.json())
        .then(data => {
            console.log('Dados inseridos:', data);
            // Limpar o formulário após a inserção bem-sucedida, se necessário
            form.reset();
        })
        .catch(error => console.error('Erro ao inserir dados:', error));
    });
});
