<?php

use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;
use Illuminate\Http\Request;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Validation\ValidationException;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        web: __DIR__.'/../routes/web.php',
        api: __DIR__.'/../routes/api.php',
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
        ->withMiddleware(function (Middleware $middleware): void {
            $middleware->alias([
                'admin' => \App\Http\Middleware\PeranAdmin::class,
                'ability' => \Laravel\Sanctum\Http\Middleware\CheckAbilities::class,
            ]);

            $middleware->redirectGuestsTo(function (Request $request) {
                if ($request->is('api/*')) {
                    return null;
                }

                return route('login');
            });
    }) 
   ->withExceptions(function (Exceptions $exceptions) {

        // API selalu mengembalikan JSON
        $exceptions->shouldRenderJsonWhen(function (Request $request, Throwable $e) {
            return $request->is('api/*') || $request->expectsJson();
        });

        // Jika belum login / token tidak valid
        $exceptions->render(function (AuthenticationException $e, Request $request) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Token tidak valid atau belum dikirim',
                ], 401);
            }
        });

        // Jika validasi gagal
        $exceptions->render(function (ValidationException $e, Request $request) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Validasi gagal.',
                    'errors' => $e->errors(),
                ], 422);
            }
        });
    })
    ->create();