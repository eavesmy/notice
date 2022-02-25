/*
# File Name: common.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 2022年02月25日 星期五 15时05分54秒
*/

package service

type Service interface {
	// Send msg to your robot.
	Send(string) (int, error)
}
