```
CREATE TABLE
  `food` (
    `Id` int NOT NULL AUTO_INCREMENT,
    `food_id` int default null,
    `is_reserved` boolean DEFAULT false,
    `order_id` int DEFAULT NULL,
    PRIMARY KEY (`Id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 121 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
```

```
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
INSERT INTO `food` (`food_id`, `is_reserved`, `order_id`) VALUES (1, false, NULL);
```