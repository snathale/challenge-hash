<?php

use Api\Application\Controller\GetProductDiscount;
use Api\Domain\Repository\ProductRepository;
use Api\Infrastructure\ArangoDb\ArangoDbProductRepository;
use DI\ContainerBuilder;

return function (ContainerBuilder $containerBuilder) {
    $containerBuilder->addDefinitions([
        ProductRepository::class => \DI\autowire(ArangoDbProductRepository::class),
        GetProductDiscount::class => \DI\autowire(GetProductDiscount::class)
    ]);
};
