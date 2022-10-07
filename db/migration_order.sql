-- creating order table in postgres DATABASE
CREATE TABLE orders(
                      id SERIAL NOT NULL PRIMARY KEY,
                      ordered_at DATE NOT NULL DEFAULT CURRENT_DATE,
                      customer_id SERIAL NOT NULL,
                      CONSTRAINT fk_user FOREIGN KEY(customer_id) REFERENCES users(id) ON DELETE SET NULL 
);