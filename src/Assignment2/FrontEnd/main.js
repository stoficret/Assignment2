async function displayAPIData() {
    //get API data
    const response = await fetch("http://localhost:8011/students");
    data = await response.json();
  
    //generate HTML code
    const tableData = data
      .map(function (value) {
        return `<tr>
              <td>${value.StudentID}</td>
              <td>${value.Name}</td>
              <td>${value.DateofBirth}</td>
              <td>${value.Address}</td>
              <td>${value.PhoneNumber}</td>
          </tr>`;
      })
      .join("");
  
    //set tableBody to new HTML code
    const tabBody = document.querySelector("#tabBody");
    tabBody.innerHTML = tableData;
}
  
  displayAPIData();
  