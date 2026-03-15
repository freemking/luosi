const productsData = {
    categories: [
        {
            id: 'bolts',
            name: 'Bolts',
            description: 'High quality bolts in various standards and specifications',
            icon: '🔩'
        },
        {
            id: 'nuts',
            name: 'Nuts',
            description: 'Complete range of nuts matching all bolt sizes',
            icon: '🔘'
        },
        {
            id: 'washers',
            name: 'Washers',
            description: 'Flat washers, spring washers and specialty washers',
            icon: '⚫'
        },
        {
            id: 'screws',
            name: 'Screws',
            description: 'Machine screws, self-tapping screws and specialty screws',
            icon: '🌀'
        },
        {
            id: 'threaded-rods',
            name: 'Threaded Rods',
            description: 'Full threaded rods in various materials and lengths',
            icon: '📏'
        },
        {
            id: 'custom',
            name: 'Custom Parts',
            description: 'Non-standard custom fasteners manufactured to your requirements',
            icon: '🔧'
        }
    ],
    standards: [
        { id: 'din', name: 'DIN' },
        { id: 'iso', name: 'ISO' },
        { id: 'ansi', name: 'ANSI' },
        { id: 'jis', name: 'JIS' },
        { id: 'bs', name: 'BS' }
    ],
    materials: [
        { id: 'carbon-steel', name: 'Carbon Steel' },
        { id: 'stainless-steel', name: 'Stainless Steel' },
        { id: 'alloy-steel', name: 'Alloy Steel' },
        { id: 'brass', name: 'Brass' }
    ],
    products: [
        {
            id: 1,
            name: 'Hex Head Bolts',
            slug: 'hex-head-bolts-din-931',
            category: 'bolts',
            description: 'Hexagon head bolts with partial thread, available in various grades and sizes.',
            standards: ['DIN 931', 'ISO 4014'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M6 - M64',
            lengthRange: '30mm - 500mm',
            grades: ['4.8', '8.8', '10.9', '12.9', 'A2-70', 'A4-80'],
            surfaceTreatment: ['Plain', 'Zinc Plated', 'Hot Dip Galvanized', 'Black Oxide', 'Dacromet'],
            icon: '🔩',
            details: {
                fullDescription: 'DIN 931 hex head bolts are one of our most popular products. These bolts feature a hexagon-shaped head designed to be tightened with a wrench, providing high torque capability. The partial thread design allows for threading into a tapped hole or to be used with a nut.',
                applications: ['Industrial Machinery', 'Construction', 'Automotive Industry', 'Equipment Manufacturing', 'Electrical Equipment', 'Railway', 'Oil & Gas', 'Marine Engineering'],
                specifications: {
                    standard: 'DIN 931 / ISO 4014',
                    productName: 'Hexagon Head Bolts with Partial Thread',
                    nominalDiameter: 'M6, M8, M10, M12, M16, M20, M24, M30, M36, M42, M48, M56, M64',
                    pitch: 'Coarse: M6(1), M8(1.25), M10(1.5), M12(1.75), M16(2), M20(2.5), M24(3)',
                    lengthRange: '30mm to 500mm',
                    materials: ['Carbon Steel (C1010, C1045)', 'Alloy Steel (40Cr)', 'Stainless Steel (A2, A4, 304, 316)'],
                    propertyClasses: ['4.8', '5.8', '8.8', '10.9', '12.9', 'A2-70', 'A4-80'],
                    surfaceTreatments: ['Plain (Self-color)', 'Zinc Plated (Blue White, Yellow)', 'Hot Dip Galvanized (HDG)', 'Black Oxide', 'Dacromet', 'Mechanical Galvanizing', 'Xylan Coating']
                },
                packaging: {
                    standard: 'Small items are packed in carton boxes, then into wooden pallets. Larger sizes are packed in wooden cases or according to customer requirements.',
                    options: [
                        'Polybag + Label + Carton + Pallet',
                        'Bulk packaging in big bags',
                        'Custom packaging with customer brand',
                        'Special packaging requirements available'
                    ],
                    delivery: [
                        'Standard items in stock: 7-15 days',
                        'Custom production: 20-35 days depending on quantity',
                        'Urgent orders can be accommodated'
                    ],
                    shipping: 'By sea, by air, or by international courier according to customer requirements.'
                }
            },
            relatedProducts: [2, 3, 4, 5]
        },
        {
            id: 2,
            name: 'Hex Head Cap Screws',
            slug: 'hex-head-cap-screws-din-933',
            category: 'bolts',
            description: 'Full thread hexagon head cap screws in carbon steel and stainless steel.',
            standards: ['DIN 933', 'ISO 4017'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M3 - M48',
            lengthRange: '6mm - 200mm',
            grades: ['4.8', '8.8', '10.9', 'A2-70', 'A4-80'],
            surfaceTreatment: ['Plain', 'Zinc Plated', 'Hot Dip Galvanized', 'Black Oxide'],
            icon: '🔩',
            details: {
                fullDescription: 'DIN 933 hex head cap screws with full thread are ideal for applications where the bolt must extend through the entire fastened component. The full thread provides maximum holding power.',
                applications: ['Machine Building', 'Automotive', 'Equipment Assembly', 'Structural Applications', 'General Industrial'],
                specifications: {
                    standard: 'DIN 933 / ISO 4017',
                    productName: 'Hexagon Head Cap Screws with Full Thread',
                    nominalDiameter: 'M3 to M48',
                    lengthRange: '6mm to 200mm',
                    materials: ['Carbon Steel', 'Stainless Steel'],
                    propertyClasses: ['4.8', '8.8', '10.9', 'A2-70', 'A4-80']
                }
            },
            relatedProducts: [1, 3, 5]
        },
        {
            id: 3,
            name: 'Hex Nuts',
            slug: 'hex-nuts-din-934',
            category: 'nuts',
            description: 'Standard hexagon nuts in various grades matching all bolt sizes.',
            standards: ['DIN 934', 'ISO 4032'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M3 - M64',
            grades: ['4', '8', 'A2-70', 'A4-80'],
            surfaceTreatment: ['Plain', 'Zinc Plated', 'Hot Dip Galvanized'],
            icon: '🔘',
            details: {
                fullDescription: 'DIN 934 hexagon nuts are the most common type of nut used with bolts and studs. They are manufactured to precise tolerances for reliable performance.',
                applications: ['All general industrial applications', 'Construction', 'Machinery', 'Automotive'],
                specifications: {
                    standard: 'DIN 934 / ISO 4032',
                    productName: 'Hexagon Nuts',
                    nominalDiameter: 'M3 to M64'
                }
            },
            relatedProducts: [1, 2, 7]
        },
        {
            id: 4,
            name: 'Hex Flange Bolts',
            slug: 'hex-flange-bolts-din-6921',
            category: 'bolts',
            description: 'Hexagon flange bolts with integral washer facing for increased bearing surface.',
            standards: ['DIN 6921', 'ISO 4162'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M6 - M24',
            grades: ['8.8', '10.9', 'A2-70'],
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '🔩',
            details: {
                fullDescription: 'DIN 6921 hex flange bolts combine the bolt and washer into one component, reducing assembly time and providing a larger bearing surface.',
                applications: ['Automotive', 'Heavy Equipment', 'Structural Applications'],
                specifications: {
                    standard: 'DIN 6921 / ISO 4162',
                    nominalDiameter: 'M6 to M24'
                }
            },
            relatedProducts: [1, 2, 3]
        },
        {
            id: 5,
            name: 'Carriage Bolts',
            slug: 'carriage-bolts-din-603',
            category: 'bolts',
            description: 'Round head square neck bolts ideal for wood and timber connections.',
            standards: ['DIN 603', 'ISO 8678'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M5 - M20',
            lengthRange: '20mm - 200mm',
            grades: ['4.8', '8.8'],
            surfaceTreatment: ['Zinc Plated', 'Hot Dip Galvanized'],
            icon: '🔩',
            details: {
                fullDescription: 'DIN 603 carriage bolts with round head and square neck are designed for wood connections where the square neck prevents rotation.',
                applications: ['Wood Construction', 'Furniture', 'Timber Structures', 'Outdoor Projects'],
                specifications: {
                    standard: 'DIN 603 / ISO 8678',
                    nominalDiameter: 'M5 to M20'
                }
            },
            relatedProducts: [1, 3]
        },
        {
            id: 6,
            name: 'Hex Nyloc Nuts',
            slug: 'hex-nyloc-nuts-din-985',
            category: 'nuts',
            description: 'Nylon insert lock nuts for secure locking and vibration resistance.',
            standards: ['DIN 985', 'ISO 7040'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M3 - M36',
            grades: ['5', '8', 'A2-70'],
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '🔘',
            relatedProducts: [1, 3]
        },
        {
            id: 7,
            name: 'Flat Washers',
            slug: 'flat-washers-din-125',
            category: 'washers',
            description: 'Precision manufactured flat washers for uniform load distribution.',
            standards: ['DIN 125', 'ISO 7089'],
            material: ['carbon-steel', 'stainless-steel', 'brass'],
            sizeRange: 'M3 - M64',
            surfaceTreatment: ['Plain', 'Zinc Plated'],
            icon: '⚫',
            details: {
                fullDescription: 'DIN 125 flat washers distribute the clamping load over a larger area, protecting the workpiece surface and providing better load distribution.',
                applications: ['General Industrial', 'Construction', 'Electrical', 'Machinery'],
                specifications: {
                    standard: 'DIN 125 / ISO 7089',
                    nominalDiameter: 'M3 to M64'
                }
            },
            relatedProducts: [1, 3]
        },
        {
            id: 8,
            name: 'Spring Lock Washers',
            slug: 'spring-lock-washers-din-127',
            category: 'washers',
            description: 'Helical spring lock washers for effective locking and anti-vibration performance.',
            standards: ['DIN 127'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M3 - M48',
            surfaceTreatment:['Plain', 'Zinc Plated', 'Black Oxide'],
            icon: '⚫',
            relatedProducts: [7]
        },
        {
            id: 9,
            name: 'Socket Head Cap Screws',
            slug: 'socket-head-cap-screws-din-912',
            category: 'screws',
            description: 'Allen socket head cap screws for high torque applications.',
            standards: ['DIN 912', 'ISO 4762'],
            material: ['alloy-steel', 'stainless-steel'],
            sizeRange: 'M3 - M48',
            lengthRange: '6mm - 300mm',
            grades: ['12.9', 'A2-70', 'A4-80'],
            surfaceTreatment: ['Black Oxide', 'Plain'],
            icon: '🌀',
            relatedProducts: [1, 2]
        },
        {
            id: 10,
            name: 'Threaded Rods',
            slug: 'threaded-rods-din-975',
            category: 'threaded-rods',
            description: 'Full threaded rods in various lengths and materials for construction use.',
            standards: ['DIN 975'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M3 - M64',
            lengthRange: '1m - 3m standard',
            grades: ['4.8', '8.8', 'A2-70'],
            surfaceTreatment: ['Zinc Plated', 'Plain', 'Hot Dip Galvanized'],
            icon: '📏',
            relatedProducts: [1, 3]
        },
        {
            id: 11,
            name: 'Hex Flange Nuts',
            slug: 'hex-flange-nuts-din-6923',
            category: 'nuts',
            description: 'Flange lock nuts with integral serrated washer for increased grip.',
            standards: ['DIN 6923'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M6 - M30',
            grades: ['8', 'A2-70'],
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '🔘',
            details: {
                fullDescription: 'DIN 6923 hex flange nuts integrate a flange washer into the nut design, providing increased bearing surface and eliminates the need for a separate washer. The serrated flange creates effective locking against loosening from vibration.',
                applications: ['Automotive', 'Heavy Equipment', 'Structural Steel', 'Machinery'],
                specifications: {
                    standard: 'DIN 6923',
                    productName: 'Hexagon Flange Nuts',
                    nominalDiameter: 'M6 to M30'
                }
            },
            relatedProducts: [4, 3]
        },
        {
            id: 12,
            name: 'Wing Nuts',
            slug: 'wing-nuts-din-315',
            category: 'nuts',
            description: 'Wing nuts for hand-tightening applications where frequent adjustment is needed.',
            standards: ['DIN 315'],
            material: ['carbon-steel', 'stainless-steel', 'brass'],
            sizeRange: 'M3 - M24',
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '🔘',
            details: {
                fullDescription: 'DIN 315 wing nuts have two protruding wings that allow hand tightening and loosening without tools. Ideal for applications requiring frequent adjustments or removals.',
                applications: ['Furniture', 'Construction Scaffolding', 'Plumbing', 'Equipment Guards', 'Temporary Fastening'],
                specifications: {
                    standard: 'DIN 315',
                    productName: 'Wing Nuts',
                    nominalDiameter: 'M3 to M24'
                }
            },
            relatedProducts: [6, 3]
        },
        {
            id: 13,
            name: 'Acorn Nuts',
            slug: 'acorn-nuts-din-1587',
            category: 'nuts',
            description: 'Domed cap acorn nuts for decorative and protective finishing.',
            standards: ['DIN 1587'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M4 - M24',
            grades: ['4', 'A2-70'],
            surfaceTreatment: ['Zinc Plated', 'Polished'],
            icon: '🔘',
            details: {
                fullDescription: 'DIN 1587 acorn nuts have a domed cap that covers exposed bolt threads, providing a clean finished appearance and protecting against injury from sharp thread ends.',
                applications: ['Furniture', 'Automotive Trim', 'Railings', 'Decorative Applications', 'Exposed Fastening'],
                specifications: {
                    standard: 'DIN 1587',
                    productName: 'Acorn Cap Nuts',
                    nominalDiameter: 'M4 to M24'
                }
            },
            relatedProducts: [3, 1]
        },
        {
            id: 14,
            name: 'Tooth Lock Washers',
            slug: 'tooth-lock-washers-din-6797',
            category: 'washers',
            description: 'Internal/external tooth lock washers providing excellent locking action.',
            standards: ['DIN 6797'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M2 - M36',
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '⚫',
            details: {
                fullDescription: 'DIN 6797 tooth lock washers have serrated teeth that bite into the mating surface, providing excellent locking performance against loosening due to vibration and dynamic loads.',
                applications: ['Electrical Equipment', 'Automotive', 'Industrial Machinery', 'Applications requiring light locking'],
                specifications: {
                    standard: 'DIN 6797',
                    productName: 'Tooth Lock Washers',
                    nominalDiameter: 'M2 to M36'
                }
            },
            relatedProducts: [7, 8]
        },
        {
            id: 15,
            name: 'Fender Washers',
            slug: 'fender-washers-ansi',
            category: 'washers',
            description: 'Extra large diameter flat washers for large hole applications.',
            standards: ['ANSI'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: '#6 - 1-1/2"',
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '⚫',
            details: {
                fullDescription: 'Fender washers have a larger outer diameter relative to the inner hole size, distributing load over a larger area particularly useful on thin sheet metal and wood.',
                applications: ['Automotive Bodywork', 'Sheet Metal', 'Wood Construction', 'Plumbing'],
                specifications: {
                    standard: 'ANSI B18.22.1',
                    productName: 'Fender Washers',
                    nominalDiameter: '#6 to 1-1/2 inch'
                }
            },
            relatedProducts: [7]
        },
        {
            id: 16,
            name: 'Machine Screws Pan Head',
            slug: 'machine-screws-pan-head-din-7985',
            category: 'screws',
            description: 'Pan head machine screws with Phillips drive for general applications.',
            standards: ['DIN 7985', 'ISO 7045'],
            material: ['carbon-steel', 'stainless-steel', 'brass'],
            sizeRange: 'M1.6 - M10',
            lengthRange: '2mm - 100mm',
            surfaceTreatment: ['Zinc Plated', 'Plain', 'Nickel Plated'],
            icon: '🌀',
            details: {
                fullDescription: 'DIN 7985 pan head machine screws with Phillips cross recess drive are widely used in many industries where a neat appearance is required. The rounded head sits flush or slightly above the surface.',
                applications: ['Electronics', 'Appliances', 'Office Equipment', 'Metal Assembly'],
                specifications: {
                    standard: 'DIN 7985 / ISO 7045',
                    productName: 'Pan Head Machine Screws with Phillips Drive',
                    nominalDiameter: 'M1.6 to M10'
                }
            },
            relatedProducts: [9, 17]
        },
        {
            id: 17,
            name: 'Countersunk Machine Screws',
            slug: 'countersunk-machine-screws-din-965',
            category: 'screws',
            description: 'Flat countersunk head machine screws for flush mounting.',
            standards: ['DIN 965', 'ISO 7046'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M2 - M10',
            lengthRange: '3mm - 100mm',
            surfaceTreatment: ['Zinc Plated', 'Plain'],
            icon: '🌀',
            details: {
                fullDescription: 'DIN 965 countersunk machine screws allow flush installation where the head must sit below the surface of the workpiece. Commonly used with Phillips drive.',
                applications: ['Furniture', 'Cabinets', 'Electrical Panels', 'where flush finish is required'],
                specifications: {
                    standard: 'DIN 965 / ISO 7046',
                    productName: 'Countersunk Head Machine Screws',
                    nominalDiameter: 'M2 to M10'
                }
            },
            relatedProducts: [9, 16]
        },
        {
            id: 18,
            name: 'Self-Tapping Screws',
            slug: 'self-tapping-screws-din-7981',
            category: 'screws',
            description: 'Self-drilling tapping screws for sheet metal and wood applications.',
            standards: ['DIN 7981', 'DIN 7982'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'ST2.2 - ST6.3',
            lengthRange: '6.5mm - 100mm',
            surfaceTreatment: ['Zinc Plated', 'Phosphate'],
            icon: '🌀',
            details: {
                fullDescription: 'Self-tapping screws create their own thread when driven into pre-drilled holes in metal, wood, or plastic. Available in pan head (DIN 7981) and countersunk head (DIN 7982).',
                applications: ['Sheet Metal Fabrication', 'Construction', 'HVAC', 'Electrical Enclosures'],
                specifications: {
                    standard: 'DIN 7981 / DIN 7982',
                    productName: 'Self-Tapping Screws',
                    nominalSize: 'ST2.2 to ST6.3'
                }
            },
            relatedProducts: [16, 9]
        },
        {
            id: 19,
            name: 'Drywall Screws',
            slug: 'drywall-screws',
            category: 'screws',
            description: 'Fine thread drywall screws for gypsum board to wood or metal studs.',
            standards: ['Custom'],
            material: ['carbon-steel'],
            sizeRange: '#6 - #8',
            lengthRange: '16mm - 150mm',
            surfaceTreatment: ['Phosphate', 'Zinc Plated'],
            icon: '🌀',
            details: {
                fullDescription: 'Drywall screws with bugle head and fine or coarse thread are designed for fastening gypsum board to wood or metal studs in drywall construction.',
                applications: ['Drywall Installation', 'Interior Construction', 'Ceilings', 'Partitions'],
                specifications: {
                    productName: 'Drywall Screws',
                    sizes: '#6 (3.5mm) to #8 (4.2mm)'
                }
            },
            relatedProducts: [18]
        },
        {
            id: 20,
            name: 'Wood Screws',
            slug: 'wood-screws-din-95',
            category: 'screws',
            description: 'Traditional wood screws with countersunk head for wood joinery.',
            standards: ['DIN 95'],
            material: ['carbon-steel', 'stainless-steel', 'brass'],
            sizeRange: '1.6mm - 10mm',
            lengthRange: '6mm - 120mm',
            surfaceTreatment: ['Zinc Plated', 'Plain', 'Nickel'],
            icon: '🌀',
            details: {
                fullDescription: 'DIN 95 countersunk wood screws are designed for fastening into wood. The tapered shank and sharp threads provide excellent holding power in wood.',
                applications: ['Woodworking', 'Furniture Making', 'Cabinetry', 'Joinery'],
                specifications: {
                    standard: 'DIN 95',
                    productName: 'Countersunk Wood Screws',
                    nominalDiameter: '1.6mm to 10mm'
                }
            },
            relatedProducts: [16]
        },
        {
            id: 21,
            name: 'Set Screws',
            slug: 'set-screws-din-913',
            category: 'screws',
            description: 'Grub set screws with cup point for locking components onto shafts.',
            standards: ['DIN 913', 'DIN 914', 'DIN 915', 'DIN 916'],
            material: ['alloy-steel', 'stainless-steel'],
            sizeRange: 'M3 - M24',
            lengthRange: '4mm - 100mm',
            surfaceTreatment: ['Black Oxide'],
            icon: '🌀',
            details: {
                fullDescription: 'Set screws (grub screws) are used to secure one component against another, typically preventing relative rotation. Available with cup point (DIN 913), cone point (DIN 914), dog point (DIN 915), and flat point (DIN 916).',
                applications: ['Power Transmission', 'Gears', 'Pulleys', 'Machine Assembly'],
                specifications: {
                    standard: 'DIN 913 / DIN 914 / DIN 915 / DIN 916',
                    productName: 'Socket Set Screws',
                    nominalDiameter: 'M3 to M24'
                }
            },
            relatedProducts: [9]
        },
        {
            id: 22,
            name: 'Anchor Bolts',
            slug: 'anchor-bolts',
            category: 'bolts',
            description: 'Foundation anchor bolts for securing structures to concrete.',
            standards: ['ANSI'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: '1/2" - 2"',
            lengthRange: '100mm - 1000mm',
            grades: ['4.6', '5.8', '8.8'],
            surfaceTreatment: ['Hot Dip Galvanized', 'Zinc Plated'],
            icon: '🔩',
            details: {
                fullDescription: 'Anchor bolts are embedded in concrete to connect steel structures, columns, and equipment to the concrete foundation. Various types available including bent anchor bolts, wedge anchors, sleeve anchors.',
                applications: ['Building Foundations', 'Bridge Construction', 'Heavy Equipment Mounting', 'Structural Steel'],
                specifications: {
                    productName: 'Foundation Anchor Bolts',
                    sizes: '1/2 inch to 2 inch diameter'
                }
            },
            relatedProducts: [1, 3]
        },
        {
            id: 23,
            name: 'Eye Bolts',
            slug: 'eye-bolts-din-444',
            category: 'bolts',
            description: 'Shoulder type eye bolts for lifting and rigging applications.',
            standards: ['DIN 444'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: 'M6 - M42',
            grades: ['4.6', '8.8'],
            surfaceTreatment: ['Zinc Plated', 'Self-color'],
            icon: '🔩',
            details: {
                fullDescription: 'DIN 444 shoulder eye bolts have a forged ring for attaching ropes, cables, and chains for lifting and load bearing applications. Shouldered design provides better strength than plain pattern.',
                applications: ['Lifting', 'Rigging', 'Crane Hooks', 'Load Securing', 'Marine'],
                specifications: {
                    standard: 'DIN 444',
                    productName: 'Shoulder Eye Bolts',
                    nominalDiameter: 'M6 to M42'
                }
            },
            relatedProducts: [1]
        },
        {
            id: 24,
            name: 'U-Bolts',
            slug: 'u-bolts',
            category: 'bolts',
            description: 'U-shaped bolts for securing pipe and tube to structures.',
            standards: ['Custom'],
            material: ['carbon-steel', 'stainless-steel'],
            sizeRange: '1/4" - 1-1/2" pipe',
            surfaceTreatment: ['Zinc Plated', 'Hot Dip Galvanized'],
            icon: '🔩',
            details: {
                fullDescription: 'U-bolts are U-shaped with threads on both ends, used primarily to support pipework or tubing. Can also be used for rope and cable attachments.',
                applications: ['Plumbing', 'Pipe Support', 'Exhaust Systems', 'Construction', 'Marine'],
                specifications: {
                    productName: 'U-Bolts',
                    pipeSizes: '1/4 inch to 1-1/2 inch'
                }
            },
            relatedProducts: [1, 3]
        },
        {
            id: 25,
            name: 'Partial Threaded Studs',
            slug: 'double-end-studs-din-938',
            category: 'threaded-rods',
            description: 'Double end threaded studs for bolted connections.',
            standards: ['DIN 938', 'DIN 939'],
            material: ['carbon-steel', 'alloy-steel', 'stainless-steel'],
            sizeRange: 'M6 - M48',
            lengthRange: '50mm - 500mm',
            grades: ['4.8', '8.8', '10.9'],
            surfaceTreatment: ['Plain', 'Zinc Plated', 'Black Oxide'],
            icon: '📏',
            details: {
                fullDescription: 'DIN 938 and DIN 939 double-end threaded studs have threads on both ends for securing into a tapped base with a nut on the opposite end. Used in machinery and structural applications.',
                applications: ['Mechanical Engineering', 'Pressure Vessels', 'Flange Connections', 'Heavy Machinery'],
                specifications: {
                    standard: 'DIN 938 / DIN 939',
                    productName: 'Double End Threaded Studs',
                    nominalDiameter: 'M6 to M48'
                }
            },
            relatedProducts: [10]
        }
    ]
};

function getProducts() {
    return productsData.products;
}

function getProductBySlug(slug) {
    return productsData.products.find(p => p.slug === slug);
}

function getProductById(id) {
    return productsData.products.find(p => p.id === id);
}

function getProductsByCategory(category) {
    if (!category || category === '') {
        return productsData.products;
    }
    return productsData.products.filter(p => p.category === category);
}

function getCategories() {
    return productsData.categories;
}

function getStandards() {
    return productsData.standards;
}

function getMaterials() {
    return productsData.materials;
}

function getRelatedProducts(ids) {
    return ids.map(id => getProductById(id)).filter(Boolean);
}

window.ProductAPI = {
    getProducts,
    getProductBySlug,
    getProductById,
    getProductsByCategory,
    getCategories,
    getStandards,
    getMaterials,
    getRelatedProducts
};
