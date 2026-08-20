const emailInput = document.getElementById('email');
const passwordInput = document.getElementById('password');

async function postData(){
    const email = emailInput.value;
    const password = passwordInput.value;

    const body = {
        email: email,
        password: password
    }

    try {
        const response = await fetch('http://localhost:3000/login', {
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

async function getData(){
    try {
        const response = await fetch('http://localhost:3000/data', {
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

    const rowTemplate = `
        <tr>
            <td>${name}</td>
            <td>${email}</td>
        </tr>
    `;

    data.forEach((item) => {
        tableBody.insertAdjacentHTML('beforeend', rowTemplate);
    });
}