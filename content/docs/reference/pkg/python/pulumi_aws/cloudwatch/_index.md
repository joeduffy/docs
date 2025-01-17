---
title: Module cloudwatch
---

<div class="section" id="cloudwatch">
<h1>cloudwatch<a class="headerlink" href="#cloudwatch" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-aws/issues">pulumi/pulumi-aws repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/issues">terraform-providers/terraform-provider-aws repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_aws.cloudwatch"></span><dl class="class">
<dt id="pulumi_aws.cloudwatch.AwaitableGetLogGroupResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">AwaitableGetLogGroupResult</code><span class="sig-paren">(</span><em>arn=None</em>, <em>creation_time=None</em>, <em>name=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.AwaitableGetLogGroupResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.Dashboard">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">Dashboard</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dashboard_body=None</em>, <em>dashboard_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Dashboard resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dashboard_body</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html">documentation</a>.</li>
<li><strong>dashboard_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the dashboard.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.Dashboard.dashboard_arn">
<code class="descname">dashboard_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.dashboard_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) of the dashboard.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.Dashboard.dashboard_body">
<code class="descname">dashboard_body</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.dashboard_body" title="Permalink to this definition">¶</a></dt>
<dd><p>The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html">documentation</a>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.Dashboard.dashboard_name">
<code class="descname">dashboard_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.dashboard_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the dashboard.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.Dashboard.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dashboard_arn=None</em>, <em>dashboard_body=None</em>, <em>dashboard_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Dashboard resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dashboard_arn: The Amazon Resource Name (ARN) of the dashboard.
:param pulumi.Input[str] dashboard_body: The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html">documentation</a>.
:param pulumi.Input[str] dashboard_name: The name of the dashboard.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.Dashboard.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.Dashboard.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.Dashboard.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.EventPermission">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">EventPermission</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>action=None</em>, <em>condition=None</em>, <em>principal=None</em>, <em>statement_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>action</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The action that you are enabling the other account to perform. Defaults to <code class="docutils literal notranslate"><span class="pre">events:PutEvents</span></code>.</li>
<li><strong>condition</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.</li>
<li><strong>principal</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify <code class="docutils literal notranslate"><span class="pre">*</span></code> to permit any account to put events to your default event bus, optionally limited by <code class="docutils literal notranslate"><span class="pre">condition</span></code>.</li>
<li><strong>statement_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – An identifier string for the external account that you are granting permissions to.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventPermission.action">
<code class="descname">action</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.action" title="Permalink to this definition">¶</a></dt>
<dd><p>The action that you are enabling the other account to perform. Defaults to <code class="docutils literal notranslate"><span class="pre">events:PutEvents</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventPermission.condition">
<code class="descname">condition</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.condition" title="Permalink to this definition">¶</a></dt>
<dd><p>Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventPermission.principal">
<code class="descname">principal</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.principal" title="Permalink to this definition">¶</a></dt>
<dd><p>The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify <code class="docutils literal notranslate"><span class="pre">*</span></code> to permit any account to put events to your default event bus, optionally limited by <code class="docutils literal notranslate"><span class="pre">condition</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventPermission.statement_id">
<code class="descname">statement_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.statement_id" title="Permalink to this definition">¶</a></dt>
<dd><p>An identifier string for the external account that you are granting permissions to.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.EventPermission.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>action=None</em>, <em>condition=None</em>, <em>principal=None</em>, <em>statement_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing EventPermission resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to <code class="docutils literal notranslate"><span class="pre">events:PutEvents</span></code>.
:param pulumi.Input[dict] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
:param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify <code class="docutils literal notranslate"><span class="pre">*</span></code> to permit any account to put events to your default event bus, optionally limited by <code class="docutils literal notranslate"><span class="pre">condition</span></code>.
:param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventPermission.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventPermission.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventPermission.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.EventRule">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">EventRule</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>description=None</em>, <em>event_pattern=None</em>, <em>is_enabled=None</em>, <em>name=None</em>, <em>name_prefix=None</em>, <em>role_arn=None</em>, <em>schedule_expression=None</em>, <em>tags=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Event Rule resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>description</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The description of the rule.</li>
<li><strong>event_pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Event pattern
described a JSON object.
See full documentation of <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html">CloudWatch Events and Event Patterns</a> for details.</li>
<li><strong>is_enabled</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Whether the rule should be enabled (defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>).</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The rule’s name. By default generated by this provider.</li>
<li><strong>name_prefix</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The rule’s name. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Amazon Resource Name (ARN) associated with the role that is used for target invocation.</li>
<li><strong>schedule_expression</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The scheduling expression.
For example, <code class="docutils literal notranslate"><span class="pre">cron(0</span> <span class="pre">20</span> <span class="pre">*</span> <span class="pre">*</span> <span class="pre">?</span> <span class="pre">*)</span></code> or <code class="docutils literal notranslate"><span class="pre">rate(5</span> <span class="pre">minutes)</span></code>.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_rule.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_rule.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) of the rule.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.description">
<code class="descname">description</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.description" title="Permalink to this definition">¶</a></dt>
<dd><p>The description of the rule.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.event_pattern">
<code class="descname">event_pattern</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.event_pattern" title="Permalink to this definition">¶</a></dt>
<dd><p>Event pattern
described a JSON object.
See full documentation of <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html">CloudWatch Events and Event Patterns</a> for details.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.is_enabled">
<code class="descname">is_enabled</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.is_enabled" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether the rule should be enabled (defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The rule’s name. By default generated by this provider.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.name_prefix">
<code class="descname">name_prefix</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.name_prefix" title="Permalink to this definition">¶</a></dt>
<dd><p>The rule’s name. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.role_arn">
<code class="descname">role_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.role_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) associated with the role that is used for target invocation.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.schedule_expression">
<code class="descname">schedule_expression</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.schedule_expression" title="Permalink to this definition">¶</a></dt>
<dd><p>The scheduling expression.
For example, <code class="docutils literal notranslate"><span class="pre">cron(0</span> <span class="pre">20</span> <span class="pre">*</span> <span class="pre">*</span> <span class="pre">?</span> <span class="pre">*)</span></code> or <code class="docutils literal notranslate"><span class="pre">rate(5</span> <span class="pre">minutes)</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventRule.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>A mapping of tags to assign to the resource.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.EventRule.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>arn=None</em>, <em>description=None</em>, <em>event_pattern=None</em>, <em>is_enabled=None</em>, <em>name=None</em>, <em>name_prefix=None</em>, <em>role_arn=None</em>, <em>schedule_expression=None</em>, <em>tags=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing EventRule resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the rule.
:param pulumi.Input[str] description: The description of the rule.
:param pulumi.Input[str] event_pattern: Event pattern</p>
<blockquote>
<div>described a JSON object.
See full documentation of <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html">CloudWatch Events and Event Patterns</a> for details.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>is_enabled</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Whether the rule should be enabled (defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>).</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The rule’s name. By default generated by this provider.</li>
<li><strong>name_prefix</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The rule’s name. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Amazon Resource Name (ARN) associated with the role that is used for target invocation.</li>
<li><strong>schedule_expression</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The scheduling expression.
For example, <code class="docutils literal notranslate"><span class="pre">cron(0</span> <span class="pre">20</span> <span class="pre">*</span> <span class="pre">*</span> <span class="pre">?</span> <span class="pre">*)</span></code> or <code class="docutils literal notranslate"><span class="pre">rate(5</span> <span class="pre">minutes)</span></code>.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_rule.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_rule.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventRule.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventRule.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventRule.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.EventTarget">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">EventTarget</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>arn=None</em>, <em>batch_target=None</em>, <em>ecs_target=None</em>, <em>input=None</em>, <em>input_path=None</em>, <em>input_transformer=None</em>, <em>kinesis_target=None</em>, <em>role_arn=None</em>, <em>rule=None</em>, <em>run_command_targets=None</em>, <em>sqs_target=None</em>, <em>target_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Event Target resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Amazon Resource Name (ARN) associated of the target.</li>
<li><strong>batch_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.</li>
<li><strong>ecs_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.</li>
<li><strong>input</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Valid JSON text passed to the target.</li>
<li><strong>input_path</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The value of the <a class="reference external" href="http://goessner.net/articles/JsonPath/">JSONPath</a>
that is used for extracting part of the matched event when passing it to the target.</li>
<li><strong>input_transformer</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are providing a custom input to a target based on certain event data.</li>
<li><strong>kinesis_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if <code class="docutils literal notranslate"><span class="pre">ecs_target</span></code> is used.</li>
<li><strong>rule</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the rule you want to add targets to.</li>
<li><strong>run_command_targets</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.</li>
<li><strong>sqs_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.</li>
<li><strong>target_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The unique target assignment ID.  If missing, will generate a random, unique id.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) associated of the target.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.batch_target">
<code class="descname">batch_target</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.batch_target" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.ecs_target">
<code class="descname">ecs_target</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.ecs_target" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.input">
<code class="descname">input</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.input" title="Permalink to this definition">¶</a></dt>
<dd><p>Valid JSON text passed to the target.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.input_path">
<code class="descname">input_path</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.input_path" title="Permalink to this definition">¶</a></dt>
<dd><p>The value of the <a class="reference external" href="http://goessner.net/articles/JsonPath/">JSONPath</a>
that is used for extracting part of the matched event when passing it to the target.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.input_transformer">
<code class="descname">input_transformer</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.input_transformer" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are providing a custom input to a target based on certain event data.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.kinesis_target">
<code class="descname">kinesis_target</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.kinesis_target" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.role_arn">
<code class="descname">role_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.role_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if <code class="docutils literal notranslate"><span class="pre">ecs_target</span></code> is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.rule">
<code class="descname">rule</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.rule" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the rule you want to add targets to.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.run_command_targets">
<code class="descname">run_command_targets</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.run_command_targets" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.sqs_target">
<code class="descname">sqs_target</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.sqs_target" title="Permalink to this definition">¶</a></dt>
<dd><p>Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.EventTarget.target_id">
<code class="descname">target_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.target_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The unique target assignment ID.  If missing, will generate a random, unique id.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.EventTarget.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>arn=None</em>, <em>batch_target=None</em>, <em>ecs_target=None</em>, <em>input=None</em>, <em>input_path=None</em>, <em>input_transformer=None</em>, <em>kinesis_target=None</em>, <em>role_arn=None</em>, <em>rule=None</em>, <em>run_command_targets=None</em>, <em>sqs_target=None</em>, <em>target_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing EventTarget resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] arn: The Amazon Resource Name (ARN) associated of the target.
:param pulumi.Input[dict] batch_target: Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
:param pulumi.Input[dict] ecs_target: Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
:param pulumi.Input[str] input: Valid JSON text passed to the target.
:param pulumi.Input[str] input_path: The value of the <a class="reference external" href="http://goessner.net/articles/JsonPath/">JSONPath</a></p>
<blockquote>
<div>that is used for extracting part of the matched event when passing it to the target.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>input_transformer</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are providing a custom input to a target based on certain event data.</li>
<li><strong>kinesis_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if <code class="docutils literal notranslate"><span class="pre">ecs_target</span></code> is used.</li>
<li><strong>rule</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the rule you want to add targets to.</li>
<li><strong>run_command_targets</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.</li>
<li><strong>sqs_target</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.</li>
<li><strong>target_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The unique target assignment ID.  If missing, will generate a random, unique id.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventTarget.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.EventTarget.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.EventTarget.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.GetLogGroupResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">GetLogGroupResult</code><span class="sig-paren">(</span><em>arn=None</em>, <em>creation_time=None</em>, <em>name=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.GetLogGroupResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getLogGroup.</p>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.GetLogGroupResult.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.GetLogGroupResult.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of the Cloudwatch log group</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.GetLogGroupResult.creation_time">
<code class="descname">creation_time</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.GetLogGroupResult.creation_time" title="Permalink to this definition">¶</a></dt>
<dd><p>The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.GetLogGroupResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.GetLogGroupResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>id is the provider-assigned unique ID for this managed resource.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogDestination">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogDestination</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>name=None</em>, <em>role_arn=None</em>, <em>target_arn=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Logs destination resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A name for the log destination</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target</li>
<li><strong>target_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ARN of the target Amazon Kinesis stream or Amazon Lambda resource for the destination</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestination.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) specifying the log destination.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestination.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.name" title="Permalink to this definition">¶</a></dt>
<dd><p>A name for the log destination</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestination.role_arn">
<code class="descname">role_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.role_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestination.target_arn">
<code class="descname">target_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.target_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of the target Amazon Kinesis stream or Amazon Lambda resource for the destination</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogDestination.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>arn=None</em>, <em>name=None</em>, <em>role_arn=None</em>, <em>target_arn=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogDestination resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the log destination.
:param pulumi.Input[str] name: A name for the log destination
:param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
:param pulumi.Input[str] target_arn: The ARN of the target Amazon Kinesis stream or Amazon Lambda resource for the destination</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogDestination.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogDestination.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestination.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogDestinationPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>access_policy=None</em>, <em>destination_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Logs destination policy resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>access_policy</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy document. This is a JSON formatted string.</li>
<li><strong>destination_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A name for the subscription filter</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy.access_policy">
<code class="descname">access_policy</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy.access_policy" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy document. This is a JSON formatted string.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy.destination_name">
<code class="descname">destination_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy.destination_name" title="Permalink to this definition">¶</a></dt>
<dd><p>A name for the subscription filter</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>access_policy=None</em>, <em>destination_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogDestinationPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
:param pulumi.Input[str] destination_name: A name for the subscription filter</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogDestinationPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogDestinationPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogGroup">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogGroup</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>kms_key_id=None</em>, <em>name=None</em>, <em>name_prefix=None</em>, <em>retention_in_days=None</em>, <em>tags=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Log Group resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>kms_key_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
permissions for the CMK whenever the encrypted data is requested.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log group. If omitted, this provider will assign a random, unique name.</li>
<li><strong>name_prefix</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Creates a unique name beginning with the specified prefix. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</li>
<li><strong>retention_in_days</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Specifies the number of days
you want to retain log events in the specified log group.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_group.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_group.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) specifying the log group.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.kms_key_id">
<code class="descname">kms_key_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.kms_key_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
permissions for the CMK whenever the encrypted data is requested.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the log group. If omitted, this provider will assign a random, unique name.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.name_prefix">
<code class="descname">name_prefix</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.name_prefix" title="Permalink to this definition">¶</a></dt>
<dd><p>Creates a unique name beginning with the specified prefix. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.retention_in_days">
<code class="descname">retention_in_days</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.retention_in_days" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the number of days
you want to retain log events in the specified log group.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogGroup.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>A mapping of tags to assign to the resource.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogGroup.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>arn=None</em>, <em>kms_key_id=None</em>, <em>name=None</em>, <em>name_prefix=None</em>, <em>retention_in_days=None</em>, <em>tags=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogGroup resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the log group.
:param pulumi.Input[str] kms_key_id: The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,</p>
<blockquote>
<div>AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
permissions for the CMK whenever the encrypted data is requested.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log group. If omitted, this provider will assign a random, unique name.</li>
<li><strong>name_prefix</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Creates a unique name beginning with the specified prefix. Conflicts with <code class="docutils literal notranslate"><span class="pre">name</span></code>.</li>
<li><strong>retention_in_days</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Specifies the number of days
you want to retain log events in the specified log group.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_group.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_group.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogGroup.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogGroup.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogGroup.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogMetricFilter</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>log_group_name=None</em>, <em>metric_transformation=None</em>, <em>name=None</em>, <em>pattern=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Log Metric Filter resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>log_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log group to associate the metric filter with.</li>
<li><strong>metric_transformation</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A block defining collection of information
needed to define how metric data gets emitted. See below.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A name for the metric filter.</li>
<li><strong>pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A valid <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html">CloudWatch Logs filter pattern</a>
for extracting metric data out of ingested log events.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_metric_filter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_metric_filter.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.log_group_name">
<code class="descname">log_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.log_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the log group to associate the metric filter with.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.metric_transformation">
<code class="descname">metric_transformation</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.metric_transformation" title="Permalink to this definition">¶</a></dt>
<dd><p>A block defining collection of information
needed to define how metric data gets emitted. See below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.name" title="Permalink to this definition">¶</a></dt>
<dd><p>A name for the metric filter.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.pattern">
<code class="descname">pattern</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.pattern" title="Permalink to this definition">¶</a></dt>
<dd><p>A valid <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html">CloudWatch Logs filter pattern</a>
for extracting metric data out of ingested log events.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>log_group_name=None</em>, <em>metric_transformation=None</em>, <em>name=None</em>, <em>pattern=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogMetricFilter resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] log_group_name: The name of the log group to associate the metric filter with.
:param pulumi.Input[dict] metric_transformation: A block defining collection of information</p>
<blockquote>
<div>needed to define how metric data gets emitted. See below.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A name for the metric filter.</li>
<li><strong>pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – <p>A valid <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html">CloudWatch Logs filter pattern</a>
for extracting metric data out of ingested log events.</p>
</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_metric_filter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_metric_filter.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogMetricFilter.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogMetricFilter.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogResourcePolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>policy_document=None</em>, <em>policy_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a resource to manage a CloudWatch log resource policy.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>policy_document</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.</li>
<li><strong>policy_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Name of the resource policy.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_resource_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_resource_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy.policy_document">
<code class="descname">policy_document</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy.policy_document" title="Permalink to this definition">¶</a></dt>
<dd><p>Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy.policy_name">
<code class="descname">policy_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy.policy_name" title="Permalink to this definition">¶</a></dt>
<dd><p>Name of the resource policy.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>policy_document=None</em>, <em>policy_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogResourcePolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
:param pulumi.Input[str] policy_name: Name of the resource policy.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_resource_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_resource_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogResourcePolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogResourcePolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogStream">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogStream</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>log_group_name=None</em>, <em>name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Log Stream resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>log_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log group under which the log stream is to be created.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log stream. Must not be longer than 512 characters and must not contain <code class="docutils literal notranslate"><span class="pre">:</span></code></li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_stream.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_stream.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogStream.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The Amazon Resource Name (ARN) specifying the log stream.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogStream.log_group_name">
<code class="descname">log_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.log_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the log group under which the log stream is to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogStream.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the log stream. Must not be longer than 512 characters and must not contain <code class="docutils literal notranslate"><span class="pre">:</span></code></p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogStream.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>arn=None</em>, <em>log_group_name=None</em>, <em>name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogStream resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the log stream.
:param pulumi.Input[str] log_group_name: The name of the log group under which the log stream is to be created.
:param pulumi.Input[str] name: The name of the log stream. Must not be longer than 512 characters and must not contain <code class="docutils literal notranslate"><span class="pre">:</span></code></p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_stream.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_stream.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogStream.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogStream.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogStream.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">LogSubscriptionFilter</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>destination_arn=None</em>, <em>distribution=None</em>, <em>filter_pattern=None</em>, <em>log_group=None</em>, <em>name=None</em>, <em>role_arn=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Logs subscription filter resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>destination_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.</li>
<li><strong>distribution</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are “Random” and “ByLogStream”.</li>
<li><strong>filter_pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.</li>
<li><strong>log_group</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the log group to associate the subscription filter with</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – A name for the subscription filter</li>
<li><strong>role_arn</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use <code class="docutils literal notranslate"><span class="pre">lambda.Permission</span></code> resource for granting access from CloudWatch logs to the destination Lambda function.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.destination_arn">
<code class="descname">destination_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.destination_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.distribution">
<code class="descname">distribution</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.distribution" title="Permalink to this definition">¶</a></dt>
<dd><p>The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are “Random” and “ByLogStream”.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.filter_pattern">
<code class="descname">filter_pattern</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.filter_pattern" title="Permalink to this definition">¶</a></dt>
<dd><p>A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.log_group">
<code class="descname">log_group</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.log_group" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the log group to associate the subscription filter with</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.name" title="Permalink to this definition">¶</a></dt>
<dd><p>A name for the subscription filter</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.role_arn">
<code class="descname">role_arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.role_arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use <code class="docutils literal notranslate"><span class="pre">lambda.Permission</span></code> resource for granting access from CloudWatch logs to the destination Lambda function.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>destination_arn=None</em>, <em>distribution=None</em>, <em>filter_pattern=None</em>, <em>log_group=None</em>, <em>name=None</em>, <em>role_arn=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing LogSubscriptionFilter resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] destination_arn: The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
:param pulumi.Input[str] distribution: The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are “Random” and “ByLogStream”.
:param pulumi.Input[str] filter_pattern: A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
:param pulumi.Input[str] log_group: The name of the log group to associate the subscription filter with
:param pulumi.Input[str] name: A name for the subscription filter
:param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use <code class="docutils literal notranslate"><span class="pre">lambda.Permission</span></code> resource for granting access from CloudWatch logs to the destination Lambda function.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.LogSubscriptionFilter.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.LogSubscriptionFilter.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.cloudwatch.MetricAlarm">
<em class="property">class </em><code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">MetricAlarm</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>actions_enabled=None</em>, <em>alarm_actions=None</em>, <em>alarm_description=None</em>, <em>name=None</em>, <em>comparison_operator=None</em>, <em>datapoints_to_alarm=None</em>, <em>dimensions=None</em>, <em>evaluate_low_sample_count_percentiles=None</em>, <em>evaluation_periods=None</em>, <em>extended_statistic=None</em>, <em>insufficient_data_actions=None</em>, <em>metric_name=None</em>, <em>metric_queries=None</em>, <em>namespace=None</em>, <em>ok_actions=None</em>, <em>period=None</em>, <em>statistic=None</em>, <em>tags=None</em>, <em>threshold=None</em>, <em>treat_missing_data=None</em>, <em>unit=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides a CloudWatch Metric Alarm resource.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>actions_enabled</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Indicates whether or not actions should be executed during any changes to the alarm’s state. Defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>.</li>
<li><strong>alarm_actions</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).</li>
<li><strong>alarm_description</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The description for the alarm.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The descriptive name for the alarm. This name must be unique within the user’s AWS account</li>
<li><strong>comparison_operator</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">GreaterThanOrEqualToThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">GreaterThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanOrEqualToThreshold</span></code>.</li>
<li><strong>datapoints_to_alarm</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The number of datapoints that must be breaching to trigger the alarm.</li>
<li><strong>dimensions</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – The dimensions for this metric.  For the list of available dimensions see the AWS documentation <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">here</a>.</li>
<li><strong>evaluate_low_sample_count_percentiles</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Used only for alarms
based on percentiles. If you specify <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, the alarm state will not
change during periods with too few data points to be statistically significant.
If you specify <code class="docutils literal notranslate"><span class="pre">evaluate</span></code> or omit this parameter, the alarm will always be
evaluated and possibly change state no matter how many data points are available.
The following values are supported: <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, and <code class="docutils literal notranslate"><span class="pre">evaluate</span></code>.</li>
<li><strong>evaluation_periods</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The number of periods over which data is compared to the specified threshold.</li>
<li><strong>extended_statistic</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.</li>
<li><strong>insufficient_data_actions</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).</li>
<li><strong>metric_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name for this metric.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</li>
<li><strong>metric_queries</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – Enables you to create an alarm based on a metric math expression. You may specify at most 20.</li>
<li><strong>namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – <p>The namespace for this metric. See docs for the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html">list of namespaces</a>.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</p>
</li>
<li><strong>ok_actions</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).</li>
<li><strong>period</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The period in seconds over which the specified <code class="docutils literal notranslate"><span class="pre">stat</span></code> is applied.</li>
<li><strong>statistic</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The statistic to apply to the alarm’s associated metric.
Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">SampleCount</span></code>, <code class="docutils literal notranslate"><span class="pre">Average</span></code>, <code class="docutils literal notranslate"><span class="pre">Sum</span></code>, <code class="docutils literal notranslate"><span class="pre">Minimum</span></code>, <code class="docutils literal notranslate"><span class="pre">Maximum</span></code></li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
<li><strong>threshold</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The value against which the specified statistic is compared.</li>
<li><strong>treat_missing_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Sets how this alarm is to handle missing data points. The following values are supported: <code class="docutils literal notranslate"><span class="pre">missing</span></code>, <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, <code class="docutils literal notranslate"><span class="pre">breaching</span></code> and <code class="docutils literal notranslate"><span class="pre">notBreaching</span></code>. Defaults to <code class="docutils literal notranslate"><span class="pre">missing</span></code>.</li>
<li><strong>unit</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The unit for this metric.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_metric_alarm.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_metric_alarm.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.actions_enabled">
<code class="descname">actions_enabled</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.actions_enabled" title="Permalink to this definition">¶</a></dt>
<dd><p>Indicates whether or not actions should be executed during any changes to the alarm’s state. Defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.alarm_actions">
<code class="descname">alarm_actions</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.alarm_actions" title="Permalink to this definition">¶</a></dt>
<dd><p>The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.alarm_description">
<code class="descname">alarm_description</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.alarm_description" title="Permalink to this definition">¶</a></dt>
<dd><p>The description for the alarm.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The descriptive name for the alarm. This name must be unique within the user’s AWS account</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>The ARN of the cloudwatch metric alarm.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.comparison_operator">
<code class="descname">comparison_operator</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.comparison_operator" title="Permalink to this definition">¶</a></dt>
<dd><p>The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">GreaterThanOrEqualToThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">GreaterThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanOrEqualToThreshold</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.datapoints_to_alarm">
<code class="descname">datapoints_to_alarm</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.datapoints_to_alarm" title="Permalink to this definition">¶</a></dt>
<dd><p>The number of datapoints that must be breaching to trigger the alarm.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.dimensions">
<code class="descname">dimensions</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.dimensions" title="Permalink to this definition">¶</a></dt>
<dd><p>The dimensions for this metric.  For the list of available dimensions see the AWS documentation <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">here</a>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.evaluate_low_sample_count_percentiles">
<code class="descname">evaluate_low_sample_count_percentiles</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.evaluate_low_sample_count_percentiles" title="Permalink to this definition">¶</a></dt>
<dd><p>Used only for alarms
based on percentiles. If you specify <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, the alarm state will not
change during periods with too few data points to be statistically significant.
If you specify <code class="docutils literal notranslate"><span class="pre">evaluate</span></code> or omit this parameter, the alarm will always be
evaluated and possibly change state no matter how many data points are available.
The following values are supported: <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, and <code class="docutils literal notranslate"><span class="pre">evaluate</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.evaluation_periods">
<code class="descname">evaluation_periods</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.evaluation_periods" title="Permalink to this definition">¶</a></dt>
<dd><p>The number of periods over which data is compared to the specified threshold.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.extended_statistic">
<code class="descname">extended_statistic</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.extended_statistic" title="Permalink to this definition">¶</a></dt>
<dd><p>The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.insufficient_data_actions">
<code class="descname">insufficient_data_actions</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.insufficient_data_actions" title="Permalink to this definition">¶</a></dt>
<dd><p>The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.metric_name">
<code class="descname">metric_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.metric_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name for this metric.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.metric_queries">
<code class="descname">metric_queries</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.metric_queries" title="Permalink to this definition">¶</a></dt>
<dd><p>Enables you to create an alarm based on a metric math expression. You may specify at most 20.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.namespace">
<code class="descname">namespace</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.namespace" title="Permalink to this definition">¶</a></dt>
<dd><p>The namespace for this metric. See docs for the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html">list of namespaces</a>.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.ok_actions">
<code class="descname">ok_actions</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.ok_actions" title="Permalink to this definition">¶</a></dt>
<dd><p>The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.period">
<code class="descname">period</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.period" title="Permalink to this definition">¶</a></dt>
<dd><p>The period in seconds over which the specified <code class="docutils literal notranslate"><span class="pre">stat</span></code> is applied.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.statistic">
<code class="descname">statistic</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.statistic" title="Permalink to this definition">¶</a></dt>
<dd><p>The statistic to apply to the alarm’s associated metric.
Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">SampleCount</span></code>, <code class="docutils literal notranslate"><span class="pre">Average</span></code>, <code class="docutils literal notranslate"><span class="pre">Sum</span></code>, <code class="docutils literal notranslate"><span class="pre">Minimum</span></code>, <code class="docutils literal notranslate"><span class="pre">Maximum</span></code></p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>A mapping of tags to assign to the resource.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.threshold">
<code class="descname">threshold</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.threshold" title="Permalink to this definition">¶</a></dt>
<dd><p>The value against which the specified statistic is compared.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.treat_missing_data">
<code class="descname">treat_missing_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.treat_missing_data" title="Permalink to this definition">¶</a></dt>
<dd><p>Sets how this alarm is to handle missing data points. The following values are supported: <code class="docutils literal notranslate"><span class="pre">missing</span></code>, <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, <code class="docutils literal notranslate"><span class="pre">breaching</span></code> and <code class="docutils literal notranslate"><span class="pre">notBreaching</span></code>. Defaults to <code class="docutils literal notranslate"><span class="pre">missing</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.unit">
<code class="descname">unit</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.unit" title="Permalink to this definition">¶</a></dt>
<dd><p>The unit for this metric.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>actions_enabled=None</em>, <em>alarm_actions=None</em>, <em>alarm_description=None</em>, <em>name=None</em>, <em>arn=None</em>, <em>comparison_operator=None</em>, <em>datapoints_to_alarm=None</em>, <em>dimensions=None</em>, <em>evaluate_low_sample_count_percentiles=None</em>, <em>evaluation_periods=None</em>, <em>extended_statistic=None</em>, <em>insufficient_data_actions=None</em>, <em>metric_name=None</em>, <em>metric_queries=None</em>, <em>namespace=None</em>, <em>ok_actions=None</em>, <em>period=None</em>, <em>statistic=None</em>, <em>tags=None</em>, <em>threshold=None</em>, <em>treat_missing_data=None</em>, <em>unit=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing MetricAlarm resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm’s state. Defaults to <code class="docutils literal notranslate"><span class="pre">true</span></code>.
:param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
:param pulumi.Input[str] alarm_description: The description for the alarm.
:param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user’s AWS account
:param pulumi.Input[str] arn: The ARN of the cloudwatch metric alarm.
:param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">GreaterThanOrEqualToThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">GreaterThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanThreshold</span></code>, <code class="docutils literal notranslate"><span class="pre">LessThanOrEqualToThreshold</span></code>.
:param pulumi.Input[float] datapoints_to_alarm: The number of datapoints that must be breaching to trigger the alarm.
:param pulumi.Input[dict] dimensions: The dimensions for this metric.  For the list of available dimensions see the AWS documentation <a class="reference external" href="http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">here</a>.
:param pulumi.Input[str] evaluate_low_sample_count_percentiles: Used only for alarms</p>
<blockquote>
<div>based on percentiles. If you specify <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, the alarm state will not
change during periods with too few data points to be statistically significant.
If you specify <code class="docutils literal notranslate"><span class="pre">evaluate</span></code> or omit this parameter, the alarm will always be
evaluated and possibly change state no matter how many data points are available.
The following values are supported: <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, and <code class="docutils literal notranslate"><span class="pre">evaluate</span></code>.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>evaluation_periods</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The number of periods over which data is compared to the specified threshold.</li>
<li><strong>extended_statistic</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.</li>
<li><strong>insufficient_data_actions</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).</li>
<li><strong>metric_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – <p>The name for this metric.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</p>
</li>
<li><strong>metric_queries</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – Enables you to create an alarm based on a metric math expression. You may specify at most 20.</li>
<li><strong>namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – <p>The namespace for this metric. See docs for the <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html">list of namespaces</a>.
See docs for <a class="reference external" href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html">supported metrics</a>.</p>
</li>
<li><strong>ok_actions</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).</li>
<li><strong>period</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The period in seconds over which the specified <code class="docutils literal notranslate"><span class="pre">stat</span></code> is applied.</li>
<li><strong>statistic</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The statistic to apply to the alarm’s associated metric.
Either of the following is supported: <code class="docutils literal notranslate"><span class="pre">SampleCount</span></code>, <code class="docutils literal notranslate"><span class="pre">Average</span></code>, <code class="docutils literal notranslate"><span class="pre">Sum</span></code>, <code class="docutils literal notranslate"><span class="pre">Minimum</span></code>, <code class="docutils literal notranslate"><span class="pre">Maximum</span></code></li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags to assign to the resource.</li>
<li><strong>threshold</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The value against which the specified statistic is compared.</li>
<li><strong>treat_missing_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Sets how this alarm is to handle missing data points. The following values are supported: <code class="docutils literal notranslate"><span class="pre">missing</span></code>, <code class="docutils literal notranslate"><span class="pre">ignore</span></code>, <code class="docutils literal notranslate"><span class="pre">breaching</span></code> and <code class="docutils literal notranslate"><span class="pre">notBreaching</span></code>. Defaults to <code class="docutils literal notranslate"><span class="pre">missing</span></code>.</li>
<li><strong>unit</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The unit for this metric.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_metric_alarm.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_metric_alarm.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.cloudwatch.MetricAlarm.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.MetricAlarm.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="function">
<dt id="pulumi_aws.cloudwatch.get_log_group">
<code class="descclassname">pulumi_aws.cloudwatch.</code><code class="descname">get_log_group</code><span class="sig-paren">(</span><em>name=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.cloudwatch.get_log_group" title="Permalink to this definition">¶</a></dt>
<dd><p>Use this data source to get information about an AWS Cloudwatch Log Group</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown</a>.</div></blockquote>
</dd></dl>

</div>
