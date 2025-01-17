---
title: Module healthcare
---

<div class="section" id="healthcare">
<h1>healthcare<a class="headerlink" href="#healthcare" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-gcp/issues">pulumi/pulumi-gcp repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/issues">terraform-providers/terraform-provider-google repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_gcp.healthcare"></span><dl class="class">
<dt id="pulumi_gcp.healthcare.Dataset">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">Dataset</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>location=None</em>, <em>name=None</em>, <em>project=None</em>, <em>time_zone=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Dataset" title="Permalink to this definition">¶</a></dt>
<dd><p>A Healthcare <code class="docutils literal notranslate"><span class="pre">Dataset</span></code> is a toplevel logical grouping of <code class="docutils literal notranslate"><span class="pre">dicomStores</span></code>, <code class="docutils literal notranslate"><span class="pre">fhirStores</span></code> and <code class="docutils literal notranslate"><span class="pre">hl7V2Stores</span></code>.</p>
<p>To get more information about Dataset, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/how-tos/datasets">Creating a dataset</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Dataset.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Dataset.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.Dataset.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>location=None</em>, <em>name=None</em>, <em>project=None</em>, <em>self_link=None</em>, <em>time_zone=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Dataset.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Dataset resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs.</p>
<blockquote>
<div>If it is not provided, the provider project is used.</div></blockquote>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.Dataset.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Dataset.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Dataset.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Dataset.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DatasetIamBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>members=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DatasetIamBinding resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dataset_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.dataset_id">
<code class="descname">dataset_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.dataset_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the dataset’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>etag=None</em>, <em>members=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatasetIamBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dataset_id: The dataset ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the dataset’s IAM policy.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DatasetIamMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>member=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DatasetIamMember resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dataset_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamMember.dataset_id">
<code class="descname">dataset_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.dataset_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the dataset’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DatasetIamMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>etag=None</em>, <em>member=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatasetIamMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dataset_id: The dataset ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the dataset’s IAM policy.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DatasetIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DatasetIamMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DatasetIamPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>policy_data=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DatasetIamPolicy resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dataset_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.dataset_id">
<code class="descname">dataset_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.dataset_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The dataset ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the dataset’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset_id=None</em>, <em>etag=None</em>, <em>policy_data=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatasetIamPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dataset_id: The dataset ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the dataset’s IAM policy.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DatasetIamPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DatasetIamPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStore">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DicomStore</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStore" title="Permalink to this definition">¶</a></dt>
<dd><p>A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
(<a class="reference external" href="https://www.dicomstandard.org/about/">https://www.dicomstandard.org/about/</a>) standard for Healthcare information exchange</p>
<p>To get more information about DicomStore, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.dicomStores">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/how-tos/dicom">Creating a DICOM store</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store.html.markdown</a>.</div></blockquote>
<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DicomStore.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>self_link=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStore.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DicomStore resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DicomStore.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStore.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStore.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStore.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DicomStoreIamBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>members=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DicomStoreIamBinding resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dicom_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.dicom_store_id">
<code class="descname">dicom_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.dicom_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the DICOM store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>etag=None</em>, <em>members=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DicomStoreIamBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the DICOM store’s IAM policy.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DicomStoreIamMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>member=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DicomStoreIamMember resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dicom_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.dicom_store_id">
<code class="descname">dicom_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.dicom_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the DICOM store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>etag=None</em>, <em>member=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DicomStoreIamMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the DICOM store’s IAM policy.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.DicomStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">DicomStoreIamPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>policy_data=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a DicomStoreIamPolicy resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dicom_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.dicom_store_id">
<code class="descname">dicom_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.dicom_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The DICOM store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the DICOM store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dicom_store_id=None</em>, <em>etag=None</em>, <em>policy_data=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DicomStoreIamPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{dicom_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{dicom_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>etag</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – (Computed) The etag of the DICOM store’s IAM policy.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.DicomStoreIamPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.DicomStoreIamPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStore">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">FhirStore</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>disable_referential_integrity=None</em>, <em>disable_resource_versioning=None</em>, <em>enable_history_import=None</em>, <em>enable_update_create=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStore" title="Permalink to this definition">¶</a></dt>
<dd><p>A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (<a class="reference external" href="https://www.hl7.org/fhir/STU3/">https://www.hl7.org/fhir/STU3/</a>)
standard for Healthcare information exchange</p>
<p>To get more information about FhirStore, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/how-tos/fhir">Creating a FHIR store</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown</a>.</div></blockquote>
<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.FhirStore.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>disable_referential_integrity=None</em>, <em>disable_resource_versioning=None</em>, <em>enable_history_import=None</em>, <em>enable_update_create=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>self_link=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStore.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing FhirStore resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.FhirStore.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStore.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStore.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStore.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">FhirStoreIamBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>fhir_store_id=None</em>, <em>members=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a FhirStoreIamBinding resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>fhir_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the FHIR store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.fhir_store_id">
<code class="descname">fhir_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.fhir_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>fhir_store_id=None</em>, <em>members=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing FhirStoreIamBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the FHIR store’s IAM policy.
:param pulumi.Input[str] fhir_store_id: The FHIR store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">FhirStoreIamMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>fhir_store_id=None</em>, <em>member=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a FhirStoreIamMember resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>fhir_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the FHIR store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.fhir_store_id">
<code class="descname">fhir_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.fhir_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>fhir_store_id=None</em>, <em>member=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing FhirStoreIamMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the FHIR store’s IAM policy.
:param pulumi.Input[str] fhir_store_id: The FHIR store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.FhirStoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">FhirStoreIamPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>fhir_store_id=None</em>, <em>policy_data=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a FhirStoreIamPolicy resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>fhir_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the FHIR store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.fhir_store_id">
<code class="descname">fhir_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.fhir_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The FHIR store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>fhir_store_id=None</em>, <em>policy_data=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing FhirStoreIamPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the FHIR store’s IAM policy.
:param pulumi.Input[str] fhir_store_id: The FHIR store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{fhir_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{fhir_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.FhirStoreIamPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.FhirStoreIamPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7Store">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">Hl7Store</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>parser_config=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7Store" title="Permalink to this definition">¶</a></dt>
<dd><p>A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (<a class="reference external" href="https://www.hl7.org/hl7V2/STU3/">https://www.hl7.org/hl7V2/STU3/</a>)
standard for Healthcare information exchange</p>
<p>To get more information about Hl7V2Store, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.hl7V2Stores">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/healthcare/docs/how-tos/hl7v2">Creating a HL7v2 Store</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown</a>.</div></blockquote>
<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.Hl7Store.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dataset=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>notification_config=None</em>, <em>parser_config=None</em>, <em>self_link=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7Store.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Hl7Store resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.Hl7Store.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7Store.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7Store.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7Store.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">Hl7StoreIamBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>hl7_v2_store_id=None</em>, <em>members=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a Hl7StoreIamBinding resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>hl7_v2_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the HL7v2 store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.hl7_v2_store_id">
<code class="descname">hl7_v2_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.hl7_v2_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>hl7_v2_store_id=None</em>, <em>members=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Hl7StoreIamBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the HL7v2 store’s IAM policy.
:param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">Hl7StoreIamMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>hl7_v2_store_id=None</em>, <em>member=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a Hl7StoreIamMember resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>hl7_v2_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the HL7v2 store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.hl7_v2_store_id">
<code class="descname">hl7_v2_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.hl7_v2_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>hl7_v2_store_id=None</em>, <em>member=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Hl7StoreIamMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the HL7v2 store’s IAM policy.
:param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">healthcare.Hl7StoreIamBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.healthcare.</code><code class="descname">Hl7StoreIamPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>hl7_v2_store_id=None</em>, <em>policy_data=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Create a Hl7StoreIamPolicy resource with the given unique name, props, and options.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>hl7_v2_store_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the HL7v2 store’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.hl7_v2_store_id">
<code class="descname">hl7_v2_store_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.hl7_v2_store_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The HL7v2 store ID, in the form
<code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>hl7_v2_store_id=None</em>, <em>policy_data=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Hl7StoreIamPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the HL7v2 store’s IAM policy.
:param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form</p>
<blockquote>
<div><code class="docutils literal notranslate"><span class="pre">{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code> or
<code class="docutils literal notranslate"><span class="pre">{location_name}/{dataset_name}/{hl7_v2_store_name}</span></code>. In the second form, the provider’s
project setting will be used as a fallback.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.healthcare.Hl7StoreIamPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.healthcare.Hl7StoreIamPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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

</div>
