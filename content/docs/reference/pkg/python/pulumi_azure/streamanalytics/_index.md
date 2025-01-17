---
title: Module streamanalytics
---

<div class="section" id="streamanalytics">
<h1>streamanalytics<a class="headerlink" href="#streamanalytics" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-azure/issues">pulumi/pulumi-azure repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/issues">terraform-providers/terraform-provider-azurerm repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_azure.streamanalytics"></span><dl class="class">
<dt id="pulumi_azure.streamanalytics.AwaitableGetJobResult">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">AwaitableGetJobResult</code><span class="sig-paren">(</span><em>compatibility_level=None</em>, <em>data_locale=None</em>, <em>events_late_arrival_max_delay_in_seconds=None</em>, <em>events_out_of_order_max_delay_in_seconds=None</em>, <em>events_out_of_order_policy=None</em>, <em>job_id=None</em>, <em>location=None</em>, <em>name=None</em>, <em>output_error_policy=None</em>, <em>resource_group_name=None</em>, <em>streaming_units=None</em>, <em>transformation_query=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.AwaitableGetJobResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">FunctionJavaScriptUDF</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>inputs=None</em>, <em>name=None</em>, <em>output=None</em>, <em>resource_group_name=None</em>, <em>script=None</em>, <em>stream_analytics_job_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a JavaScript UDF Function within Stream Analytics Streaming Job.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>inputs</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – One or more <code class="docutils literal notranslate"><span class="pre">input</span></code> blocks as defined below.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the JavaScript UDF Function. Changing this forces a new resource to be created.</li>
<li><strong>output</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – An <code class="docutils literal notranslate"><span class="pre">output</span></code> blocks as defined below.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>script</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The JavaScript of this UDF Function.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.inputs">
<code class="descname">inputs</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.inputs" title="Permalink to this definition">¶</a></dt>
<dd><p>One or more <code class="docutils literal notranslate"><span class="pre">input</span></code> blocks as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the JavaScript UDF Function. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.output">
<code class="descname">output</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.output" title="Permalink to this definition">¶</a></dt>
<dd><p>An <code class="docutils literal notranslate"><span class="pre">output</span></code> blocks as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.script">
<code class="descname">script</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.script" title="Permalink to this definition">¶</a></dt>
<dd><p>The JavaScript of this UDF Function.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>inputs=None</em>, <em>name=None</em>, <em>output=None</em>, <em>resource_group_name=None</em>, <em>script=None</em>, <em>stream_analytics_job_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing FunctionJavaScriptUDF resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[list] inputs: One or more <code class="docutils literal notranslate"><span class="pre">input</span></code> blocks as defined below.
:param pulumi.Input[str] name: The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
:param pulumi.Input[dict] output: An <code class="docutils literal notranslate"><span class="pre">output</span></code> blocks as defined below.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[str] script: The JavaScript of this UDF Function.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.FunctionJavaScriptUDF.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.FunctionJavaScriptUDF.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.GetJobResult">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">GetJobResult</code><span class="sig-paren">(</span><em>compatibility_level=None</em>, <em>data_locale=None</em>, <em>events_late_arrival_max_delay_in_seconds=None</em>, <em>events_out_of_order_max_delay_in_seconds=None</em>, <em>events_out_of_order_policy=None</em>, <em>job_id=None</em>, <em>location=None</em>, <em>name=None</em>, <em>output_error_policy=None</em>, <em>resource_group_name=None</em>, <em>streaming_units=None</em>, <em>transformation_query=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getJob.</p>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.compatibility_level">
<code class="descname">compatibility_level</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.compatibility_level" title="Permalink to this definition">¶</a></dt>
<dd><p>The compatibility level for this job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.data_locale">
<code class="descname">data_locale</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.data_locale" title="Permalink to this definition">¶</a></dt>
<dd><p>The Data Locale of the Job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.events_late_arrival_max_delay_in_seconds">
<code class="descname">events_late_arrival_max_delay_in_seconds</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.events_late_arrival_max_delay_in_seconds" title="Permalink to this definition">¶</a></dt>
<dd><p>The maximum tolerable delay in seconds where events arriving late could be included.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.events_out_of_order_max_delay_in_seconds">
<code class="descname">events_out_of_order_max_delay_in_seconds</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.events_out_of_order_max_delay_in_seconds" title="Permalink to this definition">¶</a></dt>
<dd><p>The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.events_out_of_order_policy">
<code class="descname">events_out_of_order_policy</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.events_out_of_order_policy" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy which should be applied to events which arrive out of order in the input event stream.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.job_id">
<code class="descname">job_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.job_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The Job ID assigned by the Stream Analytics Job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.location">
<code class="descname">location</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.location" title="Permalink to this definition">¶</a></dt>
<dd><p>The Azure location where the Stream Analytics Job exists.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.output_error_policy">
<code class="descname">output_error_policy</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.output_error_policy" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.streaming_units">
<code class="descname">streaming_units</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.streaming_units" title="Permalink to this definition">¶</a></dt>
<dd><p>The number of streaming units that the streaming job uses.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.transformation_query">
<code class="descname">transformation_query</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.transformation_query" title="Permalink to this definition">¶</a></dt>
<dd><p>The query that will be run in the streaming job, <a class="reference external" href="https://msdn.microsoft.com/library/azure/dn834998">written in Stream Analytics Query Language (SAQL)</a>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.GetJobResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.GetJobResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>id is the provider-assigned unique ID for this managed resource.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_azure.streamanalytics.Job">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">Job</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>compatibility_level=None</em>, <em>data_locale=None</em>, <em>events_late_arrival_max_delay_in_seconds=None</em>, <em>events_out_of_order_max_delay_in_seconds=None</em>, <em>events_out_of_order_policy=None</em>, <em>location=None</em>, <em>name=None</em>, <em>output_error_policy=None</em>, <em>resource_group_name=None</em>, <em>streaming_units=None</em>, <em>tags=None</em>, <em>transformation_query=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.Job" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Job.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>compatibility_level</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are <code class="docutils literal notranslate"><span class="pre">1.0</span></code> and 1.1<a href="#id1"><span class="problematic" id="id2">``</span></a>.</li>
<li><strong>data_locale</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Specifies the Data Locale of the Job, which [should be a supported .NET Culture](<a class="reference external" href="https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx">https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx</a>).</li>
<li><strong>events_late_arrival_max_delay_in_seconds</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is``-1<code class="docutils literal notranslate"><span class="pre">(indefinite)</span> <span class="pre">to</span></code>1814399<a href="#id3"><span class="problematic" id="id4">``</span></a>(20d 23h 59m 59s).</li>
<li><strong>events_out_of_order_max_delay_in_seconds</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is``0<code class="docutils literal notranslate"><span class="pre">to</span></code>599<a href="#id5"><span class="problematic" id="id6">``</span></a>(9m 59s).</li>
<li><strong>events_out_of_order_policy</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are``Adjust<code class="docutils literal notranslate"><span class="pre">and</span></code>Drop<a href="#id7"><span class="problematic" id="id8">``</span></a>.</li>
<li><strong>location</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
<li><strong>output_error_policy</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are``Drop<code class="docutils literal notranslate"><span class="pre">and</span></code>Stop<a href="#id9"><span class="problematic" id="id10">``</span></a>.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.</li>
<li><strong>streaming_units</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Specifies the number of streaming units that the streaming job uses. Supported values are``1<code class="docutils literal notranslate"><span class="pre">,</span></code>3<code class="docutils literal notranslate"><span class="pre">,</span></code>6<code class="docutils literal notranslate"><span class="pre">and</span> <span class="pre">multiples</span> <span class="pre">of</span></code>6<code class="docutils literal notranslate"><span class="pre">up</span> <span class="pre">to</span></code>120`.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A mapping of tags assigned to the resource.</li>
<li><strong>transformation_query</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – <p>Specifies the query that will be run in the streaming job, <a class="reference external" href="https://msdn.microsoft.com/library/azure/dn834998">written in Stream Analytics Query Language (SAQL)</a>.</p>
</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_job.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_job.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.compatibility_level">
<code class="descname">compatibility_level</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.compatibility_level" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are <code class="docutils literal notranslate"><span class="pre">1.0</span></code> and 1.1`.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.data_locale">
<code class="descname">data_locale</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.data_locale" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the Data Locale of the Job, which <a class="reference external" href="https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110">should be a supported .NET Culture</a>.aspx).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.events_late_arrival_max_delay_in_seconds">
<code class="descname">events_late_arrival_max_delay_in_seconds</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.events_late_arrival_max_delay_in_seconds" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is <code class="docutils literal notranslate"><span class="pre">-1</span></code> (indefinite) to <code class="docutils literal notranslate"><span class="pre">1814399</span></code> (20d 23h 59m 59s).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.events_out_of_order_max_delay_in_seconds">
<code class="descname">events_out_of_order_max_delay_in_seconds</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.events_out_of_order_max_delay_in_seconds" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is <code class="docutils literal notranslate"><span class="pre">0</span></code> to <code class="docutils literal notranslate"><span class="pre">599</span></code> (9m 59s).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.events_out_of_order_policy">
<code class="descname">events_out_of_order_policy</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.events_out_of_order_policy" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are <code class="docutils literal notranslate"><span class="pre">Adjust</span></code> and <code class="docutils literal notranslate"><span class="pre">Drop</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.job_id">
<code class="descname">job_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.job_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The Job ID assigned by the Stream Analytics Job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.location">
<code class="descname">location</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.location" title="Permalink to this definition">¶</a></dt>
<dd><p>The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.output_error_policy">
<code class="descname">output_error_policy</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.output_error_policy" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are <code class="docutils literal notranslate"><span class="pre">Drop</span></code> and <code class="docutils literal notranslate"><span class="pre">Stop</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.streaming_units">
<code class="descname">streaming_units</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.streaming_units" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the number of streaming units that the streaming job uses. Supported values are <code class="docutils literal notranslate"><span class="pre">1</span></code>, <code class="docutils literal notranslate"><span class="pre">3</span></code>, <code class="docutils literal notranslate"><span class="pre">6</span></code> and multiples of <code class="docutils literal notranslate"><span class="pre">6</span></code> up to <code class="docutils literal notranslate"><span class="pre">120</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>A mapping of tags assigned to the resource.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.Job.transformation_query">
<code class="descname">transformation_query</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.transformation_query" title="Permalink to this definition">¶</a></dt>
<dd><p>Specifies the query that will be run in the streaming job, <a class="reference external" href="https://msdn.microsoft.com/library/azure/dn834998">written in Stream Analytics Query Language (SAQL)</a>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.Job.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>compatibility_level=None</em>, <em>data_locale=None</em>, <em>events_late_arrival_max_delay_in_seconds=None</em>, <em>events_out_of_order_max_delay_in_seconds=None</em>, <em>events_out_of_order_policy=None</em>, <em>job_id=None</em>, <em>location=None</em>, <em>name=None</em>, <em>output_error_policy=None</em>, <em>resource_group_name=None</em>, <em>streaming_units=None</em>, <em>tags=None</em>, <em>transformation_query=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Job resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] compatibility_level: Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are <code class="docutils literal notranslate"><span class="pre">1.0</span></code> and 1.1<code class="docutils literal notranslate"><span class="pre">.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">data_locale:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">Data</span> <span class="pre">Locale</span> <span class="pre">of</span> <span class="pre">the</span> <span class="pre">Job,</span> <span class="pre">which</span> <span class="pre">[should</span> <span class="pre">be</span> <span class="pre">a</span> <span class="pre">supported</span> <span class="pre">.NET</span> <span class="pre">Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[float]</span> <span class="pre">events_late_arrival_max_delay_in_seconds:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">maximum</span> <span class="pre">tolerable</span> <span class="pre">delay</span> <span class="pre">in</span> <span class="pre">seconds</span> <span class="pre">where</span> <span class="pre">events</span> <span class="pre">arriving</span> <span class="pre">late</span> <span class="pre">could</span> <span class="pre">be</span> <span class="pre">included.</span> <span class="pre">Supported</span> <span class="pre">range</span> <span class="pre">is</span></code>-1<code class="docutils literal notranslate"><span class="pre">(indefinite)</span> <span class="pre">to</span></code>1814399<code class="docutils literal notranslate"><span class="pre">(20d</span> <span class="pre">23h</span> <span class="pre">59m</span> <span class="pre">59s).</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[float]</span> <span class="pre">events_out_of_order_max_delay_in_seconds:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">maximum</span> <span class="pre">tolerable</span> <span class="pre">delay</span> <span class="pre">in</span> <span class="pre">seconds</span> <span class="pre">where</span> <span class="pre">out-of-order</span> <span class="pre">events</span> <span class="pre">can</span> <span class="pre">be</span> <span class="pre">adjusted</span> <span class="pre">to</span> <span class="pre">be</span> <span class="pre">back</span> <span class="pre">in</span> <span class="pre">order.</span> <span class="pre">Supported</span> <span class="pre">range</span> <span class="pre">is</span></code>0<code class="docutils literal notranslate"><span class="pre">to</span></code>599<code class="docutils literal notranslate"><span class="pre">(9m</span> <span class="pre">59s).</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">events_out_of_order_policy:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">policy</span> <span class="pre">which</span> <span class="pre">should</span> <span class="pre">be</span> <span class="pre">applied</span> <span class="pre">to</span> <span class="pre">events</span> <span class="pre">which</span> <span class="pre">arrive</span> <span class="pre">out</span> <span class="pre">of</span> <span class="pre">order</span> <span class="pre">in</span> <span class="pre">the</span> <span class="pre">input</span> <span class="pre">event</span> <span class="pre">stream.</span> <span class="pre">Possible</span> <span class="pre">values</span> <span class="pre">are</span></code>Adjust<code class="docutils literal notranslate"><span class="pre">and</span></code>Drop<code class="docutils literal notranslate"><span class="pre">.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">job_id:</span> <span class="pre">The</span> <span class="pre">Job</span> <span class="pre">ID</span> <span class="pre">assigned</span> <span class="pre">by</span> <span class="pre">the</span> <span class="pre">Stream</span> <span class="pre">Analytics</span> <span class="pre">Job.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">location:</span> <span class="pre">The</span> <span class="pre">Azure</span> <span class="pre">Region</span> <span class="pre">in</span> <span class="pre">which</span> <span class="pre">the</span> <span class="pre">Resource</span> <span class="pre">Group</span> <span class="pre">exists.</span> <span class="pre">Changing</span> <span class="pre">this</span> <span class="pre">forces</span> <span class="pre">a</span> <span class="pre">new</span> <span class="pre">resource</span> <span class="pre">to</span> <span class="pre">be</span> <span class="pre">created.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">name:</span> <span class="pre">The</span> <span class="pre">name</span> <span class="pre">of</span> <span class="pre">the</span> <span class="pre">Stream</span> <span class="pre">Analytics</span> <span class="pre">Job.</span> <span class="pre">Changing</span> <span class="pre">this</span> <span class="pre">forces</span> <span class="pre">a</span> <span class="pre">new</span> <span class="pre">resource</span> <span class="pre">to</span> <span class="pre">be</span> <span class="pre">created.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">output_error_policy:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">policy</span> <span class="pre">which</span> <span class="pre">should</span> <span class="pre">be</span> <span class="pre">applied</span> <span class="pre">to</span> <span class="pre">events</span> <span class="pre">which</span> <span class="pre">arrive</span> <span class="pre">at</span> <span class="pre">the</span> <span class="pre">output</span> <span class="pre">and</span> <span class="pre">cannot</span> <span class="pre">be</span> <span class="pre">written</span> <span class="pre">to</span> <span class="pre">the</span> <span class="pre">external</span> <span class="pre">storage</span> <span class="pre">due</span> <span class="pre">to</span> <span class="pre">being</span> <span class="pre">malformed</span> <span class="pre">(such</span> <span class="pre">as</span> <span class="pre">missing</span> <span class="pre">column</span> <span class="pre">values,</span> <span class="pre">column</span> <span class="pre">values</span> <span class="pre">of</span> <span class="pre">wrong</span> <span class="pre">type</span> <span class="pre">or</span> <span class="pre">size).</span> <span class="pre">Possible</span> <span class="pre">values</span> <span class="pre">are</span></code>Drop<code class="docutils literal notranslate"><span class="pre">and</span></code>Stop<code class="docutils literal notranslate"><span class="pre">.</span> 
<span class="pre">:param</span> <span class="pre">pulumi.Input[str]</span> <span class="pre">resource_group_name:</span> <span class="pre">The</span> <span class="pre">name</span> <span class="pre">of</span> <span class="pre">the</span> <span class="pre">Resource</span> <span class="pre">Group</span> <span class="pre">where</span> <span class="pre">the</span> <span class="pre">Stream</span> <span class="pre">Analytics</span> <span class="pre">Job</span> <span class="pre">should</span> <span class="pre">exist.</span> <span class="pre">Changing</span> <span class="pre">this</span> <span class="pre">forces</span> <span class="pre">a</span> <span class="pre">new</span> <span class="pre">resource</span> <span class="pre">to</span> <span class="pre">be</span> <span class="pre">created.</span>
<span class="pre">:param</span> <span class="pre">pulumi.Input[float]</span> <span class="pre">streaming_units:</span> <span class="pre">Specifies</span> <span class="pre">the</span> <span class="pre">number</span> <span class="pre">of</span> <span class="pre">streaming</span> <span class="pre">units</span> <span class="pre">that</span> <span class="pre">the</span> <span class="pre">streaming</span> <span class="pre">job</span> <span class="pre">uses.</span> <span class="pre">Supported</span> <span class="pre">values</span> <span class="pre">are</span></code>1<code class="docutils literal notranslate"><span class="pre">,</span></code>3<code class="docutils literal notranslate"><span class="pre">,</span></code>6<code class="docutils literal notranslate"><span class="pre">and</span> <span class="pre">multiples</span> <span class="pre">of</span></code>6<code class="docutils literal notranslate"><span class="pre">up</span> <span class="pre">to</span></code>120`.
:param pulumi.Input[dict] tags: A mapping of tags assigned to the resource.
:param pulumi.Input[str] transformation_query: Specifies the query that will be run in the streaming job, <a class="reference external" href="https://msdn.microsoft.com/library/azure/dn834998">written in Stream Analytics Query Language (SAQL)</a>.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_job.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_job.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.Job.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.Job.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.Job.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputBlob">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">OutputBlob</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>date_format=None</em>, <em>name=None</em>, <em>path_pattern=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>storage_account_key=None</em>, <em>storage_account_name=None</em>, <em>storage_container_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>time_format=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Output to Blob Storage.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>date_format</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Output. Changing this forces a new resource to be created.</li>
<li><strong>path_pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>storage_account_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Access Key which should be used to connect to this Storage Account.</li>
<li><strong>storage_account_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Storage Account.</li>
<li><strong>storage_container_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Container within the Storage Account.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
<li><strong>time_format</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_blob.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_blob.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.date_format">
<code class="descname">date_format</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.date_format" title="Permalink to this definition">¶</a></dt>
<dd><p>The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Output. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.path_pattern">
<code class="descname">path_pattern</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.path_pattern" title="Permalink to this definition">¶</a></dt>
<dd><p>The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.storage_account_key">
<code class="descname">storage_account_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.storage_account_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The Access Key which should be used to connect to this Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.storage_account_name">
<code class="descname">storage_account_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.storage_account_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.storage_container_name">
<code class="descname">storage_container_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.storage_container_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Container within the Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputBlob.time_format">
<code class="descname">time_format</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.time_format" title="Permalink to this definition">¶</a></dt>
<dd><p>The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.OutputBlob.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>date_format=None</em>, <em>name=None</em>, <em>path_pattern=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>storage_account_key=None</em>, <em>storage_account_name=None</em>, <em>storage_container_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>time_format=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing OutputBlob resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] date_format: The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.
:param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
:param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
:param pulumi.Input[str] storage_account_name: The name of the Storage Account.
:param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
:param pulumi.Input[str] time_format: The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_blob.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_blob.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.OutputBlob.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputBlob.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputBlob.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputEventHub">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">OutputEventHub</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>eventhub_name=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Output to an EventHub.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>eventhub_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Event Hub.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Output. Changing this forces a new resource to be created.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>servicebus_namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>shared_access_policy_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy key for the specified shared access policy.</li>
<li><strong>shared_access_policy_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_eventhub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_eventhub.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.eventhub_name">
<code class="descname">eventhub_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.eventhub_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Event Hub.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Output. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.servicebus_namespace">
<code class="descname">servicebus_namespace</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.servicebus_namespace" title="Permalink to this definition">¶</a></dt>
<dd><p>The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.shared_access_policy_key">
<code class="descname">shared_access_policy_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.shared_access_policy_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy key for the specified shared access policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.shared_access_policy_name">
<code class="descname">shared_access_policy_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.shared_access_policy_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>eventhub_name=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing OutputEventHub resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] eventhub_name: The name of the Event Hub.
:param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
:param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_eventhub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_eventhub.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.OutputEventHub.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputEventHub.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputEventHub.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputMssql">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">OutputMssql</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>database=None</em>, <em>name=None</em>, <em>password=None</em>, <em>resource_group_name=None</em>, <em>server=None</em>, <em>stream_analytics_job_name=None</em>, <em>table=None</em>, <em>user=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Output to Microsoft SQL Server Database.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Output. Changing this forces a new resource to be created.</li>
<li><strong>password</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>server</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The SQL server url. Changing this forces a new resource to be created.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
<li><strong>table</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Table in the database that the output points to. Changing this forces a new resource to be created.</li>
<li><strong>user</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_mssql.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_mssql.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Output. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.password">
<code class="descname">password</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.password" title="Permalink to this definition">¶</a></dt>
<dd><p>Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.server">
<code class="descname">server</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.server" title="Permalink to this definition">¶</a></dt>
<dd><p>The SQL server url. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.table">
<code class="descname">table</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.table" title="Permalink to this definition">¶</a></dt>
<dd><p>Table in the database that the output points to. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputMssql.user">
<code class="descname">user</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.user" title="Permalink to this definition">¶</a></dt>
<dd><p>Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.OutputMssql.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>database=None</em>, <em>name=None</em>, <em>password=None</em>, <em>resource_group_name=None</em>, <em>server=None</em>, <em>stream_analytics_job_name=None</em>, <em>table=None</em>, <em>user=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing OutputMssql resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
:param pulumi.Input[str] password: Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[str] server: The SQL server url. Changing this forces a new resource to be created.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
:param pulumi.Input[str] table: Table in the database that the output points to. Changing this forces a new resource to be created.
:param pulumi.Input[str] user: Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_mssql.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_mssql.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.OutputMssql.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputMssql.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputMssql.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">OutputServiceBusQueue</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>name=None</em>, <em>queue_name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Output to a ServiceBus Queue.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Output. Changing this forces a new resource to be created.</li>
<li><strong>queue_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Service Bus Queue.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>servicebus_namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>shared_access_policy_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy key for the specified shared access policy.</li>
<li><strong>shared_access_policy_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_queue.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_queue.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Output. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.queue_name">
<code class="descname">queue_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.queue_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Service Bus Queue.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.servicebus_namespace">
<code class="descname">servicebus_namespace</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.servicebus_namespace" title="Permalink to this definition">¶</a></dt>
<dd><p>The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.shared_access_policy_key">
<code class="descname">shared_access_policy_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.shared_access_policy_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy key for the specified shared access policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.shared_access_policy_name">
<code class="descname">shared_access_policy_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.shared_access_policy_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>name=None</em>, <em>queue_name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing OutputServiceBusQueue resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
:param pulumi.Input[str] queue_name: The name of the Service Bus Queue.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
:param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_queue.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_queue.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.OutputServiceBusQueue.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.OutputServiceBusQueue.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputBlob">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">StreamInputBlob</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>date_format=None</em>, <em>name=None</em>, <em>path_pattern=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>storage_account_key=None</em>, <em>storage_account_name=None</em>, <em>storage_container_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>time_format=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Stream Input Blob.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>date_format</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Input Blob. Changing this forces a new resource to be created.</li>
<li><strong>path_pattern</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>storage_account_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The Access Key which should be used to connect to this Storage Account.</li>
<li><strong>storage_account_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Storage Account.</li>
<li><strong>storage_container_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Container within the Storage Account.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
<li><strong>time_format</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_blob.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_blob.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.date_format">
<code class="descname">date_format</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.date_format" title="Permalink to this definition">¶</a></dt>
<dd><p>The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Input Blob. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.path_pattern">
<code class="descname">path_pattern</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.path_pattern" title="Permalink to this definition">¶</a></dt>
<dd><p>The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.storage_account_key">
<code class="descname">storage_account_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.storage_account_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The Access Key which should be used to connect to this Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.storage_account_name">
<code class="descname">storage_account_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.storage_account_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.storage_container_name">
<code class="descname">storage_container_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.storage_container_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Container within the Storage Account.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.time_format">
<code class="descname">time_format</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.time_format" title="Permalink to this definition">¶</a></dt>
<dd><p>The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>date_format=None</em>, <em>name=None</em>, <em>path_pattern=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>storage_account_key=None</em>, <em>storage_account_name=None</em>, <em>storage_container_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>time_format=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing StreamInputBlob resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] date_format: The date format. Wherever <code class="docutils literal notranslate"><span class="pre">{date}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the date format instead.
:param pulumi.Input[str] name: The name of the Stream Input Blob. Changing this forces a new resource to be created.
:param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
:param pulumi.Input[str] storage_account_name: The name of the Storage Account.
:param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
:param pulumi.Input[str] time_format: The time format. Wherever <code class="docutils literal notranslate"><span class="pre">{time}</span></code> appears in <code class="docutils literal notranslate"><span class="pre">path_pattern</span></code>, the value of this property is used as the time format instead.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_blob.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_blob.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputBlob.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputBlob.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">StreamInputEventHub</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>eventhub_consumer_group_name=None</em>, <em>eventhub_name=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Stream Input EventHub.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>eventhub_consumer_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.</li>
<li><strong>eventhub_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Event Hub.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Input EventHub. Changing this forces a new resource to be created.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>servicebus_namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>shared_access_policy_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy key for the specified shared access policy.</li>
<li><strong>shared_access_policy_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_eventhub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_eventhub.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.eventhub_consumer_group_name">
<code class="descname">eventhub_consumer_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.eventhub_consumer_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.eventhub_name">
<code class="descname">eventhub_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.eventhub_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Event Hub.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Input EventHub. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.servicebus_namespace">
<code class="descname">servicebus_namespace</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.servicebus_namespace" title="Permalink to this definition">¶</a></dt>
<dd><p>The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.shared_access_policy_key">
<code class="descname">shared_access_policy_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.shared_access_policy_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy key for the specified shared access policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.shared_access_policy_name">
<code class="descname">shared_access_policy_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.shared_access_policy_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>eventhub_consumer_group_name=None</em>, <em>eventhub_name=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>servicebus_namespace=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing StreamInputEventHub resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
:param pulumi.Input[str] eventhub_name: The name of the Event Hub.
:param pulumi.Input[str] name: The name of the Stream Input EventHub. Changing this forces a new resource to be created.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
:param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_eventhub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_eventhub.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputEventHub.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputEventHub.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub">
<em class="property">class </em><code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">StreamInputIotHub</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>endpoint=None</em>, <em>eventhub_consumer_group_name=None</em>, <em>iothub_namespace=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a Stream Analytics Stream Input IoTHub.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>endpoint</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).</li>
<li><strong>eventhub_consumer_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.</li>
<li><strong>iothub_namespace</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name or the URI of the IoT Hub.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Input IoTHub. Changing this forces a new resource to be created.</li>
<li><strong>resource_group_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</li>
<li><strong>serialization</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</li>
<li><strong>shared_access_policy_key</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy key for the specified shared access policy.</li>
<li><strong>shared_access_policy_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</li>
<li><strong>stream_analytics_job_name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Stream Analytics Job. Changing this forces a new resource to be created.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_iothub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_iothub.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.endpoint">
<code class="descname">endpoint</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.endpoint" title="Permalink to this definition">¶</a></dt>
<dd><p>The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.eventhub_consumer_group_name">
<code class="descname">eventhub_consumer_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.eventhub_consumer_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.iothub_namespace">
<code class="descname">iothub_namespace</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.iothub_namespace" title="Permalink to this definition">¶</a></dt>
<dd><p>The name or the URI of the IoT Hub.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Input IoTHub. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.resource_group_name">
<code class="descname">resource_group_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.resource_group_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.serialization">
<code class="descname">serialization</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.serialization" title="Permalink to this definition">¶</a></dt>
<dd><p>A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.shared_access_policy_key">
<code class="descname">shared_access_policy_key</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.shared_access_policy_key" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy key for the specified shared access policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.shared_access_policy_name">
<code class="descname">shared_access_policy_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.shared_access_policy_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.stream_analytics_job_name">
<code class="descname">stream_analytics_job_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.stream_analytics_job_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>endpoint=None</em>, <em>eventhub_consumer_group_name=None</em>, <em>iothub_namespace=None</em>, <em>name=None</em>, <em>resource_group_name=None</em>, <em>serialization=None</em>, <em>shared_access_policy_key=None</em>, <em>shared_access_policy_name=None</em>, <em>stream_analytics_job_name=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing StreamInputIotHub resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] endpoint: The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
:param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
:param pulumi.Input[str] iothub_namespace: The name or the URI of the IoT Hub.
:param pulumi.Input[str] name: The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
:param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
:param pulumi.Input[dict] serialization: A <code class="docutils literal notranslate"><span class="pre">serialization</span></code> block as defined below.
:param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
:param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
:param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_iothub.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_iothub.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.StreamInputIotHub.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.StreamInputIotHub.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_azure.streamanalytics.get_job">
<code class="descclassname">pulumi_azure.streamanalytics.</code><code class="descname">get_job</code><span class="sig-paren">(</span><em>name=None</em>, <em>resource_group_name=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_azure.streamanalytics.get_job" title="Permalink to this definition">¶</a></dt>
<dd><p>Use this data source to access information about an existing Stream Analytics Job.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/stream_analytics_job.html.markdown">https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/stream_analytics_job.html.markdown</a>.</div></blockquote>
</dd></dl>

</div>
