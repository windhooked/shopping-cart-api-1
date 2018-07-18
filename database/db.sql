# Assume shopping cart system allows users to shop in session and without authentication
DROP TABLE IF EXISTS Item;
DROP TABLE IF EXISTS Order;
DROP TABLE IF EXISTS Customer;
DROP TABLE IF EXISTS Promotion;

CREATE TABLE Item
(
    item_id SERIAL PRIMARY KEY,
    promo_id SERIAL,
    name VARCHAR(50) NOT NULL,
    stock INTEGER NOT NULL,
    price INTEGER NOT NULL,
    FOREIGN KEY (promo_id) REFERENCES Promotion (promo_id) ON DELETE CASCADE
);
# Fetching all available items in the store
# SELECT * FROM Item WHERE stock > 0

CREATE TABLE Order 
(
	order_id SERIAL PRIMARY KEY,
	cust_id SERIAL NOT NULL,
	item_id SERIAL NOT NULL,
	placed BOOLEAN NOT NULL DEFAULT FALSE,
	FOREIGN KEY (cust_id) REFERENCES Customer (cust_id) ON DELETE CASCADE,
	FOREIGN KEY (item_id) REFERENCES Item (item_id) ON DELETE CASCADE
);
# All cart items are from the order table mapped to a customer id with his/her item
# SELECT * FROM Order WHERE customer.cust_id = order.cust_id

CREATE TABLE Customer
(
	cust_id SERIAL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	post_address VARCHAR(50) NOT NULL
);
# How to fetch all items in a cart of a given user?
# 


CREATE TABLE Promotion
(
    promo_id SERIAL PRIMARY KEY,
    discounted_item_id SERIAL NOT NULL,
    required_no_of_items INTEGER NOT NULL,
    discounted_price INTEGER NOT NULL,
    -- discount_percentage INTEGER NOT NULL CHECK (discount_percentage >= 0 AND discount_percentage <= 100),
    FOREIGN KEY (discounted_item_id) REFERENCES Item (item_id) ON DELETE CASCADE
);
# Promotions should be applied when the customer views their cart
# Need to join Promotion and Item to apply promotions

# Note - better to load all data and then create the index

# Adding Index
# Add index on fields/columns that are commonly used in the 'WHERE' or 'Group By' clauses
# If indexing joins, index the field on the left hand side of the assignment
CREATE INDEX item_idx ON Item (item_id);