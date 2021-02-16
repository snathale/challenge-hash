<?php

use Api\Application\Settings\SettingsInterface;
use Api\Infrastructure\ArangoDb\ProductCollection;
use Api\Infrastructure\Grpc\CalculatorClient;
use ArangoDBClient\Connection;
use DI\ContainerBuilder;
use Grpc\ChannelCredentials;
use Monolog\Handler\StreamHandler;
use Monolog\Logger;
use Monolog\Processor\UidProcessor;
use Psr\Container\ContainerInterface;
use Psr\Log\LoggerInterface;

return function (ContainerBuilder $containerBuilder) {
    $containerBuilder->addDefinitions([
        LoggerInterface::class => function (ContainerInterface $c) {
            $settings = $c->get(SettingsInterface::class);

            $loggerSettings = $settings->get('logger');
            $logger = new Logger($loggerSettings['name']);

            $processor = new UidProcessor();
            $logger->pushProcessor($processor);

            $handler = new StreamHandler($loggerSettings['path'], $loggerSettings['level']);
            $logger->pushHandler($handler);

            return $logger;
        },
        ProductCollection::class => function (ContainerInterface  $c){
            $settings = $c->get(SettingsInterface::class);
            $dbSettings = $settings->get('db');
            $connection = new Connection($dbSettings);
            $collectionName = $settings->get('productCollection');
            return new ProductCollection($connection, $collectionName);
        },
        CalculatorClient::class => function(ContainerInterface $c) {
            $settings = $c->get(SettingsInterface::class);
            $host = $settings->get('calculatorHost');
            $port = $settings->get('calculatorPort');
            return new CalculatorClient($host.':'.$port, ['credentials' => ChannelCredentials::createInsecure(),]);
        }
    ]);
};
