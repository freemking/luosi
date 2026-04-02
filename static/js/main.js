document.addEventListener('DOMContentLoaded', function() {
    const mobileMenuBtn = document.querySelector('.mobile-menu-btn');
    const navMenu = document.querySelector('nav ul');
    
    if (mobileMenuBtn && navMenu) {
        mobileMenuBtn.addEventListener('click', function() {
            navMenu.classList.toggle('active');
        });
        
        const navLinks = navMenu.querySelectorAll('a');
        navLinks.forEach(link => {
            link.addEventListener('click', function() {
                navMenu.classList.remove('active');
            });
        });
    }
    
    // Mobile dropdown toggle
    const hasDropdown = document.querySelector('.has-dropdown');
    if (hasDropdown) {
        const dropdownLink = hasDropdown.querySelector('a');
        dropdownLink.addEventListener('click', function(e) {
            if (window.innerWidth <= 768) {
                e.preventDefault();
                hasDropdown.classList.toggle('active');
            }
        });
    }
    
    const thumbnails = document.querySelectorAll('.thumbnail');
    const mainImage = document.querySelector('.main-image');
    
    if (thumbnails.length > 0 && mainImage) {
        thumbnails.forEach(thumbnail => {
            thumbnail.addEventListener('click', function() {
                thumbnails.forEach(t => t.classList.remove('active'));
                this.classList.add('active');
            });
        });
    }
    
    const tabBtns = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');
    
    if (tabBtns.length > 0 && tabContents.length > 0) {
        tabBtns.forEach(btn => {
            btn.addEventListener('click', function() {
                const targetTab = this.getAttribute('data-tab');
                
                tabBtns.forEach(b => b.classList.remove('active'));
                tabContents.forEach(c => c.classList.remove('active'));
                
                this.classList.add('active');
                const content = document.getElementById(targetTab);
                if (content) {
                    content.classList.add('active');
                }
            });
        });
        
        if (tabBtns.length > 0) {
            tabBtns[0].click();
        }
    }
    
    const contactForm = document.querySelector('.contact-form form');
    if (contactForm) {
        contactForm.addEventListener('submit', function(e) {
            e.preventDefault();
            alert('Thank you for your inquiry! We will contact you soon.');
            contactForm.reset();
        });
    }
    
    const filterSelects = document.querySelectorAll('.product-filter select');
    if (filterSelects.length > 0) {
        filterSelects.forEach(select => {
            select.addEventListener('change', function() {
                console.log('Filter changed:', this.value);
            });
        });
    }
});
