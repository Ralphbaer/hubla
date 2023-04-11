import { getJwtToken } from './jwt.js';
import { showError, handleErrors, handleUnauthorized } from './error.js';

async function validateAndUploadFile(file, filename) {
    if (!file || file.type !== 'text/plain') {
        showError('Invalid file type. Please upload a .txt file.');
        return;
    }

    const jwtToken = getJwtToken();

    try {
        const formData = new FormData();
        formData.append('file', file, filename);

        const response = await fetch('http://localhost:3000/api/v1/transaction/file-transactions', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/octet-stream',
                'Content-Disposition': `attachment; filename="${filename}"`,
                'Content-Length': file.size,
                'Authorization': `Bearer ${jwtToken}`,
            },
            body: file,
        }).then(handleUnauthorized).then(handleErrors);

        if (response.status === 201) {
            const { id } = await response.json()
            return id;
        } else {
            showError('Unexpected server response. Please try again later.');
        }
    } catch (error) {
        showError(error.message);
    }
}

let uploadedFile = null;

function handleFileUpload(event) {
    uploadedFile = event.target.files[0];

    if (uploadedFile) {
        document.querySelector('.upload-label').textContent = `Selected File: ${uploadedFile.name}`;
    }
}

async function confirmFileUpload() {
    if (!uploadedFile) {
        showError('Nenhum arquivo foi selecionado. Por favor, envie um arquivo antes de confirmar.');
        return;
    }

    const fileName = uploadedFile.name;
    const id = await validateAndUploadFile(uploadedFile, fileName);
    if (id) {
        window.location.href = `http://127.0.0.1:5500/frontend/file-transaction.html?id=${id}`;
    }
}

document.addEventListener('DOMContentLoaded', () => {
    document.querySelector('#upload').addEventListener('change', handleFileUpload);
    document.querySelector('button').addEventListener('click', confirmFileUpload);
});
