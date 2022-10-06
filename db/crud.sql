-- Event Sourcing
-- CQRS

-- Membuat Table
CREATE TABLE users (
	-- serial -> int dengan bisa auto generate secara berurutan
	-- not null -> data tidak boleh kosong
	-- primary key -> antiduplicate di suatu table
    id SERIAL NOT NULL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT
);
--

-- Comment di SQL
-- CRUD process
-- Create Read Update Delete
-- "" -> column name
-- '' -> value 
--

-- Insert Data
insert into users (first_name, last_name)
values 
--	('calman', 'tara')
--	('john', 'wick')
--	('nobi', 'nobita'),
--	('john', 'smith'),
--	('okta', 'lia')
	('calboy', '')
;
--

-- Read Data
-- * -> wild card untuk mendapatkan 
-- semua column di table. PROCESSNYA MAHAL
select u.id, u.last_name 
	from users u 
;
--

-- Conditional Statement
select u.id, u.first_name, u.last_name
	from users u 
where u.first_name = 'calman'
	or u.id  = 4
;
--

-- Conditional Statement Wild Card
select u.id, u.first_name, u.last_name
	from users u 
where u.first_name not like 'cal%'
;
--

-- Conditional Statement In
select u.id, u.first_name, u.last_name
	from users u 
--where u.id = 1 or u.id = 2 or u.id = 5
	where u.id in (1,2,5)
;
--

-- Update Data
-- untuk mengupdate, 
-- dan memerlukan conditional statement
update users 
set last_name = 'tara'
--where id = 6
where first_name = 'okta'
;
--

-- Delete Data
delete from users 
where id = 4
;
--

-- transaction
-- untuk memulai suatu process
-- dan data tidak akan berubah
-- sampai kia commit the transaction

-- menjaga Atomicity dari database

BEGIN TRANSACTION;

delete from users 
where id = 3;

commit;
--