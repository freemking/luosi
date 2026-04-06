(function() {
    let scene, camera, renderer, labelRenderer, controls, currentModel;
    let isWireframe = false;
    let currentColor = '#B8C0C8';
    let isMetric = true;
    let currentModelType = 'screw';
    let envMap = null;
    let labels = [];
    
    const inchToMm = 25.4;
    
    const metricThreadData = {
        3: { pitch: 0.5, minorDia: 2.387, pitchDia: 2.675 },
        3.5: { pitch: 0.6, minorDia: 2.764, pitchDia: 3.110 },
        4: { pitch: 0.7, minorDia: 3.141, pitchDia: 3.545 },
        5: { pitch: 0.8, minorDia: 4.019, pitchDia: 4.480 },
        6: { pitch: 1.0, minorDia: 4.773, pitchDia: 5.350 },
        7: { pitch: 1.0, minorDia: 5.773, pitchDia: 6.350 },
        8: { pitch: 1.25, minorDia: 6.466, pitchDia: 7.188 },
        10: { pitch: 1.5, minorDia: 8.160, pitchDia: 9.026 },
        12: { pitch: 1.75, minorDia: 9.853, pitchDia: 10.863 },
        14: { pitch: 2.0, minorDia: 11.546, pitchDia: 12.701 },
        16: { pitch: 2.0, minorDia: 13.546, pitchDia: 14.701 },
        18: { pitch: 2.5, minorDia: 15.066, pitchDia: 16.376 },
        20: { pitch: 2.5, minorDia: 17.066, pitchDia: 18.376 },
        22: { pitch: 2.5, minorDia: 19.066, pitchDia: 20.376 },
        24: { pitch: 3.0, minorDia: 20.319, pitchDia: 22.051 }
    };
    
    const hexHeadData = {
        3: { widthAcrossFlats: 5.5, headHeight: 2 },
        4: { widthAcrossFlats: 7, headHeight: 2.8 },
        5: { widthAcrossFlats: 8, headHeight: 3.5 },
        6: { widthAcrossFlats: 10, headHeight: 4 },
        8: { widthAcrossFlats: 13, headHeight: 5.3 },
        10: { widthAcrossFlats: 16, headHeight: 6.4 },
        12: { widthAcrossFlats: 18, headHeight: 7.5 },
        14: { widthAcrossFlats: 21, headHeight: 8.8 },
        16: { widthAcrossFlats: 24, headHeight: 10 },
        18: { widthAcrossFlats: 27, headHeight: 11.5 },
        20: { widthAcrossFlats: 30, headHeight: 12.5 },
        22: { widthAcrossFlats: 34, headHeight: 14 },
        24: { widthAcrossFlats: 36, headHeight: 15 }
    };
    
    function init() {
        const container = document.getElementById('three-container');
        if (!container) return;
        
        const width = container.clientWidth;
        const height = container.clientHeight;
        
        scene = new THREE.Scene();
        scene.background = new THREE.Color(0x0f1f2e);
        
        camera = new THREE.PerspectiveCamera(45, width / height, 0.1, 1000);
        camera.position.set(50, 35, 50);
        
        renderer = new THREE.WebGLRenderer({ 
            antialias: true,
            powerPreference: 'high-performance'
        });
        renderer.setSize(width, height);
        renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
        renderer.shadowMap.enabled = true;
        renderer.shadowMap.type = THREE.PCFSoftShadowMap;
        renderer.toneMapping = THREE.ACESFilmicToneMapping;
        renderer.toneMappingExposure = 1.2;
        container.appendChild(renderer.domElement);
        
        labelRenderer = new THREE.CSS2DRenderer();
        labelRenderer.setSize(width, height);
        labelRenderer.domElement.style.position = 'absolute';
        labelRenderer.domElement.style.top = '0px';
        labelRenderer.domElement.style.pointerEvents = 'none';
        container.appendChild(labelRenderer.domElement);
        
        controls = new THREE.OrbitControls(camera, renderer.domElement);
        controls.enableDamping = true;
        controls.dampingFactor = 0.05;
        controls.minPolarAngle = 0;
        controls.maxPolarAngle = Math.PI / 2;
        controls.minDistance = 30;
        controls.maxDistance = 150;
        controls.maxTargetRadius = 50;
        
        createEnvironmentMap();
        setupLighting();
        
        const gridHelper = new THREE.GridHelper(100, 20, 0x00DEEA, 0x1f2937);
        gridHelper.position.y = -25;
        scene.add(gridHelper);
        
        createScrew();
        animate();
        
        window.addEventListener('resize', onWindowResize);
    }
    
    function createEnvironmentMap() {
        const size = 256;
        const data = new Float32Array(size * size * 4);
        
        for (let i = 0; i < size; i++) {
            for (let j = 0; j < size; j++) {
                const index = (i * size + j) * 4;
                const y = i / size;
                
                data[index] = 0.4 + y * 0.4;
                data[index + 1] = 0.45 + y * 0.35;
                data[index + 2] = 0.5 + y * 0.3;
                data[index + 3] = 1.0;
            }
        }
        
        const texture = new THREE.DataTexture(data, size, size, THREE.RGBAFormat, THREE.FloatType);
        texture.mapping = THREE.EquirectangularReflectionMapping;
        texture.needsUpdate = true;
        envMap = texture;
        scene.environment = envMap;
    }
    
    function setupLighting() {
        const ambientLight = new THREE.AmbientLight(0xffffff, 0.3);
        scene.add(ambientLight);
        
        const keyLight = new THREE.DirectionalLight(0xffffff, 1.5);
        keyLight.position.set(80, 120, 80);
        keyLight.castShadow = true;
        keyLight.shadow.mapSize.width = 2048;
        keyLight.shadow.mapSize.height = 2048;
        scene.add(keyLight);
        
        const fillLight = new THREE.DirectionalLight(0xffffff, 0.8);
        fillLight.position.set(-80, 60, -60);
        scene.add(fillLight);
        
        const rimLight = new THREE.DirectionalLight(0xffffff, 0.5);
        rimLight.position.set(0, 30, -100);
        scene.add(rimLight);
        
        const topLight = new THREE.DirectionalLight(0xffffff, 0.4);
        topLight.position.set(0, 120, 0);
        scene.add(topLight);
        
        const warmLight = new THREE.PointLight(0xFFE4C4, 0.4, 300);
        warmLight.position.set(100, 50, 0);
        scene.add(warmLight);
        
        const coolLight = new THREE.PointLight(0xE0FFFF, 0.3, 300);
        coolLight.position.set(-100, 50, 0);
        scene.add(coolLight);
    }
    
    function onWindowResize() {
        const container = document.getElementById('three-container');
        if (!container) return;
        
        const width = container.clientWidth;
        const height = container.clientHeight;
        
        camera.aspect = width / height;
        camera.updateProjectionMatrix();
        renderer.setSize(width, height);
        labelRenderer.setSize(width, height);
    }
    
    function animate() {
        requestAnimationFrame(animate);
        controls.update();
        renderer.render(scene, camera);
        labelRenderer.render(scene, camera);
    }
    
    function clearModel() {
        if (currentModel) {
            scene.remove(currentModel);
            currentModel.traverse(function(child) {
                if (child.geometry) child.geometry.dispose();
                if (child.material) child.material.dispose();
            });
        }
        labels.forEach(label => {
            scene.remove(label);
        });
        labels = [];
    }
    
    function createLabel(text, position, color) {
        color = color || '#00DEEA';
        const div = document.createElement('div');
        div.className = 'dimension-label';
        div.textContent = text;
        div.style.cssText = 'background: rgba(0,0,0,0.7); color: ' + color + '; padding: 4px 8px; border-radius: 4px; font-size: 12px; font-family: Outfit, monospace; font-weight: 600; white-space: nowrap; border: 1px solid rgba(0,222,234,0.3);';
        
        const label = new THREE.CSS2DObject(div);
        label.position.copy(position);
        scene.add(label);
        labels.push(label);
        return label;
    }
    
    function createOptimizedThreadGeometry(outerR, innerR, height, pitch, segments) {
        segments = segments || 64;
        const geometry = new THREE.CylinderGeometry(outerR, outerR, height, segments, Math.ceil(height / pitch * 4));
        const positions = geometry.attributes.position;
        
        const threadDepth = outerR - innerR;
        const profileDepth = threadDepth * 0.85;
        
        for (let i = 0; i < positions.count; i++) {
            const x = positions.getX(i);
            const y = positions.getY(i);
            const z = positions.getZ(i);
            
            const angle = y / pitch * Math.PI * 2;
            const normalizedY = Math.abs(y) / (height / 2);
            
            const threadPhase = Math.sin(angle * 2);
            const depth = profileDepth * (1 - Math.abs(threadPhase) * 0.3);
            
            const threadFactor = normalizedY < 0.85 ? 1 : (1 - normalizedY) / 0.15;
            const radius = outerR - depth * (0.5 + 0.5 * threadPhase) * threadFactor;
            
            const currentRadius = Math.sqrt(x * x + z * z);
            
            if (currentRadius > 0.001) {
                const scale = radius / currentRadius;
                positions.setX(i, x * scale);
                positions.setZ(i, z * scale);
            }
        }
        
        geometry.computeVertexNormals();
        return geometry;
    }
    
    function createMaterial(color) {
        const isDark = parseInt(color.slice(1), 16) < 0x808080;
        return new THREE.MeshStandardMaterial({
            color: color,
            metalness: 1.0,
            roughness: isDark ? 0.5 : 0.15,
            wireframe: isWireframe,
            envMap: envMap,
            envMapIntensity: isDark ? 1.5 : 2.0,
            flatShading: false,
            side: THREE.FrontSide
        });
    }
    
    function createScrew() {
        clearModel();
        currentModelType = 'screw';
        
        const diameter = parseFloat(document.getElementById('screw-diameter').value);
        const length = parseFloat(document.getElementById('screw-length').value);
        const pitch = parseFloat(document.getElementById('screw-pitch').value);
        const headType = document.getElementById('screw-head').value;
        
        const group = new THREE.Group();
        const material = createMaterial(currentColor);
        
        const scale = 0.5;
        const r = diameter * scale / 2;
        const h = length * scale;
        
        const threadData = metricThreadData[diameter] || { minorDia: diameter * 0.85, pitch: pitch };
        const minorR = threadData.minorDia * scale / 2;
        
        let headHeight;
        switch(headType) {
            case 'hex':
                const hexData = hexHeadData[diameter] || { widthAcrossFlats: r * 4, headHeight: r * 2 };
                headHeight = hexData.headHeight * scale;
                break;
            case 'socket':
                headHeight = r * 2;
                break;
            case 'flat':
                headHeight = r * 2;
                break;
            case 'pan':
                headHeight = r * 1.2;
                break;
            default:
                headHeight = r * 1.2;
        }
        
        const headGroup = createHead(headType, r, scale, material, diameter);
        headGroup.position.y = 0;
        group.add(headGroup);
        
        const transitionHeight = r * 0.3;
        const transitionGeometry = new THREE.CylinderGeometry(r, r, transitionHeight, 32);
        const transition = new THREE.Mesh(transitionGeometry, material);
        transition.position.y = -transitionHeight / 2;
        group.add(transition);
        
        const threadGeometry = createOptimizedThreadGeometry(r, minorR, h, pitch * scale, 64);
        const thread = new THREE.Mesh(threadGeometry, material);
        thread.position.y = -h / 2 - transitionHeight;
        group.add(thread);
        
        const tipHeight = r * 1.5;
        const tipGeometry = new THREE.ConeGeometry(r, tipHeight, 32);
        const tip = new THREE.Mesh(tipGeometry, material);
        tip.position.y = -h - transitionHeight - tipHeight / 2;
        tip.rotation.x = Math.PI;
        group.add(tip);
        
        const labelY = headHeight + 5;
        createLabel('M' + diameter + 'x' + length, new THREE.Vector3(r * 3, labelY, 0));
        createLabel('P' + pitch, new THREE.Vector3(r * 3, labelY - 5, 0));
        
        currentModel = group;
        scene.add(currentModel);
        
        updateSpecDisplay();
    }
    
    function createHead(type, r, scale, material, diameter) {
        const headGroup = new THREE.Group();
        
        switch(type) {
            case 'hex':
                const hexData = hexHeadData[diameter] || { widthAcrossFlats: r * 4, headHeight: r * 2 };
                const hexR = hexData.widthAcrossFlats * scale / 2;
                const hexH = hexData.headHeight * scale;
                
                const hexShape = new THREE.Shape();
                for (let i = 0; i < 6; i++) {
                    const angle = (i / 6) * Math.PI * 2 - Math.PI / 2;
                    const x = Math.cos(angle) * hexR;
                    const y = Math.sin(angle) * hexR;
                    if (i === 0) hexShape.moveTo(x, y);
                    else hexShape.lineTo(x, y);
                }
                hexShape.closePath();
                
                const extrudeSettings = { depth: hexH, bevelEnabled: false };
                const hexGeometry = new THREE.ExtrudeGeometry(hexShape, extrudeSettings);
                hexGeometry.rotateX(-Math.PI / 2);
                const hex = new THREE.Mesh(hexGeometry, material);
                hex.position.y = 0;
                headGroup.add(hex);
                
                const chamferGeometry = new THREE.CylinderGeometry(hexR * 0.85, hexR, hexH * 0.15, 6);
                const chamfer = new THREE.Mesh(chamferGeometry, material);
                chamfer.position.y = hexH - hexH * 0.075;
                headGroup.add(chamfer);
                break;
                
            case 'socket':
                const socketR = r * 1.5;
                const socketH = r * 2;
                
                const socketGeometry = new THREE.CylinderGeometry(socketR, socketR, socketH, 32);
                const socketMesh = new THREE.Mesh(socketGeometry, material);
                socketMesh.position.y = socketH / 2;
                headGroup.add(socketMesh);
                
                const socketChamferGeo = new THREE.CylinderGeometry(socketR * 0.85, socketR, socketH * 0.1, 32);
                const socketChamfer = new THREE.Mesh(socketChamferGeo, material);
                socketChamfer.position.y = socketH + socketH * 0.05;
                headGroup.add(socketChamfer);
                
                const hexDriveShape = new THREE.Shape();
                for (let i = 0; i < 6; i++) {
                    const angle = (i / 6) * Math.PI * 2 - Math.PI / 2;
                    const x = Math.cos(angle) * socketR * 0.5;
                    const y = Math.sin(angle) * socketR * 0.5;
                    if (i === 0) hexDriveShape.moveTo(x, y);
                    else hexDriveShape.lineTo(x, y);
                }
                hexDriveShape.closePath();
                
                const hexDriveGeo = new THREE.ExtrudeGeometry(hexDriveShape, { depth: socketH * 0.2, bevelEnabled: false });
                hexDriveGeo.rotateX(-Math.PI / 2);
                const hexDrive = new THREE.Mesh(hexDriveGeo, material);
                hexDrive.position.y = socketH + socketH * 0.1;
                headGroup.add(hexDrive);
                break;
                
            case 'flat':
                const flatR = r * 2.5;
                const flatH = r * 2;
                
                const flatGeometry = new THREE.CylinderGeometry(flatR, r, flatH, 32);
                const flatMesh = new THREE.Mesh(flatGeometry, material);
                flatMesh.position.y = flatH / 2;
                headGroup.add(flatMesh);
                
                const flatChamferGeo = new THREE.TorusGeometry(flatR * 0.95, r * 0.1, 8, 32);
                const flatChamfer = new THREE.Mesh(flatChamferGeo, material);
                flatChamfer.rotation.x = Math.PI / 2;
                flatChamfer.position.y = flatH;
                headGroup.add(flatChamfer);
                break;
                
            case 'pan':
                const panR = r * 2;
                
                const panGeometry = new THREE.SphereGeometry(panR, 32, 16, 0, Math.PI * 2, 0, Math.PI / 2);
                const panMesh = new THREE.Mesh(panGeometry, material);
                panMesh.scale.y = 0.6;
                panMesh.position.y = 0;
                headGroup.add(panMesh);
                
                const panFilletGeo = new THREE.TorusGeometry(r * 1.05, r * 0.15, 8, 32);
                const panFillet = new THREE.Mesh(panFilletGeo, material);
                panFillet.rotation.x = Math.PI / 2;
                panFillet.position.y = 0;
                headGroup.add(panFillet);
                break;
        }
        
        return headGroup;
    }
    
    function createNut() {
        clearModel();
        currentModelType = 'nut';
        
        const diameter = parseFloat(document.getElementById('nut-diameter').value);
        const height = parseFloat(document.getElementById('nut-height').value);
        const nutType = document.getElementById('nut-type').value;
        
        const group = new THREE.Group();
        const material = createMaterial(currentColor);
        
        const scale = 0.5;
        const r = diameter * scale / 2;
        const h = height * scale;
        
        switch(nutType) {
            case 'hex':
                const hexR = r * 2;
                const hexShape = new THREE.Shape();
                for (let i = 0; i < 6; i++) {
                    const angle = (i / 6) * Math.PI * 2 - Math.PI / 2;
                    const x = Math.cos(angle) * hexR;
                    const y = Math.sin(angle) * hexR;
                    if (i === 0) hexShape.moveTo(x, y);
                    else hexShape.lineTo(x, y);
                }
                hexShape.closePath();
                
                const holeShape = new THREE.Path();
                holeShape.absellipse(0, 0, r, r, 0, Math.PI * 2, true);
                hexShape.holes.push(holeShape);
                
                const extrudeSettings = { depth: h, bevelEnabled: false };
                const hexGeometry = new THREE.ExtrudeGeometry(hexShape, extrudeSettings);
                hexGeometry.rotateX(-Math.PI / 2);
                const hex = new THREE.Mesh(hexGeometry, material);
                hex.position.y = -h / 2;
                group.add(hex);
                break;
                
            case 'nylon':
                const lockHexShape = new THREE.Shape();
                for (let i = 0; i < 6; i++) {
                    const angle = (i / 6) * Math.PI * 2 - Math.PI / 2;
                    const x = Math.cos(angle) * r * 2;
                    const y = Math.sin(angle) * r * 2;
                    if (i === 0) lockHexShape.moveTo(x, y);
                    else lockHexShape.lineTo(x, y);
                }
                lockHexShape.closePath();
                
                const lockHole = new THREE.Path();
                lockHole.absellipse(0, 0, r, r, 0, Math.PI * 2, true);
                lockHexShape.holes.push(lockHole);
                
                const lockGeometry = new THREE.ExtrudeGeometry(lockHexShape, { depth: h * 1.2, bevelEnabled: false });
                lockGeometry.rotateX(-Math.PI / 2);
                const lockNut = new THREE.Mesh(lockGeometry, material);
                lockNut.position.y = -h * 0.6;
                group.add(lockNut);
                
                const nylonMaterial = new THREE.MeshStandardMaterial({ 
                    color: 0x3366CC, 
                    metalness: 0.0, 
                    roughness: 0.9,
                    flatShading: false
                });
                const nylonGeometry = new THREE.TorusGeometry(r * 1.1, r * 0.25, 16, 32);
                const nylon = new THREE.Mesh(nylonGeometry, nylonMaterial);
                nylon.rotation.x = Math.PI / 2;
                nylon.position.y = h * 0.35;
                group.add(nylon);
                break;
                
            case 'flange':
                const flangeHexShape = new THREE.Shape();
                for (let i = 0; i < 6; i++) {
                    const angle = (i / 6) * Math.PI * 2 - Math.PI / 2;
                    const x = Math.cos(angle) * r * 2;
                    const y = Math.sin(angle) * r * 2;
                    if (i === 0) flangeHexShape.moveTo(x, y);
                    else flangeHexShape.lineTo(x, y);
                }
                flangeHexShape.closePath();
                
                const flangeHole = new THREE.Path();
                flangeHole.absellipse(0, 0, r, r, 0, Math.PI * 2, true);
                flangeHexShape.holes.push(flangeHole);
                
                const flangeGeometry = new THREE.ExtrudeGeometry(flangeHexShape, { depth: h, bevelEnabled: false });
                flangeGeometry.rotateX(-Math.PI / 2);
                const flangeHex = new THREE.Mesh(flangeGeometry, material);
                flangeHex.position.y = -h / 2;
                group.add(flangeHex);
                
                const flangeRingGeometry = new THREE.CylinderGeometry(r * 3, r * 3, h * 0.3, 32);
                const flangeRing = new THREE.Mesh(flangeRingGeometry, material);
                flangeRing.position.y = -h / 2 - h * 0.15;
                group.add(flangeRing);
                break;
        }
        
        createLabel('M' + diameter, new THREE.Vector3(r * 3, 0, 0));
        createLabel('H' + height, new THREE.Vector3(r * 3, -5, 0));
        
        currentModel = group;
        scene.add(currentModel);
        
        updateSpecDisplay();
    }
    
    function updateSpecDisplay() {
        const specOutput = document.getElementById('spec-output');
        if (!specOutput) return;
        
        if (currentModelType === 'screw') {
            const diameter = parseFloat(document.getElementById('screw-diameter').value);
            const length = parseFloat(document.getElementById('screw-length').value);
            const pitch = parseFloat(document.getElementById('screw-pitch').value);
            const headType = document.getElementById('screw-head').value;
            
            const headNames = { hex: 'Hex Head', socket: 'Socket Head', flat: 'Flat Head', pan: 'Pan Head' };
            
            if (isMetric) {
                specOutput.textContent = 'M' + diameter + ' x ' + length + ' x ' + pitch + ' ' + headNames[headType] + ' Screw';
            } else {
                const inchDia = (diameter / inchToMm).toFixed(3);
                const inchLen = (length / inchToMm).toFixed(2);
                specOutput.textContent = inchDia + '" x ' + inchLen + '" ' + headNames[headType] + ' Screw';
            }
        } else {
            const diameter = parseFloat(document.getElementById('nut-diameter').value);
            const height = parseFloat(document.getElementById('nut-height').value);
            const nutType = document.getElementById('nut-type').value;
            
            const nutNames = { hex: 'Hex Nut', nylon: 'Nylon Lock Nut', flange: 'Flange Nut' };
            
            if (isMetric) {
                specOutput.textContent = 'M' + diameter + ' x ' + height + ' ' + nutNames[nutType];
            } else {
                const inchDia = (diameter / inchToMm).toFixed(3);
                specOutput.textContent = inchDia + '" ' + nutNames[nutType];
            }
        }
    }
    
    function updateValueDisplays() {
        if (isMetric) {
            document.getElementById('screw-diameter-val').textContent = 'M' + document.getElementById('screw-diameter').value;
            document.getElementById('screw-length-val').textContent = document.getElementById('screw-length').value + 'mm';
            document.getElementById('nut-diameter-val').textContent = 'M' + document.getElementById('nut-diameter').value;
            document.getElementById('nut-height-val').textContent = document.getElementById('nut-height').value + 'mm';
            
            document.getElementById('screw-pitch').innerHTML = '<option value="1.25">1.25mm (Standard)</option><option value="1.0">1.0mm (Fine)</option><option value="0.75">0.75mm (Extra Fine)</option>';
            document.getElementById('nut-pitch').innerHTML = '<option value="1.25">1.25mm (Standard)</option><option value="1.0">1.0mm (Fine)</option><option value="0.75">0.75mm (Extra Fine)</option>';
        } else {
            const screwDia = (parseFloat(document.getElementById('screw-diameter').value) / inchToMm).toFixed(3);
            const screwLen = (parseFloat(document.getElementById('screw-length').value) / inchToMm).toFixed(2);
            const nutDia = (parseFloat(document.getElementById('nut-diameter').value) / inchToMm).toFixed(3);
            const nutH = (parseFloat(document.getElementById('nut-height').value) / inchToMm).toFixed(2);
            
            document.getElementById('screw-diameter-val').textContent = screwDia + '"';
            document.getElementById('screw-length-val').textContent = screwLen + '"';
            document.getElementById('nut-diameter-val').textContent = nutDia + '"';
            document.getElementById('nut-height-val').textContent = nutH + '"';
            
            document.getElementById('screw-pitch').innerHTML = '<option value="1.588">16 TPI</option><option value="2.117">12 TPI</option><option value="2.822">9 TPI</option>';
            document.getElementById('nut-pitch').innerHTML = '<option value="1.588">16 TPI</option><option value="2.117">12 TPI</option><option value="2.822">9 TPI</option>';
        }
    }
    
    document.addEventListener('DOMContentLoaded', function() {
        init();
        
        if (!document.getElementById('three-container')) return;
        
        document.getElementById('show-screw').addEventListener('click', function() {
            document.getElementById('show-screw').classList.add('active');
            document.getElementById('show-nut').classList.remove('active');
            document.getElementById('screw-controls').style.display = 'block';
            document.getElementById('nut-controls').style.display = 'none';
            document.getElementById('model-type-display').textContent = 'SYM FASTENER';
            createScrew();
        });
        
        document.getElementById('show-nut').addEventListener('click', function() {
            document.getElementById('show-nut').classList.add('active');
            document.getElementById('show-screw').classList.remove('active');
            document.getElementById('nut-controls').style.display = 'block';
            document.getElementById('screw-controls').style.display = 'none';
            document.getElementById('model-type-display').textContent = 'SYM FASTENER';
            createNut();
        });
        
        document.getElementById('unit-metric').addEventListener('click', function() {
            document.getElementById('unit-metric').classList.add('active');
            document.getElementById('unit-imperial').classList.remove('active');
            isMetric = true;
            updateValueDisplays();
            updateSpecDisplay();
        });
        
        document.getElementById('unit-imperial').addEventListener('click', function() {
            document.getElementById('unit-imperial').classList.add('active');
            document.getElementById('unit-metric').classList.remove('active');
            isMetric = false;
            updateValueDisplays();
            updateSpecDisplay();
        });
        
        ['screw-diameter', 'screw-length', 'screw-pitch', 'screw-head'].forEach(function(id) {
            document.getElementById(id).addEventListener('input', function() {
                updateValueDisplays();
                createScrew();
            });
            document.getElementById(id).addEventListener('change', function() {
                updateValueDisplays();
                createScrew();
            });
        });
        
        ['nut-diameter', 'nut-height', 'nut-pitch', 'nut-type'].forEach(function(id) {
            document.getElementById(id).addEventListener('input', function() {
                updateValueDisplays();
                if (currentModelType === 'nut') createNut();
            });
            document.getElementById(id).addEventListener('change', function() {
                updateValueDisplays();
                if (currentModelType === 'nut') createNut();
            });
        });
        
        document.querySelectorAll('.color-btn').forEach(function(btn) {
            btn.addEventListener('click', function() {
                document.querySelectorAll('.color-btn').forEach(function(b) { b.classList.remove('active'); });
                btn.classList.add('active');
                currentColor = btn.dataset.color;
                if (currentModelType === 'screw') createScrew();
                else createNut();
            });
        });
        
        document.getElementById('reset-view').addEventListener('click', function() {
            camera.position.set(50, 35, 50);
            controls.reset();
        });
        
        document.getElementById('toggle-wireframe').addEventListener('click', function() {
            isWireframe = !isWireframe;
            document.getElementById('toggle-wireframe').classList.toggle('active');
            if (currentModelType === 'screw') createScrew();
            else createNut();
        });
        
        document.getElementById('export-model').addEventListener('click', function() {
            const exporter = new THREE.GLTFExporter();
            exporter.parse(scene, function(result) {
                const output = result instanceof ArrayBuffer ? result : JSON.stringify(result, null, 2);
                const blob = new Blob([output], { type: 'application/octet-stream' });
                const url = URL.createObjectURL(blob);
                const a = document.createElement('a');
                a.href = url;
                a.download = 'fastener_model.glb';
                a.click();
                URL.revokeObjectURL(url);
            }, { binary: true });
        });
    });
})();
