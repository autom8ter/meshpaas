<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Hpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * AppRef is a generic structure for looking up apps
 *
 * Generated from protobuf message <code>hpaas.AppRef</code>
 */
class AppRef extends \Google\Protobuf\Internal\Message
{
    /**
     * namespace is the k8s namespace the app/release belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 1 [(.validator.field) = {</code>
     */
    private $namespace = '';
    /**
     * name is the name of the app/release
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
     */
    private $name = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $namespace
     *           namespace is the k8s namespace the app/release belongs to(autocreated)
     *     @type string $name
     *           name is the name of the app/release
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * namespace is the k8s namespace the app/release belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getNamespace()
    {
        return $this->namespace;
    }

    /**
     * namespace is the k8s namespace the app/release belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setNamespace($var)
    {
        GPBUtil::checkString($var, True);
        $this->namespace = $var;

        return $this;
    }

    /**
     * name is the name of the app/release
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name is the name of the app/release
     *
     * Generated from protobuf field <code>string name = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

}

