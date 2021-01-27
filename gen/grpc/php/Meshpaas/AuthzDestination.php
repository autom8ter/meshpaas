<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * the destination of an authorization decision
 *
 * Generated from protobuf message <code>hpaas.AuthzDestination</code>
 */
class AuthzDestination extends \Google\Protobuf\Internal\Message
{
    /**
     * restricts access to one of many paths
     *
     * Generated from protobuf field <code>repeated string allow_paths = 2;</code>
     */
    private $allow_paths;
    /**
     * restricts access to one of many hosts
     *
     * Generated from protobuf field <code>repeated string allow_hosts = 3;</code>
     */
    private $allow_hosts;
    /**
     * restricts access to one of many methods
     *
     * Generated from protobuf field <code>repeated string allow_methods = 4;</code>
     */
    private $allow_methods;
    /**
     * restricts access to one of many ports
     *
     * Generated from protobuf field <code>repeated string allow_ports = 5;</code>
     */
    private $allow_ports;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_paths
     *           restricts access to one of many paths
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_hosts
     *           restricts access to one of many hosts
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_methods
     *           restricts access to one of many methods
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_ports
     *           restricts access to one of many ports
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * restricts access to one of many paths
     *
     * Generated from protobuf field <code>repeated string allow_paths = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowPaths()
    {
        return $this->allow_paths;
    }

    /**
     * restricts access to one of many paths
     *
     * Generated from protobuf field <code>repeated string allow_paths = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowPaths($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_paths = $arr;

        return $this;
    }

    /**
     * restricts access to one of many hosts
     *
     * Generated from protobuf field <code>repeated string allow_hosts = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowHosts()
    {
        return $this->allow_hosts;
    }

    /**
     * restricts access to one of many hosts
     *
     * Generated from protobuf field <code>repeated string allow_hosts = 3;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowHosts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_hosts = $arr;

        return $this;
    }

    /**
     * restricts access to one of many methods
     *
     * Generated from protobuf field <code>repeated string allow_methods = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowMethods()
    {
        return $this->allow_methods;
    }

    /**
     * restricts access to one of many methods
     *
     * Generated from protobuf field <code>repeated string allow_methods = 4;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowMethods($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_methods = $arr;

        return $this;
    }

    /**
     * restricts access to one of many ports
     *
     * Generated from protobuf field <code>repeated string allow_ports = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowPorts()
    {
        return $this->allow_ports;
    }

    /**
     * restricts access to one of many ports
     *
     * Generated from protobuf field <code>repeated string allow_ports = 5;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowPorts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_ports = $arr;

        return $this;
    }

}

