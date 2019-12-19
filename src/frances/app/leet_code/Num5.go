package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Member struct {
	Name string
	DepId int
	CenterId int
}

type DepMember struct {
	DepId int
	Members []Member
}

type CenterMember struct {
	CenterId int
	Members []Member
}

type Staff struct {
	memberInfo map[string]Member
	depMember map[int]DepMember
	centerMember map[int]CenterMember
}

func NewStaff(filePath string) *Staff {
	staff := &Staff{
		memberInfo:make(map[string]Member),
		depMember:make(map[int]DepMember),
		centerMember:make(map[int]CenterMember),
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return staff
	}
	defer file.Close()
	//br := bufio.NewReader(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		info := strings.Split(string(line), " ")
		if len(info) != 3{
			continue
		}
		name := info[0]
		depId,_ := strconv.Atoi(info[1])
		centerId,_ := strconv.Atoi(info[2])
		member := Member{
			Name:name,
			DepId:depId,
			CenterId:centerId,
		}
		staff.memberInfo[name] = member

		depMembers := staff.depMember[depId].Members
		depMembers = append(depMembers, member)
		depMem := DepMember{
			DepId:depId,
			Members:depMembers,
		}
		staff.depMember[depId]= depMem

		centerMembers := staff.centerMember[centerId].Members
		centerMembers = append(centerMembers, member)
		centerMem := CenterMember{
			CenterId:depId,
			Members:centerMembers,
		}
		staff.centerMember[centerId]= centerMem

	}

	return staff
}

func (s *Staff) GetMemberInfo(name string) Member {
	return s.memberInfo[name]
}

func (s *Staff) GetMembersOfDepart(depId int) []Member {
	return s.depMember[depId].Members
}

func (s *Staff) GetMembersOfCenter(centerId int) []Member {
	return s.centerMember[centerId].Members
}

func main() {
	staff := NewStaff("/home/jabinhuang/we-go/src/frances/app/leet_code/staff.txt")
	fmt.Printf("GetMemberInfo:%+v\n", staff.GetMemberInfo("lark"))
	fmt.Printf("GetMembersOfDepart(10):%+v\n", staff.GetMembersOfDepart(10))
	fmt.Printf("GetMembersOfCenter(102):%+v\n", staff.GetMembersOfCenter(102))
}
