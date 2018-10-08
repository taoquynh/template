package model

import (
	"time"
)

type User struct {
	// Tên bảng
	TableName struct{} `json:"table_name" sql:"test.users"`

	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique" valid:"email~Email không hợp lệ"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`

	// Ảnh đại diện
	Avatar string `json:"current_page"`

	// Slug user
	Slug string `json:"slug"`

	// Mảng các quyền hạn (Role) của User (ví dụ [3, 5] tức là User có Role Id là 3 và 5).
	// Một User có thể có nhiều Role, nếu Null tức là User không có quyền gì
	Roles []int32 `json:"roles" pg:",array"`

	// Email đã xác nhận (kích hoạt) hay chưa
	// True là đã xác nhận, False là chưa xác nhận, mặc định là False
	EmailConfirmed bool `json:"email_confirmed" sql:"default:false"`

	// Token để xác thực Email
	VerifyEmailToken string `json:"verify_email_token"`

	// Thời gian hiệu lực của Token xác thực email
	VerifyEmailTokenEnd time.Time `json:"verify_email_token_end"`

	// Token để reset password
	ResetPasswordToken string `json:"reset_password_token"`

	// Thời gian hiệu lực của Token reset password
	ResetPasswordTokenEnd time.Time `json:"reset_password_token_end"`

	// Số lần đăng nhập sai, mặc định là 0
	AccessFailedCount int32 `json:"access_failed_count" sql:"default:0"`

	// Thời gian khóa đăng nhập, mặc định là Null tức là không khóa
	// Sau X số lần đăng nhập sai có thể khóa User đăng nhập trong khoảng thời gian Y
	LockoutEnd time.Time `json:"lockout_end"`

	// Ngày tài khoản được tạo
	CreatedAt time.Time `json:"created_at" sql:"default:now()"`

	// Id người tạo tài khoản, Null là người dùng tự đăng ký tài khoản
	CreatedBy string `json:"created_by"`

	// Ngày gần nhất tài khoản cập nhật thông tin
	ModifiedAt time.Time `json:"modified_at" sql:"default:now()"`

	// Người cập nhật thông tin tài khoản gần nhất
	ModifiedBy string `json:"modified_by"`

	// Trạng thái của User
	// True là active, False là unactive, mặc định là True
	UserStatus bool `json:"user_status" sql:"default:true"`
}

type CreateUser struct {
	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`

	// Ảnh đại diện
	Avatar string `json:"avatar"`

	// Mảng các quyền hạn (Role) của User (ví dụ [3, 5] tức là User có Role Id là 3 và 5).
	// Một User có thể có nhiều Role, nếu Null tức là User không có quyền gì
	Roles []int32 `json:"roles" pg:",array"`
}

type UpdateUser struct {
	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`

	// Ảnh đại diện
	Avatar string `json:"avatar"`

	// Mảng các quyền hạn (Role) của User (ví dụ [3, 5] tức là User có Role Id là 3 và 5).
	// Một User có thể có nhiều Role, nếu Null tức là User không có quyền gì
	Roles []int32 `json:"roles" pg:",array"`

	// Trạng thái của User
	// True là active, False là unactive, mặc định là True
	UserStatus bool `json:"user_status" sql:"default:true"`
}

