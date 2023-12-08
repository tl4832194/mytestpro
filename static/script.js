function addPerson() {
    var form = document.getElementById("personForm");
    var formData = new FormData(form);

    fetch("/add", {
        method: "POST",
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            form.reset();
        })
        .catch(error => console.error("Error:", error));
}

function queryData() {
    fetch("/query")
        .then(response => response.json())
        .then(data => {
            var resultDiv = document.getElementById("result");
            resultDiv.innerHTML = "";

            if (data.length === 0) {
                resultDiv.innerHTML = "No data available.";
                return;
            }

            var table = document.createElement("table");
            table.border = "1";

            var headerRow = table.insertRow(0);
            var idHeader = headerRow.insertCell(0);
            var nameHeader = headerRow.insertCell(1);
            var ageHeader = headerRow.insertCell(2);
            var addressHeader = headerRow.insertCell(3);

            idHeader.innerHTML = "<b>ID</b>";
            nameHeader.innerHTML = "<b>Name</b>";
            ageHeader.innerHTML = "<b>Age</b>";
            addressHeader.innerHTML = "<b>Address</b>";

            for (var i = 0; i < data.length; i++) {
                var row = table.insertRow(i + 1);
                var idCell = row.insertCell(0);
                var nameCell = row.insertCell(1);
                var ageCell = row.insertCell(2);
                var addressCell = row.insertCell(3);

                idCell.innerHTML = data[i].ID;
                nameCell.innerHTML = data[i].Name;
                ageCell.innerHTML = data[i].Age;
                addressCell.innerHTML = data[i].Address;
            }

            resultDiv.appendChild(table);
        })
        .catch(error => console.error("Error:", error));
}
