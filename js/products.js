document.addEventListener('DOMContentLoaded', function() {
    const params = ProductRouter.getQueryParams();
    const currentCategory = params.category || '';
    
    const categorySelect = document.getElementById('category');
    const standardSelect = document.getElementById('standard');
    const materialSelect = document.getElementById('material');
    const productsGrid = document.getElementById('products-grid');
    const footerProductLinks = document.getElementById('footer-product-links');

    const categories = ProductAPI.getCategories();
    const standards = ProductAPI.getStandards();
    const materials = ProductAPI.getMaterials();

    categories.forEach(cat => {
        const option = document.createElement('option');
        option.value = cat.id;
        option.textContent = cat.name;
        if (cat.id === currentCategory) {
            option.selected = true;
        }
        categorySelect.appendChild(option);
    });

    standards.forEach(std => {
        const option = document.createElement('option');
        option.value = std.id;
        option.textContent = std.name;
        standardSelect.appendChild(option);
    });

    materials.forEach(mat => {
        const option = document.createElement('option');
        option.value = mat.id;
        option.textContent = mat.name;
        materialSelect.appendChild(option);
    });

    categories.slice(0, 4).forEach(cat => {
        const li = document.createElement('li');
        li.innerHTML = `<a href="products.html?category=${cat.id}">${cat.name}</a>`;
        footerProductLinks.appendChild(li);
    });

    function updateSEO() {
        const pageTitle = document.getElementById('page-title');
        const pageDescription = document.getElementById('page-description');
        const pageHeaderTitle = document.getElementById('page-header-title');
        const pageHeaderSubtitle = document.getElementById('page-header-subtitle');
        const breadcrumbCurrent = document.getElementById('breadcrumb-current');

        if (currentCategory) {
            const category = categories.find(c => c.id === currentCategory);
            if (category) {
                pageTitle.textContent = `${category.name} - High Quality Industrial ${category.name} | Fastener Pro`;
                pageDescription.content = `Browse our selection of high-quality ${category.name.toLowerCase}. All international standards available. Factory direct supply.`;
                pageHeaderTitle.textContent = category.name;
                pageHeaderSubtitle.textContent = category.description;
                breadcrumbCurrent.textContent = category.name;
            }
        } else {
            pageTitle.textContent = 'Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Fastener Pro';
            pageDescription.content = 'Browse our complete range of industrial fasteners including bolts, nuts, washers, screws, threaded rods and custom parts. All international standards available.';
            pageHeaderTitle.textContent = 'Our Products';
            pageHeaderSubtitle.textContent = 'Comprehensive selection of high-quality industrial fasteners';
            breadcrumbCurrent.textContent = 'Products';
        }
    }

    function renderProducts() {
        const selectedCategory = categorySelect.value;
        const selectedStandard = standardSelect.value;
        const selectedMaterial = materialSelect.value;

        let products = ProductAPI.getProducts();

        if (selectedCategory) {
            products = products.filter(p => p.category === selectedCategory);
        }

        productsGrid.innerHTML = '';

        if (products.length === 0) {
            productsGrid.innerHTML = '<p style="grid-column: 1/-1; text-align: center; padding: 40px; color: var(--gray-color);">No products found matching your criteria.</p>';
            return;
        }

        products.forEach(product => {
            const html = Templates.productCard(product);
            productsGrid.insertAdjacentHTML('beforeend', html);
        });
    }

    updateSEO();
    renderProducts();

    categorySelect.addEventListener('change', function() {
        renderProducts();
        updateSEO();
    });

    standardSelect.addEventListener('change', renderProducts);
    materialSelect.addEventListener('change', renderProducts);
});
