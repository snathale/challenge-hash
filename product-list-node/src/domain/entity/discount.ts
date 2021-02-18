export default class Discount {
    public percentage: number
    public value_in_cents: number

    constructor(percentage: number, value_in_cents: number) {
        this.percentage = percentage
        this.value_in_cents = value_in_cents
    }
}