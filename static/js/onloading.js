const img = document.getElementById('catImage');
const loading = document.getElementById('loading');

img.onload = function() {
    loading.style.display = 'none';
    img.style.display = 'block'; 
};


const navItems = document.querySelectorAll('.nav-item');

   console.log(navItems)
    navItems.forEach(item => {
        item.addEventListener('click', function() {
           
            navItems.forEach(i => i.classList.remove('active'));
            this.classList.add('active');
        });
    });


// showing page parts

function showVoting() {
    document.getElementById('breedSearchSection').style.display = 'none';
    document.getElementById('randomCatImageSection').style.display = 'block';
    document.getElementById('footerSection').style.display='flex';
}

function showBreedSearch() {
    document.getElementById('breedSearchSection').style.display = 'block';
    document.getElementById('randomCatImageSection').style.display = 'none';
    document.getElementById('footerSection').style.display='none';
}

    