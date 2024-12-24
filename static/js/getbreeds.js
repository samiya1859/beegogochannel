    document.addEventListener("DOMContentLoaded", function () {
    const breedSearchInput = document.getElementById("breedSearch");
    const breedDropdown = document.getElementById("breedDropdown");
    const searchInput = breedSearchInput; 
    const dropdownList = breedDropdown; 
    const breeds = []; 

    function populateDropdown(items) {
        dropdownList.innerHTML = ''; 
        items.forEach(item => {
            const div = document.createElement('div');
            div.textContent = item.name; // Use breed name here
            div.className = 'dropdown-item';
            div.onclick = () => {
                searchInput.value = item.name; // Set input to breed name
                dropdownList.classList.remove('show'); // Hide dropdown
                fetchBreedData(item.id); // Fetch breed details on selection
                fetchBreedImages(item.id); // Fetch breed images
            };
            dropdownList.appendChild(div);
        });
    }

    // Filter items based on search text
    function filterItems(searchText) {
        return breeds.filter(item =>
            item.name.toLowerCase().includes(searchText.toLowerCase())
        );
    }
    // Fetch the breed data from the server
    async function fetchBreeds() {
        try {
            const response = await fetch(`/breed/`);
            if (!response.ok) {
                throw new Error('Error fetching breeds');
            }
            const breeds = await response.json();
            console.log("Breeds fetched:", breeds); // Debugging log

            const breedDropdown = document.getElementById("breedDropdown");
            breedDropdown.innerHTML = ''; // Clear any existing items

            breeds.forEach(breed => {
                const li = document.createElement("li");
                const a = document.createElement("a");
                a.href = "#";
                a.textContent = breed.name;
                a.setAttribute("data-breed-id", breed.id);

                a.addEventListener("click", function (event) {
                    event.preventDefault();
                    const breedID = a.getAttribute("data-breed-id");
                    console.log("Breed ID clicked:", breedID); // Debugging log
                    if (breedID) {
                        fetchBreedData(breedID);  // Fetch breed data
                        fetchBreedImages(breedID); // Fetch breed images
                    } else {
                        console.error("Breed ID is undefined");
                    }
                });

                li.appendChild(a);
                breedDropdown.appendChild(li);
            });

            // Show the dropdown if breeds are fetched
            if (breeds.length > 0) {
                breedDropdown.classList.remove("hidden");
            } else {
                console.warn("No breeds found to display.");
            }
        } catch (error) {
            console.error("Error fetching breeds:", error);
        }
    }

    // Fetch and display breed data (description)
    async function fetchBreedData(breedID) {
        try {
            const response = await fetch(`/breed/${breedID}`);
            if (!response.ok) {
                throw new Error(`Error fetching breed data for ID ${breedID}`);
            }

            const breedData = await response.json();
            console.log("Breed data fetched:", breedData);

            // Check for the breed description before displaying it
            if (breedData && breedData.description) {
                document.getElementById("breedDescription").innerText = breedData.description;
            } else {
                document.getElementById("breedDescription").innerText = "Description not available.";
            }
        } catch (error) {
            console.error("Error fetching breed data:", error);
        }
    }

    // Fetch and display breed images
    async function fetchBreedImages(breedID) {
        try {
            const response = await fetch(`/breed/images/${breedID}`);
            if (!response.ok) {
                throw new Error(`Error fetching breed images for ID ${breedID}`);
            }
            const imagesData = await response.json();

            // Display the breed images
            const imageContainer = document.getElementById("imageContainer");
            imageContainer.innerHTML = ''; 

            if (imagesData.length > 0) {
                imagesData.forEach(image => {
                    const img = document.createElement("img");
                    img.src = image.url;
                    img.alt = `Image of breed ${breedID}`;
                    img.classList.add("rounded", "shadow-lg");
                    imageContainer.appendChild(img);
                });
            } else {
                imageContainer.innerText = "No images available.";
            }
        } catch (error) {
            console.error("Error fetching breed images:", error);
        }
    }

    // Event listener for input field to show the dropdown
    searchInput.addEventListener('focus', () => {
        if (!searchInput.value) {
            searchInput.value = breeds[0]?.name || ''; // Set default value if empty
        }
        dropdownList.classList.add('show'); // Show the dropdown
    });

    // Event listener for input field to filter dropdown items as the user types
    searchInput.addEventListener('input', (e) => {
        const filtered = filterItems(e.target.value); // Filter breeds based on input
        populateDropdown(filtered); // Update the dropdown with filtered breeds
        dropdownList.classList.add('show'); // Ensure dropdown is visible
    });

    // Close dropdown when clicking outside the input or dropdown
    document.addEventListener('click', (e) => {
        if (!dropdownList.contains(e.target) && e.target !== searchInput) {
            dropdownList.classList.remove('show');
        }
    });

    // Initially fetch and populate the breed list
    fetchBreeds();
});
