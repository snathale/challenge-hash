import Discount from "./discount";

export default class Product {
    public id: string
    public price_in_cents: number
    public title: string
    public description: string
    public discount: Discount

    constructor(id: string, price_in_cents: number, title: string, description: string) {
        this.id = id
        this.price_in_cents = price_in_cents
        this.title = title
        this.description = description
        this.discount = new Discount(0, 0)
    }

    setDiscount(discount: Discount) {
        this.discount = discount
    }
    
}