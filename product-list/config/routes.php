<?php

use Api\Application\Handler\GetProductDiscount;
use Api\Application\Middleware\HeaderMiddleware;
use Slim\App;

return function (App $app) {
    $app->get('/product', GetProductDiscount::class)
        ->addMiddleware(new HeaderMiddleware());
};
