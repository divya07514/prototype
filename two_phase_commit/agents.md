```
CREATE TABLE
  `agents` (
    `Id` int NOT NULL AUTO_INCREMENT,
    `is_reserved` boolean default false,
    `order_id` int DEFAULT NULL,
    PRIMARY KEY (`Id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 121 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
```

```
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);
INSERT INTO `agents` (`is_reserved`, `order_id`) VALUES (false, NULL);

```