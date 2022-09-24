# Create a Pack

## Anatomy of a Pack



### Actions

The **actions** folder contains action script files and action metadata files. See [Actions]() and [Workflows]() for specifics on writing actions. Since metadata files and workflow definitions can both be written as YAML, it's good practice to put the workflow definitions in a separate directory. Note that the `lib` sub-folder is often used for storing common Python code used by pack actions.

### Rules

```shell
# contents of rules/
rules/
   rule1.yaml
   rule2.yaml
```

The `rules` folder contains rules. StackStorm uses rules and workflows to capture operational patterns as automations. Rules map triggers to actions (or workflows), apply matching criteria, and map trigger payloads to action inputs.

#### Rule structure

Rules are defined in YAML. Rule definition structure, as well as required and optional elements are listed below:

```yaml
---
    name: "rule_name"                      # required
    pack: "examples"                       # optional
    description: "Rule description."       # optional
    enabled: true                          # required

    trigger:                               # required
        type: "trigger_type_ref"

    criteria:                              # optional
        trigger.payload_parameter_name1:
            type: "regex"
            pattern : "^value$"
        trigger.payload_parameter_name2:
            type: "iequals"
            pattern : "watchevent"

    action:                                # required
        ref: "action_ref"
        parameters:                        # optional
            foo: "bar"
            baz: "{{ trigger.payload_parameter_1 }}"
```

The generic form of a rule is:

* The `name` of the rule.
* The `pack` that the rule belongs to. `default` is assumed if a pack is not specified.
* The `description` of the rule.
* The `enabled` state of a rule (`true` or `false`)
* The type of `trigger` emitted from sensors to monitor, and optionally parameters associated with that trigger.
* An optional set of `criteria`, consisting of:
    * An attribute of the trigger payload.
    * The `type` of criteria comparison.
    * The `pattern` to match against.
* The `action` to execute when a rule is matched, consisting of:
    * The `ref` (action/workflow) to execute.
    * An optional set of `parameters` to pass to the action execution