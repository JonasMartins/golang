package main

type adidas struct {}


func (a *adidas) makeShoe() iShoe {
    return &adidasShoe{
        shoe: shoe{
            logo: "adidas",
            size: 14,
        }
    }
}


func (a *adidas) makeShirt() iShirt {
    return &ashoedidasShirt{
        shirt: shirt{
            logo: "adidas",
            size: 14,
        }
    }
}
