<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

/**
 * Protobuf type <code>hpaas.AuthAction</code>
 */
class AuthAction
{
    /**
     * Allow a request only if it matches the rules. This is the default type.
     *
     * Generated from protobuf enum <code>ALLOW = 0;</code>
     */
    const ALLOW = 0;
    /**
     * Deny a request if it matches any of the rules.
     *
     * Generated from protobuf enum <code>DENY = 1;</code>
     */
    const DENY = 1;
    /**
     * Audit a request if it matches any of the rules.
     *
     * Generated from protobuf enum <code>AUDIT = 2;</code>
     */
    const AUDIT = 2;
}

