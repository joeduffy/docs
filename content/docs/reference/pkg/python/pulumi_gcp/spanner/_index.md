---
title: Module spanner
---

<div class="section" id="spanner">
<h1>spanner<a class="headerlink" href="#spanner" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-gcp/issues">pulumi/pulumi-gcp repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/issues">terraform-providers/terraform-provider-google repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_gcp.spanner"></span><dl class="class">
<dt id="pulumi_gcp.spanner.Database">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">Database</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>ddls=None</em>, <em>instance=None</em>, <em>name=None</em>, <em>project=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Database" title="Permalink to this definition">¶</a></dt>
<dd><p>A Cloud Spanner Database which is hosted on a Spanner instance.</p>
<p>To get more information about Database, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/spanner/">Official Documentation</a></li>
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
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.Database.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.Database.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.Database.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>ddls=None</em>, <em>instance=None</em>, <em>name=None</em>, <em>project=None</em>, <em>state=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Database.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Database resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs.</p>
<blockquote>
<div>If it is not provided, the provider project is used.</div></blockquote>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.Database.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Database.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.Database.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Database.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">DatabaseIAMBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>database=None</em>, <em>instance=None</em>, <em>members=None</em>, <em>project=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your database using <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>database</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner database.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner instance the database belongs to.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.database">
<code class="descname">database</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.database" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner database.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the database’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner instance the database belongs to.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>database=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>members=None</em>, <em>project=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatabaseIAMBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] database: The name of the Spanner database.
:param pulumi.Input[str] etag: (Computed) The etag of the database’s IAM policy.
:param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it</p>
<blockquote>
<div>is not provided, the provider project is used.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">DatabaseIAMMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>database=None</em>, <em>instance=None</em>, <em>member=None</em>, <em>project=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your database using <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>database</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner database.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner instance the database belongs to.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.database">
<code class="descname">database</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.database" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner database.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the database’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner instance the database belongs to.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>database=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>member=None</em>, <em>project=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatabaseIAMMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] database: The name of the Spanner database.
:param pulumi.Input[str] etag: (Computed) The etag of the database’s IAM policy.
:param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it</p>
<blockquote>
<div>is not provided, the provider project is used.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">DatabaseIAMPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>database=None</em>, <em>instance=None</em>, <em>policy_data=None</em>, <em>project=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your database using <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.DatabaseIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>database</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner database.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the Spanner instance the database belongs to.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.database">
<code class="descname">database</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.database" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner database.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the database’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the Spanner instance the database belongs to.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>database=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>policy_data=None</em>, <em>project=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing DatabaseIAMPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] database: The name of the Spanner database.
:param pulumi.Input[str] etag: (Computed) The etag of the database’s IAM policy.
:param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
:param pulumi.Input[str] policy_data: The policy data generated by</p>
<blockquote>
<div>a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.DatabaseIAMPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.DatabaseIAMPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.Instance">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">Instance</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>config=None</em>, <em>display_name=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>num_nodes=None</em>, <em>project=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Instance" title="Permalink to this definition">¶</a></dt>
<dd><p>An isolated set of Cloud Spanner resources on which databases can be
hosted.</p>
<p>To get more information about Instance, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/spanner/">Official Documentation</a></li>
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
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.Instance.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.Instance.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.Instance.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>config=None</em>, <em>display_name=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>num_nodes=None</em>, <em>project=None</em>, <em>state=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Instance.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Instance resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs.</p>
<blockquote>
<div>If it is not provided, the provider project is used.</div></blockquote>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.Instance.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Instance.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.Instance.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.Instance.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMBinding">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">InstanceIAMBinding</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>instance=None</em>, <em>members=None</em>, <em>project=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your instance using <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the instance.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_binding.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the instance’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the instance.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>members=None</em>, <em>project=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing InstanceIAMBinding resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the instance’s IAM policy.
:param pulumi.Input[str] instance: The name of the instance.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it</p>
<blockquote>
<div>is not provided, the provider project is used.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_binding.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_binding.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMBinding.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMBinding.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMMember">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">InstanceIAMMember</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>instance=None</em>, <em>member=None</em>, <em>project=None</em>, <em>role=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your instance using <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the instance.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_member.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the instance’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the instance.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.role">
<code class="descname">role</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.role" title="Permalink to this definition">¶</a></dt>
<dd><p>The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>member=None</em>, <em>project=None</em>, <em>role=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing InstanceIAMMember resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the instance’s IAM policy.
:param pulumi.Input[str] instance: The name of the instance.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it</p>
<blockquote>
<div>is not provided, the provider project is used.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>role</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The role that should be applied. Only one
<code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> can be used per role. Note that custom roles must be of the format
<code class="docutils literal notranslate"><span class="pre">[projects|organizations]/{parent-name}/roles/{role-name}</span></code>.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_member.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_member.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.InstanceIAMMember.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMMember.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMMember.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.spanner.</code><code class="descname">InstanceIAMPolicy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>instance=None</em>, <em>policy_data=None</em>, <em>project=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy" title="Permalink to this definition">¶</a></dt>
<dd><p>Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:</p>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.</li>
</ul>
<blockquote>
<div><strong>Warning:</strong> It’s entirely possibly to lock yourself out of your instance using <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code>. Any permissions granted by default will be removed unless you include them in your config.</div></blockquote>
<ul class="simple">
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code>: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.</li>
<li><code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code>: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.</li>
</ul>
<blockquote>
<div><p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMPolicy</span></code> <strong>cannot</strong> be used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> and <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> or they will fight over what your policy should be.</p>
<p><strong>Note:</strong> <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMBinding</span></code> resources <strong>can be</strong> used in conjunction with <code class="docutils literal notranslate"><span class="pre">spanner.InstanceIAMMember</span></code> resources <strong>only if</strong> they do not grant privilege to the same role.</p>
</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>instance</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the instance.</li>
<li><strong>policy_data</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.etag">
<code class="descname">etag</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.etag" title="Permalink to this definition">¶</a></dt>
<dd><p>(Computed) The etag of the instance’s IAM policy.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.instance">
<code class="descname">instance</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.instance" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the instance.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.policy_data">
<code class="descname">policy_data</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.policy_data" title="Permalink to this definition">¶</a></dt>
<dd><p>The policy data generated by
a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>etag=None</em>, <em>instance=None</em>, <em>policy_data=None</em>, <em>project=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing InstanceIAMPolicy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] etag: (Computed) The etag of the instance’s IAM policy.
:param pulumi.Input[str] instance: The name of the instance.
:param pulumi.Input[str] policy_data: The policy data generated by</p>
<blockquote>
<div>a <code class="docutils literal notranslate"><span class="pre">organizations.getIAMPolicy</span></code> data source.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.spanner.InstanceIAMPolicy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.spanner.InstanceIAMPolicy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
