<?php

declare(strict_types=1);

use App\Application\Actions\User\ListUsersAction;
use App\Application\Actions\User\ViewUserAction;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\App;
use Slim\Interfaces\RouteCollectorProxyInterface as Group;

return function (App $app) {
    $app->options('/{routes:.*}', function (Request $request, Response $response) {
        // CORS Pre-Flight OPTIONS Request Handler
        return $response;
    });

    $app->get('/', function (Request $request, Response $response) {
        $response->getBody()->write('Landing page');
        return $response;
    });

    $app->get('/roadmap', function (Request $request, Response $response) {
        $response->getBody()->write('Roadmap!');
        return $response;
    });

    $app->get('/suggestions', function (Request $request, Response $response) {
        $response->getBody()->write('Suggestions!');
        return $response;
    });

    $app->get('/suggestions/:id', function (Request $request, Response $response) {
        $response->getBody()->write('Suggestion #ID');
        return $response;
    });

    $app->get('/changelog', function (Request $request, Response $response) {
        $response->getBody()->write('ChangeLog');
        return $response;
    });

    $app->group('/users', function (Group $group) {
        $group->get('', ListUsersAction::class);
        $group->get('/{id}', ViewUserAction::class);
    });
};
