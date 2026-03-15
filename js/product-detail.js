document.addEventListener('DOMContentLoaded', function() {
    const params = ProductRouter.getQueryParams();
    const slug = params.slug || 'hex-head-bolts-din-931';
    
    const product = ProductAPI.getProductBySlug(slug);
    const productContent = document.getElementById('product-content');
    const relatedProductsContainer = document.getElementById('related-products');
    const footerProductLinks = document.getElementById('footer-product-links');

    const categories = ProductAPI.getCategories();
    categories.slice(0, 4).forEach(cat => {
        const li = document.createElement('li');
        li.innerHTML = `<a href="products.html?category=${cat.id}">${cat.name}</a>`;
        footerProductLinks.appendChild(li);
    });

    if (!product) {
        productContent.innerHTML = `
            <div style="text-align: center; padding: 60px 20px;">
                <h2 style="font-size: 28px; margin-bottom: 20px; color: var(--dark-color);">Product Not Found</h2>
                <p style="color: var(--gray-color); margin-bottom: 30px;">Sorry, the product you are looking for does not exist.</p>
                <a href="products.html" class="btn">Browse All Products</a>
            </div>
        `;
        return;
    }

    updateSEO(product);
    updateBreadcrumb(product);

    const html = Templates.productDetail(product);
    productContent.innerHTML = html;

    if (product.relatedProducts && product.relatedProducts.length > 0) {
        const related = ProductAPI.getRelatedProducts(product.relatedProducts).filter(Boolean);
        related.forEach(r => {
            const html = Templates.relatedProductCard(r);
            relatedProductsContainer.insertAdjacentHTML('beforeend', html);
        });
    }

    function updateSEO(product) {
        const pageTitle = document.getElementById('page-title');
        const pageDescription = document.getElementById('page-description');
        const pageKeywords = document.getElementById('page-keywords');
        const ogTitle = document.getElementById('og-title');
        const ogDescription = document.getElementById('og-description');

        const standards = product.standards.join(', ');
        pageTitle.textContent = `${product.name} ${standards} - High Quality Industrial Fasteners | Fastener Pro`;
        pageDescription.content = `${product.description} Available in various materials and grades. Factory direct supply from professional fastener manufacturer.`;
        pageKeywords.content = `${product.slug.replace(/-/g, ', ')}, ${standards}, ${product.category}, industrial fasteners`;
        ogTitle.textContent = pageTitle.textContent;
        ogDescription.content = pageDescription.content;
    }

    function updateBreadcrumb(product) {
        const breadcrumbCurrent = document.getElementById('breadcrumb-current');
        breadcrumbCurrent.textContent = product.name;
    }
});
