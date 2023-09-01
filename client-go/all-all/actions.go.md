# File: client-go/testing/actions.go

client-go/testing/actions.go文件是client-go项目中的一个测试文件，用于定义与资源操作相关的测试接口和结构体。

- ListRestrictions、WatchRestrictions：分别表示资源操作的List和Watch限制，用于验证是否满足执行操作的条件。
- Action、GenericAction：表示资源操作的抽象接口，包含基础方法。
- GetAction、ListAction、CreateAction、UpdateAction、DeleteAction、DeleteCollectionAction、PatchAction、WatchAction、ProxyGetAction：分别表示对应资源操作的接口，继承自Action。
- ActionImpl、GenericActionImpl：表示资源操作的实现接口，包含具体操作的方法。
- GetActionImpl、ListActionImpl、CreateActionImpl、UpdateActionImpl、PatchActionImpl、DeleteActionImpl、DeleteCollectionActionImpl、WatchActionImpl、ProxyGetActionImpl：表示对应资源操作的实现接口，继承自ActionImpl，并重写了具体操作的方法。

这些结构体和接口用于测试各种资源操作的逻辑，例如获取资源、列表资源、创建资源、更新资源、删除资源等。它们提供了对资源操作的抽象和实现，方便在测试中验证操作的正确性和预期结果。

下面是一些常用的函数作用介绍：

- NewRootGetAction、NewGetAction、NewGetSubresourceAction、NewRootGetSubresourceAction、NewRootListAction、NewListAction、NewRootCreateAction、NewCreateAction、NewRootCreateSubresourceAction、NewCreateSubresourceAction、NewRootUpdateAction、NewUpdateAction、NewRootPatchAction、NewPatchAction、NewRootPatchSubresourceAction、NewPatchSubresourceAction、NewRootUpdateSubresourceAction、NewUpdateSubresourceAction、NewRootDeleteAction、NewRootDeleteActionWithOptions、NewRootDeleteSubresourceAction、NewDeleteAction、NewDeleteActionWithOptions、NewDeleteSubresourceAction、NewRootDeleteCollectionAction、NewDeleteCollectionAction、NewRootWatchAction：用于创建对应资源操作的实例。
- ExtractFromListOptions：从ListOptions中提取出操作的参数。
- NewWatchAction、NewProxyGetAction：创建Watch和ProxyGet操作的实例。
- GetNamespace、GetVerb、GetResource、GetSubresource、Matches：获取操作相关的信息。
- DeepCopy、GetValue、GetName、GetKind、GetListRestrictions、GetObject、GetPatch、GetPatchType、GetDeleteOptions、GetWatchRestrictions、GetScheme、GetPort、GetPath、GetParams：获取对应操作的参数和属性。

这些函数用于设置和获取操作的参数和属性，方便进行测试和验证。
