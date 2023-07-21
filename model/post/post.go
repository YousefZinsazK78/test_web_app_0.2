package post

import "fmt"

type Post struct {
	Id    int
	Title string
	Body  string
}

func (p Post) String() string {
	return fmt.Sprintf("id: %v - title: %s - body: %s \n", p.Id, p.Title, p.Body)
}
