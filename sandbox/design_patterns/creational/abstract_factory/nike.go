package main

type nike struct{}

func (n *nike) makeShoe() iShoe {

    return &nikeShoe{
        shoe: shoe{
            logo: "nike",
            size: 14,
        }
    }
}


func (n *nike) makeShirt() iShirt {

    return &nikeShirth{
        shirt: shirt{
            logo: "nike",
            size: 14,
        }
    }
}
