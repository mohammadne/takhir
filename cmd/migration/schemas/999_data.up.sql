-- -- Insert translations for categories
-- INSERT INTO translations (language, translation) VALUES 
--     ('en', 'Laptops'),
--     ('fa', 'لپ‌تاپ‌ها'),
--     ('en', 'Mobiles'),
--     ('fa', 'موبایل‌ها'),
--     ('en', 'Books'),
--     ('fa', 'کتاب‌ها');

-- -- Insert categories
-- INSERT INTO categories DEFAULT VALUES;
-- INSERT INTO categories DEFAULT VALUES;
-- INSERT INTO categories DEFAULT VALUES;

-- -- Link categories with translations using inner queries
-- INSERT INTO categories_translations (category_id, translation_id) VALUES
--     (1, (SELECT id FROM translations WHERE translation = 'Laptops')),
--     (1, (SELECT id FROM translations WHERE translation = 'لپ‌تاپ‌ها')),
--     (2, (SELECT id FROM translations WHERE translation = 'Mobiles')),
--     (2, (SELECT id FROM translations WHERE translation = 'موبایل‌ها')),
--     (3, (SELECT id FROM translations WHERE translation = 'Books')),
--     (3, (SELECT id FROM translations WHERE translation = 'کتاب‌ها'));

-- -- Insert translations for product names and descriptions
-- INSERT INTO translations (language, translation) VALUES 
--     -- Laptops
--     ('en', 'Dell XPS 15'), ('fa', 'دل ایکس‌پی‌اس ۱۵'),
--     ('en', 'A powerful laptop with high-end specs'), ('fa', 'یک لپ‌تاپ قدرتمند با مشخصات بالا'),
--     ('en', 'MacBook Pro 14'), ('fa', 'مک‌بوک پرو ۱۴'),
--     ('en', 'Apple’s premium laptop for professionals'), ('fa', 'لپ‌تاپ حرفه‌ای اپل برای افراد متخصص'),

--     -- Mobiles
--     ('en', 'iPhone 15'), ('fa', 'آیفون ۱۵'),
--     ('en', 'The latest iPhone with advanced features'), ('fa', 'جدیدترین آیفون با ویژگی‌های پیشرفته'),
--     ('en', 'Samsung Galaxy S23'), ('fa', 'سامسونگ گلکسی S23'),
--     ('en', 'A flagship Android phone from Samsung'), ('fa', 'یک گوشی پرچم‌دار اندرویدی از سامسونگ'),

--     -- Books
--     ('en', 'The Pragmatic Programmer'), ('fa', 'برنامه‌نویس عمل‌گرا'),
--     ('en', 'A must-read book for every programmer'), ('fa', 'یک کتاب ضروری برای هر برنامه‌نویس'),
--     ('en', 'Clean Code'), ('fa', 'کد تمیز'),
--     ('en', 'Best practices for writing clean and efficient code'), ('fa', 'بهترین روش‌ها برای نوشتن کد تمیز و بهینه');

-- -- Insert products using inner queries for category_id
-- INSERT INTO products (category_id) VALUES 
--     ((SELECT id FROM categories LIMIT 1 OFFSET 0)), -- Laptops
--     ((SELECT id FROM categories LIMIT 1 OFFSET 0)), 
--     ((SELECT id FROM categories LIMIT 1 OFFSET 1)), -- Mobiles
--     ((SELECT id FROM categories LIMIT 1 OFFSET 1)), 
--     ((SELECT id FROM categories LIMIT 1 OFFSET 2)), -- Books
--     ((SELECT id FROM categories LIMIT 1 OFFSET 2));

