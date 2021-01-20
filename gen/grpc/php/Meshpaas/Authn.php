<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Authn holds authentication options for an application
 *
 * Generated from protobuf message <code>meshpaas.Authn</code>
 */
class Authn extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .meshpaas.AuthnRule rules = 1;</code>
     */
    private $rules;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Meshpaas\AuthnRule[]|\Google\Protobuf\Internal\RepeatedField $rules
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.AuthnRule rules = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRules()
    {
        return $this->rules;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.AuthnRule rules = 1;</code>
     * @param \Meshpaas\AuthnRule[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRules($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Meshpaas\AuthnRule::class);
        $this->rules = $arr;

        return $this;
    }

}

