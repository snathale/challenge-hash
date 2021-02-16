<?php

namespace Api\Domain\Entity;


use JsonSerializable;

class Product implements JsonSerializable  {

    private string $id;

    private int $priceInCents;

    private string $title;

    private string $description;

    private ?Discount $discount;

    /**
     * Product constructor.
     * @param string $id
     * @param int $priceInCents
     * @param string $title
     * @param string $description
     * @param Discount|null $discount
     */
    public function __construct(string $id, int $priceInCents, string $title, string $description, Discount $discount = null)
    {
        $this->id = $id;
        $this->priceInCents = $priceInCents;
        $this->title = $title;
        $this->description = $description;
        $this->discount = empty($discount) ? new Discount(0, 0) : null;
    }

    /**
     * @return string
     */
    public function getId(): string
    {
        return $this->id;
    }

    /**
     * @param Discount|null $discount
     */
    public function setDiscount(?Discount $discount): void
    {
        $this->discount = $discount;
    }

    public function jsonSerialize() :array
    {
        return [
            'id' => $this->id,
            'price_in_cents' => $this->priceInCents,
            'title' => $this->title,
            'description' => $this->description,
            'discount' => $this->discount->jsonSerialize()
        ];
    }
}