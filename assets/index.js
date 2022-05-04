document.body.style.backgroundColor = "#D6CFC7";
document.body.style.backgroundColor = "#787276";
document.body.style.color = "#351E10";
birdTable = document.querySelector("table")

/*
Use the browsers `fetch` API to make a GET call to /bird
We expect the response to be a JSON list of birds, of the
form :
[
    {"species":"...","description":"..."},
    {"species":"...","description":"..."}
]
*/
fetch("/bird")
    .then(response => response.json())
    .then(birdList => {
    //Once we fetch the list, we iterate over it
    birdList.forEach(bird => {
        // Create the table row
        row = document.createElement("tr")

        // Create the table data elements for the species and
                // description columns
        species = document.createElement("td")
        species.innerHTML = bird.species
        description = document.createElement("td")
        description.innerHTML = bird.description

        // Add the data elements to the row
        row.appendChild(species)
        row.appendChild(description)
        // Finally, add the row element to the table itself
        birdTable.appendChild(row)
    })
})