<?php


namespace Api\Application\Middleware;


use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Psr\Http\Server\MiddlewareInterface as Middleware;
use Psr\Http\Server\RequestHandlerInterface as RequestHandler;
use Slim\Psr7\Response as SlimResponse;

class JSendResponseMiddleware implements Middleware
{

    public function process(Request $request, RequestHandler $handler): Response
    {
        var_dump('ater');
        $response = $handler->handle($request);
        $response->getBody()->write(' AFTER');
        return $response;

        // TODO: Get response to convert do jsendResponse
    }
}