-- -- Link product translations using inner queries
-- INSERT INTO products_translations (type, product_id, translation_id) VALUES
--     -- Dell XPS 15
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 0), (SELECT id FROM translations WHERE translation = 'Dell XPS 15')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 0), (SELECT id FROM translations WHERE translation = 'A powerful laptop with high-end specs')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 0), (SELECT id FROM translations WHERE translation = 'دل ایکس‌پی‌اس ۱۵')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 0), (SELECT id FROM translations WHERE translation = 'یک لپ‌تاپ قدرتمند با مشخصات بالا')),

--     -- MacBook Pro 14
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 1), (SELECT id FROM translations WHERE translation = 'MacBook Pro 14')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 1), (SELECT id FROM translations WHERE translation = 'Apple’s premium laptop for professionals')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 1), (SELECT id FROM translations WHERE translation = 'مک‌بوک پرو ۱۴')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 1), (SELECT id FROM translations WHERE translation = 'لپ‌تاپ حرفه‌ای اپل برای افراد متخصص')),

--     -- iPhone 15
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 2), (SELECT id FROM translations WHERE translation = 'iPhone 15')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 2), (SELECT id FROM translations WHERE translation = 'The latest iPhone with advanced features')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 2), (SELECT id FROM translations WHERE translation = 'آیفون ۱۵')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 2), (SELECT id FROM translations WHERE translation = 'جدیدترین آیفون با ویژگی‌های پیشرفته')),

--     -- Samsung Galaxy S23
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 3), (SELECT id FROM translations WHERE translation = 'Samsung Galaxy S23')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 3), (SELECT id FROM translations WHERE translation = 'A flagship Android phone from Samsung')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 3), (SELECT id FROM translations WHERE translation = 'سامسونگ گلکسی S23')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 3), (SELECT id FROM translations WHERE translation = 'یک گوشی پرچم‌دار اندرویدی از سامسونگ')),

--     -- The Pragmatic Programmer
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 4), (SELECT id FROM translations WHERE translation = 'The Pragmatic Programmer')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 4), (SELECT id FROM translations WHERE translation = 'A must-read book for every programmer')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 4), (SELECT id FROM translations WHERE translation = 'برنامه‌نویس عمل‌گرا')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 4), (SELECT id FROM translations WHERE translation = 'یک کتاب ضروری برای هر برنامه‌نویس')),

--     -- Clean Code
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 5), (SELECT id FROM translations WHERE translation = 'Clean Code')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 5), (SELECT id FROM translations WHERE translation = 'Best practices for writing clean and efficient code')),
--     ('name', (SELECT id FROM products LIMIT 1 OFFSET 5), (SELECT id FROM translations WHERE translation = 'کد تمیز')),
--     ('description', (SELECT id FROM products LIMIT 1 OFFSET 5), (SELECT id FROM translations WHERE translation = 'بهترین روش‌ها برای نوشتن کد تمیز و بهینه'));

-- -- Insert inventories using inner queries
-- INSERT INTO inventories (product_id, stock, price) VALUES
--     ((SELECT id FROM products LIMIT 1 OFFSET 0), 10, 1500),
--     ((SELECT id FROM products LIMIT 1 OFFSET 1), 8, 2000),
--     ((SELECT id FROM products LIMIT 1 OFFSET 2), 15, 999),
--     ((SELECT id FROM products LIMIT 1 OFFSET 3), 12, 799),
--     ((SELECT id FROM products LIMIT 1 OFFSET 4), 20, 50),
--     ((SELECT id FROM products LIMIT 1 OFFSET 5), 25, 45);

-- -- -- retrieve products of 'Laptop' category
-- -- SELECT t1.translation FROM translations t1
-- -- INNER JOIN products_translations pt ON pt.translation_id = t1.id
-- -- INNER JOIN products p ON p.id = pt.product_id
-- -- INNER JOIN categories c ON c.id = p.category_id
-- -- INNER JOIN categories_translations ct ON ct.category_id = c.id
-- -- INNER JOIN translations t2 ON t2.id = ct.translation_id
-- -- WHERE t2.translation = 'Laptops' AND pt.type = 'name' and t1.language = 'en';
