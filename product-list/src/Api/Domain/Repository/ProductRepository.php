<?php


namespace Api\Domain\Repository;


use Api\Domain\Entity\Product;
use Api\Infrastructure\ArangoDb\RepositoryError;

interface ProductRepository
{
    /**
     * @return Product[]
     * @throws RepositoryError
     */
    public function getAllProducts(): array;
}