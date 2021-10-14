package mycopier

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name     string
	Birthday *time.Time
	NickName string
	Role     string
	Age      int32
	FakeAge  *int32
	Notes    []string
	flags    []byte
}

func (user User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name       string
	Birthday   *time.Time
	NickName   *string
	Age        int64
	FakeAge    int
	EmployeeID int64
	DoubleAge  int32
	SuperRule  string
	Notes      []*string
	flags      []byte
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func checkEmployee(employee Employee, user User, t *testing.T, testCase string) {
	if employee.Name != user.Name {
		t.Errorf("%v: Name haven't been copied correctly.", testCase)
	}
	if employee.NickName == nil || *employee.NickName != user.NickName {
		t.Errorf("%v: NickName haven't been copied correctly.", testCase)
	}
	if employee.Birthday == nil && user.Birthday != nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if employee.Birthday != nil && user.Birthday == nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if employee.Birthday != nil && user.Birthday != nil &&
		!employee.Birthday.Equal(*(user.Birthday)) {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if employee.Age != int64(user.Age) {
		t.Errorf("%v: Age haven't been copied correctly.", testCase)
	}
	if user.FakeAge != nil && employee.FakeAge != int(*user.FakeAge) {
		t.Errorf("%v: FakeAge haven't been copied correctly.", testCase)
	}
	if employee.DoubleAge != user.DoubleAge() {
		t.Errorf("%v: Copy from method doesn't work", testCase)
	}
	if employee.SuperRule != "Super "+user.Role {
		t.Errorf("%v: Copy to method doesn't work", testCase)
	}

	if len(employee.Notes) != len(user.Notes) {
		t.Fatalf("%v: Copy from slice doesn't work, employee notes len: %v, user: %v", testCase, len(employee.Notes), len(user.Notes))
	}

	for idx, note := range user.Notes {
		if note != *employee.Notes[idx] {
			t.Fatalf("%v: Copy from slice doesn't work, notes idx: %v employee: %v user: %v", testCase, idx, *employee.Notes[idx], note)
		}
	}
}

func checkEmployee2(employee Employee, user *User, t *testing.T, testCase string) {
	if user == nil {
		if employee.Name != "" || employee.NickName != nil || employee.Birthday != nil || employee.Age != 0 ||
			employee.DoubleAge != 0 || employee.FakeAge != 0 || employee.SuperRule != "" || employee.Notes != nil {
			t.Errorf("%v : employee should be empty", testCase)
		}
		return
	}

	checkEmployee(employee, *user, t, testCase)
}

func TestCopySameStructWithPointerField(t *testing.T) {
	var fakeAge int32 = 12
	var curerntTime time.Time = time.Now()
	user := &User{
		Birthday: &curerntTime,
		Name:     "cannon",
		Age:      18,
		FakeAge:  &fakeAge,
		Role:     "Admin",
		Notes:    []string{"hello world", "welcome"},
		flags:    []byte{'x'},
	}
	newUser := &User{}
	err := Copy(newUser, user)
	if err != nil {
		t.Errorf("复制出错err is %v", err)
	}
	if user.Birthday == newUser.Birthday {
		t.Errorf("TestCopySameStructWithPointerField: copy Birthday failed since they need to have different address")
	}

	if user.FakeAge == newUser.FakeAge {
		t.Errorf("TestCopySameStructWithPointerField: copy FakeAge failed since they need to have different address")
	}

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", newUser)
}
