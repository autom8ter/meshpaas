<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>meshpaas.AuthzDestination</code>
 */
class AuthzDestination extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
     */
    private $hosts;
    /**
     * Generated from protobuf field <code>repeated string ports = 2;</code>
     */
    private $ports;
    /**
     * Generated from protobuf field <code>repeated string methods = 3;</code>
     */
    private $methods;
    /**
     * Generated from protobuf field <code>repeated string paths = 4;</code>
     */
    private $paths;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $hosts
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $ports
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $methods
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $paths
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHosts()
    {
        return $this->hosts;
    }

    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHosts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->hosts = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string ports = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPorts()
    {
        return $this->ports;
    }

    /**
     * Generated from protobuf field <code>repeated string ports = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPorts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->ports = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string methods = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMethods()
    {
        return $this->methods;
    }

    /**
     * Generated from protobuf field <code>repeated string methods = 3;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMethods($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->methods = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string paths = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPaths()
    {
        return $this->paths;
    }

    /**
     * Generated from protobuf field <code>repeated string paths = 4;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPaths($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->paths = $arr;

        return $this;
    }

}

