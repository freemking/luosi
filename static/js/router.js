const ProductRouter = {
    routes: {
        home: '/',
        products: '/products.html',
        product: '/product.html',
        about: '/about.html',
        contact: '/contact.html'
    },

    getQueryParams() {
        const params = {};
        const queryString = window.location.search.substring(1);
        const pairs = queryString.split('&');
        for (let i = 0; i < pairs.length; i++) {
            if (!pairs[i]) continue;
            const pair = pairs[i].split('=');
            params[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
        }
        return params;
    },

    getCurrentCategory() {
        const params = this.getQueryParams();
        return params.category || '';
    },

    getCurrentProductSlug() {
        const params = this.getQueryParams();
        return params.slug || '';
    },

    navigateToProductList(category = '') {
        let url = 'products.html';
        if (category) {
            url += `?category=${encodeURIComponent(category)}`;
        }
        window.location.href = url;
    },

    navigateToProductDetail(slug) {
        window.location.href = `product.html?slug=${encodeURIComponent(slug)}`;
    }
};

window.ProductRouter = ProductRouter;
