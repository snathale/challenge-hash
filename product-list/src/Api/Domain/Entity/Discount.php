<?php


namespace Api\Domain\Entity;


class Discount implements \JsonSerializable
{
    private float $percentage;

    private int $valueInCents;

    /**
     * Discount constructor.
     * @param float $percentage
     * @param int $valueInCents
     */
    public function __construct(float $percentage, int $valueInCents)
    {
        $this->percentage = $percentage;
        $this->valueInCents = $valueInCents;
    }

    public function jsonSerialize():array
    {
        return [
            'percentage' => round($this->percentage, 2),
            'value_in_cents' => $this->valueInCents
        ];
    }
}