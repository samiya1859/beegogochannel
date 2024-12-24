document.addEventListener("DOMContentLoaded", function () {
    const searchInput = document.getElementById("breedSearch");
    const dropdownList = document.getElementById("breedDropdown");
    const clearButton = document.getElementById("clearButton");
    const breeds = []; 

    function populateDropdown(items) {
        dropdownList.innerHTML = '';
        items.forEach(item => {
            const div = document.createElement('div');
            div.textContent = item.name;
            div.className = 'dropdown-item';
            
            // Add selected class if this is the current value
            if (item.name === searchInput.value) {
                div.classList.add('selected');
            }
            
            div.onclick = () => {
                searchInput.value = item.name;
                dropdownList.classList.add('hidden');
                clearButton.classList.add('show');
                fetchBreedData(item.id);
                fetchBreedImages(item.id);
            };
            dropdownList.appendChild(div);
        });
        dropdownList.classList.remove('hidden');
    }

    // Handle clear button
    clearButton.addEventListener('click', () => {
        searchInput.value = '';
        clearButton.classList.remove('show');
        dropdownList.classList.add('hidden');
    });

    // Show/hide clear button based on input
    searchInput.addEventListener('input', (e) => {
        clearButton.classList.toggle('show', e.target.value.length > 0);
        const filtered = filterItems(e.target.value);
        populateDropdown(filtered);
    });

    // Show dropdown on focus
    searchInput.addEventListener('focus', () => {
        const filtered = filterItems(searchInput.value);
        populateDropdown(filtered);
    });

    // Hide dropdown when clicking outside
    document.addEventListener('click', (e) => {
        if (!dropdownList.contains(e.target) && e.target !== searchInput) {
            dropdownList.classList.add('hidden');
        }
    });

    // Filter items based on search text
    function filterItems(searchText) {
        return breeds.filter(item =>
            item.name.toLowerCase().includes(searchText.toLowerCase())
        );
    }

    // Your existing fetch functions remain the same
    async function fetchBreeds() {
        try {
            const response = await fetch(`/breed/`);
            if (!response.ok) throw new Error('Error fetching breeds');
            const data = await response.json();
            breeds.push(...data); // Store breeds in the array
            // Initially populate dropdown if needed
            if (breeds.length > 0) {
                populateDropdown(breeds);
            }
        } catch (error) {
            console.error("Error fetching breeds:", error);
        }
    }

    // Initially fetch breeds
    fetchBreeds();
});