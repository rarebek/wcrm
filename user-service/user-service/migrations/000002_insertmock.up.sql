INSERT INTO owners (id, full_name, company_name, email, password, avatar, tax, created_at, updated_at)
VALUES 
    ('11111111-1111-1111-1111-111111111111', 'John Doe', 'Doe-Corporation', 'john.doe@example.com', 'password123', 'avatar1.jpg', 20, NOW(), NOW()),
    ('22222222-2222-2222-2222-222222222222', 'Jane Smith', 'Smith-Enterprises', 'jane.smith@example.com', 'password456', 'avatar2.jpg', 15, NOW(), NOW()),
    ('33333333-3333-3333-3333-333333333333', 'Michael Johnson', 'Johnson-Co', 'michael.johnson@example.com', 'password789', 'avatar3.jpg', 25, NOW(), NOW()),
    ('44444444-4444-4444-4444-444444444444', 'Emily Brown', 'Brown-Industries', 'emily.brown@example.com', 'passwordabc', 'avatar4.jpg', 18, NOW(), NOW()),
    ('55555555-5555-5555-5555-555555555555', 'David Wilson', 'Wilson-Corp', 'david.wilson@example.com', 'passworddef', 'avatar5.jpg', 22, NOW(), NOW());

-- Mock data for workers table
INSERT INTO workers (id, full_name, login_key, password, owner_id, created_at, updated_at)
VALUES 
    ('11111111-1111-1111-1111-111111111112', 'Alice Johnson', 'alice123', 'workerpass123', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
    ('22222222-2222-2222-2222-222222222223', 'Bob Williams', 'bob456', 'workerpass456', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
    ('33333333-3333-3333-3333-333333333334', 'Carol Davis', 'carol789', 'workerpass789', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
    ('44444444-4444-4444-4444-444444444445', 'Ethan Martin', 'ethanabc', 'workerpassabc', '33333333-3333-3333-3333-333333333333', NOW(), NOW()),
    ('55555555-5555-5555-5555-555555555556', 'Grace Anderson', 'gracedef', 'workerpassdef', '33333333-3333-3333-3333-333333333333', NOW(), NOW());

-- Mock data for geolocations table
INSERT INTO geolocations (latitude, longitude, owner_id)
VALUES 
    (40.7128, -74.0060, '11111111-1111-1111-1111-111111111111'),
    (34.0522, -118.2437, '22222222-2222-2222-2222-222222222222'),
    (51.5074, -0.1278, '33333333-3333-3333-3333-333333333333'),
    (48.8566, 2.3522, '44444444-4444-4444-4444-444444444444'),
    (37.7749, -122.4194, '55555555-5555-5555-5555-555555555555');