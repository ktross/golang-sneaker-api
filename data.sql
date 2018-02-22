CREATE TABLE `sneakers` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR(64) NOT NULL
);

CREATE TABLE `true_to_size` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `sneaker_id` INTEGER NOT NULL,
  `size` INTEGER NOT NULL,
  FOREIGN KEY (sneaker_id) REFERENCES sneakers(id)
);

INSERT INTO `sneakers` (name)
VALUES
  ('Jordan 3 Retro Black Cement (2018)'),
  ('Jordan 1 Retro High NRG Patent Gold Toe'),
  ('Air Foamposite One Big Bang'),
  ('adidas Yeezy Boost 350 V2 Blue Tint');

INSERT INTO `true_to_size` (sneaker_id, size)
VALUES
  (1, 1),
  (1, 2),
  (1, 2),
  (1, 3),
  (1, 2),
  (1, 3),
  (1, 2),
  (1, 2),
  (1, 3),
  (1, 4),
  (1, 2),
  (1, 5),
  (1, 2),
  (1, 3),
  (2, 1),
  (2, 2),
  (2, 2),
  (2, 3),
  (2, 2),
  (2, 3),
  (2, 2),
  (2, 2),
  (2, 3),
  (2, 4),
  (2, 2),
  (2, 5),
  (2, 2),
  (2, 3),
  (2, 2),
  (3, 3),
  (3, 1),
  (3, 3),
  (3, 1),
  (3, 2),
  (4, 5),
  (4, 5),
  (4, 5),
  (4, 5),
  (4, 5);