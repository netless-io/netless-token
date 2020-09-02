<?php

$finder = PhpCsFixer\Finder::create()
    ->exclude('vendor')
    ->in([__DIR__.'/src/', __DIR__.'/tests/'])
;

return PhpCsFixer\Config::create()
    ->setRules([
        '@PSR4' => true,
    ])
    ->setFinder($finder)
    ;
