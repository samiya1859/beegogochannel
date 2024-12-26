function addToFavorites(imageID) {
    // Assuming 'my-user-1234' is the user ID; you may want to change this dynamically
    const subID = "my-user-1234";  

    // Prepare the data to be sent to the server
    const data = {
        image_id: imageID,
        sub_id: subID
    };
    console.log(data)
    // Make the POST request to your backend
    fetch('/fav', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())  // Parse the JSON response
    .then(data => {
        // Check if the response status is success
        if (data.status === 'success') {
            alert("Image added to favorites!");
            console.log("SUCCESSS!!!")
        } else {
            alert("Failed to add image to favorites: " + data.message);
        }
    })
    .catch(error => {
        console.error("Error adding to favorites:", error);
        alert("Error adding image to favorites.");
    });
}



// for getting all favs
// Function to fetch favorites and display them
// function fetchFavorites() {
//     fetch('/fav/getall')  
//         .then(response => response.json())
//         .then(favorites => {
//             const favContainer = document.getElementById("fav-image-container");
//             console.log(favContainer)
//             favContainer.innerHTML = ""; 

//             // Check if there are favorites to display
//             if (favorites && favorites.length > 0) {
//                 favorites.forEach(fav => {
//                     const favItem = document.createElement("div");
//                     favItem.classList.add("fav-item");
//                     favItem.innerHTML = `
//                         <img src="${fav.image.url}" alt="Favorite Image"/>
//                     `;
//                     favContainer.appendChild(favItem);
//                 });
//             } else {
//                 favContainer.innerHTML = "<p>No favorites found.</p>";
//             }

//             // Display the container once the data is loaded
//             document.getElementById("fav-container").style.display = "block";
//         })
//         .catch(err => console.error('Error fetching favorites:', err));
// }

// Function to switch between grid and scroll layout
function switchLayout(layout) {
    const favContainer = document.getElementById("fav-image-container");
    
    if (layout === "grid") {
        favContainer.classList.remove("scroll-container");
        favContainer.classList.add("grid-container");
    } else {
        favContainer.classList.remove("grid-container");
        favContainer.classList.add("scroll-container");
    }
}



