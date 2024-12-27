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
    .then(response => response.json())  
    .then(data => {
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



// Function to toggle between Grid and Scroll layout
function switchLayout(layout) {
    const favImageContainer = document.getElementById('fav-image-container');
    favImageContainer.classList.remove('grid-container', 'scroll-container');
    if (layout === 'grid') {
        favImageContainer.classList.add('grid-container');
    } else if (layout === 'scroll') {
        favImageContainer.classList.add('scroll-container');
    }
}

// Function to fetch and load favorite images
function loadFavorites() {
    console.log("Loading favorites...");  // This will show in the console if called

    fetch('https://api.thecatapi.com/v1/favourites', {
        method: 'GET',
        headers: {
            'x-api-key': 'live_aEIcWNMCOCKINdpArjjNnm54ivWft2t1E2ZiBWTPsjuhWXPjq5Ih8NhFzZUqzwHW' // Replace with your actual API key
        }
    })
    .then(response => {
        console.log("Response Status: ", response.status);
        if (!response.ok) {
            throw new Error("Network response was not ok.");
        }
        return response.json();
    })
    .then(data => {
        console.log("Data received: ", data);  // Check if data is coming through

        const favContainer = document.getElementById("fav-image-container");
        if (!favContainer) {
            console.error("Favorite image container not found.");
            return;
        }
        
        favContainer.innerHTML = ""; 
        
        data.forEach(fav => {
            const favItem = document.createElement("div");
            favItem.classList.add("fav-item");
            
            const img = document.createElement("img");
            img.src = fav.image.url;
            img.alt = "Favorite Cat Image";
            
            favItem.appendChild(img);
            favContainer.appendChild(favItem);
        });

        // Switch to the grid layout
        switchLayout('grid');
    })
    .catch(error => {
        console.error('Error fetching favorites:', error);
    });
}

function showFavorites() {
    document.getElementById('breedSearchSection').style.display = 'none';
    document.getElementById('randomCatImageSection').style.display = 'none';
    document.getElementById('footerSection').style.display='none';
    document.getElementById('fav-container').style.display='block';
    loadFavorites();
}