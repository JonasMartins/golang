package main

type iShoe interface {

    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int

}

type shoe struct {
    logo string
    size int
}

func (s *shoe) setLogo(logo string) {
    s.logo = logo
}

func (s *shoe) getLogo() {
    return s.logo
}


func (s *shoe) setSize(size int) {
    s.size = size
}

func (s *shoe) getSize() {
    return s.size
}

