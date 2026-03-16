document.addEventListener('DOMContentLoaded', function() {
    const categoriesContainer = document.getElementById('categories-grid');
    const featuredContainer = document.getElementById('featured-products');

    const categories = ProductAPI.getCategories();
    const allProducts = ProductAPI.getProducts();
    const featuredProducts = allProducts.slice(0, 4);

    categories.forEach(category => {
        const html = Templates.categoryCard(category);
        categoriesContainer.insertAdjacentHTML('beforeend', html);
    });

    featuredProducts.forEach(product => {
        const html = Templates.productCard(product);
        featuredContainer.insertAdjacentHTML('beforeend', html);
    });
});
