<?php


use Api\Application\Settings\Settings;
use Api\Application\Settings\SettingsInterface;
use ArangoDBClient\ConnectionOptions;
use DI\ContainerBuilder;
use Monolog\Logger;

return function (ContainerBuilder $containerBuilder) {
    $containerBuilder->addDefinitions([
        SettingsInterface::class => function () {
            return new Settings([
                'displayErrorDetails' => getenv('DISPLAY_ERROR_DETAILS') ? getenv('DISPLAY_ERROR_DETAILS') : true,
                'logger' => [
                    'name' => 'slim-config',
                    'path' => 'php://stdout',
                    'level' => Logger::DEBUG,
                ],
                'db' => [
                    ConnectionOptions::OPTION_DATABASE => getenv('DB_NAME') ? getenv('DB_NAME') : 'dummy_discount_db',
                    ConnectionOptions::OPTION_ENDPOINT => getenv('DB_ENDPOINT') ? getenv('DB_ENDPOINT') : 'tcp://arangodb.svc:8529',
                    ConnectionOptions::OPTION_AUTH_TYPE => 'Basic',
                    ConnectionOptions::OPTION_AUTH_USER => getenv('DB_USER') ? getenv('DB_USER') : 'root',
                    ConnectionOptions::OPTION_AUTH_PASSWD => getenv('DB_PASSWORD') ? getenv('DB_PASSWORD') : 'dummy_passowrd',
                    ConnectionOptions::OPTION_CONNECTION => 'Keep-Alive',
                    ConnectionOptions::OPTION_TIMEOUT => 3,
                    ConnectionOptions::OPTION_RECONNECT => true,
                    ConnectionOptions::OPTION_CREATE => true,
                ],
                'productCollection' => getenv('PRODUCT_COLLECTION_NAME') ? getenv('PRODUCT_COLLECTION_NAME') : 'product',
                'calculatorHost' => getenv('CALCULATOR_HOST') ? getenv('CALCULATOR_HOST') : 'calculator.svc',
                'calculatorPort' => getenv('CALCULATOR_PORT') ? getenv('CALCULATOR_PORT') : 3000,
            ]);
        }
    ]);
};
