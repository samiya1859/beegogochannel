document.addEventListener("DOMContentLoaded", function () {
    const searchInput = document.getElementById("breedSearch");
    const dropdownList = document.getElementById("breedDropdown");
    const clearButton = document.getElementById("clearButton");
    const imageContainer = document.getElementById("imageContainer");
    const breedContent = document.getElementById("breedContent");
    const breeds = []; // Array to store breed data

    // Populate the dropdown with filtered items
    function populateDropdown(items) {
        dropdownList.innerHTML = '';
        items.forEach(item => {
            const div = document.createElement('div');
            div.textContent = item.name;
            div.className = 'dropdown-item';
            
            if (item.name === searchInput.value) {
                div.classList.add('selected');
            }
            
            div.onclick = () => {
                searchInput.value = item.name;
                dropdownList.classList.add('hidden');
                clearButton.classList.add('show');
                breedContent.classList.remove('hidden'); // Show breed content div
                
                // Fetch data and images for the selected breed
                fetchBreedData(item.id);
                fetchBreedImages(item.id);
            };
            dropdownList.appendChild(div);
        });
        dropdownList.classList.remove('hidden');
    }

    // Handle the clear button
    clearButton.addEventListener('click', () => {
        searchInput.value = '';
        clearButton.classList.remove('show');
        dropdownList.classList.add('hidden');
        imageContainer.innerHTML = ''; // Clear images
        breedContent.classList.add('hidden'); // Hide breed content div
    });

    // Show/hide clear button and filter items based on input
    searchInput.addEventListener('input', (e) => {
        clearButton.classList.toggle('show', e.target.value.length > 0);
        const filtered = filterItems(e.target.value);
        populateDropdown(filtered);
    });

    // Show the dropdown on focus
    searchInput.addEventListener('focus', () => {
        const filtered = filterItems(searchInput.value);
        populateDropdown(filtered);
    });

    // Hide the dropdown when clicking outside
    document.addEventListener('click', (e) => {
        if (!dropdownList.contains(e.target) && e.target !== searchInput) {
            dropdownList.classList.add('hidden');
        }
    });

    // Filter breeds based on search text
    function filterItems(searchText) {
        return breeds.filter(item =>
            item.name.toLowerCase().includes(searchText.toLowerCase())
        );
    }

    // Fetch breed data and populate dropdown
    async function fetchBreeds() {
        try {
            const response = await fetch(`/breed/`);
            if (!response.ok) throw new Error('Error fetching breeds');
            const data = await response.json();
            breeds.push(...data); 
            if (breeds.length > 0) {
                populateDropdown(breeds);
            }
        } catch (error) {
            console.error("Error fetching breeds:", error);
        }
    }

    // Fetch breed-specific data
    async function fetchBreedData(breedID) {
        try {
            const response = await fetch(`/breed/${breedID}`);
            if (!response.ok) {
                throw new Error(`Error fetching data for breed ID ${breedID}`);
            }
            const breedData = await response.json();

            // Update the breed information section
            const breedTitle = document.getElementById("breedTitle");
            const breedOrigin = document.getElementById("breedOrigin");
            const breedId = document.getElementById("breedId");
            const breedDescription = document.getElementById("breedDescription");
            const wikiLink = document.querySelector(".wikipedia-link");

            breedTitle.textContent = breedData.name;
            breedOrigin.innerHTML = `
            (${breedData.origin})
            `;
            breedId.textContent = breedData.id;
            breedDescription.textContent = breedData.description;
            wikiLink.href = breedData.wikipedia_url;
            wikiLink.target = "_blank";
        } catch (error) {
            console.error("Error fetching breed data:", error);
        }
    }

    // Fetch breed-specific images
    async function fetchBreedImages(breedID) {
        console.log(`Fetching breed images for: ${breedID}`);
        try {
            const response = await fetch(`/breed/images/${breedID}`);
            if (!response.ok) {
                throw new Error(`Error fetching breed images for ID ${breedID}`);
            }
            const imagesData = await response.json();
            imageContainer.innerHTML = ''; // Clear previous images

            if (imagesData.length > 0) {
                imagesData.forEach((image, index) => {
                    console.log("Fetched image", image);
                    const li = document.createElement("li");
                    li.classList.add("slider-item"); // Add class for slider item
                    if (index === 0) {
                        li.classList.add("slide-visible");
                    } 
                    
                    const img = document.createElement("img");
                    img.src = image.url;
                    img.alt = `Image of breed ${breedID}`;
                    img.loading = "lazy";
                    li.appendChild(img);

                    // Append the slide to the container
                    imageContainer.appendChild(li);
                });
           
            } else {
                imageContainer.innerText = "No images available.";
            }
        } catch (error) {
            console.error("Error fetching breed images:", error);
        }
    }
    

    // Initial fetch of breed data
    fetchBreeds();
});
