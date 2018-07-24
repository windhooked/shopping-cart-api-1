DROP TABLE IF EXISTS Customer CASCADE;
DROP TABLE IF EXISTS Item CASCADE;
DROP TABLE IF EXISTS Promotion CASCADE;
DROP TABLE IF EXISTS Purchase_Order CASCADE;

CREATE TABLE Customer
(
    cust_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    post_address VARCHAR(50) NOT NULL
);

CREATE TABLE Item
(
    item_id SERIAL PRIMARY KEY,
    promo_id INTEGER,
    name VARCHAR(50) NOT NULL,
    stock INTEGER NOT NULL,
    price INTEGER NOT NULL
);

CREATE TABLE Promotion
(
    promo_id SERIAL PRIMARY KEY,
    required_item_id INTEGER NOT NULL,
    required_quantity INTEGER NOT NULL,
    discount_percentage INTEGER NOT NULL CHECK (discount_percentage >= 0 AND discount_percentage <= 100)
);
-- Promotions should be applied when the customer views their cart
-- Need to join Promotion and Item to apply promotions

CREATE TABLE Purchase_Order 
(
    purchase_order_id SERIAL PRIMARY KEY,
    cust_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    dispatched BOOLEAN NOT NULL DEFAULT FALSE
);
-- All cart items are from the PurchaseOrder table mapped to a customer id with his/her item
-- SELECT * FROM PurchaseOrder WHERE customer.cust_id = PurchaseOrder.cust_id

-- Adding foreign keys after table generation
ALTER TABLE Item ADD FOREIGN KEY (promo_id) REFERENCES Promotion (promo_id) ON DELETE CASCADE;
ALTER TABLE Promotion ADD FOREIGN KEY (required_item_id) REFERENCES Item (item_id) ON DELETE CASCADE;
ALTER TABLE Purchase_Order ADD FOREIGN KEY (cust_id) REFERENCES Customer (cust_id) ON DELETE CASCADE;
ALTER TABLE Purchase_Order ADD FOREIGN KEY (item_id) REFERENCES Item (item_id) ON DELETE CASCADE;

-- Adding test data
INSERT INTO Item (name, stock, price) VALUES ('Belts', 10, 20);
INSERT INTO Item (name, stock, price) VALUES ('Shirts', 5, 60);
INSERT INTO Item (name, stock, price) VALUES ('Suits', 2, 300);
INSERT INTO Item (name, stock, price) VALUES ('Trousers', 4, 70);
INSERT INTO Item (name, stock, price) VALUES ('Shoes', 1, 120);
INSERT INTO Item (name, stock, price) VALUES ('Ties', 8, 20);

INSERT INTO Promotion (required_item_id, required_quantity, discount_percentage) VALUES (4, 2, 15);
INSERT INTO Promotion (required_item_id, required_quantity, discount_percentage) VALUES (2, 2, 25);
INSERT INTO Promotion (required_item_id, required_quantity, discount_percentage) VALUES (2, 3, 50);

INSERT INTO Customer (first_name, last_name, post_address) VALUES ('John', 'Doe', 'Ryde, NSW');
INSERT INTO Customer (first_name, last_name, post_address) VALUES ('Bob', 'Williams', 'Liverpool, NSW');

UPDATE Item SET promo_id = 1 WHERE name = 'Belts';
UPDATE Item SET promo_id = 1 WHERE name = 'Shoes';
UPDATE Item SET promo_id = 2 WHERE name = 'Shirts';
UPDATE Item SET promo_id = 3 WHERE name = 'Ties';

-- Adding Index
-- Note: better to load all data and then create the index. Use Explain Analyse to detect bottlenecks in query.
-- Add index on fields/columns that are commonly used in the 'WHERE' or 'Group By' clauses
-- If indexing joins, index the field on the left hand side of the assignment
-- CREATE INDEX item_idx ON Item (item_id);