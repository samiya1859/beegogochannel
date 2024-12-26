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

<script src="https://cdn.jsdelivr.net/npm/swiffy-slider@1.6.0/dist/js/swiffy-slider.min.js" crossorigin="anonymous" defer></script>
<link href="https://cdn.jsdelivr.net/npm/swiffy-slider@1.6.0/dist/css/swiffy-slider.min.css" rel="stylesheet" crossorigin="anonymous">

  
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
                <div class="image-container">
                    <div class="swiffy-slider slider-item-ratio slider-item-ratio-16x9 slider-nav-animation slider-nav-animation-fadein slider-item-first-visible" id="swiffy-animation">
                        <ul class="slider-container" id="imageContainer">
                        </ul>
                    
                        <button type="button" class="slider-nav" aria-label="Go to previous"></button>
                        <button type="button" class="slider-nav slider-nav-next" aria-label="Go to next"></button>
                    
                        <div class="slider-indicators">
                            <button aria-label="Go to slide" class="active"></button>
                            <button aria-label="Go to slide"></button>
                            <button aria-label="Go to slide"></button>
                            <button aria-label="Go to slide"></button>
                            <button aria-label="Go to slide"></button>
                            <button aria-label="Go to slide"></button>
                        </div>
                    </div>
                </div>
        
                <!-- Breed Information -->
                <div class="breed-info">
                    <div class="breednames">
                        <h2 id="breedTitle" class="breed-title"></h2>
                        <span id="breedOrigin"></span>
                        <span id="breedId"></span>
                    </div>
                    
                    <p id="breedDescription" class="breed-description"></p>
                    <div class="source-link">
                        <a href="#" class="wikipedia-link">WIKIPEDIA</a>
                    </div>
                </div>
            </div>
        </div> 
        
        <!-- Favorites Section -->
        <div id="fav-container"  style="display: none;">
            <div class="view-options">
                <button class="active" id="grid-view" onclick="switchLayout('grid')">
                    <i class="ri-layout-grid-fill"></i> 
                </button>
                <button id="scroll-view" onclick="switchLayout('scroll')">
                    <i class="ri-layout-horizontal-line"></i> 
                </button>
            </div>
            <div id="fav-image-container" class="grid-container">
                {{if .Favorites}}
                    {{range .Favorites}}
                        <div class="fav-item">
                            <img src="{{.Image.URL}}" alt="Favorite Image"/>
                           
                        </div>
                    {{end}}
                {{else}}
                    <p>No favorites found.</p>
                {{end}}
            </div>
        </div>

        <!-- Footer should be outside breedSearchSection -->
        <div class="footer" id="footerSection">
            <div class="heart-container">
                <i class="ri-heart-2-line" id="heart-icon" onclick="addToFavorites('{{.ImageID}}')"></i>
                <span class="heart-animation"></span>
            </div>
            <div class="thumbs" style="gap: 10px;">
                <i class="ri-thumb-up-line" id="thumb-up" onclick="sendVote('{{.ImageID}}',true)"></i>
                <i class="ri-thumb-down-line" id="thumb-down" onclick="sendVote('{{.ImageID}}', false)"></i>
            </div>
        </div>

    </div>
    
    
    <script src="static/js/onloading.js"></script>
    <script src="static/js/breedfetching.js"></script>
    <script src="static/js/voting.js"></script>
    <script src="static/js/makefav.js"></script>
    
</body>
</html>
