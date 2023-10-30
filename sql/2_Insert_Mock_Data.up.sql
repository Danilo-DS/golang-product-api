-- CATEGORY DATA
INSERT INTO CATEGORY (NAME, DESCRIPTION, ACTIVE) VALUES ('Fruits', 'This food group offers antioxidants (vitamins A, C and E) which are important nutrients for protecting cells, as they combat the action of free radicals.', TRUE);
INSERT INTO CATEGORY (NAME, DESCRIPTION, ACTIVE) VALUES ('Legumes', 'This group includes grains such as: beans, lentils, chickpeas, soybeans and oilseeds. Rich in fiber, they contribute to the proper functioning of the intestine.', FALSE);
INSERT INTO CATEGORY (NAME, DESCRIPTION, ACTIVE) VALUES ('Sugars', 'This group is the “ugly duckling” of the pyramid. Composed of sugar, honey, sweets and sugary products such as chocolate milk. They are poor in nutrients, have no fiber, and their consumption should be sporadic.', TRUE);
INSERT INTO CATEGORY (NAME, DESCRIPTION, ACTIVE) VALUES ('Meat and eggs', 'This is the main group of protein sources of animal origin, essential for the formation of tissues, enzymes, antibodies, vitamins B6 and B12.', TRUE);
INSERT INTO CATEGORY (NAME, DESCRIPTION, ACTIVE) VALUES ('Milk and Commitment', 'The foods in this group are also sources of protein, in addition to being rich in calcium, a fundamental nutrient for the constitution of our bones and teeth.', TRUE);

-- PRODUCT DATA
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Banana', 1,'MS1CYW5hbmEtMS0yMDIzLTEwLTI3IDAxOjE4OjA0Ljc0OTIyNTggLTAzMDAgLTAzIG09KzAuMDAyMDc2ODAx', 1);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Apple', 0.99,'Mi1BcHBsZS0xLTIwMjMtMTAtMjcgMDE6MTg6MDQuNzU5MTcyNSAtMDMwMCAtMDMgbT0rMC4wMTIwMjM1MDE=', 1);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Orange', 0.59,'My1PcmFuZ2UtMS0yMDIzLTEwLTI3IDAxOjE4OjA0Ljc1OTE3MjUgLTAzMDAgLTAzIG09KzAuMDEyMDIzNTAx', 1);

INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Cucumber', 0.39,'NC1DdWN1bWJlci0yLTIwMjMtMTAtMjcgMDE6MTg6MDQuNzU5MTcyNSAtMDMwMCAtMDMgbT0rMC4wMTIwMjM1MDE=', 2);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Carrot', 1.09,'NS1DYXJyb3QtMi0yMDIzLTEwLTI3IDAxOjE4OjA0Ljc1OTcwMDMgLTAzMDAgLTAzIG09KzAuMDEyNTUxMzAx', 2);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Pumpkin', 3.19,'Ni1QdW1wa2luLTItMjAyMy0xMC0yNyAwMToxODowNC43NTk3MDAzIC0wMzAwIC0wMyBtPSswLjAxMjU1MTMwMQ==', 2);

INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Chicken Fillet', 7,'Ny1DaGlja2VuIEZpbGxldC0zLTIwMjMtMTAtMjcgMDE6MTg6MDQuNzU5NzAwMyAtMDMwMCAtMDMgbT0rMC4wMTI1NTEzMDE=', 3);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Egg', 0.15,'OC1FZ2ctMy0yMDIzLTEwLTI3IDAxOjE4OjA0Ljc1OTcwMDMgLTAzMDAgLTAzIG09KzAuMDEyNTUxMzAx', 3);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Chicken Thigh', 5.89,'OS1DaGlja2VuIFRoaWdoLTMtMjAyMy0xMC0yNyAwMToxODowNC43NTk3MDAzIC0wMzAwIC0wMyBtPSswLjAxMjU1MTMwMQ==', 3);

INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Cereal Bar', 2.29,'MTAtQ2VyZWFsIEJhci00LTIwMjMtMTAtMjcgMDE6MTg6MDQuNzYwMjI1NiAtMDMwMCAtMDMgbT0rMC4wMTMwNzY2MDE=', 4);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Soy cheese', 14.69,'MTEtU295IGNoZWVzZS00LTIwMjMtMTAtMjcgMDE6MTg6MDQuNzYwMjI1NiAtMDMwMCAtMDMgbT0rMC4wMTMwNzY2MDE=', 4);
INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES ('Granola', 19.5,'MTItR3Jhbm9sYS00LTIwMjMtMTAtMjcgMDE6MTg6MDQuNzYwMjI1NiAtMDMwMCAtMDMgbT0rMC4wMTMwNzY2MDE=', 4);
