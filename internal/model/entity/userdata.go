// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Userdata is the golang structure for table userdata.
type Userdata struct {
	Id        uint   `json:"id"        orm:"id"         ` //
	Username  string `json:"username"  orm:"username"   ` //
	Email     string `json:"email"     orm:"email"      ` //
	Stackdata string `json:"stackdata" orm:"stackdata"  ` //
	CreatedAt uint   `json:"createdAt" orm:"created_at" ` //
	UpdatedAt uint   `json:"updatedAt" orm:"updated_at" ` //
}
