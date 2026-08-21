const nomeInput = document.getElementById('nome');
const emailInput = document.getElementById('email');
const senhaInput = document.getElementById('senha');

async function login(){
    const body = {
        email: document.getElementById('email').value,
        senha: document.getElementById('senha').value
    }

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        });

        if (!response.ok) {
            throw new Error('Erro na solicitação: ' + response.status);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

async function cadastro(){
    const body = {
        nome: document.getElementById('nome').value,
        email: document.getElementById('email').value,
        senha: document.getElementById('senha').value
    }

    try {
        const response = await fetch('http://localhost:8080/cadastro', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        });
        if (!response.ok) {
            throw new Error('Erro na solicitação: ' + response.status);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

async function buscarUsuarios(){
    try {
        const response = await fetch('http://localhost:8080/usuarios', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Erro na solicitação: ' + response.status);
        }

        const data = await response.json();
        console.log(data);

        carregarGrid(data);
    } catch (error) {
        console.error('Error:', error);
    }
}

carregarGrid = (data) =>  {
    const tableBody = document.getElementById('tableBody');

    tableBody.innerHTML = '';

    data.forEach((item) => {
        const rowTemplate = `
            <tr>
                <td>${item.id}</td>
                <td>${item.nome}</td>
                <td>${item.email}</td>
            </tr>
        `;
        tableBody.insertAdjacentHTML('beforeend', rowTemplate);
    });
}