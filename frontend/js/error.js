export function showError(message) {
    const toast = document.getElementById('toast');
    toast.innerText = message;
    toast.classList.add('show');
    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}

export async function handleErrors(response) {
    if (!response.ok) {
        const errorData = await response.json();
        const errorCode = errorData.err_code;
        let errorMessage;

        switch (errorCode) {
            case "ErrFileMetadataAlreadyExists":
                errorMessage = 'Arquivo já enviado.';
                break;
            case "ErrScanningFile":
                errorMessage = 'Ocorreu um erro ao ler o arquivo.';
                break;
            case "ErrParsingParsingLine":
                errorMessage = 'Ocorreu um erro ao interpretar as linhas do arquivo. Cheque novamente a formatação, espaçamento e posicionamento dos dados.';
                break;
            case "ErrInvalidEmailOrPassword":
                errorMessage = 'Usuário não existe ou credenciais estão incorretas';
                break;
            case "ErrOnlyTxtAreAccepted":
                errorMessage = 'Apenas arquivos .txt são aceitos';
            case "ErrProvideAFileOrEnsureNotEmpty":
                errorMessage = 'Arquivo vazio. Envie um arquivo com conteúdo.';
                break;
            default:
                errorMessage = "Ocorreu um erro no servidor.";
                break;
        }

        throw new Error(errorMessage);
    }

    return response;
}

export function handleUnauthorized(response) {
    if (response.status === 401) {
        localStorage.removeItem('jwt');

        // Redirect the user to the login page
        window.location.href = '/frontend/login.html';
    }

    return response;
}