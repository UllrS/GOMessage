package repository

import (
	"MessageGO/models"
	"database/sql"
	"fmt"
	"time"
)

//Записывает модель Message в базу данных
func PostMessage(message *models.Message) (int64, error) {
	fmt.Println("Order post")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	date := time.Now().Unix()
	result, err := db.Exec("INSERT INTO messages (body, sender, recipient, attachedpath, attachedname, time) VALUES (?,?,?,?,?,?)", message.Body, message.Sender, message.Recipient, message.AttachedPath, message.AttachedName, date)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

//Возвращает все сообщения между пользователями user и companion
func GetMessages(user, companion, startTime string) ([]*models.Message, error) {
	fmt.Println("GET messages")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, body, sender, recipient, attachedpath, attachedname, time FROM messages WHERE sender = ? AND recipient = ? AND time >= ? OR recipient = ? AND sender = ? AND time >= ? ORDER BY time", user, companion, startTime, user, companion, startTime)
	// rows, err := db.Query("SELECT id, body, sender, recipient, time FROM messages WHERE sender = ? AND recipient = ? ORDER BY time", sender, recipient)
	if err != nil {
		return nil, err
	}

	var list []*models.Message
	for rows.Next() {
		newMessage := models.Message{}
		if err := rows.Scan(&newMessage.Id, &newMessage.Body, &newMessage.Sender, &newMessage.Recipient, &newMessage.AttachedPath, &newMessage.AttachedName, &newMessage.Time); err == nil {
			list = append(list, &newMessage)
		}
	}
	return list, nil

}

//Удаляет из базы данных сообщение пользователя user с идентификатором id
func DeleteMessage(id int64, user string) int64 {
	fmt.Println("message delete")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM messages WHERE id = ? AND sender = ? OR id = ? AND recipient = ?", id, user, id, user)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return rowsCount
}

//Удаляет из базы данных все сообщения между пользователем user и companion
func DeleteChat(user, companion string) (int64, error) {
	fmt.Println("chat delete")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM messages WHERE sender = ? AND recipient = ? OR recipient = ? AND sender = ?", user, companion, user, companion)
	if err != nil {
		return 0, err
	}
	deleteLines, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return deleteLines, nil
}

//Удаляет из базы данных все сообщения пользователя user
func DeleteUser(user string) (int64, error) {
	fmt.Println("message delete")
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM messages WHERE sender = ? OR recipient = ?", user, user)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	deleteLines, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return deleteLines, nil
}

// func UpdateMessage(id, message *models.Message) bool {
// 	fmt.Println("Report post")
// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	defer db.Close()
// 	date := time.Now().Unix()
// 	_, err = db.Exec("UPDATE messages SET body = ?, time = ? WHERE id = ?", message, date, id)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	return true
// }
// func GetAllOrders() []models.Order {

// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT instruction, report, comment, time FROM orders")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	var list []models.Order
// 	for rows.Next() {
// 		newOrder := models.Order{}
// 		if err := rows.Scan(&newOrder.Instruction, &newOrder.Report, &newOrder.Comment, &newOrder.Time); err == nil {
// 			list = append(list, newOrder)
// 		} else {
// 			fmt.Println(err.Error())
// 		}
// 	}
// 	return list

// }
// func GetAllOrdersCollaps() []*models.Order {

// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT COUNT(instruction), bee, instruction, report, comment, time FROM orders WHERE report='' GROUP BY instruction, bee")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	var list []*models.Order
// 	for rows.Next() {
// 		newOrder := models.Order{}
// 		if err := rows.Scan(&newOrder.Count, &newOrder.Bee, &newOrder.Instruction, &newOrder.Report, &newOrder.Comment, &newOrder.Time); err == nil {
// 			list = append(list, &newOrder)
// 		} else {
// 			fmt.Println(err.Error())
// 		}
// 	}
// 	return list

// }
// func GetAllCompleteOrdersCollaps() []*models.Order {

// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT COUNT(instruction), bee, instruction, report, comment, time FROM orders WHERE report != '' GROUP BY instruction, bee, report ORDER BY time")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	var list []*models.Order
// 	for rows.Next() {
// 		newOrder := models.Order{}
// 		if err := rows.Scan(&newOrder.Count, &newOrder.Bee, &newOrder.Instruction, &newOrder.Report, &newOrder.Comment, &newOrder.Time); err == nil {
// 			list = append(list, &newOrder)
// 		} else {
// 			fmt.Println(err.Error())
// 		}
// 	}
// 	return list

// }
// func OrderPost(order models.Order, count int) {
// 	fmt.Println("Order post")
// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()
// 	fmt.Println(count, order)
// 	if count < 1 {
// 		count = 1
// 	}
// 	beeList := GetBeeList()
// 	for _, val := range beeList {
// 		if order.Bee == val.Name || order.Bee == "" {
// 			for i := 0; i < count; i++ {
// 				db.Exec("INSERT INTO orders (bee, instruction, report, comment, time) VALUES (?,?,?,?,?)", val.Name, order.Instruction, order.Report, order.Comment, order.Time)
// 			}
// 		}
// 	}
// 	return
// }

// func GetAllOrdersfromInstr(inst string) []models.Order {

// 	db, err := sql.Open("sqlite3", "./base.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT instruction, report, comment, time FROM orders WHERE instruction = ?", inst)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	var list []models.Order
// 	for rows.Next() {
// 		newOrder := models.Order{}
// 		if err := rows.Scan(&newOrder.Instruction, &newOrder.Report, &newOrder.Comment, &newOrder.Time); err == nil {
// 			list = append(list, newOrder)
// 		} else {
// 			fmt.Println(err.Error())
// 		}
// 	}
// 	return list

// }
