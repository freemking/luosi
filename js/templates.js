const Templates = {
    productCard(product) {
        const standardsHtml = product.standards.map(std => `<span class="standard-tag">${std}</span>`).join('');
        return `
            <div class="product-card">
                <div class="product-image">${product.icon}</div>
                <div class="product-content">
                    <h3>${product.name}</h3>
                    <p>${product.description}</p>
                    <div class="product-standards">
                        ${standardsHtml}
                    </div>
                    <a href="product.html?slug=${product.slug}" class="btn btn-outline">View Details</a>
                </div>
            </div>
        `;
    },

    categoryCard(category) {
        return `
            <div class="category-card">
                <div class="category-image">${category.icon}</div>
                <div class="category-content">
                    <h3>${category.name}</h3>
                    <p>${category.description}</p>
                    <a href="products.html?category=${category.id}" class="btn btn-outline">View Products</a>
                </div>
            </div>
        `;
    },

    productDetail(product) {
        const standardsHtml = product.standards.map(std => `<span class="standard-item">${std}</span>`).join('');
        const details = product.details;
        
        let specsHtml = '';
        if (details && details.specifications) {
            const specs = details.specifications;
            specsHtml = `
                <table class="spec-table">
                    ${specs.standard ? `<tr><td>Standard:</td><td>${specs.standard}</td></tr>` : ''}
                    ${specs.productName ? `<tr><td>Product Name:</td><td>${specs.productName}</td></tr>` : ''}
                    ${specs.nominalDiameter ? `<tr><td>Nominal Diameter:</td><td>${specs.nominalDiameter}</td></tr>` : ''}
                    ${specs.pitch ? `<tr><td>Pitch:</td><td>${specs.pitch}</td></tr>` : ''}
                    ${specs.lengthRange ? `<tr><td>Length Range:</td><td>${specs.lengthRange}</td></tr>` : ''}
                    <tr><td>Materials:</td><td>${Array.isArray(product.material) ? product.material.join(', ') : product.material}</td></tr>
                    ${product.grades ? `<tr><td>Property Classes:</td><td>${product.grades.join(', ')}</td></tr>` : ''}
                    ${product.surfaceTreatment ? `<tr><td>Surface Treatments:</td><td>${product.surfaceTreatment.join(', ')}</td></tr>` : ''}
                </table>
            `;
        }

        let applicationsHtml = '';
        if (details && details.applications) {
            applicationsHtml = `
                <div class="application-list">
                    ${details.applications.map(app => `
                        <div class="application-item">
                            <span>✓</span>
                            <span>${app}</span>
                        </div>
                    `).join('')}
                </div>
            `;
        }

        let packagingHtml = '';
        if (details && details.packaging) {
            const pkg = details.packaging;
            packagingHtml = `
                <p style="color: var(--gray-color); line-height: 1.8; margin-bottom: 16px;"><strong>Standard Packaging:</strong> ${pkg.standard}</p>
                <p style="color: var(--gray-color); line-height: 1.8; margin-bottom: 16px;"><strong>Packaging Options:</strong></p>
                <ul style="color: var(--gray-color); line-height: 1.8; margin-left: 20px; margin-bottom: 16px;">
                    ${pkg.options.map(opt => `<li>${opt}</li>`).join('')}
                </ul>
                <p style="color: var(--gray-color); line-height: 1.8; margin-bottom: 16px;"><strong>Delivery Time:</strong></p>
                <ul style="color: var(--gray-color); line-height: 1.8; margin-left: 20px;">
                    ${pkg.delivery.map(d => `<li>${d}</li>`).join('')}
                </ul>
                <p style="color: var(--gray-color); line-height: 1.8; margin-top: 16px;"><strong>Shipping:</strong> ${pkg.shipping}</p>
            `;
        }

        return `
            <div class="product-detail-container">
                <div class="product-gallery">
                    <div class="main-image">${product.icon}</div>
                    <div class="thumbnail-grid">
                        <div class="thumbnail active">${product.icon}</div>
                        <div class="thumbnail">${product.icon}</div>
                        <div class="thumbnail">${product.icon}</div>
                        <div class="thumbnail">📏</div>
                    </div>
                </div>
                <div class="product-info">
                    <h1>${product.name}</h1>
                    ${product.standards.length > 0 ? `<div class="product-standard">${product.standards[0]}</div>` : ''}
                    <p>${product.description}</p>
                    ${details && details.fullDescription ? `<p>${details.fullDescription}</p>` : ''}
                    
                    <div class="available-standards">
                        <h3>Available Standards</h3>
                        <div class="standards-list">
                            ${standardsHtml}
                        </div>
                    </div>

                    <div class="product-specs">
                        <h3>Product Specifications</h3>
                        ${specsHtml}
                    </div>

                    <a href="contact.html" class="btn">Request Quote & Inquire</a>
                </div>
            </div>

            <div class="product-detail-bottom">
                <div class="tabs">
                    <ul class="tab-list">
                        <li><button class="tab-btn active" data-tab="description">Description</button></li>
                        <li><button class="tab-btn" data-tab="specifications">Specifications</button></li>
                        <li><button class="tab-btn" data-tab="applications">Applications</button></li>
                        <li><button class="tab-btn" data-tab="packaging">Packaging & Delivery</button></li>
                    </ul>
                </div>

                <div class="tab-content active" id="description">
                    <h3 style="font-size: 24px; margin-bottom: 20px; color: var(--dark-color);">Product Description</h3>
                    ${details && details.fullDescription ? `<p style="color: var(--gray-color); line-height: 1.8;">${details.fullDescription}</p>` : `<p style="color: var(--gray-color); line-height: 1.8;">${product.description}</p>`}
                </div>

                <div class="tab-content" id="specifications">
                    <h3 style="font-size: 24px; margin-bottom: 20px; color: var(--dark-color);">Technical Specifications</h3>
                    <div class="product-specs" style="margin-top: 0;">
                        ${specsHtml}
                    </div>
                </div>

                <div class="tab-content" id="applications">
                    <h3 style="font-size: 24px; margin-bottom: 20px; color: var(--dark-color);">Common Applications</h3>
                    ${applicationsHtml}
                </div>

                <div class="tab-content" id="packaging">
                    <h3 style="font-size: 24px; margin-bottom: 20px; color: var(--dark-color);">Packaging & Delivery</h3>
                    ${packagingHtml}
                </div>
            </div>
        `;
    },

    relatedProductCard(product) {
        const standardsHtml = product.standards.map(std => `<span class="standard-tag">${std}</span>`).join('');
        return `
            <div class="product-card">
                <div class="product-image">${product.icon}</div>
                <div class="product-content">
                    <h3>${product.name}</h3>
                    <p>${product.description}</p>
                    <div class="product-standards">
                        ${standardsHtml}
                    </div>
                    <a href="product.html?slug=${product.slug}" class="btn btn-outline">View Details</a>
                </div>
            </div>
        `;
    },

    filterOptions(options, currentValue) {
        return `
            <option value="">All ${options[0].id.split('-').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ')}s</option>
            ${options.map(opt => `<option value="${opt.id}" ${currentValue === opt.id ? 'selected' : ''}>${opt.name}</option>`).join('')}
        `;
    }
};

window.Templates = Templates;
