# File: rules/origin.go

在Prometheus项目中，rules/origin.go文件的作用是定义与规则相关的原始数据结构和函数。

首先是ruleOrigin结构体，它用于表示规则的源数据，包含了规则的名称、时间戳、表达式等信息。RuleDetail结构体则是对ruleOrigin的扩展，它包含了更详细的规则信息，如标签、持续时间、触发条件等。

NewRuleDetail函数是用来创建RuleDetail结构体的函数，根据给定的参数生成一个新的RuleDetail。

NewOriginContext函数用于从RuleDetail结构体创建一个原始上下文，它会将RuleDetail中的数据映射到原始上下文中，方便后续对规则进行处理。

FromOriginContext函数则是从原始上下文中创建一个RuleDetail结构体，它会将原始上下文中的数据提取出来，并填充到RuleDetail中，以便于对规则的处理和分析。

这些函数的作用是为Prometheus中的规则管理和评估提供了一种方便的方式。通过定义和操作这些数据结构和函数，Prometheus可以对规则进行存储、操作和展示。同时，这些数据结构和函数也为规则的配置和调整提供了灵活性和扩展性。
