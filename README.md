# 🐛 IssueTrackerAPI-Golang

สร้างด้วย **Go + Fiber** และเชื่อมต่อฐานข้อมูล **MySQL**

## 🛠️ Setup `.env.local` to Your Computer

1. สร้างไฟล์ `.env.local` ที่ root ของโปรเจกต์

2. คัดลอกและแก้ไขค่าตามเครื่องของคุณ เช่น:

env`
DB_USER=root
DB_PASS=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=issue_tracker
JWT_SECRET=your_jwt_secret`


## ✨ Features

- ✅ JWT Authentication
- ✅ Project CRUD
- ✅ Employee Register & Login
- ✅ Many-to-Many: Projects ↔ Employees
- ✅ Upload project image
- ✅ Export Project to Excel (.xlsx)
