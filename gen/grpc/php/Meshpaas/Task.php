<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Task is scheduled cron job
 *
 * Generated from protobuf message <code>meshpaas.Task</code>
 */
class Task extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the task
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * task project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     */
    private $project = '';
    /**
     * Generated from protobuf field <code>string image_pull_secret = 3;</code>
     */
    private $image_pull_secret = '';
    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 4 [(.validator.field) = {</code>
     */
    private $containers;
    /**
     * schedule is the cron schedule: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/
     *
     * Generated from protobuf field <code>string schedule = 7 [(.validator.field) = {</code>
     */
    private $schedule = '';
    /**
     * completions is the number of times to execute the task. If completions = 0, the task will run forever
     *
     * Generated from protobuf field <code>uint32 completions = 8;</code>
     */
    private $completions = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the task
     *     @type string $project
     *           task project
     *     @type string $image_pull_secret
     *     @type \Meshpaas\Container[]|\Google\Protobuf\Internal\RepeatedField $containers
     *     @type string $schedule
     *           schedule is the cron schedule: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/
     *     @type int $completions
     *           completions is the number of times to execute the task. If completions = 0, the task will run forever
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name of the task
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the task
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * task project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getProject()
    {
        return $this->project;
    }

    /**
     * task project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setProject($var)
    {
        GPBUtil::checkString($var, True);
        $this->project = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string image_pull_secret = 3;</code>
     * @return string
     */
    public function getImagePullSecret()
    {
        return $this->image_pull_secret;
    }

    /**
     * Generated from protobuf field <code>string image_pull_secret = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setImagePullSecret($var)
    {
        GPBUtil::checkString($var, True);
        $this->image_pull_secret = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 4 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getContainers()
    {
        return $this->containers;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 4 [(.validator.field) = {</code>
     * @param \Meshpaas\Container[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setContainers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Meshpaas\Container::class);
        $this->containers = $arr;

        return $this;
    }

    /**
     * schedule is the cron schedule: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/
     *
     * Generated from protobuf field <code>string schedule = 7 [(.validator.field) = {</code>
     * @return string
     */
    public function getSchedule()
    {
        return $this->schedule;
    }

    /**
     * schedule is the cron schedule: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/
     *
     * Generated from protobuf field <code>string schedule = 7 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setSchedule($var)
    {
        GPBUtil::checkString($var, True);
        $this->schedule = $var;

        return $this;
    }

    /**
     * completions is the number of times to execute the task. If completions = 0, the task will run forever
     *
     * Generated from protobuf field <code>uint32 completions = 8;</code>
     * @return int
     */
    public function getCompletions()
    {
        return $this->completions;
    }

    /**
     * completions is the number of times to execute the task. If completions = 0, the task will run forever
     *
     * Generated from protobuf field <code>uint32 completions = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setCompletions($var)
    {
        GPBUtil::checkUint32($var);
        $this->completions = $var;

        return $this;
    }

}

