async function validateAndUploadFile(file) {
    if (!file || file.type !== "text/plain") {
        alert("Invalid file type. Please upload a .txt file.");
        return;
    }

    try {
        console.log("TESTE")
        const formData = new FormData();
        formData.append("file", file);

        const response = await fetch("http://localhost:3000/transactions/upload", {
            method: "POST",
            body: formData
        });

        if (!response.ok) {
            throw new Error(`HTTP error ${response.status}`);
        }

        const data = await response.json();

        return data;
    } catch (error) {
        console.error("Error uploading file:", error);
        alert("Error uploading file. Please try again.");
    }
}

function populateTable(data) {
    const tableBody = document.querySelector("table tbody");
    tableBody.innerHTML = "";

    data.forEach((item) => {
        const row = document.createElement("tr");

        Object.values(item).forEach((value) => {
            const cell = document.createElement("td");
            cell.textContent = value;
            row.appendChild(cell);
        });

        tableBody.appendChild(row);
    });
}

let uploadedFile = null;

function onFileUpload(event) {
    uploadedFile = event.target.files[0];

    if (uploadedFile) {
        document.querySelector('.table-container').style.display = 'block';
    }
}

async function onFileConfirm() {
    if (!uploadedFile) {
        alert("No file has been uploaded. Please upload a file before confirming.");
        return;
    }

    const data = await validateAndUploadFile(uploadedFile);

    if (data) {
        populateTable(data);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    document.querySelector('#upload').addEventListener('change', onFileUpload);
    document.querySelector('button').addEventListener('click', onFileConfirm);
});
