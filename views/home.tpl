<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>The cat api</title>
    
    <link
    href="https://cdn.jsdelivr.net/npm/remixicon@4.5.0/fonts/remixicon.css"
    rel="stylesheet"/>

    <link rel="stylesheet" href="static/css/catproject.css">


   <!-- Tailwind CSS CDN -->
  <script src="https://cdn.tailwindcss.com"></script>

  <!-- DaisyUI CDN -->
  <script src="https://cdn.jsdelivr.net/npm/daisyui@2.45.0/dist/full.js"></script>
  
</head>
<body>
    
    <div class="main-content">
        <nav style="display: flex; gap: 30px;">
            <div style="text-align: center;" onclick="showVoting()">
                <i class="ri-arrow-up-down-line"></i>
                <p id="voting" class="nav-item">Voting</p> 
            </div>
            <div style="text-align: center;" onclick="showBreedSearch()">
                <i class="ri-search-eye-line"></i>
                <p id="breeds" class="nav-item">Breeds</p>
            </div>

            <div style="text-align: center;" onclick="showFavorites()">
                <i class="ri-heart-2-line"></i>
                <p id="favs" class="nav-item">Favs</p>
            </div>
        </nav>

        <!-- Voting Section -->
        <div id="randomCatImageSection">
            <div id="loading" class="loading-animation">
                <img src="static/img/cat.png"/>
            </div>
            <img id="catImage" src="{{.ImageURL}}" alt="Random Cat" style="display: none;"/>
        </div>
         
        <!-- Search by breed functionality -->
        <div id="breedSearchSection" style="display:none;">
            <div class="search-container">
                <div class="input-wrapper">
                <input type="text" id="breedSearch" placeholder="Select a breed..." />
                <span id="clearButton" class="clear-button hidden">×</span>
                </div>
                <ul id="breedDropdown" class="hidden"></ul>
            </div>

            <div id="breedContent" class="breed-content hidden">
                <!-- Main Image and Dots -->
                <div class="image-container">
                <div id="imageContainer" class="main-image"></div>
                <div class="dot-indicators" id="dotIndicators"></div>
            </div>

            <!-- Breed Information -->
            <div class="breed-info">
                <h2 id="breedTitle" class="breed-title"></h2>
                <p id="breedDescription" class="breed-description"></p>
                <div class="source-link">
                    <a href="#" class="wikipedia-link">WIKIPEDIA</a>
                </div>
            </div>
        </div>

        <!-- footer -->
        <div class="footer" id="footerSection">
            <div class="heart-container">
                <i class="ri-heart-2-line" id="heart-icon" onclick="addToFavorites('{{.CatImageID}}')"></i>
                <span class="heart-animation"></span>
            </div>
            <div class="thumbs" style="gap: 10px;">
                <i class="ri-thumb-up-line" id="thumb-up" onclick="sendVote('{{.CatImageID}}', true)"></i>
                <i class="ri-thumb-down-line" id="thumb-down" onclick="sendVote('{{.CatImageID}}', false)"></i>
            </div>
        </div>
    </div>
    
    
    <script src="static/js/onloading.js"></script>
    <script src="static/js/breeddropdown.js"></script>

</body>
</html>
