<?php


namespace Api\Application\Middleware;


use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Psr\Http\Server\MiddlewareInterface as Middleware;
use Psr\Http\Server\RequestHandlerInterface as RequestHandler;

class HeaderMiddleware implements Middleware
{
    public function process(Request $request, RequestHandler $handler): Response
    {
        $headers = getallheaders();
        if (!empty($headers) && isset($headers['X-User-Id'])) {
            $request = $request->withAttribute('userId', $headers['X-User-Id']);
        }
        return $handler->handle($request);
    }
}