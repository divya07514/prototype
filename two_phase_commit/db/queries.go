package db

var SelectAgent = "Select Id, is_reserved, order_id from agents where is_reserved is false and order_id is null limit 1 for update"
var BlockAgent = "Update agents set is_reserved = true where id=?"
var BookAgent = "Select Id, is_reserved, order_id from agents where is_reserved is true and order_id is null limit 1 for update"
var UpdateAgent = "Update agents set is_reserved = false , order_id=? where id=?"

var SelectFood = "Select Id, food_id, is_reserved, order_id from food where is_reserved is false and food_id=? and order_Id is null limit 1 for update"
var BlockFood = "Update food set is_reserved = true where id=?"
var BookFood = "Select Id, food_id, is_reserved, order_id from food where is_reserved is true and food_id=? and order_Id is null limit 1 for update"
var UpdateFood = "Update food set is_reserved = true, order_id = ? where id=?"
