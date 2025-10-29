package main

import (
	"fmt"
)

const MAX = 100

func main() {
	var ids [MAX]int
	var names [MAX]string
	var prices [MAX]float64
	var count int = 0
	var nextID int = 1 

	for {
		fmt.Println("\n===== COFFEE SHOP MANAGEMENT =====")
		fmt.Println("1. Thêm món mới")
		fmt.Println("2. Sửa món")
		fmt.Println("3. Xóa món")
		fmt.Println("4. Hiển thị menu")
		fmt.Println("5. Thoát")
		fmt.Print("Chọn: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if count >= MAX {
				fmt.Println("Menu đầy, không thể thêm.")
				continue
			}

			ids[count] = nextID
			nextID++

			fmt.Print("Nhập tên món: ")
			fmt.Scan(&names[count])
			fmt.Print("Nhập giá: ")
			fmt.Scan(&prices[count])

			count++
			fmt.Println(" Đã thêm món thành công!")

		case 2:
			var id int
			fmt.Print("Nhập ID món cần sửa: ")
			fmt.Scan(&id)
			index := -1
			for i := 0; i < count; i++ {
				if ids[i] == id {
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Println(" Không tìm thấy món có ID đó.")
			} else {
				fmt.Print("Nhập tên mới: ")
				fmt.Scan(&names[index])
				fmt.Print("Nhập giá mới: ")
				fmt.Scan(&prices[index])
				fmt.Println(" Đã cập nhật món thành công!")
			}

		case 3:
			var id int
			fmt.Print("Nhập ID món cần xóa: ")
			fmt.Scan(&id)
			index := -1
			for i := 0; i < count; i++ {
				if ids[i] == id {
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Println(" Không tìm thấy món có ID đó.")
			} else {
				for i := index; i < count-1; i++ {
					ids[i] = ids[i+1]
					names[i] = names[i+1]
					prices[i] = prices[i+1]
				}
				count--
				fmt.Println(" Đã xóa món thành công!")
			}

		case 4:
			if count == 0 {
				fmt.Println("Menu trống!")
			} else {
				fmt.Println("\n----- MENU HIỆN TẠI -----")
				for i := 0; i < count; i++ {
					fmt.Printf("%d. ID: %d | Tên: %s | Giá: %.2f\n",
						i+1, ids[i], names[i], prices[i])
				}
			}

		case 5:
			fmt.Println(" Thoát chương trình. Tạm biệt!")
			return

		default:
			fmt.Println(" Lựa chọn không hợp lệ.")
		}
	}
}
